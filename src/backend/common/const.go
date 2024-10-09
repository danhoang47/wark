package common

const (
	EmptyCachedValue = ""
)

const (
	userKeyPrefix = "users:"
)

func GetUserMemCachedKey(id string) (key string) {
	return userKeyPrefix + id
}
