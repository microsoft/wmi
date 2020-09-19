// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_BridgeMgmt_Transaction struct
type MDM_BridgeMgmt_Transaction struct {
	*cim.WmiInstance
}

func NewMDM_BridgeMgmt_TransactionEx1(instance *cim.WmiInstance) (newInstance *MDM_BridgeMgmt_Transaction, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_BridgeMgmt_Transaction{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_BridgeMgmt_TransactionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_BridgeMgmt_Transaction, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_BridgeMgmt_Transaction{
		WmiInstance: tmp,
	}
	return
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
