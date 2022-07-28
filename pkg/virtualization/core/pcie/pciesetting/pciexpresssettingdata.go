// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pciesetting

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type PciExpressSettingData struct {
	*v2.Msvm_PciExpressSettingData
}

func NewPciExpressSettingData(instance *wmi.WmiInstance) (*PciExpressSettingData, error) {
	wmivm, err := v2.NewMsvm_PciExpressSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PciExpressSettingData{wmivm}, nil
}
