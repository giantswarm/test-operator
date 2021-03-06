package v10

import (
	"time"

	"github.com/giantswarm/versionbundle"
)

func VersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "cloudconfig",
				Description: "Improved encryption key injection.",
				Kind:        versionbundle.KindChanged,
			},
			{
				Component:   "kvmconfig",
				Description: "Updated ingress annotations for api and etcd.",
				Kind:        versionbundle.KindChanged,
			},
		},
		Components: []versionbundle.Component{
			{
				Name:    "calico",
				Version: "3.0.2",
			},
			{
				Name:    "containerlinux",
				Version: "1632.3.0",
			},
			{
				Name:    "docker",
				Version: "17.09.0",
			},
			{
				Name:    "etcd",
				Version: "3.3.1",
			},
			{
				Name:    "coredns",
				Version: "1.0.6",
			},
			{
				Name:    "kubernetes",
				Version: "1.9.5",
			},
			{
				Name:    "nginx-ingress-controller",
				Version: "0.12.0",
			},
		},
		Dependencies: []versionbundle.Dependency{},
		Deprecated:   true,
		Name:         "kvm-operator",
		Time:         time.Date(2018, time.March, 22, 15, 30, 0, 0, time.UTC),
		Version:      "2.1.3",
		WIP:          false,
	}
}
