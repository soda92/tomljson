package main

import (
	"flag"
)

func main() {
	testtoml := flag.Bool("testtoml", false, "test toml parse")
	flag.Parse()
	if *testtoml {
		testToml()
	}

	config := MustParseJson("config.json")
	MustSaveJson(config, "config2.json")
	MustSaveToml(config, "config.toml")
}
