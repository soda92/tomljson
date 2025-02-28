package main

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
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

func MustSaveToml(config v2config, file string) {
	b, err := toml.Marshal(config)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(file, b, 0644)
	if err != nil {
		panic(err)
	}
}
