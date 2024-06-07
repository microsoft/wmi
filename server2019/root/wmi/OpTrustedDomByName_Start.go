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

// OpTrustedDomByName_Start struct
type OpTrustedDomByName_Start struct { 
	*OpTrustedDomByName
}

	func NewOpTrustedDomByName_StartEx1(instance *cim.WmiInstance) (newInstance *OpTrustedDomByName_Start, err error) {tmp, err := NewOpTrustedDomByNameEx1(instance)
		
	if err != nil { return }
	newInstance = &OpTrustedDomByName_Start {
	OpTrustedDomByName: tmp,
	}
	return
	}
	

	func NewOpTrustedDomByName_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OpTrustedDomByName_Start, err error) {tmp, err := NewOpTrustedDomByNameEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OpTrustedDomByName_Start {
	OpTrustedDomByName: tmp,
	}
	return
	}
	

