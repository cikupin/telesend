package telesend

import (
	"errors"
	"net/url"
	"strings"
)

// GetChannelID will retrieve group ID from given telegram channel url
func GetChannelID(urlStr string) (string, error) {
	id, err := getValue(urlStr)
	if err != nil {
		return "", err
	}

	splittedID := strings.Split(id, "_")
	if !strings.HasPrefix(splittedID[0], "c") {
		return "", errors.New("url is not telegram channel url")
	}
	return strings.Replace(splittedID[0], "c", "-100", -1), nil
}

// GetGroupID will retrieve group ID from given telegram group url
func GetGroupID(urlStr string) (string, error) {
	id, err := getValue(urlStr)
	if err != nil {
		return "", err
	}

	if !strings.HasPrefix(id, "g") {
		return "", errors.New("url is not telegram group url")
	}
	return strings.Replace(id, "g", "-", -1), nil
}

// getValue will extract ID from telegram query param URL
func getValue(urlStr string) (string, error) {
	urlStr = strings.Replace(urlStr, "#/im", "", -1)
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	qry := u.Query()
	if _, ok := qry["p"]; !ok {
		return "", err
	}

	splitted := strings.Split(qry.Get("p"), "?p=")
	if len(splitted) > 1 {
		return splitted[len(splitted)-1], nil
	}
	return splitted[0], nil
}
