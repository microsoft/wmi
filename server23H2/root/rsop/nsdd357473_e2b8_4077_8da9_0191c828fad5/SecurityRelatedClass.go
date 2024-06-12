// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSDD357473_E2B8_4077_8DA9_0191C828FAD5
//////////////////////////////////////////////
package nsdd357473_e2b8_4077_8da9_0191c828fad5

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
