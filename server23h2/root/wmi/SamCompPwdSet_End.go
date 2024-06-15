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

// SamCompPwdSet_End struct
type SamCompPwdSet_End struct {
	*SamCompPwdSet
}

func NewSamCompPwdSet_EndEx1(instance *cim.WmiInstance) (newInstance *SamCompPwdSet_End, err error) {
	tmp, err := NewSamCompPwdSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamCompPwdSet_End{
		SamCompPwdSet: tmp,
	}
	return
}

func NewSamCompPwdSet_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamCompPwdSet_End, err error) {
	tmp, err := NewSamCompPwdSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamCompPwdSet_End{
		SamCompPwdSet: tmp,
	}
	return
}
