package main

import "gopkg.in/yaml.v3"

type NewPort struct {
	Name   string `yaml:"name,omitempty,omitempty"`
	Number int    `yaml:"number,omitempty,omitempty"`
}

type NewService struct {
	Name string  `yaml:"name,omitempty"`
	Port NewPort `yaml:"port,omitempty"`
}

type NewBackend struct {
	Service NewService `yaml:"service,omitempty"`
}

type NewPath struct {
	Path     string     `yaml:"path,omitempty"`
	PathType string     `yaml:"pathType,omitempty"`
	Backend  NewBackend `yaml:"backend,omitempty"`
}

type NewHttp struct {
	Paths []NewPath `yaml:"paths,omitempty"`
}

type NewRules struct {
	Host string  `yaml:"host,omitempty"`
	Http NewHttp `yaml:"http,omitempty"`
}

type NewSpec struct {
	Rules []NewRules `yaml:"rules,omitempty"`
}

type NewIngress struct {
	ApiVersion string                 `yaml:"apiVersion,omitempty"`
	Kind       string                 `yaml:"kind,omitempty"`
	Metadata   map[string]interface{} `yaml:"metadata,omitempty"`
	Spec       NewSpec                `yaml:"spec,omitempty"`
}

func (n NewIngress) ToBYaml() ([]byte, error) {
	// Converte para []bytes
	bYaml, err := yaml.Marshal(n)
	if err != nil {
		panic(err.Error())
	}
	return bYaml, err
}
