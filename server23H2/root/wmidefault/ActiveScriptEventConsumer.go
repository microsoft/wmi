// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.DEFAULT
//////////////////////////////////////////////
package wmidefault

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ActiveScriptEventConsumer struct
type ActiveScriptEventConsumer struct {
	*__EventConsumer

	//
	KillTimeout uint32

	//
	Name string

	//
	ScriptFilename string

	//
	ScriptingEngine string

	//
	ScriptText string
}

func NewActiveScriptEventConsumerEx1(instance *cim.WmiInstance) (newInstance *ActiveScriptEventConsumer, err error) {
	tmp, err := New__EventConsumerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ActiveScriptEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

func NewActiveScriptEventConsumerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ActiveScriptEventConsumer, err error) {
	tmp, err := New__EventConsumerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ActiveScriptEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

// SetKillTimeout sets the value of KillTimeout for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyKillTimeout(value uint32) (err error) {
	return instance.SetProperty("KillTimeout", (value))
}

// GetKillTimeout gets the value of KillTimeout for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyKillTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("KillTimeout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyName() (value string, err error) {
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

// SetScriptFilename sets the value of ScriptFilename for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptFilename(value string) (err error) {
	return instance.SetProperty("ScriptFilename", (value))
}

// GetScriptFilename gets the value of ScriptFilename for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptFilename() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptFilename")
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

// SetScriptingEngine sets the value of ScriptingEngine for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptingEngine(value string) (err error) {
	return instance.SetProperty("ScriptingEngine", (value))
}

// GetScriptingEngine gets the value of ScriptingEngine for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptingEngine() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptingEngine")
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

// SetScriptText sets the value of ScriptText for the instance
func (instance *ActiveScriptEventConsumer) SetPropertyScriptText(value string) (err error) {
	return instance.SetProperty("ScriptText", (value))
}

// GetScriptText gets the value of ScriptText for the instance
func (instance *ActiveScriptEventConsumer) GetPropertyScriptText() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptText")
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
