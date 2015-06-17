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
	const serviceUrl = "http://302-location.com"
	redirectUrl := url.QueryEscape(_url)
	buffer := bytes.Buffer{}
	buffer.WriteString(_secret)
	buffer.WriteString(strconv.FormatInt(int64(_id), 10))
	buffer.WriteString(_url)
	hasher := sha256.New()
	hasher.Write(buffer.Bytes())
	token := (hex.EncodeToString(hasher.Sum(nil)))[0:4]
	params := fmt.Sprintf("i=%d&u=%s&t=%s", _id, redirectUrl, token)
	return fmt.Sprintf("%s/?%s", serviceUrl, params)
}