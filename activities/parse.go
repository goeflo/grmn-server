package activities

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/tormoder/fit"
)

type FitData struct {
}

func GetActivityRecords(name string) (records []*fit.RecordMsg, err error) {
	fmt.Printf("parse fit file: %v\n", name)
	testData, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	fit, err := fit.Decode(bytes.NewReader(testData))
	if err != nil {
		return nil, err
	}

	fitActivity, err := fit.Activity()
	if err != nil {
		return nil, err
	}
	fmt.Printf("created: %v, product: %v\n", fit.FileId.TimeCreated, fit.FileId.GetProduct())
	fmt.Printf("activity records: %v\n", len(fitActivity.Records))
	return fitActivity.Records, nil

	// sessions := fitActivity.Sessions
	// fmt.Printf("session messages: %v\n", len(sessions))

	// if len(sessions) == 1 {
	// 	if sessions[0].NormalizedPower != 65535 {
	// 		fmt.Printf("normalized power: %v\n", sessions[0].NormalizedPower)
	// 	} else {
	// 		fmt.Println("normalized power: %")
	// 	}
	// }

}

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
