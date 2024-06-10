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

// IsapiFilterElement struct
type IsapiFilterElement struct {
	*CollectionElement

	//
	EnableCache bool

	//
	Enabled bool

	//
	Name string

	//
	Path string

	//
	PreCondition string
}

func NewIsapiFilterElementEx1(instance *cim.WmiInstance) (newInstance *IsapiFilterElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &IsapiFilterElement{
		CollectionElement: tmp,
	}
	return
}

func NewIsapiFilterElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *IsapiFilterElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &IsapiFilterElement{
		CollectionElement: tmp,
	}
	return
}

// SetEnableCache sets the value of EnableCache for the instance
func (instance *IsapiFilterElement) SetPropertyEnableCache(value bool) (err error) {
	return instance.SetProperty("EnableCache", (value))
}

// GetEnableCache gets the value of EnableCache for the instance
func (instance *IsapiFilterElement) GetPropertyEnableCache() (value bool, err error) {
	retValue, err := instance.GetProperty("EnableCache")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *IsapiFilterElement) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *IsapiFilterElement) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetName sets the value of Name for the instance
func (instance *IsapiFilterElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *IsapiFilterElement) GetPropertyName() (value string, err error) {
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

// SetPath sets the value of Path for the instance
func (instance *IsapiFilterElement) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *IsapiFilterElement) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPreCondition sets the value of PreCondition for the instance
func (instance *IsapiFilterElement) SetPropertyPreCondition(value string) (err error) {
	return instance.SetProperty("PreCondition", (value))
}

// GetPreCondition gets the value of PreCondition for the instance
func (instance *IsapiFilterElement) GetPropertyPreCondition() (value string, err error) {
	retValue, err := instance.GetProperty("PreCondition")
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
