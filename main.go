package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"

	"flag"
)

type MyConfig struct {
	Version int
	Name    string
	Tags    []string
}

func testToml() {
	doc := `
	version = 2
	name = "go-toml"
	tags = ["go", "toml"]
	`

	var cfg MyConfig
	err := toml.Unmarshal([]byte(doc), &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("version:", cfg.Version)
	fmt.Println("name:", cfg.Name)
	fmt.Println("tags:", cfg.Tags)
}

func main() {
	testtoml := flag.Bool("testtoml", false, "test toml parse")
	if *testtoml {
		testToml()
	}

	data, _ := os.ReadFile("config.json")

	var config v2struct
	err := json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("%+v\n", config)
	}
}
