package deps

// PasswordHasher описывает интерфейс для хеширования и проверки паролей
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CompareHash(hash, password string) error
}
