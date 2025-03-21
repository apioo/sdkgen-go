// Client automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package generated

import (
	"github.com/apioo/sdkgen-go/v2"
)

type Client struct {
	internal *sdkgen.ClientAbstract
}

func (client *Client) Product() *ProductTag {
	return NewProductTag(client.internal.HttpClient, client.internal.Parser)
}

func NewClient(baseUrl string, credentials sdkgen.CredentialsInterface) (*Client, error) {
	var client, err = sdkgen.NewClient(baseUrl, credentials)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		internal: client,
	}, nil
}

func NewClientWithVersion(baseUrl string, credentials sdkgen.CredentialsInterface, version string) (*Client, error) {
	var client, err = sdkgen.NewClientWithVersion(baseUrl, credentials, version)
	if err != nil {
		return &Client{}, err
	}

	return &Client{
		internal: client,
	}, nil
}

func Build(token string) (*Client, error) {
	var credentials = sdkgen.HttpBearer{Token: token}

	return NewClientWithVersion("http://127.0.0.1:8081", credentials, "0.1.0")
}

func BuildAnonymous() (*Client, error) {
	var credentials = sdkgen.Anonymous{}

	return NewClient("http://127.0.0.1:8081", credentials)
}
