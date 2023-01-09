package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	"gopkg.in/ini.v1"
)

type Config struct {
	Url     string `ini:"sso_start_url"`
	Region  string `ini:"sso_region"`
	Account string `ini:"sso_account_id"`
	Role    string `ini:"sso_role_name"`
}

func parseValues(s []*ini.Section) []map[string]Config {
	// Make slice and store profiles
	profiles := make([]map[string]Config, 0, 10)

	for _, v := range s {
		config := new(Config)
		name := v.Name()
		entry := map[string]Config{}

		v.StrictMapTo(config)
		entry[name] = *config
		profiles = append(profiles, entry)
	}
	return profiles[1:]
}

func fuzzyFind(p []map[string]Config) int {
	idx, err := fuzzyfinder.Find(
		p,
		func(i int) string {
			key := reflect.ValueOf(p[i]).MapKeys()
			return key[0].String()
		})
	if err != nil {
		os.Exit(1)
	}
	return idx
}

func main() {
	// Open aws config file
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	region := flag.String("region", "", "Specify region")
	flag.Parse()

	cfg, err := ini.Load(homedir + "/.aws/config")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	// Get all sections from ini file
	sections := cfg.Sections()

	profiles := parseValues(sections)

	idx := fuzzyFind(profiles)

	for k, v := range profiles[idx] {
		p := strings.Split(k, " ")
		var awsProfile string
		if *region != "" {
			awsProfile = fmt.Sprint("export AWS_PROFILE=" + p[1] + "\n" + "export AWS_REGION=" + *region + "\n")
		} else {
			awsProfile = fmt.Sprint("export AWS_PROFILE=" + p[1] + "\n" + "export AWS_REGION=" + v.Region + "\n")
		}
		data := []byte(awsProfile)
		if err := ioutil.WriteFile("/tmp/aws_profile", data, 0664); err != nil {
			panic(err)
		}
	}
}
