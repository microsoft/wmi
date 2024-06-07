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

// PS_RemoteAccessAccounting struct
type PS_RemoteAccessAccounting struct { 
	*cim.WmiInstance
}

	func NewPS_RemoteAccessAccountingEx1(instance *cim.WmiInstance) (newInstance *PS_RemoteAccessAccounting, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessAccounting {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPS_RemoteAccessAccountingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PS_RemoteAccessAccounting, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PS_RemoteAccessAccounting {
	WmiInstance: tmp,
	}
	return
	}
	

// 

// <param name="ComputerName" type="string "></param>
// <param name="DisableAccountingType" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="cmdletOutput" type="RemoteAccessAccounting "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessAccounting) SetByDisableAccounting( /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* IN */ DisableAccountingType string,
 /* OUT */ cmdletOutput RemoteAccessAccounting) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByDisableAccounting" , ComputerName, PassThru, DisableAccountingType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="AccountingOnOffMsg" type="string "></param>
// <param name="ComputerName" type="string "></param>
// <param name="EnableAccountingType" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RadiusPort" type="uint16 "></param>
// <param name="RadiusScore" type="uint8 "></param>
// <param name="RadiusServer" type="string "></param>
// <param name="RadiusTimeout" type="uint32 "></param>
// <param name="SharedSecret" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccounting "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessAccounting) SetByEnableAccounting( /* IN */ ComputerName string,
 /* IN */ PassThru bool,
 /* IN */ RadiusServer string,
 /* IN */ SharedSecret string,
 /* IN */ RadiusPort uint16,
 /* IN */ RadiusScore uint8,
 /* IN */ RadiusTimeout uint32,
 /* IN */ AccountingOnOffMsg string,
 /* IN */ EnableAccountingType string,
 /* OUT */ cmdletOutput RemoteAccessAccounting) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetByEnableAccounting" , ComputerName, PassThru, RadiusServer, SharedSecret, RadiusPort, RadiusScore, RadiusTimeout, AccountingOnOffMsg, EnableAccountingType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ComputerName" type="string "></param>

// <param name="cmdletOutput" type="RemoteAccessAccounting "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *PS_RemoteAccessAccounting) Get( /* IN */ ComputerName string,
 /* OUT */ cmdletOutput RemoteAccessAccounting) (result uint32, err error) {retVal, err := instance.InvokeMethod("Get" , ComputerName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


