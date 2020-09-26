package HexToBase64

import (
	"encoding/base64"
	"encoding/hex"
)

func Convert(s string) string {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}
