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

// OutputCacheProfileSettings struct
type OutputCacheProfileSettings struct {
	*EmbeddedObject

	//
	OutputCacheProfiles []OutputCacheProfile
}

func NewOutputCacheProfileSettingsEx1(instance *cim.WmiInstance) (newInstance *OutputCacheProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OutputCacheProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewOutputCacheProfileSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OutputCacheProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OutputCacheProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetOutputCacheProfiles sets the value of OutputCacheProfiles for the instance
func (instance *OutputCacheProfileSettings) SetPropertyOutputCacheProfiles(value []OutputCacheProfile) (err error) {
	return instance.SetProperty("OutputCacheProfiles", (value))
}

// GetOutputCacheProfiles gets the value of OutputCacheProfiles for the instance
func (instance *OutputCacheProfileSettings) GetPropertyOutputCacheProfiles() (value []OutputCacheProfile, err error) {
	retValue, err := instance.GetProperty("OutputCacheProfiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(OutputCacheProfile)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " OutputCacheProfile is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, OutputCacheProfile(valuetmp))
	}

	return
}
