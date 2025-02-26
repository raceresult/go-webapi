package main

import (
	"fmt"
	"time"

	model "github.com/raceresult/go-model"
	"github.com/raceresult/go-model/variant"
	"github.com/raceresult/go-webapi"
)

func main() {
	err := demo()
	if err != nil {
		fmt.Println(err)
	}
}

func demo() error {
	// create new api
	api := webapi.NewAPI("events.raceresult.com", true, "")

	// login
	err := api.Public().Login(webapi.WithAPIKey("your_api_key"))
	if err != nil {
		return err
	}

	// log out when done
	defer func() {
		_ = api.Public().Logout()
	}()

	// get list of events
	events, err := api.Public().EventList(0, "")
	if err != nil {
		return err
	}
	fmt.Printf("Your accout has %d events.\n", len(events))

	// Create or open event
	var ea *webapi.EventAPI
	if false {
		// create new event
		date := time.Date(2021, 10, 27, 0, 0, 0, 0, time.Local)
		e, err := api.Public().CreateEvent("new event", date, 840, 0, 0, 0, 0)
		if err != nil {
			return err
		}
		ea = e
	} else {

		// open existing event
		ea = api.EventAPI("165435")
	}

	// get number of participants
	count, err := ea.Data().Count("")
	if err != nil {
		return err
	}
	fmt.Printf("This event has %d participants.\n", count)

	// create contest
	c := model.Contest{
		Name: "my new contest",
	}
	cid, err := ea.Contests().Save(c, 0)
	if err != nil {
		return err
	}
	fmt.Printf("New contest created, ID %d.\n", cid)

	// add a participant
	values := variant.VariantMap{
		"FirstName":   variant.RString("John"),
		"LastName":    variant.RString("Doe"),
		"Sex":         variant.RString("m"),
		"DateOfBirth": variant.RString("1975-01-01"),
		"Contest":     variant.RInt(cid),
	}
	if err := ea.Participants().SaveFields(webapi.Bib(12000), values, false); err != nil {
		return err
	}

	// add multiple participants
	var pp []variant.VariantMap
	for i := 12001; i <= 12005; i++ {
		p := variant.VariantMap{
			"Bib":         variant.RInt(i),
			"FirstName":   variant.RString("Test"),
			"LastName":    variant.RInt(i),
			"Sex":         variant.RString("m"),
			"DateOfBirth": variant.RString("1975-01-01"),
			"Contest":     variant.RInt(cid),
		}
		pp = append(pp, p)
	}
	if err := ea.Participants().Save(pp, false); err != nil {
		return err
	}

	// get some data
	data, err := ea.Data().List([]string{"Bib", "FirstName", "LastName", "Contest.Name"}, `[FirstName]="Mark"`, []string{"LastName"},
		0, 0, nil, "", "")
	if err != nil {
		return err
	}
	for _, record := range data {
		fmt.Printf("%+v\n", record)
	}

	// Undo changes
	// ---------------------------------------------------------------------------------------------

	// delete participants again
	if err := ea.Participants().Delete("[Bib]>=12000", webapi.Bib(0), 0); err != nil {
		return err
	}

	// delete new contest
	if err := ea.Contests().Delete(cid); err != nil {
		return err
	}

	return nil
}
