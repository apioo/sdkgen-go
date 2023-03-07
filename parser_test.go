package sdkgen

import (
	"testing"
)

type Entry struct {
	Path       string
	Parameters map[string]interface{}
	Expect     string
}

func TestUrl(t *testing.T) {
	var parser = NewParser("https://api.acme.com/")

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

func Map(key string, value interface{}) map[string]interface{} {
	var params = make(map[string]interface{})
	params[key] = value
	return params
}
