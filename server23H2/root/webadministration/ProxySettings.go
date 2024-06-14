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

// ProxySettings struct
type ProxySettings struct {
	*EmbeddedObject

	//
	AutoDetect int32

	//
	BypassOnLocal int32

	//
	ProxyAddress string

	//
	ScriptLocation string

	//
	UseSystemDefault int32
}

func NewProxySettingsEx1(instance *cim.WmiInstance) (newInstance *ProxySettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ProxySettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewProxySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProxySettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProxySettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAutoDetect sets the value of AutoDetect for the instance
func (instance *ProxySettings) SetPropertyAutoDetect(value int32) (err error) {
	return instance.SetProperty("AutoDetect", (value))
}

// GetAutoDetect gets the value of AutoDetect for the instance
func (instance *ProxySettings) GetPropertyAutoDetect() (value int32, err error) {
	retValue, err := instance.GetProperty("AutoDetect")
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

// SetBypassOnLocal sets the value of BypassOnLocal for the instance
func (instance *ProxySettings) SetPropertyBypassOnLocal(value int32) (err error) {
	return instance.SetProperty("BypassOnLocal", (value))
}

// GetBypassOnLocal gets the value of BypassOnLocal for the instance
func (instance *ProxySettings) GetPropertyBypassOnLocal() (value int32, err error) {
	retValue, err := instance.GetProperty("BypassOnLocal")
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

// SetProxyAddress sets the value of ProxyAddress for the instance
func (instance *ProxySettings) SetPropertyProxyAddress(value string) (err error) {
	return instance.SetProperty("ProxyAddress", (value))
}

// GetProxyAddress gets the value of ProxyAddress for the instance
func (instance *ProxySettings) GetPropertyProxyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyAddress")
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

// SetScriptLocation sets the value of ScriptLocation for the instance
func (instance *ProxySettings) SetPropertyScriptLocation(value string) (err error) {
	return instance.SetProperty("ScriptLocation", (value))
}

// GetScriptLocation gets the value of ScriptLocation for the instance
func (instance *ProxySettings) GetPropertyScriptLocation() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptLocation")
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

// SetUseSystemDefault sets the value of UseSystemDefault for the instance
func (instance *ProxySettings) SetPropertyUseSystemDefault(value int32) (err error) {
	return instance.SetProperty("UseSystemDefault", (value))
}

// GetUseSystemDefault gets the value of UseSystemDefault for the instance
func (instance *ProxySettings) GetPropertyUseSystemDefault() (value int32, err error) {
	retValue, err := instance.GetProperty("UseSystemDefault")
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
