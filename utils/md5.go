package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//Encrypt ..
func Encrypt(p string) string {
	h := md5.New()
	h.Write([]byte(p))
	return hex.EncodeToString(h.Sum(nil))
}
