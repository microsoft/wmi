// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// EnumTrustedDoms_End struct
type EnumTrustedDoms_End struct { 
	*EnumTrustedDoms
}

	func NewEnumTrustedDoms_EndEx1(instance *cim.WmiInstance) (newInstance *EnumTrustedDoms_End, err error) {tmp, err := NewEnumTrustedDomsEx1(instance)
		
	if err != nil { return }
	newInstance = &EnumTrustedDoms_End {
	EnumTrustedDoms: tmp,
	}
	return
	}
	

	func NewEnumTrustedDoms_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EnumTrustedDoms_End, err error) {tmp, err := NewEnumTrustedDomsEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EnumTrustedDoms_End {
	EnumTrustedDoms: tmp,
	}
	return
	}
	

