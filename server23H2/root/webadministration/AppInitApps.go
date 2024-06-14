// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppInitApps struct
type AppInitApps struct {
	*CollectionElement

	//
	HostName string

	//
	InitializationPage string
}

func NewAppInitAppsEx1(instance *cim.WmiInstance) (newInstance *AppInitApps, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AppInitApps{
		CollectionElement: tmp,
	}
	return
}

func NewAppInitAppsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppInitApps, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppInitApps{
		CollectionElement: tmp,
	}
	return
}

// SetHostName sets the value of HostName for the instance
func (instance *AppInitApps) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *AppInitApps) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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

// SetInitializationPage sets the value of InitializationPage for the instance
func (instance *AppInitApps) SetPropertyInitializationPage(value string) (err error) {
	return instance.SetProperty("InitializationPage", (value))
}

// GetInitializationPage gets the value of InitializationPage for the instance
func (instance *AppInitApps) GetPropertyInitializationPage() (value string, err error) {
	retValue, err := instance.GetProperty("InitializationPage")
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
