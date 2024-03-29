// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"fmt"
	//"log"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"

	"github.com/microsoft/wmi/pkg/virtualization/network/ethernetport"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchextension"
	"github.com/microsoft/wmi/pkg/virtualization/network/switchport"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

const (
	IsolationNone                  uint32 = 0
	IsolationNativeVirtualSubnet   uint32 = 1
	IsolationExternalVirtualSubnet uint32 = 2
	IsolationVLAN                  uint32 = 3
)

type HostComputerSystem struct {
	*v2.Msvm_ComputerSystem
}

// NewHostComputerSystem
func NewHostComputerSystem(instance *wmi.WmiInstance) (*HostComputerSystem, error) {
	wmivm, err := v2.NewMsvm_ComputerSystemEx1(instance)
	if err != nil {
		return nil, err
	}
	return &HostComputerSystem{wmivm}, nil
}

// GetHostComputerSystem gets an existing virtual machine
func GetHostComputerSystem(whost *host.WmiHost) (vm *HostComputerSystem, err error) {
	creds := whost.GetCredential()
	querytmp := query.NewWmiQuery("Msvm_ComputerSystem")
	querytmp.AddFilterWithComparer("Description", "Microsoft Virtual Machine", query.NotEquals)
	wmivm, err := v2.NewMsvm_ComputerSystemEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, querytmp)
	if err != nil {
		return
	}
	vm = &HostComputerSystem{wmivm}
	return
}

func GetDefaultEthernetPortAllocationSettingData(whost *host.WmiHost) (*switchport.EthernetPortAllocationSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_EthernetPortAllocationSettingData")
	if err != nil {
		return nil, err
	}
	return switchport.NewEthernetPortAllocationSettingData(inst)
}

func GetInternalPortAllocationSettingData(whost *host.WmiHost, switchPortName string) (epas *switchport.EthernetPortAllocationSettingData, err error) {
	epas, err = GetDefaultEthernetPortAllocationSettingData(whost)
	if err != nil {
		return nil, err
	}

	hostcm, err := GetHostComputerSystem(whost)
	if err != nil {
		return nil, err
	}
	defer hostcm.Close()

	err = epas.SetPropertyElementName(switchPortName)
	if err != nil {
		return
	}
	err = epas.SetPropertyHostResource([]string{hostcm.InstancePath()})
	if err != nil {
		return
	}

	return
}

// GetExternalPortAllocationSettingData
func GetExternalPortAllocationSettingData(whost *host.WmiHost, switchPortName string, physicalNicNames []string) (epas *switchport.EthernetPortAllocationSettingData, err error) {
	epas, err = GetDefaultEthernetPortAllocationSettingData(whost)
	if err != nil {
		return nil, err
	}

	if len(physicalNicNames) == 0 {
		err = errors.Wrapf(errors.InvalidInput, "Physical Nic Name is missing")
		return
	}
	err = epas.SetPropertyElementName(switchPortName)
	if err != nil {
		return
	}

	hresource := []string{}
	for _, nicName := range physicalNicNames {
		extPort, err1 := ethernetport.GetExternalEthernetPort(whost, nicName)
		if err1 != nil {
			err = err1
			return
		}
		defer extPort.Close()
		hresource = append(hresource, extPort.InstancePath())
	}

	err = epas.SetPropertyHostResource(hresource)
	if err != nil {
		return
	}
	return
}

func (hc *HostComputerSystem) GetInstalledEthernetSwitchExtensions() (col switchextension.InstalledEthernetSwitchExtensionCollection, err error) {
	tmpCol, err := hc.GetAllRelated("Msvm_InstalledEthernetSwitchExtension")
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			col.Close()
			col = nil
		}
	}()

	for _, inst := range tmpCol {
		tmp, err1 := switchextension.NewInstalledEthernetSwitchExtension(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, tmp)
	}

	//log.Printf("GetInstalledEthernetSwitchExtensions [%d]\n", len(col))
	return
}

