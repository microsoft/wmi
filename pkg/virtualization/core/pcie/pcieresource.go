// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

func NewPciExpress(whost *host.WmiHost) (pcie *v2.Msvm_PciExpress, err error) {
	creds := whost.GetCredential()
	pciExpressQuery := query.NewWmiQuery("Msvm_PciExpress")

	pcie, err = v2.NewMsvm_PciExpressEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, pciExpressQuery)
	return
}

func GetInstancePath(whost *host.WmiHost, deviceId string) (instancePath string, err error) {
	pcie, err := NewPciExpress(whost)
	if err != nil {
		return
	}
	defer pcie.Close()

	pcie.SetPropertyDeviceInstancePath(deviceId)
	instancePath = pcie.InstancePath()
	return
}
