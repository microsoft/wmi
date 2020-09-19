// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcNetworkSettingTask struct
type MSFT_DtcNetworkSettingTask struct {
	*cim.WmiInstance
}

func NewMSFT_DtcNetworkSettingTaskEx1(instance *cim.WmiInstance) (newInstance *MSFT_DtcNetworkSettingTask, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcNetworkSettingTask{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DtcNetworkSettingTaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DtcNetworkSettingTask, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DtcNetworkSettingTask{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcNetworkSettings "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcNetworkSettingTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput DtcNetworkSettings) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DisableNetworkAccess" type="bool "></param>
// <param name="DtcName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcNetworkSettingTask) SetByDisableNetwork( /* IN */ DtcName string,
	/* IN */ DisableNetworkAccess bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByDisableNetwork", DtcName, DisableNetworkAccess)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="AuthenticationLevel" type="string "></param>
// <param name="DtcName" type="string "></param>
// <param name="InboundTransactionsEnabled" type="bool "></param>
// <param name="LUTransactionsEnabled" type="bool "></param>
// <param name="OutboundTransactionsEnabled" type="bool "></param>
// <param name="RemoteAdministrationAccessEnabled" type="bool "></param>
// <param name="RemoteClientAccessEnabled" type="bool "></param>
// <param name="XATransactionsEnabled" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcNetworkSettingTask) SetByNetworkSettings( /* IN */ DtcName string,
	/* IN */ InboundTransactionsEnabled bool,
	/* IN */ OutboundTransactionsEnabled bool,
	/* IN */ RemoteClientAccessEnabled bool,
	/* IN */ RemoteAdministrationAccessEnabled bool,
	/* IN */ XATransactionsEnabled bool,
	/* IN */ LUTransactionsEnabled bool,
	/* IN */ AuthenticationLevel string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetByNetworkSettings", DtcName, InboundTransactionsEnabled, OutboundTransactionsEnabled, RemoteClientAccessEnabled, RemoteAdministrationAccessEnabled, XATransactionsEnabled, LUTransactionsEnabled, AuthenticationLevel)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
