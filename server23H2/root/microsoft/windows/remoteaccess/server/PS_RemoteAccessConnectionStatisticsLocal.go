// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PS_RemoteAccessConnectionStatisticsLocal struct
type PS_RemoteAccessConnectionStatisticsLocal struct {
	*cim.WmiInstance
}

func NewPS_RemoteAccessConnectionStatisticsLocalEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessConnectionStatisticsLocal, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessConnectionStatisticsLocal{
		WmiInstance: tmp,
	}
	return
}

func NewPS_RemoteAccessConnectionStatisticsLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *PS_RemoteAccessConnectionStatisticsLocal, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &PS_RemoteAccessConnectionStatisticsLocal{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccountingConnectionLocal []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatisticsLocal) GetByAccountingStatistics( /* IN */ StartDateTime string,
	/* IN */ EndDateTime string,
	/* OUT */ cmdletOutput []RemoteAccessAccountingConnectionLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByAccountingStatistics", StartDateTime, EndDateTime)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ResourceName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessMonitoringConnectionLocal []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatisticsLocal) GetByActiveStatistics( /* IN */ ResourceName string,
	/* OUT */ cmdletOutput []RemoteAccessMonitoringConnectionLocal) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetByActiveStatistics", ResourceName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
