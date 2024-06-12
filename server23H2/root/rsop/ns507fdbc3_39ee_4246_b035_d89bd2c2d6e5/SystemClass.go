// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS507FDBC3_39EE_4246_B035_D89BD2C2D6E5
//////////////////////////////////////////////
package ns507fdbc3_39ee_4246_b035_d89bd2c2d6e5

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
