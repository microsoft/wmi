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

// SamSetInfoUser_End struct
type SamSetInfoUser_End struct { 
	*SamSetInfoUser
}

	func NewSamSetInfoUser_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetInfoUser_End, err error) {tmp, err := NewSamSetInfoUserEx1(instance)
		
	if err != nil { return }
	newInstance = &SamSetInfoUser_End {
	SamSetInfoUser: tmp,
	}
	return
	}
	

	func NewSamSetInfoUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamSetInfoUser_End, err error) {tmp, err := NewSamSetInfoUserEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamSetInfoUser_End {
	SamSetInfoUser: tmp,
	}
	return
	}
	

