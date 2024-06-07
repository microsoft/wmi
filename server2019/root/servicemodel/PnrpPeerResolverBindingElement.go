// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// PnrpPeerResolverBindingElement struct
type PnrpPeerResolverBindingElement struct { 
	*PeerResolverBindingElement
}

	func NewPnrpPeerResolverBindingElementEx1(instance *cim.WmiInstance) (newInstance *PnrpPeerResolverBindingElement, err error) {tmp, err := NewPeerResolverBindingElementEx1(instance)
		
	if err != nil { return }
	newInstance = &PnrpPeerResolverBindingElement {
	PeerResolverBindingElement: tmp,
	}
	return
	}
	

	func NewPnrpPeerResolverBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PnrpPeerResolverBindingElement, err error) {tmp, err := NewPeerResolverBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PnrpPeerResolverBindingElement {
	PeerResolverBindingElement: tmp,
	}
	return
	}
	

