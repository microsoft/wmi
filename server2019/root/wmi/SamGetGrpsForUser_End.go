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

// SamGetGrpsForUser_End struct
type SamGetGrpsForUser_End struct {
	*SamGetGrpsForUser
}

func NewSamGetGrpsForUser_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetGrpsForUser_End, err error) {
	tmp, err := NewSamGetGrpsForUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetGrpsForUser_End{
		SamGetGrpsForUser: tmp,
	}
	return
}

func NewSamGetGrpsForUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetGrpsForUser_End, err error) {
	tmp, err := NewSamGetGrpsForUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetGrpsForUser_End{
		SamGetGrpsForUser: tmp,
	}
	return
}
