package webapi

import (
	"encoding/json"
	"github.com/raceresult/go-model/datetime"
	"github.com/raceresult/go-model/decimal"
	"net/url"
	"strconv"
	"time"
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
			return t.Format(time.RFC3339)

		case bool:
			if t {
				return "true"
			}
			return "false"

		case decimal.Decimal:
			return t.ToString()

		case datetime.DateTime:
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
