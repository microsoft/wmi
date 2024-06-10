// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Powershellv3
//
// ////////////////////////////////////////////
package powershellv3

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ManagedElement struct
type CIM_ManagedElement struct {
	*cim.WmiInstance

	//
	Caption string

	//
	Description string

	//
	ElementName string

	//
	InstanceID string
}

func NewCIM_ManagedElementEx1(instance *cim.WmiInstance) (newInstance *CIM_ManagedElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ManagedElement{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ManagedElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ManagedElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ManagedElement{
		WmiInstance: tmp,
	}
	return
}

// SetCaption sets the value of Caption for the instance
func (instance *CIM_ManagedElement) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", (value))
}

// GetCaption gets the value of Caption for the instance
func (instance *CIM_ManagedElement) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
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

// SetDescription sets the value of Description for the instance
func (instance *CIM_ManagedElement) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *CIM_ManagedElement) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetElementName sets the value of ElementName for the instance
func (instance *CIM_ManagedElement) SetPropertyElementName(value string) (err error) {
	return instance.SetProperty("ElementName", (value))
}

// GetElementName gets the value of ElementName for the instance
func (instance *CIM_ManagedElement) GetPropertyElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ElementName")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *CIM_ManagedElement) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *CIM_ManagedElement) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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
