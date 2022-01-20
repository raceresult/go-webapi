package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
	"github.com/raceresult/go-model/variant"
)

// Participants contains all api endpoints regarding participants
type Participants struct {
	api *EventAPI
}

// newParticipant creates a new Participant api endpoint group
func newParticipants(api *EventAPI) *Participants {
	return &Participants{
		api: api,
	}
}

// GetFields returns fields of one participant
func (q *Participants) GetFields(bib int, fields []string) (variant.VariantMap, error) {
	values := urlValues{
		"bib":    bib,
		"fields": fields,
	}
	bts, err := q.api.get("part/getfields", values)
	if err != nil {
		return nil, err
	}
	var dest variant.VariantMap
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetFieldsWithChanges presents the given changes would be applied the participant and then returns field values of this participant
func (q *Participants) GetFieldsWithChanges(bib int, fields []string, changes variant.VariantMap) (variant.VariantMap, error) {
	values := urlValues{
		"bib":    bib,
		"fields": fields,
	}
	bts, err := q.api.post("part/getfieldswithchanges", values, changes)
	if err != nil {
		return nil, err
	}
	var dest variant.VariantMap
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// SaveExpression calculates the result of an expression and saves it in the given field.
func (q *Participants) SaveExpression(bib int, field string, expression string, noHistory bool) error {
	values := urlValues{
		"bib":        bib,
		"field":      field,
		"expression": expression,
		"noHistory":  noHistory,
	}
	_, err := q.api.get("part/saveexpression", values)
	return err
}

// SaveValueArray saves multiple values for possibly different participants in one call
func (q *Participants) SaveValueArray(values []model.SaveValueArrayItem, noHistory bool) error {
	uvalues := urlValues{
		"noHistory": noHistory,
	}
	_, err := q.api.post("part/savevaluearray", uvalues, values)
	return err
}

// SaveFields saves multiple fields for one participant
func (q *Participants) SaveFields(bib int, values variant.VariantMap, noHistory bool) error {
	uvalues := urlValues{
		"bib":       bib,
		"noHistory": noHistory,
	}
	_, err := q.api.post("part/savefields", uvalues, values)
	return err
}

// Save adds or updates one or more participants
func (q *Participants) Save(values []variant.VariantMap, noHistory bool) error {
	uvalues := urlValues{
		"noHistory": noHistory,
	}
	_, err := q.api.post("part/savefields", uvalues, values)
	return err
}

// Delete deletes participants.
func (q *Participants) Delete(filter string, bib int, contest int) error {
	values := urlValues{
		"filter": filter,
	}
	if bib == 0 {
		values["bib"] = "ALL"
	} else {
		values["bib"] = bib
	}
	if contest == 0 {
		values["contest"] = "ALL"
	} else {
		values["contest"] = contest
	}
	_, err := q.api.get("part/delete", values)
	return err
}

// New create new participant and returns the bib
func (q *Participants) New(bib int, contest int, firstfree bool) (int, error) {
	values := urlValues{
		"bib":       bib,
		"contest":   contest,
		"firstfree": firstfree,
	}
	bts, err := q.api.get("part/new", values)
	if err != nil {
		return 0, err
	}
	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// EntryFee all entry fees charged to the participants with the given bibs
func (q *Participants) EntryFee(bibs []int) ([]model.EntryFeeItem, error) {
	values := urlValues{
		"bibs": intSliceToString(bibs),
	}
	bts, err := q.api.get("part/entryfee", values)
	if err != nil {
		return nil, err
	}
	var dest []model.EntryFeeItem
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// CreateBlanks creates blank participants
func (q *Participants) CreateBlanks(from, to int, contest int, skipExcluded bool) error {
	values := urlValues{
		"from":         from,
		"to":           to,
		"contest":      contest,
		"skipExcluded": skipExcluded,
	}
	_, err := q.api.get("part/clearbankinformation", values)
	return err
}

// SwapBibs swaps the bibs of two participants
func (q *Participants) SwapBibs(bib1, bib2 int) error {
	values := urlValues{
		"bib1": bib1,
		"bib2": bib2,
	}
	_, err := q.api.get("part/swapbibs", values)
	return err
}

// ResetBibs assigns new bibs
func (q *Participants) ResetBibs(sort string, firstBib int, ranges bool, filter string, noHistory bool) error {
	values := urlValues{
		"sort":      sort,
		"firstBib":  firstBib,
		"ranges":    ranges,
		"filter":    filter,
		"noHistory": noHistory,
	}
	_, err := q.api.get("part/resetbibs", values)
	return err
}

// DataManipulation changes multiple participants at the same time
func (q *Participants) DataManipulation(values map[string]string, filter string, noHistory bool) error {
	uvalues := urlValues{
		"filter":    filter,
		"noHistory": noHistory,
	}
	_, err := q.api.post("part/datamanipulation", uvalues, values)
	return err
}

// ClearBankInformation removes all banking information for the participants matching the given filters
func (q *Participants) ClearBankInformation(bib int, contest int, filter string) error {
	values := urlValues{
		"bib":     bib,
		"contest": contest,
		"filter":  filter,
	}
	_, err := q.api.get("part/clearbankinformation", values)
	return err
}

// ImportSES imports an entire SES file into the current file
func (q *Participants) ImportSES(file []byte, filter string, identity string, addParticipants bool, updateParticipants bool,
	contestFrom, contestTo int, timesFrom, timesTo int, importRawData bool) (*model.ImportResult, error) {
	values := urlValues{
		"filter":             filter,
		"identity":           identity,
		"addParticipants":    addParticipants,
		"updateParticipants": updateParticipants,
		"contestFrom":        contestFrom,
		"contestTo":          contestTo,
		"timesFrom":          timesFrom,
		"timesTo":            timesTo,
		"importRawData":      importRawData,
	}
	bts, err := q.api.post("part/importses", values, file)
	if err != nil {
		return nil, err
	}
	var dest model.ImportResult
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Import imports participants from an csv/xls/xlsx file
func (q *Participants) Import(file []byte, addParticipants bool, updateParticipants bool, colHandling int,
	identityColumns int, lang string) (*model.ImportResult, error) {
	values := urlValues{
		"addParticipants":    addParticipants,
		"updateParticipants": updateParticipants,
		"colHandling":        colHandling,
		"identityColumns":    identityColumns,
		"lang":               lang,
	}
	bts, err := q.api.post("part/import", values, file)
	if err != nil {
		return nil, err
	}
	var dest model.ImportResult
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// FreeBib returns an unused bib
func (q *Participants) FreeBib(maxBibPlus1 bool, contest int, preferred int) (int, error) {
	values := urlValues{
		"maxBibPlus1": maxBibPlus1,
		"contest":     contest,
		"preferred":   preferred,
	}
	bts, err := q.api.get("part/freebib", values)
	if err != nil {
		return 0, err
	}
	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// FrequentClubs returns the most frequent clubs containing the given wildcard
func (q *Participants) FrequentClubs(wildcard string, maxNumber int) ([]string, error) {
	values := urlValues{
		"wildcard":  wildcard,
		"maxNumber": maxNumber,
	}
	bts, err := q.api.get("part/frequentclubs", values)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}
