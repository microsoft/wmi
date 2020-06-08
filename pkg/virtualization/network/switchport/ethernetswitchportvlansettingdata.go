// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortVLANSettingData struct {
	*v2.Msvm_EthernetSwitchPortVlanSettingData
}

// NewEthernetSwitchPortVLANSettingData
func NewEthernetSwitchPortVLANSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortVLANSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortVlanSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortVLANSettingData{wmivm}, nil
}
