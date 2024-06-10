// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess.Server
//
// ////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessInboxAccountingLocal struct
type PS_RemoteAccessInboxAccountingLocal struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessInboxAccountingLocalEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessInboxAccountingLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessInboxAccountingLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessInboxAccountingLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessInboxAccountingLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessInboxAccountingLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="AccountingStatus" type="uint8 "></param>
// <param name="StoreLimit" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessInboxAccountingLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessInboxAccountingLocal) Set( /* IN */ AccountingStatus uint8,
	/* IN */ StoreLimit string,
	/* OUT */ cmdletOutput RemoteAccessInboxAccountingLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", AccountingStatus, StoreLimit)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="RemoteAccessInboxAccountingLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessInboxAccountingLocal) Get( /* OUT */ cmdletOutput RemoteAccessInboxAccountingLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessInboxAccountingLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessInboxAccountingLocal) CleanInboxAccountingStore( /* IN */ StartDateTime string,
	/* IN */ EndDateTime string,
	/* OUT */ cmdletOutput RemoteAccessInboxAccountingLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CleanInboxAccountingStore", StartDateTime, EndDateTime)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
