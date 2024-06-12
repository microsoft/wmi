// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS223D6E5D_B897_4CF4_927F_6DBB3EA73CF0
//////////////////////////////////////////////
package ns223d6e5d_b897_4cf4_927f_6dbb3ea73cf0

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __InstanceProviderRegistration struct
type __InstanceProviderRegistration struct {
	*__ObjectProviderRegistration
}

func New__InstanceProviderRegistrationEx1(instance *cim.WmiInstance) (newInstance *__InstanceProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__InstanceProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}

func New__InstanceProviderRegistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__InstanceProviderRegistration, err error) {
	tmp, err := New__ObjectProviderRegistrationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__InstanceProviderRegistration{
		__ObjectProviderRegistration: tmp,
	}
	return
}
