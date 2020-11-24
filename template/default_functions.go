package template

import (
	"strings"
	gotemplate "text/template"
	"time"
)

func GetDefaultFuncs() gotemplate.FuncMap {
	return map[string]interface{}{
		"Now":         Now,
		"NowUnix":     NowUnix,
		"NowUnixNano": NowUnixNano,
		"ToLower":     ToLower,
		"Replace":     Replace,
	}
}

func Now(vv ...interface{}) (interface{}, error) {
	format := time.RFC3339
	if len(vv) > 0 {
		fmt, isString := vv[0].(string)
		if isString {
			format = fmt
		}
	}
	return time.Now().Format(format), nil
}

func NowUnix(vv ...interface{}) (interface{}, error) {
	return time.Now().Unix(), nil
}

func NowUnixNano(vv ...interface{}) (interface{}, error) {
	return time.Now().UnixNano(), nil
}

func ToLower(vv ...interface{}) (interface{}, error) {
	if len(vv) > 0 {
		return strings.ToLower(vv[0].(string)), nil
	}
	return "", nil
}

func Replace(vv ...interface{}) (interface{}, error) {
	if len(vv) == 4 {
		return strings.Replace(vv[0].(string), vv[1].(string), vv[2].(string), vv[3].(int)), nil
	}
	return "", nil
}
