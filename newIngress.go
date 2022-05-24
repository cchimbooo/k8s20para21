package main

import "gopkg.in/yaml.v3"

type NewPort struct {
	Name   string `yaml:"name,omitempty"`
	Number int    `yaml:"number,omitempty"`
}

type NewService struct {
	Name string  `yaml:"name"`
	Port NewPort `yaml:"port"`
}

type NewBackend struct {
	Service NewService `yaml:"service"`
}

type NewPath struct {
	Path     string     `yaml:"path"`
	PathType string     `yaml:"pathType"`
	Backend  NewBackend `yaml:"backend"`
}

type NewHttp struct {
	Paths []NewPath `yaml:"paths"`
}

type NewRules struct {
	Host string  `yaml:"host"`
	Http NewHttp `yaml:"http"`
}

type NewSpec struct {
	Rules []NewRules `yaml:"rules"`
}

type NewIngress struct {
	ApiVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   map[string]interface{} `yaml:"metadata"`
	Spec       NewSpec                `yaml:"spec"`
}

func (n NewIngress) ToBYaml() ([]byte, error) {
	// Converte para []bytes
	bYaml, err := yaml.Marshal(n)
	if err != nil {
		panic(err.Error())
	}
	return bYaml, err
}
