// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPort struct {
	*v2.Msvm_EthernetSwitchPort
}

// NewEthernetSwitchPort
func NewEthernetSwitchPort(instance *wmi.WmiInstance) (*EthernetSwitchPort, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPort{wmivm}, nil
}

func (vs *EthernetSwitchPort) GetEthernetPortAllocationSettingData() (*EthernetPortAllocationSettingData, error) {
	wmipasd, err := vs.GetAllRelated("Msvm_EthernetPortAllocationSettingData")
	////	"Msvm_ElementSettingData",
	//	"SettingData", "ManagedElement")
	if err != nil {
		return nil, err
	}
	if len(wmipasd) == 0 {
		return nil, errors.Wrapf(errors.NotFound, "No Related Instances of Msvm_EthernetPortAllocationSettingData found")
	}
	defer wmipasd.Close()

	pasd, err := wmipasd[0].Clone()
	if err != nil {
		return nil, err
	}

	return NewEthernetPortAllocationSettingData(pasd)
}
