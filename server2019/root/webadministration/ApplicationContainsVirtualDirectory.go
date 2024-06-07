// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// ApplicationContainsVirtualDirectory struct
type ApplicationContainsVirtualDirectory struct { 
	*ObjectContainerAssociation
}

	func NewApplicationContainsVirtualDirectoryEx1(instance *cim.WmiInstance) (newInstance *ApplicationContainsVirtualDirectory, err error) {tmp, err := NewObjectContainerAssociationEx1(instance)
		
	if err != nil { return }
	newInstance = &ApplicationContainsVirtualDirectory {
	ObjectContainerAssociation: tmp,
	}
	return
	}
	

	func NewApplicationContainsVirtualDirectoryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ApplicationContainsVirtualDirectory, err error) {tmp, err := NewObjectContainerAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ApplicationContainsVirtualDirectory {
	ObjectContainerAssociation: tmp,
	}
	return
	}
	

