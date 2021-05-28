package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

func main() {
	files, err := filepath.Glob("./json/*.json")
	if err != nil {
		log.Fatal(err)
	}

	filePattern := regexp.MustCompile("^(.+?)\\.(.+?)\\..+")
	for _, file := range files {
		basename := filepath.Base(file)
		match := filePattern.FindStringSubmatch(basename)

		var collectionName = match[1]
		var dateString = match[2]
		jst, _ := time.LoadLocation("Asia/Tokyo")
		sensorTime, err := time.ParseInLocation("20060102-150405", dateString, jst)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var jsonObject map[string]interface{}
		json.Unmarshal(bytes, &jsonObject)
		jsonObject["date"] = sensorTime.Format(time.RFC3339)

		addToFirestore(collectionName, jsonObject)

		err = os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
