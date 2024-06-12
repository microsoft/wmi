// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS0F6E8CAD_B227_4B33_9B56_FF85C39C659C
//////////////////////////////////////////////
package ns0f6e8cad_b227_4b33_9b56_ff85c39c659c

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __MethodProviderRegistration struct
type __MethodProviderRegistration struct {
	*__ProviderRegistration
}

func New__MethodProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}

func New__MethodProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__MethodProviderRegistration, err error) {
	tmp, err := New__ProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__MethodProviderRegistration{
		__ProviderRegistration: tmp,
	}
	return
}
