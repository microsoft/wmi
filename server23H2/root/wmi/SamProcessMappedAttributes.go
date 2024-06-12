// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SamProcessMappedAttributes struct
type SamProcessMappedAttributes struct {
	*MSSAMTrace
}

func NewSamProcessMappedAttributesEx1(instance *cim.WmiInstance) (newInstance *SamProcessMappedAttributes, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamProcessMappedAttributes{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamProcessMappedAttributesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamProcessMappedAttributes, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamProcessMappedAttributes{
		MSSAMTrace: tmp,
	}
	return
}
