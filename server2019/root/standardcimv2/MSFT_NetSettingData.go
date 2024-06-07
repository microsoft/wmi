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

// MSFT_NetSettingData struct
type MSFT_NetSettingData struct { 
	*CIM_SettingData
}

	func NewMSFT_NetSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetSettingData, err error) {tmp, err := NewCIM_SettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetSettingData {
	CIM_SettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetSettingData, err error) {tmp, err := NewCIM_SettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetSettingData {
	CIM_SettingData: tmp,
	}
	return
	}
	

