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
	CredentialsInterface
	ClientId         string
	ClientSecret     string
	TokenUrl         string
	AuthorizationUrl string
	RefreshUrl       string
}

type ClientCredentials struct {
	CredentialsInterface
	ClientId     string
	ClientSecret string
	TokenUrl     string
	RefreshUrl   string
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
