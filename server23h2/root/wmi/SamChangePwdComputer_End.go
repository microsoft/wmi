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

// SamChangePwdComputer_End struct
type SamChangePwdComputer_End struct {
	*SamChangePwdComputer
}

func NewSamChangePwdComputer_EndEx1(instance *cim.WmiInstance) (newInstance *SamChangePwdComputer_End, err error) {
	tmp, err := NewSamChangePwdComputerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamChangePwdComputer_End{
		SamChangePwdComputer: tmp,
	}
	return
}

func NewSamChangePwdComputer_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamChangePwdComputer_End, err error) {
	tmp, err := NewSamChangePwdComputerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamChangePwdComputer_End{
		SamChangePwdComputer: tmp,
	}
	return
}