func (hc *HostComputerSystem) GetInstalledEthernetSwitchExtensionByName(name string) (ext *wmi.WmiInstance, err error) {
	exts, err := hc.GetInstalledEthernetSwitchExtensions()
	if err != nil {
		return
	}

	defer exts.Close()

	for _, exttmp := range exts {
		eName, err1 := exttmp.GetPropertyElementName()
		if err1 != nil {
			err = err1
			return
		}
		if eName == name {
			return exttmp.Clone()
		}
	}

	return nil, errors.Wrapf(errors.NotFound, "EthernetSwitchExtension [%s]", name)
}

func (hc *HostComputerSystem) GetDefaultPortSettingData(featureName, className string) (
	*wmi.WmiInstance, error) {

	capability, err := hc.GetFeatureCapability(featureName)
	if err != nil {
		return nil, err
	}
	defer capability.Close()

	return capability.GetRelated(className)
}

func (hc *HostComputerSystem) GetFeatureCapability(featureName string) (*v2.Msvm_EthernetSwitchFeatureCapabilities, error) {
	exts, err := hc.GetInstalledEthernetSwitchExtensions()
	if err != nil {
		return nil, err
	}
	defer exts.Close()

	for _, ext := range exts {
		fc, err := ext.GetFeatureCapabilityByName(featureName)
		if err != nil {
			if errors.IsNotFound(err) {
				continue
			}
			return nil, err
		}
		// If found, return
		return fc, nil
	}

	return nil, errors.Wrapf(errors.NotFound, "FeatureCapability [%s]", featureName)
}

func (hc *HostComputerSystem) GetDefaultEthernetSwitchPortProfileSettingData(profileName,
	vendorGuid, profileGuid string, profileData uint32) (*switchport.EthernetSwitchPortProfileSettingData, error) {
	tmp, err := hc.GetDefaultPortSettingData("Ethernet Switch Port Profile Settings", "Msvm_EthernetSwitchPortProfileSettingData")
	if err != nil {
		return nil, err
	}

	spps, err := switchport.NewEthernetSwitchPortProfileSettingData(tmp)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			spps.Close()
			spps = nil
		}
	}()

	err = spps.SetPropertyProfileName(profileName)
	if err != nil {
		return nil, err
	}
	err = spps.SetPropertyProfileId(fmt.Sprintf("{%s}", profileGuid))
	if err != nil {
		return nil, err
	}
	err = spps.SetPropertyProfileData(profileData)
	if err != nil {
		return nil, err
	}
	err = spps.SetPropertyVendorId(fmt.Sprintf("{%s}", vendorGuid))
	if err != nil {
		return nil, err
	}
	return spps, nil
}

func (hc *HostComputerSystem) GetDefaultEthernetSwitchPortVLANSettingData(vlanId uint16) (*switchport.EthernetSwitchPortVLANSettingData, error) {
	tmp, err := hc.GetDefaultPortSettingData("Ethernet Switch Port VLAN Settings", "Msvm_EthernetSwitchPortVLANSettingData")
	if err != nil {
		return nil, err
	}

	spps, err := switchport.NewEthernetSwitchPortVLANSettingData(tmp)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			spps.Close()
			spps = nil
		}
	}()

	err = spps.SetPropertyAccessVlanId(vlanId)
	if err != nil {
		return nil, err
	}
	err = spps.SetPropertyOperationMode(1)
	if err != nil {
		return nil, err
	}

	return spps, nil
}

func (hc *HostComputerSystem) GetDefaultEthernetSwitchPortIsolationSettingData(vlanId uint32) (*switchport.EthernetSwitchPortIsolationSettingData, error) {
	tmp, err := hc.GetDefaultPortSettingData("Ethernet Switch Port Isolation Settings", "Msvm_EthernetSwitchPortIsolationSettingData")
	if err != nil {
		return nil, err
	}

	spps, err := switchport.NewEthernetSwitchPortIsolationSettingData(tmp)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			spps.Close()
			spps = nil
		}
	}()

	err = spps.SetPropertyDefaultIsolationId(vlanId)
	if err != nil {
		return nil, err
	}

	// VLAN 3
	err = spps.SetPropertyIsolationMode(IsolationVLAN)
	if err != nil {
		return nil, err
	}

	err = spps.SetPropertyAllowUntaggedTraffic(false)
	if err != nil {
		return nil, err
	}

	err = spps.SetPropertyEnableMultiTenantStack(false)
	if err != nil {
		return nil, err
	}

	return spps, nil
}
