// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_EthernetPort struct
type CIM_EthernetPort struct {
	*CIM_NetworkPort

	//
	Capabilities []EthernetPort_Capabilities

	//
	CapabilityDescriptions []string

	//
	EnabledCapabilities []EthernetPort_EnabledCapabilities

	//
	MaxDataSize uint32

	//
	OtherEnabledCapabilities []string
}

func NewCIM_EthernetPortEx1(instance *cim.WmiInstance) (newInstance *CIM_EthernetPort, err error) {
	tmp, err := NewCIM_NetworkPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_EthernetPort{
		CIM_NetworkPort: tmp,
	}
	return
}

func NewCIM_EthernetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_EthernetPort, err error) {
	tmp, err := NewCIM_NetworkPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_EthernetPort{
		CIM_NetworkPort: tmp,
	}
	return
}

// SetCapabilities sets the value of Capabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyCapabilities(value []EthernetPort_Capabilities) (err error) {
	return instance.SetProperty("Capabilities", (value))
}

// GetCapabilities gets the value of Capabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyCapabilities() (value []EthernetPort_Capabilities, err error) {
	retValue, err := instance.GetProperty("Capabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, EthernetPort_Capabilities(valuetmp))
	}

	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_EthernetPort) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", (value))
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_EthernetPort) GetPropertyCapabilityDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CapabilityDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetEnabledCapabilities sets the value of EnabledCapabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyEnabledCapabilities(value []EthernetPort_EnabledCapabilities) (err error) {
	return instance.SetProperty("EnabledCapabilities", (value))
}

// GetEnabledCapabilities gets the value of EnabledCapabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyEnabledCapabilities() (value []EthernetPort_EnabledCapabilities, err error) {
	retValue, err := instance.GetProperty("EnabledCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, EthernetPort_EnabledCapabilities(valuetmp))
	}

	return
}

// SetMaxDataSize sets the value of MaxDataSize for the instance
func (instance *CIM_EthernetPort) SetPropertyMaxDataSize(value uint32) (err error) {
	return instance.SetProperty("MaxDataSize", (value))
}

// GetMaxDataSize gets the value of MaxDataSize for the instance
func (instance *CIM_EthernetPort) GetPropertyMaxDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDataSize")
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

// SetOtherEnabledCapabilities sets the value of OtherEnabledCapabilities for the instance
func (instance *CIM_EthernetPort) SetPropertyOtherEnabledCapabilities(value []string) (err error) {
	return instance.SetProperty("OtherEnabledCapabilities", (value))
}

// GetOtherEnabledCapabilities gets the value of OtherEnabledCapabilities for the instance
func (instance *CIM_EthernetPort) GetPropertyOtherEnabledCapabilities() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherEnabledCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
