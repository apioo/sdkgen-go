package sdkgen

type CredentialsInterface interface {
}

type ApiKey struct {
	CredentialsInterface
	Token string
	Name  string
	In    string
}

type OAuth2 struct {
	CredentialsInterface
	ClientId         string
	ClientSecret     string
	TokenUrl         string
	AuthorizationUrl string
	TokenStore       TokenStoreInterface
	Scopes           []string
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
