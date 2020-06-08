// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortRdmaSettingData struct {
	*v2.Msvm_EthernetSwitchPortRdmaSettingData
}

// NewEthernetSwitchPortRdmaSettingData
func NewEthernetSwitchPortRdmaSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortRdmaSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortRdmaSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortRdmaSettingData{wmivm}, nil
}
