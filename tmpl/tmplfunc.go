// Package tmpl provides ...
package tmpl

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var TplFuncMap = make(template.FuncMap)

func init() {
	TplFuncMap["dateformat"] = DateFormat
	TplFuncMap["str2html"] = Str2html
	TplFuncMap["join"] = StringsJoin
	TplFuncMap["isnotzero"] = IsNotZero
	TplFuncMap["base64img"] = Base64Img
}

func Str2html(raw string) template.HTML {
	return template.HTML(raw)
}

// DateFormat takes a time and a layout string and returns a string with the formatted date. Used by the template parser as "dateformat"
func DateFormat(t time.Time, layout string) string {
	return t.Format(layout)
}

func StringsJoin(a []string, sep string) string {
	return strings.Join(a, sep)
}

func IsNotZero(t time.Time) bool {
	return !t.IsZero()
}

func Base64Img(domain, avatar string) string {
	resp, err := http.Get("https://" + domain + "/static/img/" + avatar)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}

	return "data:" + resp.Header.Get("content-type") + ";base64," + base64.StdEncoding.EncodeToString(data)
}
