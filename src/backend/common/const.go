package common

const (
	EmptyCachedValue = "EMPTY_VALUE"
)

const (
	userKeyPrefix = "users:"
)

func GetUserMemCachedKey(id string) (key string) {
	return userKeyPrefix + id
}
