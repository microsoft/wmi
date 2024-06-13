// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// XmlSerializerOperationBehavior struct
type XmlSerializerOperationBehavior struct {
	*Behavior

	// Defines the style of the SOAP message.
	Style string

	// Specifies the SOAP encoding style.
	Use string
}

func NewXmlSerializerOperationBehaviorEx1(instance *cim.WmiInstance) (newInstance *XmlSerializerOperationBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &XmlSerializerOperationBehavior{
		Behavior: tmp,
	}
	return
}

func NewXmlSerializerOperationBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *XmlSerializerOperationBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &XmlSerializerOperationBehavior{
		Behavior: tmp,
	}
	return
}

// SetStyle sets the value of Style for the instance
func (instance *XmlSerializerOperationBehavior) SetPropertyStyle(value string) (err error) {
	return instance.SetProperty("Style", (value))
}

// GetStyle gets the value of Style for the instance
func (instance *XmlSerializerOperationBehavior) GetPropertyStyle() (value string, err error) {
	retValue, err := instance.GetProperty("Style")
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

// SetUse sets the value of Use for the instance
func (instance *XmlSerializerOperationBehavior) SetPropertyUse(value string) (err error) {
	return instance.SetProperty("Use", (value))
}

// GetUse gets the value of Use for the instance
func (instance *XmlSerializerOperationBehavior) GetPropertyUse() (value string, err error) {
	retValue, err := instance.GetProperty("Use")
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
