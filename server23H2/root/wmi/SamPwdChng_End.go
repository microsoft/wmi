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

// SamPwdChng_End struct
type SamPwdChng_End struct {
	*SamPwdChng
}

func NewSamPwdChng_EndEx1(instance *cim.WmiInstance) (newInstance *SamPwdChng_End, err error) {
	tmp, err := NewSamPwdChngEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamPwdChng_End{
		SamPwdChng: tmp,
	}
	return
}

func NewSamPwdChng_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamPwdChng_End, err error) {
	tmp, err := NewSamPwdChngEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamPwdChng_End{
		SamPwdChng: tmp,
	}
	return
}
