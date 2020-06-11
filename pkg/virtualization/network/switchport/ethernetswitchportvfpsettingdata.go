// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortVfpSettingData struct {
	*v2.Msvm_EthernetSwitchPortVfpSettingData
}

// NewEthernetSwitchPortVfpSettingData
func NewEthernetSwitchPortVfpSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortVfpSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortVfpSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortVfpSettingData{wmivm}, nil
}
