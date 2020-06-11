// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortProfileSettingData struct {
	*v2.Msvm_EthernetSwitchPortProfileSettingData
}

// NewEthernetSwitchPortProfileSettingData
func NewEthernetSwitchPortProfileSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortProfileSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortProfileSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortProfileSettingData{wmivm}, nil
}
