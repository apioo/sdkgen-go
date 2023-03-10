package sdkgen

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
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

func (parser *Parser) Query(parameters map[string]interface{}) url.Values {
	var result = url.Values{}
	for name, value := range parameters {
		if value == "" {
			continue
		}

		result.Add(name, ToString(value))
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
			if pos != -1 {
				name = part[1:pos]
			} else {
				name = part[1:]
			}
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
	if value == nil {
		return ""
	}

	if reflect.TypeOf(value).Name() == "string" {
		return value.(string)
	} else if reflect.TypeOf(value).Name() == "float32" || reflect.TypeOf(value).Name() == "float64" {
		return fmt.Sprintf("%g", value)
	} else if reflect.TypeOf(value).Name() == "int" {
		return strconv.FormatInt(int64(value.(int)), 10)
	} else if reflect.TypeOf(value).Name() == "int8" {
		return strconv.FormatInt(int64(value.(int8)), 10)
	} else if reflect.TypeOf(value).Name() == "int16" {
		return strconv.FormatInt(int64(value.(int16)), 10)
	} else if reflect.TypeOf(value).Name() == "int32" {
		return strconv.FormatInt(int64(value.(int32)), 10)
	} else if reflect.TypeOf(value).Name() == "int64" {
		return strconv.FormatInt(value.(int64), 10)
	} else if reflect.TypeOf(value).Name() == "bool" {
		if value.(bool) {
			return "1"
		} else {
			return "0"
		}
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
