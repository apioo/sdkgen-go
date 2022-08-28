package sdkgen

type CredentialsInterface interface {
}

type ApiKey struct {
	CredentialsInterface
	Token string
	Name  string
	In    string
}

type AuthorizationCode struct {
	OAuth2Abstract
}

type ClientCredentials struct {
	OAuth2Abstract
}

type HttpBasic struct {
	CredentialsInterface

	UserName string
	Password string
}

type HttpBearer struct {
	CredentialsInterface
	Token string
}

type OAuth2Abstract struct {
	CredentialsInterface
	ClientId         string
	ClientSecret     string
	TokenUrl         string
	AuthorizationUrl string
}
