package webapi

import (
	"bytes"
	"encoding/json"
)

// EventAPI contains all api functions for a specific event
type EventAPI struct {
	eventID string
	api     *API
}

// NewEventAPI creates a new EventAPI object
func NewEventAPI(eventID string, api *API) *EventAPI {
	q := &EventAPI{
		eventID: eventID,
		api:     api,
	}
	return q
}

// AgeGroups contains all api endpoints regarding age groups
func (q *EventAPI) AgeGroups() *AgeGroups {
	return newAgeGroups(q)
}

// Archives contains all api endpoints regarding archives
func (q *EventAPI) Archives() *Archives {
	return newArchives(q)
}

// Backup contains all api endpoints regarding backup
func (q *EventAPI) Backup() *Backup {
	return newBackup(q)
}

// BibRanges contains all api endpoints regarding bib ranges
func (q *EventAPI) BibRanges() *BibRanges {
	return newBibRanges(q)
}

// Certificates contains all api endpoints regarding certificates
func (q *EventAPI) Certificates() *Certificates {
	return newCertificates(q)
}

// CertificateSets contains all api endpoints regarding certificate sets
func (q *EventAPI) CertificateSets() *CertificateSets {
	return newCertificateSets(q)
}

// Chat contains all api endpoints regarding chat
func (q *EventAPI) Chat() *Chat {
	return newChat(q)
}

// ChipFile contains all api endpoints regarding chip file
func (q *EventAPI) ChipFile() *ChipFile {
	return newChipFile(q)
}

// Contests contains all api endpoints regarding contests
func (q *EventAPI) Contests() *Contests {
	return newContests(q)
}

// CustomFields contains all api endpoints regarding custom fields
func (q *EventAPI) CustomFields() *CustomFields {
	return newCustomFields(q)
}

// Data contains all api endpoints regarding data
func (q *EventAPI) Data() *Data {
	return newData(q)
}

// Dependencies contains all api endpoints regarding dependencies
func (q *EventAPI) Dependencies() *Dependencies {
	return newDependencies(q)
}

// EmailTemplates contains all api endpoints regarding email templates
func (q *EventAPI) EmailTemplates() *EmailTemplates {
	return newEmailTemplates(q)
}

// EntryFees contains all api endpoints regarding entry fees
func (q *EventAPI) EntryFees() *EntryFees {
	return newEntryFees(q)
}

// Exporters contains all api endpoints regarding exporters
func (q *EventAPI) Exporters() *Exporters {
	return newExporters(q)
}

// File contains all api endpoints regarding the event file
func (q *EventAPI) File() *File {
	return newFile(q)
}

// Forwarding contains all api endpoints regarding online forwarding
func (q *EventAPI) Forwarding() *Forwarding {
	return newForwarding(q)
}

// History contains all api endpoints regarding history entries
func (q *EventAPI) History() *History {
	return newHistory(q)
}

// Information contains all api endpoints regarding information
func (q *EventAPI) Information() *Information {
	return newInformation(q)
}

// Kiosks contains all api endpoints regarding kiosks
func (q *EventAPI) Kiosks() *Kiosks {
	return newKiosks(q)
}

// Labels contains all api endpoints regarding labels
func (q *EventAPI) Labels() *Labels {
	return newLabels(q)
}

// Lists contains all api endpoints regarding lists
func (q *EventAPI) Lists() *Lists {
	return newLists(q)
}

// OverwriteValues contains all api endpoints regarding timing point rules
func (q *EventAPI) OverwriteValues() *OverwriteValues {
	return newOverwriteValues(q)
}

// Participants contains all api endpoints regarding participants
func (q *EventAPI) Participants() *Participants {
	return newParticipants(q)
}

// Pictures contains all api endpoints regarding pictures
func (q *EventAPI) Pictures() *Pictures {
	return newPictures(q)
}

// Rankings contains all api endpoints regarding rankings
func (q *EventAPI) Rankings() *Rankings {
	return newRankings(q)
}

// RawData contains all api endpoints regarding raw data entries
func (q *EventAPI) RawData() *RawData {
	return newRawData(q)
}

// RawDataRules contains all api endpoints regarding raw data rules
func (q *EventAPI) RawDataRules() *RawDataRules {
	return newRawDataRules(q)
}

// Results contains all api endpoints regarding results
func (q *EventAPI) Results() *Results {
	return newResults(q)
}

// Settings contains all api endpoints regarding settings
func (q *EventAPI) Settings() *Settings {
	return newSettings(q)
}

// SimpleAPI contains all api endpoints regarding SimpleAPI
func (q *EventAPI) SimpleAPI() *SimpleAPI {
	return newSimpleAPI(q)
}

// Splits contains all api endpoints regarding splits
func (q *EventAPI) Splits() *Splits {
	return newSplits(q)
}

// Statistics contains all api endpoints regarding statistics
func (q *EventAPI) Statistics() *Statistics {
	return newStatistics(q)
}

// Synchronization contains all api endpoints regarding synchronization
func (q *EventAPI) Synchronization() *Synchronization {
	return newSynchronization(q)
}

// TeamScores contains all api endpoints regarding teamscores
func (q *EventAPI) TeamScores() *TeamScores {
	return newTeamScores(q)
}

// Times contains all api endpoints regarding times
func (q *EventAPI) Times() *Times {
	return newTimes(q)
}

// TimingPoints contains all api endpoints regarding timing points
func (q *EventAPI) TimingPoints() *TimingPoints {
	return newTimingPoints(q)
}

// Vouchers contains all api endpoints regarding vouchers
func (q *EventAPI) Vouchers() *Vouchers {
	return newVouchers(q)
}

// WebHooks contains all api endpoints regarding web hooks
func (q *EventAPI) WebHooks() *WebHooks {
	return newWebHooks(q)
}

// TimingPointRules contains all api endpoints regarding timing point rules
func (q *EventAPI) TimingPointRules() *TimingPointRules {
	return newTimingPointRules(q)
}

// UserDefinedFields contains all api endpoints regarding user defined fields
func (q *EventAPI) UserDefinedFields() *UserDefinedFields {
	return newUserDefinedFields(q)
}

// EventID returns the id of this event
func (q *EventAPI) EventID() string {
	return q.eventID
}

// MultiRequest returns several resources in one call
func (q *EventAPI) MultiRequest(requests []string) (map[string]interface{}, error) {
	req, _ := json.Marshal(requests)
	bts, err := q.api.post(q.eventID, "multirequest", nil, "application/json", bytes.NewReader(req))
	if err != nil {
		return nil, err
	}
	var dest map[string]interface{}
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Get makes a GET request on the server
func (q *EventAPI) get(cmd string, values urlValues) ([]byte, error) {
	return q.api.get(q.eventID, cmd, values)
}

// Get makes a GET request on the server
func (q *EventAPI) Get(requestURI string) ([]byte, error) {
	return q.api.get(q.eventID, requestURI, nil)
}

// Post makes a POST request on the server
func (q *EventAPI) post(cmd string, values urlValues, data interface{}) ([]byte, error) {
	return q.api.post(q.eventID, cmd, values, "", data)
}

// Post makes a POST request on the server
func (q *EventAPI) Post(requestURI string, data interface{}) ([]byte, error) {
	return q.api.post(q.eventID, requestURI, nil, "", data)
}
