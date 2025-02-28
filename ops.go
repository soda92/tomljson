package main

import (
	"encoding/json"
	"os"
)

func MustParseJson(file string) v2config {
	data, _ := os.ReadFile(file)

	var config v2config
	err := json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func MustSaveJson(config v2config, file string) {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	} else {
		err := os.WriteFile(file, b, 0644)
		if err != nil {
			panic(err)
		}
	}
}
