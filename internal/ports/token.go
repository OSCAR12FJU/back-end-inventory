package ports

type TokenGenerator interface {
	GenerateToken(email string) (string, error)
}
