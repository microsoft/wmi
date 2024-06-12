// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NSA684FA6E_0838_4146_8046_48F98754B07A
//////////////////////////////////////////////
package nsa684fa6e_0838_4146_8046_48f98754b07a

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
