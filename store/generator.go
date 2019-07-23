package store

import {
	"crypto/md5",
	"encoding/hex"
}

func shortLinkGenerator(key string) string {
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}