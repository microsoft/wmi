// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.msdtc
//
// ////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTransactionTask struct
type MSFT_DtcTransactionTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcTransactionTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcTransactionTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcTransactionTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcTransactionTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcTransactionTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcTransactionInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput []DtcTransactionInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Abort" type="bool "></param>
// <param name="DtcName" type="string "></param>
// <param name="TransactionId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionTask) SetByAbortSet( /* IN */ DtcName string,
	/* IN */ TransactionId string,
	/* IN */ Abort bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByAbortSet", DtcName, TransactionId, Abort)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Commit" type="bool "></param>
// <param name="DtcName" type="string "></param>
// <param name="TransactionId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionTask) SetByCommitSet( /* IN */ DtcName string,
	/* IN */ TransactionId string,
	/* IN */ Commit bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByCommitSet", DtcName, TransactionId, Commit)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DtcName" type="string "></param>
// <param name="Forget" type="bool "></param>
// <param name="TransactionId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionTask) SetByForgetSet( /* IN */ DtcName string,
	/* IN */ TransactionId string,
	/* IN */ Forget bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByForgetSet", DtcName, TransactionId, Forget)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DtcName" type="string "></param>
// <param name="Trace" type="bool "></param>
// <param name="TransactionId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionTask) SetByTraceSet( /* IN */ DtcName string,
	/* IN */ TransactionId string,
	/* IN */ Trace bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByTraceSet", DtcName, TransactionId, Trace)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
