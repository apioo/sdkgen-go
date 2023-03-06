package sdkgen

import (
	"encoding/json"
	"github.com/go-chrono/chrono"
	"reflect"
	"strings"
	"time"
)

type Parser struct {
	BaseUrl string
}

func (parser *Parser) Url(path string, parameters map[string]interface{}) string {
	return parser.BaseUrl + "/" + parser.SubstituteParameters(path, parameters)
}

func (parser *Parser) Parse(data string, model *interface{}) error {
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		return err
	}
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

		value, ok := parameters[name]
		if ok {
			part = ToString(value)
		}

		result = append(result, part)
	}

	return strings.Join(result, "/")
}

func ToString(value interface{}) string {
	if reflect.TypeOf(value).Name() == "string" {
		return value.(string)
	} else if reflect.TypeOf(value).Name() == "float32" || reflect.TypeOf(value).Name() == "float64" {
		return value.(string)
	} else if reflect.TypeOf(value).Name() == "int" || reflect.TypeOf(value).Name() == "int8" || reflect.TypeOf(value).Name() == "int16" || reflect.TypeOf(value).Name() == "int32" || reflect.TypeOf(value).Name() == "int64" {
		return value.(string)
	} else if reflect.TypeOf(value).Name() == "bool" {
		if value.(bool) {
			return "1"
		} else {
			return "0"
		}
	} else if reflect.TypeOf(value).Name() == "chrono.LocalDate" {
		return value.(chrono.LocalDate).String()
	} else if reflect.TypeOf(value).Name() == "chrono.LocalTime" {
		return value.(chrono.LocalTime).String()
	} else if reflect.TypeOf(value).Name() == "chrono.LocalDateTime" {
		return value.(chrono.LocalDateTime).String()
	} else if reflect.TypeOf(value).Name() == "time.Time" {
		return value.(time.Time).Format(time.RFC3339)
	} else {
		return ""
	}
}

func NewParser(baseUrl string) *Parser {
	return &Parser{
		BaseUrl: strings.TrimRight(baseUrl, "/"),
	}
}
