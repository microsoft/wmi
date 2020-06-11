// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortOffloadSettingData struct {
	*v2.Msvm_EthernetSwitchPortOffloadSettingData
}

// NewEthernetSwitchPortOffloadSettingData
func NewEthernetSwitchPortOffloadSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortOffloadSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortOffloadSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortOffloadSettingData{wmivm}, nil
}
