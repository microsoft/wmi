// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamQueryInfoDom_End struct
type SamQueryInfoDom_End struct {
	*SamQueryInfoDom
}

func NewSamQueryInfoDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamQueryInfoDom_End, err error) {
	tmp, err := NewSamQueryInfoDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoDom_End{
		SamQueryInfoDom: tmp,
	}
	return
}

func NewSamQueryInfoDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamQueryInfoDom_End, err error) {
	tmp, err := NewSamQueryInfoDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamQueryInfoDom_End{
		SamQueryInfoDom: tmp,
	}
	return
}
