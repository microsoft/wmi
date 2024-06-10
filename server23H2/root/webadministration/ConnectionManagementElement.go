// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ConnectionManagementElement struct
type ConnectionManagementElement struct {
	*CollectionElement

	//
	Address string

	//
	MaxConnection int32
}

func NewConnectionManagementElementEx1(instance *cim.WmiInstance) (newInstance *ConnectionManagementElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConnectionManagementElement{
		CollectionElement: tmp,
	}
	return
}

func NewConnectionManagementElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConnectionManagementElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConnectionManagementElement{
		CollectionElement: tmp,
	}
	return
}

// SetAddress sets the value of Address for the instance
func (instance *ConnectionManagementElement) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *ConnectionManagementElement) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
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

// SetMaxConnection sets the value of MaxConnection for the instance
func (instance *ConnectionManagementElement) SetPropertyMaxConnection(value int32) (err error) {
	return instance.SetProperty("MaxConnection", (value))
}

// GetMaxConnection gets the value of MaxConnection for the instance
func (instance *ConnectionManagementElement) GetPropertyMaxConnection() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxConnection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
