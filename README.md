🔐 AuthService – Reusable Authentication Microservice in Go
AuthService is a ready-to-use authentication microservice written in Go. It supports both REST and gRPC interfaces, and can be easily integrated into any project that requires user registration, login, and access control.

📌 Features
User registration and login via email & password
JWT-based access/refresh token generation
Token refresh and revocation
gRPC methods for inter-service token validation
REST API for frontend clients
Database migrations via golang-migrate
Follows Clean Architecture principles

🧱 Tech Stack
Go (Golang)
PostgreSQL
gRPC
JWT (JSON Web Tokens)
Echo / Fiber / Chi (framework of choice)
golang-migrate
Docker + Docker Compose

🚀 Planned Improvements
Two-Factor Authentication (2FA / MFA)
OAuth 2.0 / Google Sign-In support
Role-based Access Control (RBAC)
Login attempt limits & brute-force protection

🔐 AuthService – многоразовый микросервис авторизации на Go
AuthService — это готовый к использованию микросервис авторизации, написанный на Go, с поддержкой REST API и gRPC. Он может быть легко интегрирован в любые проекты, требующие регистрацию, аутентификацию и проверку прав доступа пользователей.

📌 Возможности
Регистрация и вход по email/паролю
Выдача access/refresh токенов (JWT)
Обновление и отзыв токенов
gRPC-методы для межсервисной проверки токенов
REST API для фронтенда
Поддержка миграций через golang-migrate
Архитектура в духе Clean Architecture

🧱 Используемые технологии
Golang
PostgreSQL
gRPC
JWT
Echo/Fiber/Chi (на выбор)
golang-migrate
Docker + Docker Compose

🚀 Планируемые улучшения
MFA / двухфакторная авторизация
Поддержка OAuth 2.0 / Google Auth
Role-based Access Control (RBAC)
Лимиты попыток входа / защита от brute-force
