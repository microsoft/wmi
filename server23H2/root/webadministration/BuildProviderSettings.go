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

// BuildProviderSettings struct
type BuildProviderSettings struct {
	*EmbeddedObject

	//
	BuildProviders []BuildProvider
}

func NewBuildProviderSettingsEx1(instance *cim.WmiInstance) (newInstance *BuildProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &BuildProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewBuildProviderSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *BuildProviderSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &BuildProviderSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetBuildProviders sets the value of BuildProviders for the instance
func (instance *BuildProviderSettings) SetPropertyBuildProviders(value []BuildProvider) (err error) {
	return instance.SetProperty("BuildProviders", (value))
}

// GetBuildProviders gets the value of BuildProviders for the instance
func (instance *BuildProviderSettings) GetPropertyBuildProviders() (value []BuildProvider, err error) {
	retValue, err := instance.GetProperty("BuildProviders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(BuildProvider)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " BuildProvider is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, BuildProvider(valuetmp))
	}

	return
}
