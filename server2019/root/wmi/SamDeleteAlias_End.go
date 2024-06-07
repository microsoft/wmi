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

// SamDeleteAlias_End struct
type SamDeleteAlias_End struct { 
	*SamDeleteAlias
}

	func NewSamDeleteAlias_EndEx1(instance *cim.WmiInstance) (newInstance *SamDeleteAlias_End, err error) {tmp, err := NewSamDeleteAliasEx1(instance)
		
	if err != nil { return }
	newInstance = &SamDeleteAlias_End {
	SamDeleteAlias: tmp,
	}
	return
	}
	

	func NewSamDeleteAlias_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SamDeleteAlias_End, err error) {tmp, err := NewSamDeleteAliasEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SamDeleteAlias_End {
	SamDeleteAlias: tmp,
	}
	return
	}
	

