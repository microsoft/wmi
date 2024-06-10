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

// PS_RemoteAccessConnectionStatisticsSummary struct
type PS_RemoteAccessConnectionStatisticsSummary struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessConnectionStatisticsSummaryEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessConnectionStatisticsSummary, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessConnectionStatisticsSummary{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessConnectionStatisticsSummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessConnectionStatisticsSummary, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessConnectionStatisticsSummary{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ComputerName" type="string "></param>
// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccountingSummary "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatisticsSummary) GetByAccountingStatistics( /* IN */ StartDateTime string,
	/* IN */ EndDateTime string,
	/* IN */ ComputerName string,
	/* OUT */ cmdletOutput RemoteAccessAccountingSummary) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByAccountingStatistics", StartDateTime, EndDateTime, ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessMonitoringSummary "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatisticsSummary) GetByActiveStatistics( /* IN */ ComputerName string,
	/* OUT */ cmdletOutput RemoteAccessMonitoringSummary) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByActiveStatistics", ComputerName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
