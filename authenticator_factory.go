package sdkgen

import (
	"errors"
	"reflect"
)

func AuthenticatorFactory(credentials CredentialsInterface) (AuthenticatorInterface, error) {
	if reflect.TypeOf(credentials).Name() == "HttpBasic" {
		return &HttpBasicAuthenticator{
			Credentials: credentials.(HttpBasic),
		}, nil
	} else if reflect.TypeOf(credentials).Name() == "HttpBearer" {
		return &HttpBearerAuthenticator{
			Credentials: credentials.(HttpBearer),
		}, nil
	} else if reflect.TypeOf(credentials).Name() == "ApiKey" {
		return &ApiKeyAuthenticator{
			Credentials: credentials.(ApiKey),
		}, nil
	} else if reflect.TypeOf(credentials).Name() == "OAuth2" {
		return &OAuth2Authenticator{
			Credentials: credentials.(OAuth2),
		}, nil
	}

	return nil, errors.New("unknown credentials type")
}
