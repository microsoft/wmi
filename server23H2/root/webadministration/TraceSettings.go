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

// TraceSettings struct
type TraceSettings struct {
	*EmbeddedObject

	//
	AutoFlush bool

	//
	IndentSize int32

	//
	Listeners TraceListenerSettings

	//
	UseGlobalLock bool
}

func NewTraceSettingsEx1(instance *cim.WmiInstance) (newInstance *TraceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TraceSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewTraceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TraceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TraceSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAutoFlush sets the value of AutoFlush for the instance
func (instance *TraceSettings) SetPropertyAutoFlush(value bool) (err error) {
	return instance.SetProperty("AutoFlush", (value))
}

// GetAutoFlush gets the value of AutoFlush for the instance
func (instance *TraceSettings) GetPropertyAutoFlush() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoFlush")
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

// SetIndentSize sets the value of IndentSize for the instance
func (instance *TraceSettings) SetPropertyIndentSize(value int32) (err error) {
	return instance.SetProperty("IndentSize", (value))
}

// GetIndentSize gets the value of IndentSize for the instance
func (instance *TraceSettings) GetPropertyIndentSize() (value int32, err error) {
	retValue, err := instance.GetProperty("IndentSize")
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

// SetListeners sets the value of Listeners for the instance
func (instance *TraceSettings) SetPropertyListeners(value TraceListenerSettings) (err error) {
	return instance.SetProperty("Listeners", (value))
}

// GetListeners gets the value of Listeners for the instance
func (instance *TraceSettings) GetPropertyListeners() (value TraceListenerSettings, err error) {
	retValue, err := instance.GetProperty("Listeners")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(TraceListenerSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " TraceListenerSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = TraceListenerSettings(valuetmp)

	return
}

// SetUseGlobalLock sets the value of UseGlobalLock for the instance
func (instance *TraceSettings) SetPropertyUseGlobalLock(value bool) (err error) {
	return instance.SetProperty("UseGlobalLock", (value))
}

// GetUseGlobalLock gets the value of UseGlobalLock for the instance
func (instance *TraceSettings) GetPropertyUseGlobalLock() (value bool, err error) {
	retValue, err := instance.GetProperty("UseGlobalLock")
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
