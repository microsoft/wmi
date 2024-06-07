// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSE415D675_0F49_4CFF_A478_A2849757FFEE.Computer
//////////////////////////////////////////////
package computer
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_PolmkrLocalUsersSetting struct
type RSOP_PolmkrLocalUsersSetting struct { 
	*RSOP_PolmkrLocalUsersAndGroupsSetting
}

	func NewRSOP_PolmkrLocalUsersSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrLocalUsersSetting, err error) {tmp, err := NewRSOP_PolmkrLocalUsersAndGroupsSettingEx1(instance)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrLocalUsersSetting {
	RSOP_PolmkrLocalUsersAndGroupsSetting: tmp,
	}
	return
	}
	

	func NewRSOP_PolmkrLocalUsersSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_PolmkrLocalUsersSetting, err error) {tmp, err := NewRSOP_PolmkrLocalUsersAndGroupsSettingEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_PolmkrLocalUsersSetting {
	RSOP_PolmkrLocalUsersAndGroupsSetting: tmp,
	}
	return
	}
	

