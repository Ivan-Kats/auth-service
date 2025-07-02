package mock

import (
	"auth-service/internal/usecase/deps"
	"fmt"
	"github.com/google/uuid"
)

type MockTokenService struct{}

func NewMockTokenService() deps.TokenService {
	return &MockTokenService{}
}

func (m *MockTokenService) GenerateTokens(userID uuid.UUID) (string, string, error) {
	// Просто возвращаем строки с userID
	access := fmt.Sprintf("access-token-%s", userID.String())
	refresh := fmt.Sprintf("refresh-token-%s", userID.String())
	return access, refresh, nil
}
