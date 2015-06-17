package location302

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"bytes"
	"net/url"
)

func GetLink(_id int, _secret string, _url string) string {
	const serviceUrl = "http://302-location302.com"
	redirectUrl := url.QueryEscape(_url)
	buffer := bytes.Buffer{}
	buffer.WriteString(_secret)
	buffer.WriteString(strconv.Itoa(_id))
	buffer.WriteString(_url)
	converted := []byte(buffer.String())
	hasher := sha256.New()
	hasher.Write(converted)
	token := (hex.EncodeToString(hasher.Sum(nil)))
	token = token[0:4]
	params := fmt.Sprintf("i=%d&u=%s&t=%s", _id, redirectUrl, token)
	return fmt.Sprintf("%s/?%s", serviceUrl, params)
}