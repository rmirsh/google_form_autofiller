package formservice

import "github.com/rmirsh/google_form_autofiller/internal/domain/form"

type Fetcher interface {
	Fetch(url string) (string, error)
}

type Parser interface {
	Parse(html string) (form.Form, error)
}

type Filler interface {
	Fill(field Form.Field) any
}

type Submitter interface {
	Submit(url string, payload map[string]string) error
}
