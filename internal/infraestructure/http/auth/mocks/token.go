package authMock

import "github.com/stretchr/testify/mock"

type AuthMock struct {
	mock.Mock
}

func NewAuthMock() *AuthMock {
	return &AuthMock{}
}

func (auth *AuthMock) GenerateToken() (string, error) {
	args := auth.Called()

	return args.Get(0).(string), args.Error(1)
}
