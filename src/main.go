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

	filePattern := regexp.MustCompile("^(.+?)\\..+")
	for _, file := range files {
		fmt.Println(file)
		match := filePattern.FindStringSubmatch(file)
		fmt.Println(match[1])

		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		var result map[string]interface{}
		json.Unmarshal(bytes, &result)
		result["date"] = time.Now().Format(time.RFC3339)
		fmt.Println(result)
	}
}
