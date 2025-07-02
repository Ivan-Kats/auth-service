package deps

import "github.com/google/uuid"

// TokenService описывает интерфейс для генерации access и refresh токенов
type TokenService interface {
	GenerateTokens(userID uuid.UUID) (accessToken string, refreshToken string, err error)
	//ParseAccessToken(token string) (userID uuid.UUID, err error)
	//InvalidateRefreshToken(refresh string) error
}
