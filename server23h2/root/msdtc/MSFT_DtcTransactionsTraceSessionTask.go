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
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTransactionsTraceSessionTask struct
type MSFT_DtcTransactionsTraceSessionTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcTransactionsTraceSessionTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcTransactionsTraceSessionTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsTraceSessionTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcTransactionsTraceSessionTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcTransactionsTraceSessionTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsTraceSessionTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSessionTask) Stop() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Stop")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSessionTask) Write() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Write")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSessionTask) Start() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Start")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="cmdletOutput" type="DtcTransactionsTraceSession "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSessionTask) Get( /* OUT */ cmdletOutput DtcTransactionsTraceSession) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BufferCount" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSessionTask) Set( /* IN */ BufferCount uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set", BufferCount)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
