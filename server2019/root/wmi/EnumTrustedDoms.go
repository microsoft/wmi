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

// EnumTrustedDoms struct
type EnumTrustedDoms struct { 
	*MSLSATrace
}

	func NewEnumTrustedDomsEx1(instance *cim.WmiInstance) (newInstance *EnumTrustedDoms, err error) {tmp, err := NewMSLSATraceEx1(instance)
		
	if err != nil { return }
	newInstance = &EnumTrustedDoms {
	MSLSATrace: tmp,
	}
	return
	}
	

	func NewEnumTrustedDomsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EnumTrustedDoms, err error) {tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EnumTrustedDoms {
	MSLSATrace: tmp,
	}
	return
	}
	

