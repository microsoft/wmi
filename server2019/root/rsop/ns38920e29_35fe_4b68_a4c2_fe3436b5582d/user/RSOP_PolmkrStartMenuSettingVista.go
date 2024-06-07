// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS38920E29_35FE_4B68_A4C2_FE3436B5582D.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrStartMenuSettingVista struct
type RSOP_PolmkrStartMenuSettingVista struct { 
	*RSOP_PolmkrProSetting
}

	func NewRSOP_PolmkrStartMenuSettingVistaEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrStartMenuSettingVista, err error) {tmp, err := NewRSOP_PolmkrProSettingEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrStartMenuSettingVista {
	RSOP_PolmkrProSetting: tmp,
	}
	return
	}
	

	func NewRSOP_PolmkrStartMenuSettingVistaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_PolmkrStartMenuSettingVista, err error) {tmp, err := NewRSOP_PolmkrProSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrStartMenuSettingVista {
	RSOP_PolmkrProSetting: tmp,
	}
	return
	}
	

