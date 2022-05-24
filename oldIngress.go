package main

import "strconv"

func (i IngressStruct) ConverterParaNovo() NewIngress {
	n := NewIngress{
		ApiVersion: "networking.k8s.io/v1",
		Kind:       "Ingress",
		Metadata:   i.Metadata,
	}

	ns := NewSpec{Rules: make([]NewRules, len(i.Spec.Rules))}

	for k, v := range i.Spec.Rules {
		ns.Rules[k] = v.ToNewRules()
	}
	n.Spec = ns
	return n

}

type OldBackend struct {
	ServiceName string `yaml:"serviceName"`
	ServicePort string `yaml:"servicePort"`
}

func (o OldBackend) ToNewBackend() NewBackend {
	if o.ServiceName == "ssl-redirect" {
		return NewBackend{
			Service: NewService{
				Name: "ssl-redirect",
				Port: NewPort{
					Name: o.ServicePort,
				},
			},
		}
	}
	nw := NewBackend{
		Service: NewService{
			Name: o.ServiceName,
			Port: NewPort{},
		},
	}
	num, err := strconv.Atoi(o.ServicePort)
	if err != nil {
		panic(err)
	}

	nw.Service.Port.Number = num
	return nw
}

type OldPath struct {
	Path    string     `yaml:"path"`
	Backend OldBackend `yaml:"backend"`
}

func (o OldPath) ToNewPath() NewPath {
	np := NewPath{
		Path:     o.Path,
		PathType: "ImplementationSpecific",
		Backend:  o.Backend.ToNewBackend(),
	}
	if np.Backend.Service.Name == "ssl-redirect" {
		np.Path = "/"
		np.PathType = "Prefix"
	}
	return np
}

type OldHttp struct {
	Paths []OldPath `yaml:"paths"`
}

type OldRules struct {
	Host string  `yaml:"host"`
	Http OldHttp `yaml:"http"`
}

func (o OldRules) ToNewRules() NewRules {
	nr := NewRules{
		Host: o.Host,
		Http: NewHttp{
			Paths: make([]NewPath, len(o.Http.Paths)),
		},
	}

	for k, v := range o.Http.Paths {
		nr.Http.Paths[k] = v.ToNewPath()
	}
	return nr
}

type OldSpec struct {
	Rules []OldRules `yaml:"rules"`
}

type IngressStruct struct {
	ApiVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Metadata   map[string]interface{} `yaml:"metadata"`
	Spec       OldSpec                `yaml:"spec"`
}
