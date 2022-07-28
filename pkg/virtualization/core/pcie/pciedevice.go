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

type PcieDevice struct {
	*PciExpress
}

func NewPcieDeviceEx1(instance *wmi.WmiInstance) (*PcieDevice, error) {
	wmivm, err := NewPciExpress(instance)
	if err != nil {
		return nil, err
	}
	return &PcieDevice{wmivm}, nil
}

func NewPcieDeviceEx6(wmihost *host.WmiHost) (device *v2.Msvm_PciExpress, err error) {
	creds := wmihost.GetCredential()
	pciExpressQuery := query.NewWmiQuery("Msvm_PciExpress")

	device, err = v2.NewMsvm_PciExpressEx6(wmihost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, pciExpressQuery)
	return
}

func (pcie *PcieDevice) CloneEx1() (*PcieDevice, error) {
	tmp, err := pcie.Clone()
	if err != nil {
		return nil, err
	}
	return NewPcieDeviceEx1(tmp)
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

// GetVirtualMachine gets the VM that the pci express is attached to
func (pcie *PcieDevice) GetVirtualMachine() (instance *wmi.WmiInstance, err error) {
	vmsetting, err := pcie.GetRelated("Msvm_VirtualSystemSettingData")
	if err != nil {
		return
	}
	defer vmsetting.Close()
	return vmsetting.GetRelated("Msvm_ComputerSystem")
}
