// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortAclSettingData struct {
	*v2.Msvm_EthernetSwitchPortAclSettingData
}

// NewEthernetSwitchPortAclSettingData
func NewEthernetSwitchPortAclSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortAclSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortAclSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortAclSettingData{wmivm}, nil
}
