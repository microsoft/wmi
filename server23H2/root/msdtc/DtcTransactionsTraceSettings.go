// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DtcTransactionsTraceSettings struct
type DtcTransactionsTraceSettings struct {
	*cim.WmiInstance

	//
	AbortedTransactionsTracingEnabled bool

	//
	AllTransactionsTracingEnabled bool

	//
	LongLivedTransactionsTracingEnabled bool
}

func NewDtcTransactionsTraceSettingsEx1(instance *cim.WmiInstance) (newInstance *DtcTransactionsTraceSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsTraceSettings{
		WmiInstance: tmp,
	}
	return
}

func NewDtcTransactionsTraceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcTransactionsTraceSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsTraceSettings{
		WmiInstance: tmp,
	}
	return
}

// SetAbortedTransactionsTracingEnabled sets the value of AbortedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyAbortedTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("AbortedTransactionsTracingEnabled", (value))
}

// GetAbortedTransactionsTracingEnabled gets the value of AbortedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyAbortedTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AbortedTransactionsTracingEnabled")
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

// SetAllTransactionsTracingEnabled sets the value of AllTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyAllTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("AllTransactionsTracingEnabled", (value))
}

// GetAllTransactionsTracingEnabled gets the value of AllTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyAllTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AllTransactionsTracingEnabled")
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

// SetLongLivedTransactionsTracingEnabled sets the value of LongLivedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyLongLivedTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("LongLivedTransactionsTracingEnabled", (value))
}

// GetLongLivedTransactionsTracingEnabled gets the value of LongLivedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyLongLivedTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("LongLivedTransactionsTracingEnabled")
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
