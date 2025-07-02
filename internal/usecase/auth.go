package usecase

import (
	"context"
	"github.com/google/uuid"
)

type AuthUsecase interface {
	Register(ctx context.Context, email, password string) error
	Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, err error)
	ValidateToken(ctx context.Context, accessToken string) (uuid.UUID, error)
	RefreshTokens(ctx context.Context, refreshToken string) (newAccess string, newRefresh string, err error)
	Logout(ctx context.Context, refreshToken string) error
}
