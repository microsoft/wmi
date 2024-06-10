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

// SamGetDisplayEnumIdx_End struct
type SamGetDisplayEnumIdx_End struct {
	*SamGetDisplayEnumIdx
}

func NewSamGetDisplayEnumIdx_EndEx1(instance *cim.WmiInstance) (newInstance *SamGetDisplayEnumIdx_End, err error) {
	tmp, err := NewSamGetDisplayEnumIdxEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetDisplayEnumIdx_End{
		SamGetDisplayEnumIdx: tmp,
	}
	return
}

func NewSamGetDisplayEnumIdx_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetDisplayEnumIdx_End, err error) {
	tmp, err := NewSamGetDisplayEnumIdxEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetDisplayEnumIdx_End{
		SamGetDisplayEnumIdx: tmp,
	}
	return
}
