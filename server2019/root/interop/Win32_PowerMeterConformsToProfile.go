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

// Win32_PowerMeterConformsToProfile struct
type Win32_PowerMeterConformsToProfile struct { 
	*CIM_ElementConformsToProfile
}

	func NewWin32_PowerMeterConformsToProfileEx1(instance *cim.WmiInstance) (newInstance *Win32_PowerMeterConformsToProfile, err error) {tmp, err := NewCIM_ElementConformsToProfileEx1(instance)
		
	if err != nil { return }
	newInstance = &Win32_PowerMeterConformsToProfile {
	CIM_ElementConformsToProfile: tmp,
	}
	return
	}
	

	func NewWin32_PowerMeterConformsToProfileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_PowerMeterConformsToProfile, err error) {tmp, err := NewCIM_ElementConformsToProfileEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_PowerMeterConformsToProfile {
	CIM_ElementConformsToProfile: tmp,
	}
	return
	}
	

