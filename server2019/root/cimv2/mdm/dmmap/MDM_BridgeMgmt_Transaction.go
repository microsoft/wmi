// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_BridgeMgmt_Transaction struct
type MDM_BridgeMgmt_Transaction struct {
	cim.WmiInstance
}

//

// <param name="requestId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_BridgeMgmt_Transaction) TransactionBegin( /* IN */ requestId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TransactionBegin", requestId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="requestId" type="string "></param>

// <param name="failedRequestDetails" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="transactionMiResult" type="uint32 "></param>
func (instance *MDM_BridgeMgmt_Transaction) TransactionEnd( /* IN */ requestId string,
	/* OUT */ transactionMiResult uint32,
	/* OUT */ failedRequestDetails string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("TransactionEnd", requestId)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
