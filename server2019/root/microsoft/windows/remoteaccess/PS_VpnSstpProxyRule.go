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

// PS_VpnSstpProxyRule struct
type PS_VpnSstpProxyRule struct { 
	*cim.WmiInstance
}

	func NewPS_VpnSstpProxyRuleEx1(instance *cim.WmiInstance) (newInstance *PS_VpnSstpProxyRule, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_VpnSstpProxyRule {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_VpnSstpProxyRuleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_VpnSstpProxyRule, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_VpnSstpProxyRule {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="PassThru" type="bool "></param>
// <param name="Rule" type="VpnSstpProxyRule []"></param>

// <param name="cmdletOutput" type="VpnSstpProxyRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnSstpProxyRule) Add( /* IN */ Rule []VpnSstpProxyRule,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput []VpnSstpProxyRule) (result uint32, err error) {retVal, err := instance.InvokeMethod("Add" , Rule, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="GatewayAddress" type="string []"></param>
// <param name="TenantID" type="string "></param>

// <param name="cmdletOutput" type="VpnSstpProxyRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnSstpProxyRule) New( /* IN */ TenantID string,
 /* IN */ GatewayAddress []string,
 /* OUT */ cmdletOutput VpnSstpProxyRule) (result uint32, err error) {retVal, err := instance.InvokeMethod("New" , TenantID, GatewayAddress)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="GatewayAddress" type="string []"></param>
// <param name="PassThru" type="bool "></param>
// <param name="TenantID" type="string "></param>

// <param name="cmdletOutput" type="VpnSstpProxyRule "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnSstpProxyRule) Set( /* IN */ TenantID string,
 /* IN */ GatewayAddress []string,
 /* IN */ PassThru bool,
 /* OUT */ cmdletOutput VpnSstpProxyRule) (result uint32, err error) {retVal, err := instance.InvokeMethod("Set" , TenantID, GatewayAddress, PassThru)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="cmdletOutput" type="VpnSstpProxyRule []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnSstpProxyRule) Get( /* OUT */ cmdletOutput []VpnSstpProxyRule) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Force" type="bool "></param>
// <param name="TenantID" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_VpnSstpProxyRule) Remove( /* IN */ TenantID []string,
 /* IN */ Force bool) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Remove" , TenantID, Force);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


