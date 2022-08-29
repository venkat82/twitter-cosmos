package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserFeedKeyPrefix is the prefix to retrieve all UserFeed
	UserFeedKeyPrefix = "UserFeed/value/"
)

// UserFeedKey returns the store key to retrieve a UserFeed from the index fields
func UserFeedKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
