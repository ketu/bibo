package bibo

import (
	"gopkg.in/yaml.v2"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

var (
	defaultFilename = "bibo.yaml"
)

type Configuration struct {
	host              string
	Debug             bool                   `json:"debug,omitempty" yaml:"Debug" toml:"Debug"`
	Charset           string                 `json:"charset,omitempty" yaml:"Charset" toml:"Charset"`
	DatabaseHost      string                 `yaml:"DatabaseHost" toml:"DatabaseHost"`
	DatabaseName      string                 `yaml:"DatabaseName" toml:"DatabaseName"`
	DatabaseUser      string                 `yaml:"DatabaseUser" toml:"DatabaseUser"`
	DatabasePasswd    string                 `yaml:"DatabasePasswd" toml:"DatabasePasswd"`
	RemoteAddrHeaders map[string]bool        `json:"remoteAddrHeaders,omitempty" yaml:"RemoteAddrHeaders" toml:"RemoteAddrHeaders"`
	Other             map[string]interface{} `json:"other,omitempty" yaml:"Other" toml:"Other"`
}

func (c Configuration) getHost() string {
	return c.host
}

func parseConfigurationFile(filename string) (Configuration, error) {
	c := Configuration{}
	file, _ := filepath.Abs(filename)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal([]byte(data), &c)
	return c, err
}

func DefaultConfiguration() Configuration {
	c, err := parseConfigurationFile(defaultFilename)
	if err != nil {
		fmt.Println(err)
	}
	return c
}
