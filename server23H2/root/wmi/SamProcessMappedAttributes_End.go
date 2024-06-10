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

// SamProcessMappedAttributes_End struct
type SamProcessMappedAttributes_End struct {
	*SamProcessMappedAttributes
}

func NewSamProcessMappedAttributes_EndEx1(instance *cim.WmiInstance) (newInstance *SamProcessMappedAttributes_End, err error) {
	tmp, err := NewSamProcessMappedAttributesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamProcessMappedAttributes_End{
		SamProcessMappedAttributes: tmp,
	}
	return
}

func NewSamProcessMappedAttributes_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamProcessMappedAttributes_End, err error) {
	tmp, err := NewSamProcessMappedAttributesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamProcessMappedAttributes_End{
		SamProcessMappedAttributes: tmp,
	}
	return
}
