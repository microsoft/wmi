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

// SamEnumGrpsInDom_End struct
type SamEnumGrpsInDom_End struct {
	*SamEnumGrpsInDom
}

func NewSamEnumGrpsInDom_EndEx1(instance *cim.WmiInstance) (newInstance *SamEnumGrpsInDom_End, err error) {
	tmp, err := NewSamEnumGrpsInDomEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamEnumGrpsInDom_End{
		SamEnumGrpsInDom: tmp,
	}
	return
}

func NewSamEnumGrpsInDom_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamEnumGrpsInDom_End, err error) {
	tmp, err := NewSamEnumGrpsInDomEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamEnumGrpsInDom_End{
		SamEnumGrpsInDom: tmp,
	}
	return
}
