// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// PS_VpnConnectionTrigger struct
type PS_VpnConnectionTrigger struct { 
	*cim.WmiInstance
}

	func NewPS_VpnConnectionTriggerEx1(instance *cim.WmiInstance) (newInstance *PS_VpnConnectionTrigger, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_VpnConnectionTrigger {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_VpnConnectionTriggerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_VpnConnectionTrigger, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_VpnConnectionTrigger {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ConnectionName" type="string "></param>

// <param name="cmdletOutput" type="VpnConnectionTrigger "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnConnectionTrigger) Get( /* IN */ ConnectionName string,
 /* OUT */ cmdletOutput VpnConnectionTrigger) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , ConnectionName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


