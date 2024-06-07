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

// PS_DANetworkLocationServer struct
type PS_DANetworkLocationServer struct { 
	*cim.WmiInstance
}

	func NewPS_DANetworkLocationServerEx1(instance *cim.WmiInstance) (newInstance *PS_DANetworkLocationServer, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_DANetworkLocationServer {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_DANetworkLocationServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_DANetworkLocationServer, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_DANetworkLocationServer {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="CheckReachability" type="bool "></param>
// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Url" type="string "></param>

// <param name="cmdletOutput" type="DANetworkLocationServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DANetworkLocationServer) SetByExternalServer( /* IN */ ComputerName string,
 /* IN */ Force bool,
 /* IN */ PassThru bool,
 /* IN */ Url string,
 /* IN */ CheckReachability bool,
 /* OUT */ cmdletOutput DANetworkLocationServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByExternalServer" , ComputerName, Force, PassThru, Url, CheckReachability)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Certificate" type="uint8 []"></param>
// <param name="ComputerName" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="NlsOnDAServer" type="bool "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="DANetworkLocationServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DANetworkLocationServer) SetByOnDAServer( /* IN */ ComputerName string,
 /* IN */ Force bool,
 /* IN */ PassThru bool,
 /* IN */ Certificate []uint8,
 /* IN */ NlsOnDAServer bool,
 /* OUT */ cmdletOutput DANetworkLocationServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByOnDAServer" , ComputerName, Force, PassThru, Certificate, NlsOnDAServer)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="CheckReachability" type="bool "></param>
// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="DANetworkLocationServer "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_DANetworkLocationServer) Get( /* IN */ CheckReachability bool,
 /* IN */ ComputerName string,
 /* OUT */ cmdletOutput DANetworkLocationServer) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , CheckReachability, ComputerName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


