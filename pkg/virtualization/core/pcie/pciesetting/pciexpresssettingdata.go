// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pciesetting

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

func NewPciExpressSettingDataEx1(instance *wmi.WmiInstance) (*PciExpressSettingData, error) {
	wmivm, err := v2.NewMsvm_PciExpressSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PciExpressSettingData{wmivm}, nil
}

func NewPciExpressSettingDataEx6(whost *host.WmiHost) (*PciExpressSettingData, error) {
	creds := whost.GetCredential()
	settingDataQuery := query.NewWmiQuery("Msvm_PciExpressSettingData")
	wmivm, err := v2.NewMsvm_PciExpressSettingDataEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, settingDataQuery)
	if err != nil {
		return nil, err
	}
	return &PciExpressSettingData{wmivm}, nil
}
