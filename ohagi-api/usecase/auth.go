package usecase

type Auth interface {
	GenarateToken(req LoginRequest, res LoginResponse) error
}

func NewAuth() Auth {
	return &login{}
}

type login struct {
}

func (l *login) GenarateToken(req LoginRequest, res LoginResponse) error {
	return res.RenderLoginResult(req.UserID)
}

type LoginRequest struct {
	UserID string
}

type LoginResponse interface {
	RenderLoginResult(petsToken string) error
	RenderFailure(err error) error
}
