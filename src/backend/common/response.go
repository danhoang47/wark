package common

const (
	ErrorCodeSuccess = iota
)

type Response[T comparable] struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
	Data      T      `json:"data"`
	Status    int    `json:"status"`
}
