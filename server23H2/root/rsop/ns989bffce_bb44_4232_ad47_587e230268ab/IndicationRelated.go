// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS989BFFCE_BB44_4232_AD47_587E230268AB
//////////////////////////////////////////////
package ns989bffce_bb44_4232_ad47_587e230268ab

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
