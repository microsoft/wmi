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

// PS_VpnTrafficSelector struct
type PS_VpnTrafficSelector struct { 
	*cim.WmiInstance
}

	func NewPS_VpnTrafficSelectorEx1(instance *cim.WmiInstance) (newInstance *PS_VpnTrafficSelector, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_VpnTrafficSelector {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_VpnTrafficSelectorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_VpnTrafficSelector, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_VpnTrafficSelector {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="IPAddressRange" type="string []"></param>
// <param name="PortRange" type="uint32 []"></param>
// <param name="ProtocolId" type="uint32 "></param>
// <param name="TsPayloadId" type="uint16 "></param>
// <param name="Type" type="uint32 "></param>

// <param name="cmdletOutput" type="VpnTrafficSelector "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnTrafficSelector) New( /* IN */ IPAddressRange []string,
 /* IN */ PortRange []uint32,
 /* IN */ ProtocolId uint32,
 /* IN */ TsPayloadId uint16,
 /* IN */ Type uint32,
 /* OUT */ cmdletOutput VpnTrafficSelector) (result uint32, err error) {retVal, err := instance.InvokeMethod("New" , IPAddressRange, PortRange, ProtocolId, TsPayloadId, Type)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


