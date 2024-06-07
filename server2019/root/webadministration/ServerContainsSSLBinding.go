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

// ServerContainsSSLBinding struct
type ServerContainsSSLBinding struct { 
	*ObjectContainerAssociation
}

	func NewServerContainsSSLBindingEx1(instance *cim.WmiInstance) (newInstance *ServerContainsSSLBinding, err error) {tmp, err := NewObjectContainerAssociationEx1(instance)
		
	if err != nil { return }
	newInstance = &ServerContainsSSLBinding {
	ObjectContainerAssociation: tmp,
	}
	return
	}
	

	func NewServerContainsSSLBindingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ServerContainsSSLBinding, err error) {tmp, err := NewObjectContainerAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ServerContainsSSLBinding {
	ObjectContainerAssociation: tmp,
	}
	return
	}
	

