// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RegisteredProfile struct
type RegisteredProfile struct { 
	*CIM_RegisteredProfile
}

	func NewRegisteredProfileEx1(instance *cim.WmiInstance) (newInstance *RegisteredProfile, err error) {tmp, err := NewCIM_RegisteredProfileEx1(instance)
		
	if err != nil { return }
	newInstance = &RegisteredProfile {
	CIM_RegisteredProfile: tmp,
	}
	return
	}
	

	func NewRegisteredProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RegisteredProfile, err error) {tmp, err := NewCIM_RegisteredProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RegisteredProfile {
	CIM_RegisteredProfile: tmp,
	}
	return
	}
	

