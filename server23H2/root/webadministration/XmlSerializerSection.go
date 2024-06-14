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

// XmlSerializerSection struct
type XmlSerializerSection struct {
	*ConfigurationSection

	//
	CheckDeserializeAdvances bool
}

func NewXmlSerializerSectionEx1(instance *cim.WmiInstance) (newInstance *XmlSerializerSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &XmlSerializerSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewXmlSerializerSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *XmlSerializerSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &XmlSerializerSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCheckDeserializeAdvances sets the value of CheckDeserializeAdvances for the instance
func (instance *XmlSerializerSection) SetPropertyCheckDeserializeAdvances(value bool) (err error) {
	return instance.SetProperty("CheckDeserializeAdvances", (value))
}

// GetCheckDeserializeAdvances gets the value of CheckDeserializeAdvances for the instance
func (instance *XmlSerializerSection) GetPropertyCheckDeserializeAdvances() (value bool, err error) {
	retValue, err := instance.GetProperty("CheckDeserializeAdvances")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
