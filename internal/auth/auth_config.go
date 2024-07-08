package auth

type AuthConfig struct {
	JWTSecret string
}

func GetAuthConfig() *AuthConfig {
	return &AuthConfig{}

}
