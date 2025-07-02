package mock

import (
	"context"
	"errors"
	"github.com/Ivan-Kats/auth-service/internal/model"
	"sync"

	"github.com/google/uuid"
)

type UserRepoMock struct {
	mu    sync.RWMutex
	users map[string]*model.User // key: email
}

func NewUserRepoMock() *UserRepoMock {
	return &UserRepoMock{
		users: make(map[string]*model.User),
	}
}

func (r *UserRepoMock) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return errors.New("user already exists")
	}

	r.users[user.Email] = user
	return nil
}

func (r *UserRepoMock) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[email]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepoMock) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
