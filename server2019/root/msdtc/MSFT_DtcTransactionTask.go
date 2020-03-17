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

// MSFT_DtcTransactionTask struct
type MSFT_DtcTransactionTask struct {
	cim.WmiInstance
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
