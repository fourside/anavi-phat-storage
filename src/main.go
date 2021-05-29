package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

		jstTimeString, err := parseTime(dateString)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var jsonObject map[string]interface{}
		json.Unmarshal(bytes, &jsonObject)
		jsonObject["date"] = jstTimeString

		addToFirestore(collectionName, jsonObject)

		err = os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
