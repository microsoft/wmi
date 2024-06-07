// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_RemoteAccessAccountingStatisticsSummary struct
type PS_RemoteAccessAccountingStatisticsSummary struct { 
	*cim.WmiInstance
}

	func NewPS_RemoteAccessAccountingStatisticsSummaryEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessAccountingStatisticsSummary, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessAccountingStatisticsSummary {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_RemoteAccessAccountingStatisticsSummaryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_RemoteAccessAccountingStatisticsSummary, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessAccountingStatisticsSummary {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccountingConnectionSummaryLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessAccountingStatisticsSummary) GetByAccountingStatistics( /* IN */ StartDateTime string,
 /* IN */ EndDateTime string,
 /* OUT */ cmdletOutput RemoteAccessAccountingConnectionSummaryLocal) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByAccountingStatistics" , StartDateTime, EndDateTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="cmdletOutput" type="RemoteAccessActiveConnectionSummaryLocal "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessAccountingStatisticsSummary) GetByActiveStatistics( /* OUT */ cmdletOutput RemoteAccessActiveConnectionSummaryLocal) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByActiveStatistics" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


