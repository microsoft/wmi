// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualEthernetSwitch struct
type Msvm_VirtualEthernetSwitch struct {
	*CIM_ComputerSystem

	//
	MaxIOVOffloads uint32

	//
	MaxVMQOffloads uint32
}

func NewMsvm_VirtualEthernetSwitchEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualEthernetSwitch, err error) {
	tmp, err := NewCIM_ComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitch{
		CIM_ComputerSystem: tmp,
	}
	return
}

func NewMsvm_VirtualEthernetSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualEthernetSwitch, err error) {
	tmp, err := NewCIM_ComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualEthernetSwitch{
		CIM_ComputerSystem: tmp,
	}
	return
}

// SetMaxIOVOffloads sets the value of MaxIOVOffloads for the instance
func (instance *Msvm_VirtualEthernetSwitch) SetPropertyMaxIOVOffloads(value uint32) (err error) {
	return instance.SetProperty("MaxIOVOffloads", (value))
}

// GetMaxIOVOffloads gets the value of MaxIOVOffloads for the instance
func (instance *Msvm_VirtualEthernetSwitch) GetPropertyMaxIOVOffloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxIOVOffloads")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMaxVMQOffloads sets the value of MaxVMQOffloads for the instance
func (instance *Msvm_VirtualEthernetSwitch) SetPropertyMaxVMQOffloads(value uint32) (err error) {
	return instance.SetProperty("MaxVMQOffloads", (value))
}

// GetMaxVMQOffloads gets the value of MaxVMQOffloads for the instance
func (instance *Msvm_VirtualEthernetSwitch) GetPropertyMaxVMQOffloads() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxVMQOffloads")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
func (instance *Msvm_VirtualEthernetSwitch) GetRelatedEthernetSwitchBandwidthData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchBandwidthData")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedTransparentBridgingService() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_TransparentBridgingService")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedEthernetSwitchOperationalData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchOperationalData")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedEthernetSwitchHardwareOffloadData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchHardwareOffloadData")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedEthernetSwitchExtension() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_EthernetSwitchExtension")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedResourcePool() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ResourcePool")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedVirtualEthernetSwitchSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchSettingData")
}

func (instance *Msvm_VirtualEthernetSwitch) GetRelatedEthernetSwitchPort() (value []*cim.WmiInstance, err error) {
	return instance.GetAllRelated("Msvm_EthernetSwitchPort")
}
