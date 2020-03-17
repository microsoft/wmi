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

// DtcTransactionsTraceSession struct
type DtcTransactionsTraceSession struct {
	cim.WmiInstance

	//
	BufferCount uint32

	//
	SessionStatus string
}

// SetBufferCount sets the value of BufferCount for the instance
func (instance *DtcTransactionsTraceSession) SetPropertyBufferCount(value uint32) (err error) {
	return instance.SetProperty("BufferCount", value)
}

// GetBufferCount gets the value of BufferCount for the instance
func (instance *DtcTransactionsTraceSession) GetPropertyBufferCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSessionStatus sets the value of SessionStatus for the instance
func (instance *DtcTransactionsTraceSession) SetPropertySessionStatus(value string) (err error) {
	return instance.SetProperty("SessionStatus", value)
}

// GetSessionStatus gets the value of SessionStatus for the instance
func (instance *DtcTransactionsTraceSession) GetPropertySessionStatus() (value string, err error) {
	retValue, err := instance.GetProperty("SessionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
