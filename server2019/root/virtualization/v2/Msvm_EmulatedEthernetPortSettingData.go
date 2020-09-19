// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EmulatedEthernetPortSettingData struct
type Msvm_EmulatedEthernetPortSettingData struct {
	*CIM_EthernetPortAllocationSettingData

	//
	ClusterMonitored bool

	//
	StaticMacAddress bool
}

func NewMsvm_EmulatedEthernetPortSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EmulatedEthernetPortSettingData, err error) {
	tmp, err := NewCIM_EthernetPortAllocationSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EmulatedEthernetPortSettingData{
		CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
}

func NewMsvm_EmulatedEthernetPortSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EmulatedEthernetPortSettingData, err error) {
	tmp, err := NewCIM_EthernetPortAllocationSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EmulatedEthernetPortSettingData{
		CIM_EthernetPortAllocationSettingData: tmp,
	}
	return
}

// SetClusterMonitored sets the value of ClusterMonitored for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) SetPropertyClusterMonitored(value bool) (err error) {
	return instance.SetProperty("ClusterMonitored", (value))
}

// GetClusterMonitored gets the value of ClusterMonitored for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) GetPropertyClusterMonitored() (value bool, err error) {
	retValue, err := instance.GetProperty("ClusterMonitored")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetStaticMacAddress sets the value of StaticMacAddress for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) SetPropertyStaticMacAddress(value bool) (err error) {
	return instance.SetProperty("StaticMacAddress", (value))
}

// GetStaticMacAddress gets the value of StaticMacAddress for the instance
func (instance *Msvm_EmulatedEthernetPortSettingData) GetPropertyStaticMacAddress() (value bool, err error) {
	retValue, err := instance.GetProperty("StaticMacAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
func (instance *Msvm_EmulatedEthernetPortSettingData) GetRelatedAllocationCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_AllocationCapabilities")
}
