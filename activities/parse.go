package activities

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/tormoder/fit"
)

type FitData struct {
	Date            time.Time
	Product         string
	NormalizedPower int
	MovingTime      uint32
	Distance        uint32
	Records         []*fit.RecordMsg
}

func (d *FitData) String() string {
	return fmt.Sprintf("date: %v; product: %v, moving time: %v, distance: %v, normalized power: %v\n",
		d.Date, d.Product, d.MovingTime, d.Distance, d.NormalizedPower)
}

func GetActivitySummary(name string) (summary *FitData, err error) {
	fmt.Printf("parse fit file: %v\n", name)
	fitFileData, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	fit, err := fit.Decode(bytes.NewReader(fitFileData))
	if err != nil {
		return nil, err
	}

	summary = &FitData{
		Date:    fit.FileId.TimeCreated,
		Product: fmt.Sprintf("%v", fit.FileId.GetProduct()),
	}

	fitActivity, err := fit.Activity()
	if err != nil {
		return nil, err
	}
	sessions := fitActivity.Sessions

	if len(sessions) == 1 {
		if sessions[0].NormalizedPower != 65535 {
			summary.NormalizedPower = int(sessions[0].NormalizedPower)
		}
		summary.MovingTime = sessions[0].TotalTimerTime
		summary.Distance = sessions[0].TotalDistance
	}

	summary.Records = fitActivity.Records
	fmt.Printf("%v\n", summary)

	return summary, nil
}

// func GetActivityRecords(name string) (records []*fit.RecordMsg, err error) {
// 	fmt.Printf("parse fit file: %v\n", name)
// 	testData, err := os.ReadFile(name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fit, err := fit.Decode(bytes.NewReader(testData))
// 	if err != nil {
// 		return nil, err
// 	}

// 	fitActivity, err := fit.Activity()
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Printf("created: %v, product: %v\n", fit.FileId.TimeCreated, fit.FileId.GetProduct())
// 	fmt.Printf("activity records: %v\n", len(fitActivity.Records))
// 	return fitActivity.Records, nil

// }

func ParseFast(name string) (activity *Activity, err error) {

	testData, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		return nil, err
	}

	// fitActivity, err := fit.Activity()
	// if err != nil {
	// 	return nil, err
	// }

	// for _, record := range fitActivity.Records {
	// }
	// 	fmt.Println(record.PositionLat)
	// 	fmt.Println(record.PositionLong)
	// 	fmt.Println(record.Power)
	// 	break
	// }

	//fmt.Printf("created %v, product %v\n", fit.FileId.TimeCreated, fit.FileId.GetProduct())
	activity = &Activity{
		Filename: path.Base(name),
		Time:     fit.FileId.TimeCreated,
		//Records:  fitActivity.Records,
	}

	return activity, nil
}
