package activities

import (
	"fmt"
	"testing"
)

func TestActivitiesList(t *testing.T) {
	list := GetListOfActivities("test")
	if len(list) != 3 {
		t.Errorf("should find 3 activities")
	}

	containsDate := false
	containsTime := false
	for _, item := range list {
		fmt.Printf("date %v time %v\n", item.Date, item.Time)
		if item.Date == "2019-01-18" {
			containsDate = true
		}
		if item.Time == "06:17:21" {
			containsTime = true
		}
	}
	if containsDate == false {
		t.Errorf("list of activities should contain date '2019-01-18'")
	}
	if containsTime == false {
		t.Errorf("list of activities should contain time '06:17:21'")
	}
}
