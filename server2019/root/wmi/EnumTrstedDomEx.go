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

// EnumTrstedDomEx struct
type EnumTrstedDomEx struct { 
	*MSLSATrace
}

	func NewEnumTrstedDomExEx1(instance *cim.WmiInstance) (newInstance *EnumTrstedDomEx, err error) {tmp, err := NewMSLSATraceEx1(instance)
		
	if err != nil { return }
	newInstance = &EnumTrstedDomEx {
	MSLSATrace: tmp,
	}
	return
	}
	

	func NewEnumTrstedDomExEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *EnumTrstedDomEx, err error) {tmp, err := NewMSLSATraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &EnumTrstedDomEx {
	MSLSATrace: tmp,
	}
	return
	}
	

