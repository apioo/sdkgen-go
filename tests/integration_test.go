package tests

import (
	"encoding/json"
	"github.com/apioo/sdkgen-go"
	"github.com/apioo/sdkgen-go/tests/generated"
	"net/url"
	"strings"
	"testing"
)

func TestClientGetAll(t *testing.T) {
	client, _ := generated.Build("my_token")

	response, _ := client.Product().GetAll(8, 16, "foobar")

	headers := *response.Headers
	args := *response.Args

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "GET")
	AssertEquals(t, args["startIndex"], "8")
	AssertEquals(t, args["count"], "16")
	AssertEquals(t, args["search"], "foobar")
}

func TestClientCreate(t *testing.T) {
	client, _ := generated.Build("my_token")

	payload := NewPayload()
	response, _ := client.Product().Create(payload)

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"dateString\":\"2024-09-22\",\"dateTimeString\":\"2024-09-22T10:09:00\",\"timeString\":\"10:09:00\",\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientUpdate(t *testing.T) {
	client, _ := generated.Build("my_token")

	payload := NewPayload()
	response, _ := client.Product().Update(1, payload)

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "PUT")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"dateString\":\"2024-09-22\",\"dateTimeString\":\"2024-09-22T10:09:00\",\"timeString\":\"10:09:00\",\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientPatch(t *testing.T) {
	client, _ := generated.Build("my_token")

	payload := NewPayload()
	response, _ := client.Product().Patch(1, payload)

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "PATCH")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"dateString\":\"2024-09-22\",\"dateTimeString\":\"2024-09-22T10:09:00\",\"timeString\":\"10:09:00\",\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientDelete(t *testing.T) {
	client, _ := generated.Build("my_token")

	response, _ := client.Product().Delete(1)

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "DELETE")
}

func TestClientBinary(t *testing.T) {
	client, _ := generated.Build("my_token")

	var payload = []byte{0x66, 0x6F, 0x6F, 0x62, 0x61, 0x72}

	response, _ := client.Product().Binary(payload)

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, response.Data, "foobar")
}

func TestClientForm(t *testing.T) {
	client, _ := generated.Build("my_token")

	var payload = url.Values{}
	payload.Set("foo", "bar")

	response, _ := client.Product().Form(payload)

	headers := *response.Headers
	form := *response.Form

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, form["foo"], "bar")
}

func TestClientJson(t *testing.T) {
	client, _ := generated.Build("my_token")

	payload := map[string]string{"string": "bar"}

	response, _ := client.Product().Json(payload)

	headers := *response.Headers
	json := *response.Json

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, json.String, "bar")
}

func TestClientMultipart(t *testing.T) {
	client, _ := generated.Build("my_token")

	var payload = &sdkgen.Multipart{}
	payload.AddFile("foo", "upload.txt", strings.NewReader("foobar"))

	response, _ := client.Product().Multipart(payload)

	headers := *response.Headers
	files := *response.Files

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, files["foo"], "foobar")
}

func TestClientText(t *testing.T) {
	client, _ := generated.Build("my_token")

	response, _ := client.Product().Text("foobar")

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, response.Data, "foobar")
}

func TestClientXml(t *testing.T) {
	client, _ := generated.Build("my_token")

	response, _ := client.Product().Xml("<foo>bar</foo>")

	headers := *response.Headers

	AssertEquals(t, headers["Authorization"], "Bearer my_token")
	AssertEquals(t, headers["Accept"], "application/json")
	AssertEquals(t, headers["User-Agent"], "SDKgen Client v2.0")
	AssertEquals(t, response.Method, "POST")
	AssertEquals(t, response.Data, "<foo>bar</foo>")
}

func AssertEquals(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func AssertJson(t *testing.T, model *generated.TestRequest, want string) {
	raw, _ := json.Marshal(model)
	var got = string(raw)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func NewPayload() generated.TestRequest {
	var objectFoo = generated.TestObject{
		Id:   1,
		Name: "foo",
	}

	var objectBar = generated.TestObject{
		Id:   2,
		Name: "bar",
	}

	var arrayScalar = []string{"foo", "bar"}
	var arrayObject = []generated.TestObject{objectFoo, objectBar}

	var mapScalar = make(map[string]string)
	mapScalar["foo"] = "bar"
	mapScalar["bar"] = "foo"

	var mapObject = make(map[string]generated.TestObject)
	mapObject["foo"] = objectFoo
	mapObject["bar"] = objectBar

	return generated.TestRequest{
		Int:            1337,
		Float:          13.37,
		String:         "foobar",
		Bool:           true,
		DateString:     "2024-09-22",
		DateTimeString: "2024-09-22T10:09:00",
		TimeString:     "10:09:00",
		ArrayScalar:    arrayScalar,
		ArrayObject:    arrayObject,
		MapScalar:      &mapScalar,
		MapObject:      &mapObject,
		Object:         &objectFoo,
	}
}
