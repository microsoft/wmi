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

// ConnectionStringSettings struct
type ConnectionStringSettings struct {
	*CollectionElement

	//
	ConnectionString string

	//
	Name string

	//
	ProviderName string
}

func NewConnectionStringSettingsEx1(instance *cim.WmiInstance) (newInstance *ConnectionStringSettings, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ConnectionStringSettings{
		CollectionElement: tmp,
	}
	return
}

func NewConnectionStringSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ConnectionStringSettings, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ConnectionStringSettings{
		CollectionElement: tmp,
	}
	return
}

// SetConnectionString sets the value of ConnectionString for the instance
func (instance *ConnectionStringSettings) SetPropertyConnectionString(value string) (err error) {
	return instance.SetProperty("ConnectionString", (value))
}

// GetConnectionString gets the value of ConnectionString for the instance
func (instance *ConnectionStringSettings) GetPropertyConnectionString() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionString")
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
func (instance *ConnectionStringSettings) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ConnectionStringSettings) GetPropertyName() (value string, err error) {
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

// SetProviderName sets the value of ProviderName for the instance
func (instance *ConnectionStringSettings) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", (value))
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *ConnectionStringSettings) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
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
