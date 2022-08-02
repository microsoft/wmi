// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

func NewPcieDeviceEx6(wmihost *host.WmiHost) (device *v2.Msvm_PciExpress, err error) {
	creds := wmihost.GetCredential()
	pciExpressQuery := query.NewWmiQuery("Msvm_PciExpress")

	device, err = v2.NewMsvm_PciExpressEx6(wmihost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, pciExpressQuery)
	return
}

func GetHostResource(whost *host.WmiHost, locationPath string) (hostResource string, err error) {
	device, err := NewPcieDeviceEx6(whost)
	if err != nil {
		return
	}
	device.SetPropertyLocationPath(locationPath)
	hostResource = device.InstancePath()
	return
}
