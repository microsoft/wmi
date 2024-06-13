// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Group struct
type CIM_Group struct {
	*CIM_Collection

	//
	BusinessCategory string

	//
	CommonName string

	//
	CreationClassName string

	//
	Name string
}

func NewCIM_GroupEx1(instance *cim.WmiInstance) (newInstance *CIM_Group, err error) {
	tmp, err := NewCIM_CollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Group{
		CIM_Collection: tmp,
	}
	return
}

func NewCIM_GroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Group, err error) {
	tmp, err := NewCIM_CollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Group{
		CIM_Collection: tmp,
	}
	return
}

// SetBusinessCategory sets the value of BusinessCategory for the instance
func (instance *CIM_Group) SetPropertyBusinessCategory(value string) (err error) {
	return instance.SetProperty("BusinessCategory", (value))
}

// GetBusinessCategory gets the value of BusinessCategory for the instance
func (instance *CIM_Group) GetPropertyBusinessCategory() (value string, err error) {
	retValue, err := instance.GetProperty("BusinessCategory")
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

// SetCommonName sets the value of CommonName for the instance
func (instance *CIM_Group) SetPropertyCommonName(value string) (err error) {
	return instance.SetProperty("CommonName", (value))
}

// GetCommonName gets the value of CommonName for the instance
func (instance *CIM_Group) GetPropertyCommonName() (value string, err error) {
	retValue, err := instance.GetProperty("CommonName")
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

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Group) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Group) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetName sets the value of Name for the instance
func (instance *CIM_Group) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *CIM_Group) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
