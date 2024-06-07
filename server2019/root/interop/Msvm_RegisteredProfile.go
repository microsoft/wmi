// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_RegisteredProfile struct
type Msvm_RegisteredProfile struct { 
	*CIM_RegisteredProfile
}

	func NewMsvm_RegisteredProfileEx1(instance *cim.WmiInstance) (newInstance *Msvm_RegisteredProfile, err error) {tmp, err := NewCIM_RegisteredProfileEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_RegisteredProfile {
	CIM_RegisteredProfile: tmp,
	}
	return
	}
	

	func NewMsvm_RegisteredProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_RegisteredProfile, err error) {tmp, err := NewCIM_RegisteredProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_RegisteredProfile {
	CIM_RegisteredProfile: tmp,
	}
	return
	}
	
func  (instance* Msvm_RegisteredProfile) GetRelatedRegisteredProfile() (value *cim.WmiInstance, err error) {
		 return instance.GetRelated("Msvm_RegisteredProfile"); 
	}
	

