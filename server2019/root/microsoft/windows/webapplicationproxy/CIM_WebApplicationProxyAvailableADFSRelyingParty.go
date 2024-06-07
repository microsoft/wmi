// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.WebApplicationProxy
//////////////////////////////////////////////
package webapplicationproxy
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
)

// CIM_WebApplicationProxyAvailableADFSRelyingParty struct
type CIM_WebApplicationProxyAvailableADFSRelyingParty struct { 
	*cim.WmiInstance
}

	func NewCIM_WebApplicationProxyAvailableADFSRelyingPartyEx1(instance *cim.WmiInstance) (newInstance *CIM_WebApplicationProxyAvailableADFSRelyingParty, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &CIM_WebApplicationProxyAvailableADFSRelyingParty {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewCIM_WebApplicationProxyAvailableADFSRelyingPartyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_WebApplicationProxyAvailableADFSRelyingParty, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_WebApplicationProxyAvailableADFSRelyingParty {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="cmdletOutput" type="RelyingPartyMetadata []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_WebApplicationProxyAvailableADFSRelyingParty) Get( /* OUT */ cmdletOutput []RelyingPartyMetadata) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


