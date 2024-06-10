// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_LogicalPort struct
type CIM_LogicalPort struct {
	*CIM_LogicalDevice

	//
	MaxSpeed uint64

	//
	OtherPortType string

	//
	PortType uint16

	//
	RequestedSpeed uint64

	//
	Speed uint64

	//
	UsageRestriction uint16
}

func NewCIM_LogicalPortEx1(instance *cim.WmiInstance) (newInstance *CIM_LogicalPort, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalPort{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_LogicalPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_LogicalPort, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_LogicalPort{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetMaxSpeed sets the value of MaxSpeed for the instance
func (instance *CIM_LogicalPort) SetPropertyMaxSpeed(value uint64) (err error) {
	return instance.SetProperty("MaxSpeed", (value))
}

// GetMaxSpeed gets the value of MaxSpeed for the instance
func (instance *CIM_LogicalPort) GetPropertyMaxSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetOtherPortType sets the value of OtherPortType for the instance
func (instance *CIM_LogicalPort) SetPropertyOtherPortType(value string) (err error) {
	return instance.SetProperty("OtherPortType", (value))
}

// GetOtherPortType gets the value of OtherPortType for the instance
func (instance *CIM_LogicalPort) GetPropertyOtherPortType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherPortType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPortType sets the value of PortType for the instance
func (instance *CIM_LogicalPort) SetPropertyPortType(value uint16) (err error) {
	return instance.SetProperty("PortType", (value))
}

// GetPortType gets the value of PortType for the instance
func (instance *CIM_LogicalPort) GetPropertyPortType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortType")
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

// SetRequestedSpeed sets the value of RequestedSpeed for the instance
func (instance *CIM_LogicalPort) SetPropertyRequestedSpeed(value uint64) (err error) {
	return instance.SetProperty("RequestedSpeed", (value))
}

// GetRequestedSpeed gets the value of RequestedSpeed for the instance
func (instance *CIM_LogicalPort) GetPropertyRequestedSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestedSpeed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSpeed sets the value of Speed for the instance
func (instance *CIM_LogicalPort) SetPropertySpeed(value uint64) (err error) {
	return instance.SetProperty("Speed", (value))
}

// GetSpeed gets the value of Speed for the instance
func (instance *CIM_LogicalPort) GetPropertySpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("Speed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUsageRestriction sets the value of UsageRestriction for the instance
func (instance *CIM_LogicalPort) SetPropertyUsageRestriction(value uint16) (err error) {
	return instance.SetProperty("UsageRestriction", (value))
}

// GetUsageRestriction gets the value of UsageRestriction for the instance
func (instance *CIM_LogicalPort) GetPropertyUsageRestriction() (value uint16, err error) {
	retValue, err := instance.GetProperty("UsageRestriction")
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
