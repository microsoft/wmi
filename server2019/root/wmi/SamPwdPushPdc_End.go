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

// SamPwdPushPdc_End struct
type SamPwdPushPdc_End struct { 
	*SamPwdPushPdc
}

	func NewSamPwdPushPdc_EndEx1(instance *cim.WmiInstance) (newInstance *SamPwdPushPdc_End, err error) {tmp, err := NewSamPwdPushPdcEx1(instance)
		
	if err != nil { return }
	newInstance = &SamPwdPushPdc_End {
	SamPwdPushPdc: tmp,
	}
	return
	}
	

	func NewSamPwdPushPdc_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamPwdPushPdc_End, err error) {tmp, err := NewSamPwdPushPdcEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamPwdPushPdc_End {
	SamPwdPushPdc: tmp,
	}
	return
	}
	

