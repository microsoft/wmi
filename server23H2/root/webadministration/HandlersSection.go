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

// HandlersSection struct
type HandlersSection struct {
	*ConfigurationSectionWithCollection

	//
	AccessPolicy int32

	//
	Handlers []HandlerAction
}

func NewHandlersSectionEx1(instance *cim.WmiInstance) (newInstance *HandlersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HandlersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewHandlersSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HandlersSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HandlersSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAccessPolicy sets the value of AccessPolicy for the instance
func (instance *HandlersSection) SetPropertyAccessPolicy(value int32) (err error) {
	return instance.SetProperty("AccessPolicy", (value))
}

// GetAccessPolicy gets the value of AccessPolicy for the instance
func (instance *HandlersSection) GetPropertyAccessPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("AccessPolicy")
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

// SetHandlers sets the value of Handlers for the instance
func (instance *HandlersSection) SetPropertyHandlers(value []HandlerAction) (err error) {
	return instance.SetProperty("Handlers", (value))
}

// GetHandlers gets the value of Handlers for the instance
func (instance *HandlersSection) GetPropertyHandlers() (value []HandlerAction, err error) {
	retValue, err := instance.GetProperty("Handlers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(HandlerAction)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " HandlerAction is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, HandlerAction(valuetmp))
	}

	return
}
