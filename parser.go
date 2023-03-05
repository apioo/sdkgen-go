package sdkgen

import "strings"

type Parser struct {
	BaseUrl string
}

func (parser *Parser) Url(path string, parameters map[string]interface{}) string {
	return parser.BaseUrl + "/" + parser.SubstituteParameters(path, parameters)
}

func (parser *Parser) Parse(data string) interface{} {
	return nil
}

func (parser *Parser) Query(parameters map[string]interface{}) interface{} {
	var result map[string]string
	for name, value := range parameters {
		if value == "" {
			continue
		}

		result[name] = ToString(value)
	}

	return result
}

func (parser *Parser) SubstituteParameters(path string, parameters map[string]interface{}) string {
	var parts = strings.Split(path, "/")
	var result []string

	for _, part := range parts {
		if part == "" {
			continue
		}

		var name string
		if strings.HasPrefix(part, ":") {
			name = part[1:]
		} else if strings.HasPrefix(part, "$") {
			var pos = strings.Index(part, "<")
			name = part[1 : pos-1]
		} else if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			name = part[1 : len(part)-1]
		}

		if name != "" {
			part = parameters[name].(string)
		}

		result = append(result, part)
	}

	return path + "/" + strings.Join(result, "/")
}

func ToString(value interface{}) string {
	return value.(string)
}

func NewParser(baseUrl string) *Parser {
	return &Parser{
		BaseUrl: strings.TrimRight(baseUrl, "/"),
	}
}
