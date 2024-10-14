package tokenprovider

import "time"

type Provider interface {
	Generate(string, time.Duration) (string, error)
	Verify(string) (string, error)
}
