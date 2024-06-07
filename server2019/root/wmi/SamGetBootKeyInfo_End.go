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

// SamGetBootKeyInfo_End struct
type SamGetBootKeyInfo_End struct { 
	*SamGetBootKeyInfo
}

	func NewSamGetBootKeyInfo_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetBootKeyInfo_End, err error) {tmp, err := NewSamGetBootKeyInfoEx1(instance)
		
	if err != nil { return }
	newInstance = &SamGetBootKeyInfo_End {
	SamGetBootKeyInfo: tmp,
	}
	return
	}
	

	func NewSamGetBootKeyInfo_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamGetBootKeyInfo_End, err error) {tmp, err := NewSamGetBootKeyInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamGetBootKeyInfo_End {
	SamGetBootKeyInfo: tmp,
	}
	return
	}
	

