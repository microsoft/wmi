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

// SamOpenGrp_End struct
type SamOpenGrp_End struct {
	*SamOpenGrp
}

func NewSamOpenGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamOpenGrp_End, err error) {
	tmp, err := NewSamOpenGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamOpenGrp_End{
		SamOpenGrp: tmp,
	}
	return
}

func NewSamOpenGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamOpenGrp_End, err error) {
	tmp, err := NewSamOpenGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamOpenGrp_End{
		SamOpenGrp: tmp,
	}
	return
}
