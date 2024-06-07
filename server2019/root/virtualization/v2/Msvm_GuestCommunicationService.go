// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_GuestCommunicationService struct
type Msvm_GuestCommunicationService struct { 
	*Msvm_GuestService
}

	func NewMsvm_GuestCommunicationServiceEx1(instance *cim.WmiInstance) (newInstance *Msvm_GuestCommunicationService, err error) {tmp, err := NewMsvm_GuestServiceEx1(instance)
		
	if err != nil { return }
	newInstance = &Msvm_GuestCommunicationService {
	Msvm_GuestService: tmp,
	}
	return
	}
	

	func NewMsvm_GuestCommunicationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Msvm_GuestCommunicationService, err error) {tmp, err := NewMsvm_GuestServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Msvm_GuestCommunicationService {
	Msvm_GuestService: tmp,
	}
	return
	}
	

