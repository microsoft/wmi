// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetPortAllocationSettingData struct {
	*v2.Msvm_EthernetPortAllocationSettingData
}

// NewEthernetPortAllocationSettingData
func NewEthernetPortAllocationSettingData(instance *wmi.WmiInstance) (*EthernetPortAllocationSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetPortAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetPortAllocationSettingData{wmivm}, nil
}
