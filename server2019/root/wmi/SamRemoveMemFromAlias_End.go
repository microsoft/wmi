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

// SamRemoveMemFromAlias_End struct
type SamRemoveMemFromAlias_End struct { 
	*SamRemoveMemFromAlias
}

	func NewSamRemoveMemFromAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamRemoveMemFromAlias_End, err error) {tmp, err := NewSamRemoveMemFromAliasEx1(instance)
		
	if err != nil { return }
	newInstance = &SamRemoveMemFromAlias_End {
	SamRemoveMemFromAlias: tmp,
	}
	return
	}
	

	func NewSamRemoveMemFromAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamRemoveMemFromAlias_End, err error) {tmp, err := NewSamRemoveMemFromAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamRemoveMemFromAlias_End {
	SamRemoveMemFromAlias: tmp,
	}
	return
	}
	

