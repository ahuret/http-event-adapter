package template

import (
	gotemplate "text/template"
	"time"
)

func GetDefaultFuncs() gotemplate.FuncMap {
	return map[string]interface{}{
		"Now":     Now,
		"UnixNow": UnixNow,
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

func UnixNow(vv ...interface{}) (interface{}, error) {
	return time.Now().UnixNano(), nil
}
