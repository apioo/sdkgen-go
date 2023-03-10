package sdkgen

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClientGetAll(t *testing.T) {
	client, _ := Build("my_token")

	response, _ := client.GetAll(8, 16, "foobar")

	AssertEquals(t, response.Headers["Authorization"], "Bearer my_token")
	AssertEquals(t, response.Headers["Accept"], "application/json")
	AssertEquals(t, response.Headers["User-Agent"], "SDKgen Client v1.0")
	AssertEquals(t, response.Method, "GET")
	AssertEquals(t, response.Args["startIndex"], "8")
	AssertEquals(t, response.Args["count"], "16")
	AssertEquals(t, response.Args["search"], "foobar")
	AssertJson(t, response.Json, "{\"int\":0,\"float\":0,\"string\":\"\",\"bool\":false,\"arrayScalar\":null,\"arrayObject\":null,\"mapScalar\":null,\"mapObject\":null,\"object\":{\"id\":0,\"name\":\"\"}}")
}

func TestClientCreate(t *testing.T) {
	client, _ := Build("my_token")

	payload := NewPayload()
	response, _ := client.Create(payload)

	AssertEquals(t, response.Headers["Authorization"], "Bearer my_token")
	AssertEquals(t, response.Headers["Accept"], "application/json")
	AssertEquals(t, response.Headers["User-Agent"], "SDKgen Client v1.0")
	AssertEquals(t, response.Method, "POST")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientUpdate(t *testing.T) {
	client, _ := Build("my_token")

	payload := NewPayload()
	response, _ := client.Update(1, payload)

	AssertEquals(t, response.Headers["Authorization"], "Bearer my_token")
	AssertEquals(t, response.Headers["Accept"], "application/json")
	AssertEquals(t, response.Headers["User-Agent"], "SDKgen Client v1.0")
	AssertEquals(t, response.Method, "PUT")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientPatch(t *testing.T) {
	client, _ := Build("my_token")

	payload := NewPayload()
	response, _ := client.Patch(1, payload)

	AssertEquals(t, response.Headers["Authorization"], "Bearer my_token")
	AssertEquals(t, response.Headers["Accept"], "application/json")
	AssertEquals(t, response.Headers["User-Agent"], "SDKgen Client v1.0")
	AssertEquals(t, response.Method, "PATCH")
	AssertJson(t, response.Json, "{\"int\":1337,\"float\":13.37,\"string\":\"foobar\",\"bool\":true,\"arrayScalar\":[\"foo\",\"bar\"],\"arrayObject\":[{\"id\":1,\"name\":\"foo\"},{\"id\":2,\"name\":\"bar\"}],\"mapScalar\":{\"bar\":\"foo\",\"foo\":\"bar\"},\"mapObject\":{\"bar\":{\"id\":2,\"name\":\"bar\"},\"foo\":{\"id\":1,\"name\":\"foo\"}},\"object\":{\"id\":1,\"name\":\"foo\"}}")
}

func TestClientDelete(t *testing.T) {
	client, _ := Build("my_token")

	response, _ := client.Delete(1)

	AssertEquals(t, response.Headers["Authorization"], "Bearer my_token")
	AssertEquals(t, response.Headers["Accept"], "application/json")
	AssertEquals(t, response.Headers["User-Agent"], "SDKgen Client v1.0")
	AssertEquals(t, response.Method, "DELETE")
	AssertJson(t, response.Json, "{\"int\":0,\"float\":0,\"string\":\"\",\"bool\":false,\"arrayScalar\":null,\"arrayObject\":null,\"mapScalar\":null,\"mapObject\":null,\"object\":{\"id\":0,\"name\":\"\"}}")
}

func AssertEquals(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func AssertJson(t *testing.T, model TestRequest, want string) {
	raw, _ := json.Marshal(model)
	var got = string(raw)
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func NewPayload() TestRequest {
	var objectFoo = TestObject{
		Id:   1,
		Name: "foo",
	}

	var objectBar = TestObject{
		Id:   2,
		Name: "bar",
	}

	var arrayScalar = []string{"foo", "bar"}
	var arrayObject = []TestObject{objectFoo, objectBar}

	var mapScalar = make(map[string]string)
	mapScalar["foo"] = "bar"
	mapScalar["bar"] = "foo"

	var mapObject = make(map[string]TestObject)
	mapObject["foo"] = objectFoo
	mapObject["bar"] = objectBar

	return TestRequest{
		Int:         1337,
		Float:       13.37,
		String:      "foobar",
		Bool:        true,
		ArrayScalar: arrayScalar,
		ArrayObject: arrayObject,
		MapScalar:   mapScalar,
		MapObject:   mapObject,
		Object:      objectFoo,
	}
}

//
// --------------------------------------------------- GENERATED -------------------------------------------------------
//

// Client automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type Client struct {
	internal *ClientAbstract
}

// GetAll Returns a collection
func (client *Client) GetAll(startIndex int, count int, search string) (TestResponse, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/anything", pathParams))
	if err != nil {
		return TestResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return TestResponse{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TestResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return TestResponse{}, errors.New("could not read response body")
		}

		var response TestResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TestResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return TestResponse{}, errors.New("the server returned an unknown status code")
	}
}

// Create Creates a new product
func (client *Client) Create(payload TestRequest) (TestResponse, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/anything", pathParams))
	if err != nil {
		return TestResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return TestResponse{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return TestResponse{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TestResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return TestResponse{}, errors.New("could not read response body")
		}

		var response TestResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TestResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return TestResponse{}, errors.New("the server returned an unknown status code")
	}
}

// Update Updates an existing product
func (client *Client) Update(id int, payload TestRequest) (TestResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/anything/:id", pathParams))
	if err != nil {
		return TestResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return TestResponse{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", u.String(), reqBody)
	if err != nil {
		return TestResponse{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TestResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return TestResponse{}, errors.New("could not read response body")
		}

		var response TestResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TestResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return TestResponse{}, errors.New("the server returned an unknown status code")
	}
}

// Patch Patches an existing product
func (client *Client) Patch(id int, payload TestRequest) (TestResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/anything/:id", pathParams))
	if err != nil {
		return TestResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return TestResponse{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PATCH", u.String(), reqBody)
	if err != nil {
		return TestResponse{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TestResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return TestResponse{}, errors.New("could not read response body")
		}

		var response TestResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TestResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return TestResponse{}, errors.New("the server returned an unknown status code")
	}
}

// Delete Deletes an existing product
func (client *Client) Delete(id int) (TestResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["id"] = id

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/anything/:id", pathParams))
	if err != nil {
		return TestResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return TestResponse{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TestResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return TestResponse{}, errors.New("could not read response body")
		}

		var response TestResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TestResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return TestResponse{}, errors.New("the server returned an unknown status code")
	}
}

func Build(token string) (*Client, error) {
	var credentials = HttpBearer{Token: token}

	client, err := NewClient("http://127.0.0.1:8081", credentials)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		internal: client,
	}, nil
}

// test_map_object automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type TestMapObject = map[string]TestObject

// test_map_scalar automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type TestMapScalar = map[string]string

// test_object automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type TestObject struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// test_request automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type TestRequest struct {
	Int         int           `json:"int"`
	Float       float64       `json:"float"`
	String      string        `json:"string"`
	Bool        bool          `json:"bool"`
	ArrayScalar []string      `json:"arrayScalar"`
	ArrayObject []TestObject  `json:"arrayObject"`
	MapScalar   TestMapScalar `json:"mapScalar"`
	MapObject   TestMapObject `json:"mapObject"`
	Object      TestObject    `json:"object"`
}

// test_response automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

type TestResponse struct {
	Args    TestMapScalar `json:"args"`
	Headers TestMapScalar `json:"headers"`
	Json    TestRequest   `json:"json"`
	Method  string        `json:"method"`
}
