// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS5B5202B1_2FC1_4BB0_A3B5_EB3B3EAE1626.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __SecurityRelatedClass struct
type __SecurityRelatedClass struct {
	*cim.WmiInstance
}

func New__SecurityRelatedClassEx1(instance *cim.WmiInstance) (newInstance *__SecurityRelatedClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &__SecurityRelatedClass{
		WmiInstance: tmp,
	}
	return
}

func New__SecurityRelatedClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__SecurityRelatedClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__SecurityRelatedClass{
		WmiInstance: tmp,
	}
	return
}
