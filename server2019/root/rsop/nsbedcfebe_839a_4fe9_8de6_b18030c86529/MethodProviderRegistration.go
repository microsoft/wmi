// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NSBEDCFEBE_839A_4FE9_8DE6_B18030C86529
//////////////////////////////////////////////
package nsbedcfebe_839a_4fe9_8de6_b18030c86529
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __MethodProviderRegistration struct
type __MethodProviderRegistration struct { 
	*__ProviderRegistration
}

	func New__MethodProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__MethodProviderRegistration, err error) {tmp, err := New__ProviderRegistrationEx1(instance)
		
	if err != nil { return }
	newInstance = &__MethodProviderRegistration {
	__ProviderRegistration: tmp,
	}
	return
	}
	

	func New__MethodProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__MethodProviderRegistration, err error) {tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__MethodProviderRegistration {
	__ProviderRegistration: tmp,
	}
	return
	}
	

