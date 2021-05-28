package main

import (
	"encoding/json"
	"fmt"
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

		jst, _ := time.LoadLocation("Asia/Tokyo")
		sensorTime, err := time.ParseInLocation("20060102-150405", match[2], jst)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var result map[string]interface{}
		json.Unmarshal(bytes, &result)
		result["date"] = sensorTime.Format(time.RFC3339)

		fmt.Println(match[1])
		fmt.Println(result)
	}
}
