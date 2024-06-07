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

// PS_DAAppServerConnection struct
type PS_DAAppServerConnection struct { 
	*cim.WmiInstance
}

	func NewPS_DAAppServerConnectionEx1(instance *cim.WmiInstance) (newInstance *PS_DAAppServerConnection, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_DAAppServerConnection {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_DAAppServerConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_DAAppServerConnection, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_DAAppServerConnection {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="ConnectionType" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="TrafficProtection" type="string "></param>

// <param name="cmdletOutput" type="DAAppServerConnection "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DAAppServerConnection) Set( /* IN */ ConnectionType string,
 /* IN */ TrafficProtection string,
 /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput DAAppServerConnection) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , ConnectionType, TrafficProtection, ComputerName, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


