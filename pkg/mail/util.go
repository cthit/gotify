package mail

import "encoding/base64"

func EncodeHeader(header string) string {
	return "=?utf-8?B?" + base64.RawURLEncoding.EncodeToString([]byte(header)) + "?="
}
