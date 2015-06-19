package location302

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"bytes"
	"net/url"
	"net/http"
)


type Location struct {
	id int
	secret string
	url string
}

// Simple method build Location and return builded link
func GetLink(id int, secret string, url string) string {
	loc := Location{}
	loc.SetId(id)
	loc.SetSecret(secret)
	loc.SetUrl(url)
	return loc.GetLink()
}

// Getter and Setters START

func (l Location) Id() int {
	return l.id
}

func (l *Location) SetId(value int) {
	l.id = value
}

func (l Location) Secret() string {
	return l.secret
}

func (l *Location) SetSecret(value string) {
	l.secret = value
}

func (l Location) Url() string {
	return l.url
}

func (l *Location) SetUrl(value string) {
	l.url = value
}

// Getter and Setters END


func (l *Location) GetLink() string {
	return l.generateLink(l.id, l.secret, l.url)
}

func NewLocation(id int, secret string, url string) *Location {
	var loc = new(Location)
	loc.SetId(id)
	loc.SetSecret(secret)
	loc.SetUrl(url)
	return loc
}

func New() *Location {
	var loc = &Location{}
	return loc
}

func (l Location) Verify() (bool, error) {
	client := &http.Client{}
	request, err := http.NewRequest("HEAD", l.GetLink(), nil)
	if err != nil {
		return false, err
	}
	request.Header.Set("User-Agent", "Location302 Bot v 1.0")
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()
	if response.StatusCode == 200{
		return true, nil
	}else {
		return false, fmt.Errorf("Status code is %d", response.StatusCode)
	}
}

// Fluent API methods START

func (l *Location) WithId(value int) *Location {
	l.id = value
	return l
}

func (l *Location) WithSecret(value string) *Location {
	l.secret = value
	return l
}

func (l *Location) WithUrl(value string) *Location {
	l.url = value
	return l
}

// Fluent API methods END

//private

func (l Location) generateLink(_id int, _secret string, _url string) string {
	buffer := concat(_secret, strconv.FormatInt(int64(_id), 10), _url)
	token := l.getToken(buffer)
	return token
}

func (l Location) prepareUrl() string {
	redirectUrl := url.QueryEscape(l.url)
	return redirectUrl
}

func (l Location) getServiceUrl() string {
	return "http://302-location.com"
}

// Concatenating slice of string in one string
func concat(strings ...string) bytes.Buffer {
	buffer := bytes.Buffer{}
	for _, str := range strings {
		buffer.WriteString(str)
	}
	return buffer
}

func (l *Location) getToken(buffer bytes.Buffer) string {
	hasher := sha256.New()
	hasher.Write(buffer.Bytes())
	token := (hex.EncodeToString(hasher.Sum(nil)))[0:4]
	params := fmt.Sprintf("i=%d&u=%s&t=%s", l.id, l.prepareUrl(), token)
	return fmt.Sprintf("%s/?%s", l.getServiceUrl(), params)
}