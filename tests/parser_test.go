package tests

import (
	"github.com/apioo/sdkgen-go"
	"github.com/apioo/sdkgen-go/tests/generated"
	"testing"
	"time"
)

type Entry struct {
	Path       string
	Parameters map[string]interface{}
	Expect     string
}

func TestUrl(t *testing.T) {
	var parser = sdkgen.NewParser("https://api.acme.com/")

	var tests []Entry
	tests = append(tests, Entry{Path: "/foo/bar", Parameters: nil, Expect: "https://api.acme.com/foo/bar"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo"})
	tests = append(tests, Entry{Path: "/foo/$bar<[0-9]+>", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo"})
	tests = append(tests, Entry{Path: "/foo/$bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo"})
	tests = append(tests, Entry{Path: "/foo/{bar}", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo"})
	tests = append(tests, Entry{Path: "/foo/:bar/bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo/bar"})
	tests = append(tests, Entry{Path: "/foo/$bar<[0-9]+>/bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo/bar"})
	tests = append(tests, Entry{Path: "/foo/$bar/bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo/bar"})
	tests = append(tests, Entry{Path: "/foo/{bar}/bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo/bar"})

	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", nil), Expect: "https://api.acme.com/foo/"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", 1337), Expect: "https://api.acme.com/foo/1337"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", 13.37), Expect: "https://api.acme.com/foo/13.37"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", true), Expect: "https://api.acme.com/foo/1"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", false), Expect: "https://api.acme.com/foo/0"})
	tests = append(tests, Entry{Path: "/foo/:bar", Parameters: Map("bar", "foo"), Expect: "https://api.acme.com/foo/foo"})

	for _, test := range tests {
		got := parser.Url(test.Path, test.Parameters)
		want := test.Expect
		if got != test.Expect {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
}

func TestQuery(t *testing.T) {
	var parser = sdkgen.NewParser("https://api.acme.com/")

	var test = generated.TestObject{
		Name: "foo",
	}

	var parameters = make(map[string]interface{})
	parameters["null"] = nil
	parameters["int"] = 1337
	parameters["float"] = 13.37
	parameters["true"] = true
	parameters["false"] = false
	parameters["string"] = "foo"
	parameters["date"] = time.Date(2023, 2, 21, 0, 0, 0, 0, time.UTC)
	parameters["datetime"] = time.Date(2023, 2, 21, 19, 19, 0, 0, time.UTC)
	parameters["time"] = time.Date(1970, 1, 1, 19, 19, 0, 0, time.UTC)
	parameters["args"] = test

	var result = parser.QueryWithStruct(parameters, []string{"args"})

	AssertEquals(t, result.Get("int"), "1337")
	AssertEquals(t, result.Get("float"), "13.37")
	AssertEquals(t, result.Get("true"), "1")
	AssertEquals(t, result.Get("false"), "0")
	AssertEquals(t, result.Get("string"), "foo")
	AssertEquals(t, result.Get("date"), "2023-02-21")
	AssertEquals(t, result.Get("datetime"), "2023-02-21T19:19:00Z")
	AssertEquals(t, result.Get("time"), "19:19:00")
	AssertEquals(t, result.Get("name"), "foo")
}

func Map(key string, value interface{}) map[string]interface{} {
	var params = make(map[string]interface{})
	params[key] = value
	return params
}
