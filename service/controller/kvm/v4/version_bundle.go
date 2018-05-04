package v3

import (
	"time"

	"github.com/giantswarm/versionbundle"
)

func VersionBundle() versionbundle.Bundle {
	return versionbundle.Bundle{
		Changelogs: []versionbundle.Changelog{
			{
				Component:   "test-operator",
				Description: "Installed chart-operator in kube-system namespace.",
				Kind:        versionbundle.KindChanged,
			},
			{
				Component:   "test-operator",
				Description: "Removed misleading component reference to kvm-operator.",
				Kind:        versionbundle.KindFixed,
			},
		},
		Components:   []versionbundle.Component{},
		Dependencies: []versionbundle.Dependency{},
		Deprecated:   false,
		Name:         "test-operator",
		Provider:     "kvm",
		Time:         time.Date(2018, time.April, 26, 12, 00, 0, 0, time.UTC),
		Version:      "0.3.0",
		WIP:          true,
	}
}
