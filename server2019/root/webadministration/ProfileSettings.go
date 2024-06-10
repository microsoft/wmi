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

// ProfileSettings struct
type ProfileSettings struct {
	*EmbeddedObject

	//
	Profiles []ProfileElement
}

func NewProfileSettingsEx1(instance *cim.WmiInstance) (newInstance *ProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewProfileSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProfileSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProfileSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetProfiles sets the value of Profiles for the instance
func (instance *ProfileSettings) SetPropertyProfiles(value []ProfileElement) (err error) {
	return instance.SetProperty("Profiles", (value))
}

// GetProfiles gets the value of Profiles for the instance
func (instance *ProfileSettings) GetPropertyProfiles() (value []ProfileElement, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ProfileElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ProfileElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ProfileElement(valuetmp))
	}

	return
}
