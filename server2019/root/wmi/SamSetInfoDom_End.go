// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamSetInfoDom_End struct
type SamSetInfoDom_End struct {
	*SamSetInfoDom
}

func NewSamSetInfoDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetInfoDom_End, err error) {
	tmp, err := NewSamSetInfoDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoDom_End{
		SamSetInfoDom: tmp,
	}
	return
}

func NewSamSetInfoDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetInfoDom_End, err error) {
	tmp, err := NewSamSetInfoDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoDom_End{
		SamSetInfoDom: tmp,
	}
	return
}
