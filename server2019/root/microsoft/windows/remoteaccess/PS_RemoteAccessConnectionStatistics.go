// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_RemoteAccessConnectionStatistics struct
type PS_RemoteAccessConnectionStatistics struct { 
	*cim.WmiInstance
}

	func NewPS_RemoteAccessConnectionStatisticsEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessConnectionStatistics, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessConnectionStatistics {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_RemoteAccessConnectionStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_RemoteAccessConnectionStatistics, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessConnectionStatistics {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccountingConnection []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatistics) GetByAccountingStatistics( /* IN */ ComputerName string,
 /* IN */ StartDateTime string,
 /* IN */ EndDateTime string,
 /* OUT */ cmdletOutput []RemoteAccessAccountingConnection) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByAccountingStatistics" , ComputerName, StartDateTime, EndDateTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="ResourceName" type="string "></param>
// <param name="RoutingDomain" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessMonitoringConnection []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConnectionStatistics) GetByActiveStatistics( /* IN */ ComputerName string,
 /* IN */ ResourceName string,
 /* IN */ RoutingDomain string,
 /* OUT */ cmdletOutput []RemoteAccessMonitoringConnection) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByActiveStatistics" , ComputerName, ResourceName, RoutingDomain)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


