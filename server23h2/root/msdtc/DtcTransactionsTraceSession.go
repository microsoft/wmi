// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// DtcTransactionsTraceSession struct
type DtcTransactionsTraceSession struct {
	*cim.WmiInstance

	//
	BufferCount uint32

	//
	SessionStatus string
}

func NewDtcTransactionsTraceSessionEx1(instance *cim.WmiInstance) (newInstance *DtcTransactionsTraceSession, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsTraceSession{
		WmiInstance: tmp,
	}
	return
}

func NewDtcTransactionsTraceSessionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcTransactionsTraceSession, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsTraceSession{
		WmiInstance: tmp,
	}
	return
}

// SetBufferCount sets the value of BufferCount for the instance
func (instance *DtcTransactionsTraceSession) SetPropertyBufferCount(value uint32) (err error) {
	return instance.SetProperty("BufferCount", (value))
}

// GetBufferCount gets the value of BufferCount for the instance
func (instance *DtcTransactionsTraceSession) GetPropertyBufferCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferCount")
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

// SetSessionStatus sets the value of SessionStatus for the instance
func (instance *DtcTransactionsTraceSession) SetPropertySessionStatus(value string) (err error) {
	return instance.SetProperty("SessionStatus", (value))
}

// GetSessionStatus gets the value of SessionStatus for the instance
func (instance *DtcTransactionsTraceSession) GetPropertySessionStatus() (value string, err error) {
	retValue, err := instance.GetProperty("SessionStatus")
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
