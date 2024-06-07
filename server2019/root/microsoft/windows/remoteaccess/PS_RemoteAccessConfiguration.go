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

// PS_RemoteAccessConfiguration struct
type PS_RemoteAccessConfiguration struct { 
	*cim.WmiInstance
}

	func NewPS_RemoteAccessConfigurationEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessConfiguration, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_RemoteAccessConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_RemoteAccessConfiguration, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessConfiguration {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="cmdletOutput" type="RemoteAccessConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConfiguration) Get( /* OUT */ cmdletOutput RemoteAccessConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Configuration" type="RemoteAccessConfiguration "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RestartServiceIfRequired" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessConfiguration) Set( /* IN */ Configuration RemoteAccessConfiguration,
 /* IN */ RestartServiceIfRequired bool,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput RemoteAccessConfiguration) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , Configuration, RestartServiceIfRequired, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


