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

// SamUserPwdSet_End struct
type SamUserPwdSet_End struct { 
	*SamUserPwdSet
}

	func NewSamUserPwdSet_EndEx1(instance *cim.WmiInstance) (newInstance *SamUserPwdSet_End, err error) {tmp, err := NewSamUserPwdSetEx1(instance)
		
	if err != nil { return }
	newInstance = &SamUserPwdSet_End {
	SamUserPwdSet: tmp,
	}
	return
	}
	

	func NewSamUserPwdSet_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamUserPwdSet_End, err error) {tmp, err := NewSamUserPwdSetEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamUserPwdSet_End {
	SamUserPwdSet: tmp,
	}
	return
	}
	

