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

// SamNameById_End struct
type SamNameById_End struct {
	*SamNameById
}

func NewSamNameById_EndEx1(instance *cim.WmiInstance) (newInstance *SamNameById_End, err error) {
	tmp, err := NewSamNameByIdEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamNameById_End{
		SamNameById: tmp,
	}
	return
}

func NewSamNameById_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamNameById_End, err error) {
	tmp, err := NewSamNameByIdEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamNameById_End{
		SamNameById: tmp,
	}
	return
}
