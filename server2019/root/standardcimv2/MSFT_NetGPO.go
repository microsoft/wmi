// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetGPO struct
type MSFT_NetGPO struct { 
	*CIM_SettingData
}

	func NewMSFT_NetGPOEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetGPO, err error) {tmp, err := NewCIM_SettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetGPO {
	CIM_SettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetGPOEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetGPO, err error) {tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetGPO {
	CIM_SettingData: tmp,
	}
	return
	}
	

// 

// <param name="DomainController" type="string "></param>
// <param name="PolicyStore" type="string "></param>

// <param name="GPOSession" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetGPO) Open( /* IN */ PolicyStore string,
 /* IN */ DomainController string,
 /* OUT */ GPOSession string) (result uint32, err error) {retVal, err := instance.InvokeMethod("Open" , PolicyStore, DomainController)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="GPOSession" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetGPO) Save( /* IN */ GPOSession string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Save" , GPOSession);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


