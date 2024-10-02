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

// SamCompCreate_End struct
type SamCompCreate_End struct {
	*SamCompCreate
}

func NewSamCompCreate_EndEx1(instance *cim.WmiInstance) (newInstance *SamCompCreate_End, err error) {
	tmp, err := NewSamCompCreateEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamCompCreate_End{
		SamCompCreate: tmp,
	}
	return
}

func NewSamCompCreate_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamCompCreate_End, err error) {
	tmp, err := NewSamCompCreateEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamCompCreate_End{
		SamCompCreate: tmp,
	}
	return
}
