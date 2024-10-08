package tokenprovider

type Provider interface {
	Generate(id string) string
	Verify(token string) (string, error)
}
