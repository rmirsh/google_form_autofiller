package parser

import (
	"encoding/json"
	"errors"
	"regexp"
)

const DataVarName = "FB_PUBLIC_LOAD_DATA_"

func ExtractFormData(html string) (any, error) {
	re := regexp.MustCompile(`var\s+` + DataVarName + `\s*=\s*(.*?);`)
	matches := re.FindStringSubmatch(html)
	if len(matches) < 2 {
		return nil, errors.New("can't find form data variable")
	}

	var data any
	if err := json.Unmarshal([]byte(matches[1]), &data); err != nil {
		return nil, err
	}

	return data, nil
}
