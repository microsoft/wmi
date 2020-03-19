// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_StorageFaultEvent struct
type WSP_StorageFaultEvent struct {
	*MSFT_StorageFaultEvent
}

func NewWSP_StorageFaultEventEx1(instance *cim.WmiInstance) (newInstance *WSP_StorageFaultEvent, err error) {
	tmp, err := NewMSFT_StorageFaultEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageFaultEvent{
		MSFT_StorageFaultEvent: tmp,
	}
	return
}

func NewWSP_StorageFaultEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_StorageFaultEvent, err error) {
	tmp, err := NewMSFT_StorageFaultEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_StorageFaultEvent{
		MSFT_StorageFaultEvent: tmp,
	}
	return
}

// This method manually fires alerts

// <param name="FaultAlert" type="WSP_StorageFaultEvent ">Copy of the alert payload to be fired</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_StorageFaultEvent) FireAlert( /* IN */ FaultAlert WSP_StorageFaultEvent) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FireAlert", FaultAlert)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
