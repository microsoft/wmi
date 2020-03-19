// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTransactionsTraceSettingTask struct
type MSFT_DtcTransactionsTraceSettingTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcTransactionsTraceSettingTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcTransactionsTraceSettingTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsTraceSettingTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcTransactionsTraceSettingTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcTransactionsTraceSettingTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionsTraceSettingTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="cmdletOutput" type="DtcTransactionsTraceSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSettingTask) Get( /* OUT */ cmdletOutput DtcTransactionsTraceSettings) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AllTransactionsTracingEnabled" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSettingTask) SetByTraceAllParameterSet( /* IN */ AllTransactionsTracingEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByTraceAllParameterSet", AllTransactionsTracingEnabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AbortedTransactionsTracingEnabled" type="bool "></param>
// <param name="LongLivedTransactionsTracingEnabled" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsTraceSettingTask) SetByTraceSelectedParameterSet( /* IN */ AbortedTransactionsTracingEnabled bool,
	/* IN */ LongLivedTransactionsTracingEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByTraceSelectedParameterSet", AbortedTransactionsTracingEnabled, LongLivedTransactionsTracingEnabled)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
