// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortIsolationSettingData struct {
	*v2.Msvm_EthernetSwitchPortIsolationSettingData
}

// NewEthernetSwitchPortIsolationSettingData
func NewEthernetSwitchPortIsolationSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortIsolationSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortIsolationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortIsolationSettingData{wmivm}, nil
}
