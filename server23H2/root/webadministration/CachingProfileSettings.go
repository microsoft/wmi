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

// CachingProfileSettings struct
type CachingProfileSettings struct {
	*EmbeddedObject

	//
	Profiles []CachingProfileElement
}

func NewCachingProfileSettingsEx1(instance *cim.WmiInstance) (newInstance *CachingProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CachingProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewCachingProfileSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CachingProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CachingProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetProfiles sets the value of Profiles for the instance
func (instance *CachingProfileSettings) SetPropertyProfiles(value []CachingProfileElement) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *CachingProfileSettings) GetPropertyProfiles() (value []CachingProfileElement, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(CachingProfileElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " CachingProfileElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, CachingProfileElement(valuetmp))
	}

	return
}
