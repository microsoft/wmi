// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// WSP_HealthActionEvent struct
type WSP_HealthActionEvent struct {
	*MSFT_HealthActionEvent
}

func NewWSP_HealthActionEventEx1(instance *cim.WmiInstance) (newInstance *WSP_HealthActionEvent, err error) {
	tmp, err := NewMSFT_HealthActionEventEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WSP_HealthActionEvent{
		MSFT_HealthActionEvent: tmp,
	}
	return
}

func NewWSP_HealthActionEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WSP_HealthActionEvent, err error) {
	tmp, err := NewMSFT_HealthActionEventEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WSP_HealthActionEvent{
		MSFT_HealthActionEvent: tmp,
	}
	return
}

// This method manually fires alerts

// <param name="HealthActionAlert" type="WSP_HealthActionEvent ">Copy of the alert payload to be fired</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WSP_HealthActionEvent) FireAlert( /* IN */ HealthActionAlert WSP_HealthActionEvent) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FireAlert", HealthActionAlert)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
