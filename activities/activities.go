package activities

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Activity struct {
	Filename string
	Time     time.Time
}

var Activities map[string]*Activity

// Start scan directory dir for fit files, parse activity and add to Activities
func Start(dir string) error {
	Activities = make(map[string]*Activity)
	readActivitiesFromFile()
	activitiyFilenameList := getListOfActivityFiles(dir)
	for _, filename := range activitiyFilenameList {

		if isParsed(path.Base(filename)) {
			continue
		}

		ac, err := ParseFast(filepath.Join(dir, filename))
		if err != nil {
			log.Fatalf("can not parse fit file '%v'", err)
		}
		Activities[path.Base(filename)] = ac
	}
	fmt.Printf("found %v fit files\n", len(Activities))
	writeActivitiesToFile()
	return nil
}

func isParsed(name string) bool {
	_, ok := Activities[name]
	return ok
}

func readActivitiesFromFile() {
	b, err := os.ReadFile("activities.json")
	if err != nil {
		fmt.Printf("can not read activities data\n")
		return
	}

	if err := json.Unmarshal(b, &Activities); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %v activities from file\n", len(Activities))
}

func writeActivitiesToFile() {
	b, _ := json.Marshal(Activities)
	_ = os.WriteFile("activities.json", b, 0644)
}

// GetListOfActivities get the list of activity fit filenames
func GetListOfActivities(dir string) []string {

	fileList := []string{}
	for _, activity := range Activities {
		fileList = append(fileList, activity.Filename)
	}
	return fileList
}

// func extractDate(name string) string {
// 	return name[:10]
// }

// func extractTime(name string) string {
// 	return name[11:19]
// }

func getListOfActivityFiles(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fileList := []string{}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".fit") {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList
}
