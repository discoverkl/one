package one

import (
	"crypto/md5"
	"encoding/hex"
	"hash/fnv"
	"strings"

	"github.com/google/uuid"
)

func encode(msg string) string {
	return b32StdEncoding.EncodeToString([]byte(msg))
}

func decode(msg string) (string, error) {
	data, err := b32StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Hash(id string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))
	return h.Sum32()
}

func HashString(id string) string {
	return HashStringLen(id, 7)
}

func HashStringLen(id string, maxLen int) string {
	if b32StdPadding == b32NoPadding {
		panic("HashLen need padding")
	}
	// md5Data := md5.New().Sum([]byte(id))
	output := encode(string(md5Hash([]byte(id))))
	if maxLen > 0 && len(output) > maxLen {
		output = output[:maxLen]
	}
	return output
}

func md5Hash(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

func Md5(data []byte) string {
	return hex.EncodeToString(md5Hash(data))
}

func Guid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func ShortID() string {
	guid := Guid()
	return HashString(guid)
}
