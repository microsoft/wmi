// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DtcTransactionsTraceSettings struct
type DtcTransactionsTraceSettings struct {
	cim.WmiInstance

	//
	AbortedTransactionsTracingEnabled bool

	//
	AllTransactionsTracingEnabled bool

	//
	LongLivedTransactionsTracingEnabled bool
}

// SetAbortedTransactionsTracingEnabled sets the value of AbortedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyAbortedTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("AbortedTransactionsTracingEnabled", value)
}

// GetAbortedTransactionsTracingEnabled gets the value of AbortedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyAbortedTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AbortedTransactionsTracingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllTransactionsTracingEnabled sets the value of AllTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyAllTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("AllTransactionsTracingEnabled", value)
}

// GetAllTransactionsTracingEnabled gets the value of AllTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyAllTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AllTransactionsTracingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLongLivedTransactionsTracingEnabled sets the value of LongLivedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) SetPropertyLongLivedTransactionsTracingEnabled(value bool) (err error) {
	return instance.SetProperty("LongLivedTransactionsTracingEnabled", value)
}

// GetLongLivedTransactionsTracingEnabled gets the value of LongLivedTransactionsTracingEnabled for the instance
func (instance *DtcTransactionsTraceSettings) GetPropertyLongLivedTransactionsTracingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("LongLivedTransactionsTracingEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
