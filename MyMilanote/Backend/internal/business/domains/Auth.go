package domains

type AuthTokens struct {
	AccessToken  string
	RefreshToken string
}

type LoginData struct {
	//TODO
}

type RegistrationData struct {
	//TODO
}

type IAuthService interface {
	Login(data *LoginData) (*AuthTokens, error)
	AuthTokens(tokens *AuthTokens) error
	UpdateTokens(token *AuthTokens) (*AuthTokens, error)
	Register(data *RegistrationData) error //TODO: Должен в Redis сохранить токен регистрации и послать его в почтовый сервис
	AckRegistration(key string) error      //TODO: Придумать данные для сохранения
}

type IAuthRepository interface {
	GetAuthData(data *LoginData) (*User, error)
}
