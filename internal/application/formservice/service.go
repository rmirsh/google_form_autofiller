package formservice

import (
	"fmt"
	"strings"
)

type Service struct {
	fetcher   Fetcher
	parser    Parser
	filler    Filler
	submitter Submitter
}

func NewService(fetcher Fetcher, parser Parser, filler Filler, submitter Submitter) *Service {
	return &Service{
		fetcher:   fetcher,
		parser:    parser,
		filler:    filler,
		submitter: submitter,
	}
}

func (s *Service) AutoFillAndSubmit(formURL string) error {
	html, err := s.fetcher.Fetch(formURL)
	if err != nil {
		return fmt.Errorf("failed to fetch form: %w", err)
	}

	formData, err := s.parser.Parse(html)
	if err != nil {
		return fmt.Errorf("failed to parse form: %w", err)
	}

	payload := make(map[string]string)
	for _, field := range formData.Fields {
		value := s.filler.Fill(field)

		switch v := value.(type) {
		case string:
			payload["entry."+field.ID] = v
		case []string:
			for _, val := range v {
				payload["entry."+field.ID] = val
			}
		default:
			payload["entry."+field.ID] = fmt.Sprintf("%v", v)
		}
	}

	submitURL := buildSubmitURL(formURL)

	return s.submitter.Submit(submitURL, payload)
}

func buildSubmitURL(formURL string) string {
	if strings.HasSuffix(formURL, "/viewform") {
		return strings.TrimSuffix(formURL, "/viewform") + "/formResponse"
	}
	return formURL
}
