// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamRemoveMemFromForeignDom_End struct
type SamRemoveMemFromForeignDom_End struct {
	*SamRemoveMemFromForeignDom
}

func NewSamRemoveMemFromForeignDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamRemoveMemFromForeignDom_End, err error) {
	tmp, err := NewSamRemoveMemFromForeignDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamRemoveMemFromForeignDom_End{
		SamRemoveMemFromForeignDom: tmp,
	}
	return
}

func NewSamRemoveMemFromForeignDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamRemoveMemFromForeignDom_End, err error) {
	tmp, err := NewSamRemoveMemFromForeignDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamRemoveMemFromForeignDom_End{
		SamRemoveMemFromForeignDom: tmp,
	}
	return
}
