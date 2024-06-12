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

// SamIdByName_End struct
type SamIdByName_End struct {
	*SamIdByName
}

func NewSamIdByName_EndEx1(instance *cim.WmiInstance) (newInstance *SamIdByName_End, err error) {
	tmp, err := NewSamIdByNameEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamIdByName_End{
		SamIdByName: tmp,
	}
	return
}

func NewSamIdByName_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamIdByName_End, err error) {
	tmp, err := NewSamIdByNameEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamIdByName_End{
		SamIdByName: tmp,
	}
	return
}
