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

// SamOpenDomain_End struct
type SamOpenDomain_End struct {
	*SamOpenDomain
}

func NewSamOpenDomain_EndEx1(instance *cim.WmiInstance) (newInstance *SamOpenDomain_End, err error) {
	tmp, err := NewSamOpenDomainEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamOpenDomain_End{
		SamOpenDomain: tmp,
	}
	return
}

func NewSamOpenDomain_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamOpenDomain_End, err error) {
	tmp, err := NewSamOpenDomainEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamOpenDomain_End{
		SamOpenDomain: tmp,
	}
	return
}
