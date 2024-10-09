package tokenprovider

type Provider interface {
	Generate(id string) (string, error)
	Verify(token string) (string, error)
}
