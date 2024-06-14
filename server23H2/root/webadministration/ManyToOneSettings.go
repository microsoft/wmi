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

// ManyToOneSettings struct
type ManyToOneSettings struct {
	*EmbeddedObject

	//
	ManyToOneMappings []ManyToOneCertificateMappingElement
}

func NewManyToOneSettingsEx1(instance *cim.WmiInstance) (newInstance *ManyToOneSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ManyToOneSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewManyToOneSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ManyToOneSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ManyToOneSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetManyToOneMappings sets the value of ManyToOneMappings for the instance
func (instance *ManyToOneSettings) SetPropertyManyToOneMappings(value []ManyToOneCertificateMappingElement) (err error) {
	return instance.SetProperty("ManyToOneMappings", (value))
}

// GetManyToOneMappings gets the value of ManyToOneMappings for the instance
func (instance *ManyToOneSettings) GetPropertyManyToOneMappings() (value []ManyToOneCertificateMappingElement, err error) {
	retValue, err := instance.GetProperty("ManyToOneMappings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ManyToOneCertificateMappingElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ManyToOneCertificateMappingElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ManyToOneCertificateMappingElement(valuetmp))
	}

	return
}
