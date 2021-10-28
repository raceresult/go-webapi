package webapi

import (
	"encoding/json"
	"strconv"
	"strings"
)

func parseJsonStringArr(bts []byte) ([]string, error) {
	var dest []string
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

func parseJsonIntArr(bts []byte) ([]int, error) {
	var dest []int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

func intSliceToString(ints []int) string {
	if len(ints) == 0 {
		return ""
	}
	var arr []string
	for _, c := range ints {
		arr = append(arr, strconv.Itoa(c))
	}
	return strings.Join(arr, ",")
}
