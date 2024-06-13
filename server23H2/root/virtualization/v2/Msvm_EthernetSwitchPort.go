// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPort struct
type Msvm_EthernetSwitchPort struct {
	*CIM_EthernetPort

	//
	IOVOffloadUsage uint32

	//
	VMQOffloadUsage uint32
}

func NewMsvm_EthernetSwitchPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPort, err error) {
	tmp, err := NewCIM_EthernetPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPort, err error) {
	tmp, err := NewCIM_EthernetPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPort{
		CIM_EthernetPort: tmp,
	}
	return
}

// SetIOVOffloadUsage sets the value of IOVOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) SetPropertyIOVOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("IOVOffloadUsage", (value))
}

// GetIOVOffloadUsage gets the value of IOVOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) GetPropertyIOVOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("IOVOffloadUsage")
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

// SetVMQOffloadUsage sets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) SetPropertyVMQOffloadUsage(value uint32) (err error) {
	return instance.SetProperty("VMQOffloadUsage", (value))
}

// GetVMQOffloadUsage gets the value of VMQOffloadUsage for the instance
func (instance *Msvm_EthernetSwitchPort) GetPropertyVMQOffloadUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("VMQOffloadUsage")
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
func (instance *Msvm_EthernetSwitchPort) GetRelatedVirtualEthernetSwitch() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitch")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedLANEndpoint() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_LANEndpoint")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetSwitchPortOffloadData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPortOffloadData")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetSwitchPortBandwidthData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchPortBandwidthData")
}

func (instance *Msvm_EthernetSwitchPort) GetRelatedEthernetPortAllocationSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetPortAllocationSettingData")
}
