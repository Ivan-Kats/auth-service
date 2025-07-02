package main

import (
	"fmt"
	"github.com/Ivan-Kats/auth-service/internal/repository/mock"
	"github.com/Ivan-Kats/auth-service/internal/usecase"
)

func main() {
	userRepo := mock.NewUserRepoMock()
	authUC := usecase.NewAuthUsecase(userRepo)

	err := authUC.Register(ctx, "test@example.com", "pass123")
	if err != nil {
		fmt.Println("Register failed:", err)
	}

	access, refresh, err := authUC.Login(ctx, "test@example.com", "pass123")
	if err != nil {
		fmt.Println("Login failed:", err)
	} else {
		fmt.Println("Access:", access, "Refresh:", refresh)
	}
}
