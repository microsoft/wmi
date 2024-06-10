// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34
//
// ////////////////////////////////////////////
package ns9606b916_b14a_4eb6_84f0_165f56e1dd34

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SystemClass struct
type __SystemClass struct {
	*cim.WmiInstance
}

func New__SystemClassEx1(instance *cim.WmiInstance) (newInstance *__SystemClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &__SystemClass{
		WmiInstance: tmp,
	}
	return
}

func New__SystemClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SystemClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SystemClass{
		WmiInstance: tmp,
	}
	return
}
