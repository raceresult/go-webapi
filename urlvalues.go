package webapi

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/raceresult/go-model/decimal"
	"github.com/raceresult/go-model/vbdate"
)

// urlValue helps to convert different types to url parameters

type urlValue interface{}

type urlValues map[string]urlValue

func (q urlValues) URLEncode() string {
	toString := func(v urlValue) string {
		switch t := v.(type) {
		case string:
			return t

		case int:
			return strconv.Itoa(t)

		case time.Time:
			return t.Format("2006-01-02 15:04:05")

		case bool:
			if t {
				return "true"
			}
			return "false"

		case decimal.Decimal:
			return t.ToString()

		case vbdate.VBDate:
			return t.ToString()

		case []string:
			bts, _ := json.Marshal(t)
			return string(bts)

		default:
			return ""
		}
	}

	values := url.Values{}
	for k, v := range q {
		values.Set(k, toString(v))
	}
	return values.Encode()
}
