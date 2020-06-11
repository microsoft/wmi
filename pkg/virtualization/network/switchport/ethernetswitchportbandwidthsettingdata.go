// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchport

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchPortBandwidthSettingData struct {
	*v2.Msvm_EthernetSwitchPortBandwidthSettingData
}

// NewEthernetSwitchPortBandwidthSettingData
func NewEthernetSwitchPortBandwidthSettingData(instance *wmi.WmiInstance) (*EthernetSwitchPortBandwidthSettingData, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchPortBandwidthSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchPortBandwidthSettingData{wmivm}, nil
}
