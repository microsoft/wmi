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

// PS_RemoteAccessUserActivity struct
type PS_RemoteAccessUserActivity struct { 
	*cim.WmiInstance
}

	func NewPS_RemoteAccessUserActivityEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessUserActivity, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessUserActivity {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_RemoteAccessUserActivityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_RemoteAccessUserActivity, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessUserActivity {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="EndDateTime" type="string "></param>
// <param name="HostIPAddress" type="string "></param>
// <param name="StartDateTime" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivity []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivity) GetByHostIP( /* IN */ ComputerName string,
 /* IN */ HostIPAddress string,
 /* IN */ EndDateTime string,
 /* IN */ StartDateTime string,
 /* OUT */ cmdletOutput []RemoteAccessUserActivity) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByHostIP" , ComputerName, HostIPAddress, EndDateTime, StartDateTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="SessionId" type="uint64 "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivity []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivity) GetBySessionId( /* IN */ ComputerName string,
 /* IN */ SessionId uint64,
 /* OUT */ cmdletOutput []RemoteAccessUserActivity) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetBySessionId" , ComputerName, SessionId)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>
// <param name="EndDateTime" type="string "></param>
// <param name="StartDateTime" type="string "></param>
// <param name="UserName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessUserActivity []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessUserActivity) GetByUserName( /* IN */ ComputerName string,
 /* IN */ UserName string,
 /* IN */ EndDateTime string,
 /* IN */ StartDateTime string,
 /* OUT */ cmdletOutput []RemoteAccessUserActivity) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetByUserName" , ComputerName, UserName, EndDateTime, StartDateTime)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


