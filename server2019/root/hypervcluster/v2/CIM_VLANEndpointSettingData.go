// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VLANEndpointSettingData struct
type CIM_VLANEndpointSettingData struct {
	*CIM_SettingData

	// The access VLAN for the referenced VLANEndpoint.
	AccessVLAN uint16

	// Default value for the native VLAN on this trunk endpoint/port. This property is applicable only when the endpoint is operating in trunking mode (determined by examining the OperationalEndpointMode property).
	DefaultVLAN uint16

	// VLAN Id that is used to tag untagged traffic received on this trunk endpoint/port. This property is applicable only when the endpoint is operating in trunking mode (determined by examining the SwitchEndpointMode property).
	NativeVLAN uint16

	// If a VLAN Id is part of this array, then the system MAY prune that VLAN on the related trunk endpoint/port. This property is applicable only when the endpoint is operating in trunking mode (determined by examining the OperationalEndpointMode property).
	PruneEligibleVLANList []uint16

	// If a VLAN Id is part of this array, then the system MAY trunk the traffic on the related endpoint/port. This property is applicable only when the endpoint is operating in trunking mode (determined by examining the OperationalEndpointMode property).
	TrunkedVLANList []uint16
}

func NewCIM_VLANEndpointSettingDataEx1(instance *cim.WmiInstance) (newInstance *CIM_VLANEndpointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VLANEndpointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

func NewCIM_VLANEndpointSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VLANEndpointSettingData, err error) {
	tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VLANEndpointSettingData{
		CIM_SettingData: tmp,
	}
	return
}

// SetAccessVLAN sets the value of AccessVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) SetPropertyAccessVLAN(value uint16) (err error) {
	return instance.SetProperty("AccessVLAN", (value))
}

// GetAccessVLAN gets the value of AccessVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) GetPropertyAccessVLAN() (value uint16, err error) {
	retValue, err := instance.GetProperty("AccessVLAN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDefaultVLAN sets the value of DefaultVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) SetPropertyDefaultVLAN(value uint16) (err error) {
	return instance.SetProperty("DefaultVLAN", (value))
}

// GetDefaultVLAN gets the value of DefaultVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) GetPropertyDefaultVLAN() (value uint16, err error) {
	retValue, err := instance.GetProperty("DefaultVLAN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetNativeVLAN sets the value of NativeVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) SetPropertyNativeVLAN(value uint16) (err error) {
	return instance.SetProperty("NativeVLAN", (value))
}

// GetNativeVLAN gets the value of NativeVLAN for the instance
func (instance *CIM_VLANEndpointSettingData) GetPropertyNativeVLAN() (value uint16, err error) {
	retValue, err := instance.GetProperty("NativeVLAN")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetPruneEligibleVLANList sets the value of PruneEligibleVLANList for the instance
func (instance *CIM_VLANEndpointSettingData) SetPropertyPruneEligibleVLANList(value []uint16) (err error) {
	return instance.SetProperty("PruneEligibleVLANList", (value))
}

// GetPruneEligibleVLANList gets the value of PruneEligibleVLANList for the instance
func (instance *CIM_VLANEndpointSettingData) GetPropertyPruneEligibleVLANList() (value []uint16, err error) {
	retValue, err := instance.GetProperty("PruneEligibleVLANList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTrunkedVLANList sets the value of TrunkedVLANList for the instance
func (instance *CIM_VLANEndpointSettingData) SetPropertyTrunkedVLANList(value []uint16) (err error) {
	return instance.SetProperty("TrunkedVLANList", (value))
}

// GetTrunkedVLANList gets the value of TrunkedVLANList for the instance
func (instance *CIM_VLANEndpointSettingData) GetPropertyTrunkedVLANList() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TrunkedVLANList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}
