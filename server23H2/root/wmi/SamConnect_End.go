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

// SamConnect_End struct
type SamConnect_End struct {
	*SamConnect
}

func NewSamConnect_EndEx1(instance *cim.WmiInstance) (newInstance *SamConnect_End, err error) {
	tmp, err := NewSamConnectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamConnect_End{
		SamConnect: tmp,
	}
	return
}

func NewSamConnect_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamConnect_End, err error) {
	tmp, err := NewSamConnectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamConnect_End{
		SamConnect: tmp,
	}
	return
}
