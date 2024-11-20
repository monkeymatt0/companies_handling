package main

import (
	"companies_handling/internal"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func main() {

	// Open the config.yaml file for configurations
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f := strings.Split(wd, "/")
	f = f[:len(f)-1]             // Removing the last dir in the past
	f = append(f, "config.yaml") // Adding the file name as leaf of the path
	cp := strings.Join(f, "/")   // Join to have the proper position
	cf, err := os.Open(cp)
	if err != nil {
		panic(err)
	}
	defer cf.Close()

	cb, err := io.ReadAll(cf)
	if err != nil {
		panic(err)
	}
	// Decode yaml file
	var config internal.Config
	if err := yaml.Unmarshal(cb, &config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}
