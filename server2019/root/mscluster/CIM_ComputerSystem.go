// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.MSCluster
//
// ////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ComputerSystem struct
type CIM_ComputerSystem struct {
	*CIM_System

	//
	Dedicated []uint16

	//
	IdentifyingDescriptions []string

	//
	OtherIdentifyingInfo []string
}

func NewCIM_ComputerSystemEx1(instance *cim.WmiInstance) (newInstance *CIM_ComputerSystem, err error) {
	tmp, err := NewCIM_SystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ComputerSystem{
		CIM_System: tmp,
	}
	return
}

func NewCIM_ComputerSystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ComputerSystem, err error) {
	tmp, err := NewCIM_SystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ComputerSystem{
		CIM_System: tmp,
	}
	return
}

// SetDedicated sets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) SetPropertyDedicated(value []uint16) (err error) {
	return instance.SetProperty("Dedicated", (value))
}

// GetDedicated gets the value of Dedicated for the instance
func (instance *CIM_ComputerSystem) GetPropertyDedicated() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Dedicated")
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

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_ComputerSystem) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", (value))
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_ComputerSystem) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
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

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_ComputerSystem) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", (value))
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_ComputerSystem) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
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
