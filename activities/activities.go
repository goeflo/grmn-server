package activities

import (
	"io/ioutil"
	"log"
	"strings"
)

type Activity struct {
	Filename string
	Date     string
	Time     string
}

func GetListOfActivities(dir string) []Activity {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	fileList := []Activity{}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".fit") {
			newActivity := Activity{
				Filename: file.Name(),
				Date:     extractDate(file.Name()),
				Time:     extractTime(file.Name()),
			}
			fileList = append(fileList, newActivity)
		}
	}
	return fileList
}

func extractDate(name string) string {
	return name[:10]
}

func extractTime(name string) string {
	return name[11:19]
}
