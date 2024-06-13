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

// SamOpenUser_End struct
type SamOpenUser_End struct {
	*SamOpenUser
}

func NewSamOpenUser_EndEx1(instance *cim.WmiInstance) (newInstance *SamOpenUser_End, err error) {
	tmp, err := NewSamOpenUserEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamOpenUser_End{
		SamOpenUser: tmp,
	}
	return
}

func NewSamOpenUser_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamOpenUser_End, err error) {
	tmp, err := NewSamOpenUserEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamOpenUser_End{
		SamOpenUser: tmp,
	}
	return
}
