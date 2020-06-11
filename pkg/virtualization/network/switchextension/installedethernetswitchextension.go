// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchextension

import (
	// 	"log"

	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type InstalledEthernetSwitchExtension struct {
	*v2.Msvm_InstalledEthernetSwitchExtension
}

// NewInstalledEthernetSwitchExtension
func NewInstalledEthernetSwitchExtension(instance *wmi.WmiInstance) (*InstalledEthernetSwitchExtension, error) {
	wmivm, err := v2.NewMsvm_InstalledEthernetSwitchExtensionEx1(instance)
	if err != nil {
		return nil, err
	}
	return &InstalledEthernetSwitchExtension{wmivm}, nil
}

func (iese *InstalledEthernetSwitchExtension) GetFeatureSettingData() (wmi.WmiInstanceCollection, error) {
	return iese.GetAllRelated("Msvm_FeatureSettingData") //,
	//"Msvm_FeatureSettingsDefineCapabilities",
	//"PartComponent", "GroupComponent")
}

func (iese *InstalledEthernetSwitchExtension) GetFeatureCapabilities() (wmi.WmiInstanceCollection, error) {
	return iese.GetAllRelated("Msvm_EthernetSwitchFeatureCapabilities") //,
	//	"Msvm_EthernetSwitchExtensionCapabilities",
	//		"Capabilities", "ManagedElement")
}

func (iese *InstalledEthernetSwitchExtension) GetFeatureCapabilityByName(name string) (*v2.Msvm_EthernetSwitchFeatureCapabilities, error) {
	caps, err := iese.GetFeatureCapabilities()
	if err != nil {
		return nil, err
	}
	defer caps.Close()

	// log.Printf("FeatureCapabilities [%d]\n", len(caps))

	for _, capability := range caps {
		capNameVar, err := capability.GetProperty("ElementName")
		if err != nil {
			return nil, err
		}
		capName := capNameVar.(string)
		if capName == name {
			foundInstance, err := capability.Clone()
			if err != nil {
				return nil, err
			}
			return v2.NewMsvm_EthernetSwitchFeatureCapabilitiesEx1(foundInstance)
		}
	}

	return nil, errors.Wrapf(errors.NotFound, "Unable to find Feature Capability [%s]", name)
}
