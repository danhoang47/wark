package common

const (
	ErrorCodeSuccess = iota
)

type Response struct {
	Message   string      `json:"message"`
	ErrorCode int         `json:"errorCode"`
	Data      interface{} `json:"data,omitempty"`
	Status    int         `json:"status"`
	// Limit      int         `json:"limit,omitempty"`
	// NextCursor string      `json:"nextCursor,omitempty"`
}
