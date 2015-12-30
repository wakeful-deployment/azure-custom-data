package main

import (
	"encoding/base64"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
)

const defaultXmlPath = "/var/lib/waagent/ovf-env.xml"

type LinuxProvisioningConfigurationSet struct {
	CustomData string
}

type ProvisioningSection struct {
	LinuxProvisioningConfigurationSet LinuxProvisioningConfigurationSet
}

type Environment struct {
	ProvisioningSection ProvisioningSection
}

func readFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return dat
}

func parseXML(data []byte) Environment {
	e := Environment{}

	err := xml.Unmarshal(data, &e)

	if err != nil {
		panic(err)
	}

	return e
}

func decodeCustomData(data string) string {
	result, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		panic(err)
	}

	return string(result)
}

func readAndOutput(path string) error {
	e := parseXML(readFile(path))
	result := decodeCustomData(e.ProvisioningSection.LinuxProvisioningConfigurationSet.CustomData)

	fmt.Println(result)

	return nil
}

func main() {
	path := flag.String("path", "", "[optional] Path to xml file")
	flag.Parse()

	if *path == "" {
		readAndOutput(defaultXmlPath)
	} else {
		readAndOutput(*path)
	}
}
