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

// SourceElement struct
type SourceElement struct {
	*CollectionElement

	//
	Listeners SourceListenerSettings

	//
	Name string

	//
	SwitchName string

	//
	SwitchType string

	//
	SwitchValue string
}

func NewSourceElementEx1(instance *cim.WmiInstance) (newInstance *SourceElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SourceElement{
		CollectionElement: tmp,
	}
	return
}

func NewSourceElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SourceElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SourceElement{
		CollectionElement: tmp,
	}
	return
}

// SetListeners sets the value of Listeners for the instance
func (instance *SourceElement) SetPropertyListeners(value SourceListenerSettings) (err error) {
	return instance.SetProperty("Listeners", (value))
}

// GetListeners gets the value of Listeners for the instance
func (instance *SourceElement) GetPropertyListeners() (value SourceListenerSettings, err error) {
	retValue, err := instance.GetProperty("Listeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(SourceListenerSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " SourceListenerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = SourceListenerSettings(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *SourceElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SourceElement) GetPropertyName() (value string, err error) {
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

// SetSwitchName sets the value of SwitchName for the instance
func (instance *SourceElement) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *SourceElement) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
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

// SetSwitchType sets the value of SwitchType for the instance
func (instance *SourceElement) SetPropertySwitchType(value string) (err error) {
	return instance.SetProperty("SwitchType", (value))
}

// GetSwitchType gets the value of SwitchType for the instance
func (instance *SourceElement) GetPropertySwitchType() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchType")
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

// SetSwitchValue sets the value of SwitchValue for the instance
func (instance *SourceElement) SetPropertySwitchValue(value string) (err error) {
	return instance.SetProperty("SwitchValue", (value))
}

// GetSwitchValue gets the value of SwitchValue for the instance
func (instance *SourceElement) GetPropertySwitchValue() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchValue")
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
