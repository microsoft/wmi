// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type PciExpressSettingData struct {
	*v2.Msvm_PciExpressSettingData
}

// NewPciExpressSettingData
func NewPciExpressSettingData(instance *wmi.WmiInstance) (*PciExpressSettingData, error) {
	wmivm, err := v2.NewMsvm_PciExpressSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PciExpressSettingData{wmivm}, nil
}

func GetDefaultPciExpressSettingData(whost *host.WmiHost) (*PciExpressSettingData, error) {
	creds := whost.GetCredential()
	wmivm, err := v2.NewMsvm_PciExpressSettingDataEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query.NewWmiQuery("Msvm_PciExpressSettingData"))
	if err != nil {
		return nil, err
	}
	return &PciExpressSettingData{wmivm}, nil
}
