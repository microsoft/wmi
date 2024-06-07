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

// SamOpenUser struct
type SamOpenUser struct { 
	*MSSAMTrace
}

	func NewSamOpenUserEx1(instance *cim.WmiInstance) (newInstance *SamOpenUser, err error) {tmp, err := NewMSSAMTraceEx1(instance)
		
	if err != nil { return }
	newInstance = &SamOpenUser {
	MSSAMTrace: tmp,
	}
	return
	}
	

	func NewSamOpenUserEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamOpenUser, err error) {tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamOpenUser {
	MSSAMTrace: tmp,
	}
	return
	}
	

