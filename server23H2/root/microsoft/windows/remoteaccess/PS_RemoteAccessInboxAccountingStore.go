// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessInboxAccountingStore struct
type PS_RemoteAccessInboxAccountingStore struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessInboxAccountingStoreEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessInboxAccountingStore, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessInboxAccountingStore{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessInboxAccountingStoreEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessInboxAccountingStore, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessInboxAccountingStore{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="StoreLimit" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessInboxAccounting "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessInboxAccountingStore) Set( /* IN */ StoreLimit string,
	/* IN */ ComputerName string,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput RemoteAccessInboxAccounting) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set", StoreLimit, ComputerName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>
// <param name="EndDateTime" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessInboxAccounting "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessInboxAccountingStore) Clear( /* IN */ ComputerName string,
	/* IN */ StartDateTime string,
	/* IN */ EndDateTime string,
	/* IN */ Force bool,
	/* IN */ PassThru bool,
	/* OUT */ cmdletOutput RemoteAccessInboxAccounting) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Clear", ComputerName, StartDateTime, EndDateTime, Force, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
