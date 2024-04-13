package wh

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"trl/conf"
	"trl/srv/errs"
	"trl/srv/logger"

	"golang.org/x/net/html"
)

var baseURL *url.URL

func LoadWordData(word string) (string, string) {
	body, err := getHtml(word)
	if err != nil {
		errs.LogError(err)
		return "", ""
	}
	doc, _ := html.Parse(bytes.NewReader(body))
	translation := extractTranslation(doc)
	transcription := extractTranscription(doc)

	return translation, transcription
}

func getURL(word string) string {
	if baseURL == nil {
		_url, err := url.Parse(conf.Settings.Wh.Url)
		errs.LogFatalIfError(err)
		baseURL = _url
	}
	return baseURL.JoinPath(word).String()
}

func getHtml(word string) ([]byte, error) {
	response, err := makeRequest(getURL(word))
	if err != nil {
		logger.Warning(fmt.Sprintf("an error occurred while receiving `%s` data", word))
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Warning(fmt.Sprintf("an error occurred while reading `%s` data", word))
		return nil, err
	}
	return body, nil
}

func makeRequest(url string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", getUserAgent())
	return client.Do(req)
}

func extractTranscription(data *html.Node) string {
	node := getElementById(data, "us_tr_sound")
	if node == nil {
		logger.Warning("can not find element with id `us_tr_sound`")
		return ""
	}
	node = getElementByClass(node, "transcription")
	if node == nil {
		logger.Warning("can not find element with class `transcription`")
		return ""
	}
	return strings.Trim(strings.Trim(node.FirstChild.Data, " "), "|")
}

func extractTranslation(data *html.Node) string {
	node := getElementById(data, "content_in_russian")
	if node == nil {
		logger.Warning("can not find element with id `content_in_russian`")
		return ""
	}
	node = getElementByClass(node, "t_inline_en")
	if node == nil {
		logger.Warning("can not find element with class `t_inline_en`")
		return ""
	}
	return node.FirstChild.Data
}
