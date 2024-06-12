// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSD7887C64_536D_4BB2_8B3E_EC03C01AABCC
//////////////////////////////////////////////
package nsd7887c64_536d_4bb2_8b3e_ec03c01aabcc

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// __IndicationRelated struct
type __IndicationRelated struct {
	*__SystemClass
}

func New__IndicationRelatedEx1(instance *cim.WmiInstance) (newInstance *__IndicationRelated, err error) {
	tmp, err := New__SystemClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &__IndicationRelated{
		__SystemClass: tmp,
	}
	return
}

func New__IndicationRelatedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *__IndicationRelated, err error) {
	tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &__IndicationRelated{
		__SystemClass: tmp,
	}
	return
}
