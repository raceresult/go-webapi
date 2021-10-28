package webapi

import (
	"encoding/json"
	"strconv"
	"strings"

	model "github.com/raceresult/go-model"
)

// Vouchers contains all api endpoints regarding vouchers
type Vouchers struct {
	api *EventAPI
}

// newVouchers creates a new Vouchers api endpoint group
func newVouchers(api *EventAPI) *Vouchers {
	return &Vouchers{
		api: api,
	}
}

// Get returns one or all vouchers
func (q *Vouchers) Get(code string) ([]model.Voucher, error) {
	values := urlValues{
		"code": code,
	}
	bts, err := q.api.Get("vouchers/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.Voucher
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes vouchers
func (q *Vouchers) Delete(ids []int) error {
	sids := make([]string, 0, len(ids))
	for _, id := range ids {
		sids = append(sids, strconv.Itoa(id))
	}
	_, err := q.api.Post("vouchers/delete", nil, strings.Join(sids, ";"))
	return err
}

// Save saves vouchers and returns the IDs
func (q *Vouchers) Save(items []model.Voucher) ([]int, error) {
	bts, err := q.api.Post("vouchers/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
