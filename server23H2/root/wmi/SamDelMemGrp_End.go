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

// SamDelMemGrp_End struct
type SamDelMemGrp_End struct {
	*SamDelMemGrp
}

func NewSamDelMemGrp_EndEx1(instance *cim.WmiInstance) (newInstance *SamDelMemGrp_End, err error) {
	tmp, err := NewSamDelMemGrpEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamDelMemGrp_End{
		SamDelMemGrp: tmp,
	}
	return
}

func NewSamDelMemGrp_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamDelMemGrp_End, err error) {
	tmp, err := NewSamDelMemGrpEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamDelMemGrp_End{
		SamDelMemGrp: tmp,
	}
	return
}
