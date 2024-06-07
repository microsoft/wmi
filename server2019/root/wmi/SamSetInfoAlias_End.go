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

// SamSetInfoAlias_End struct
type SamSetInfoAlias_End struct { 
	*SamSetInfoAlias
}

	func NewSamSetInfoAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetInfoAlias_End, err error) {tmp, err := NewSamSetInfoAliasEx1(instance)
		
	if err != nil { return }
	newInstance = &SamSetInfoAlias_End {
	SamSetInfoAlias: tmp,
	}
	return
	}
	

	func NewSamSetInfoAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamSetInfoAlias_End, err error) {tmp, err := NewSamSetInfoAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamSetInfoAlias_End {
	SamSetInfoAlias: tmp,
	}
	return
	}
	

