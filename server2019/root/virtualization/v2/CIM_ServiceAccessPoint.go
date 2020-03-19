// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ServiceAccessPoint struct
type CIM_ServiceAccessPoint struct {
	*CIM_EnabledLogicalElement

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	// The CreationClassName of the scoping System.
	SystemCreationClassName string

	// The Name of the scoping System.
	SystemName string
}

func NewCIM_ServiceAccessPointEx1(instance *cim.WmiInstance) (newInstance *CIM_ServiceAccessPoint, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAccessPoint{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_ServiceAccessPointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ServiceAccessPoint, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAccessPoint{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_ServiceAccessPoint) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_ServiceAccessPoint) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_ServiceAccessPoint) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_ServiceAccessPoint) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_ServiceAccessPoint) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_ServiceAccessPoint) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
