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

// SamSetBootKeyInfo_End struct
type SamSetBootKeyInfo_End struct {
	*SamSetBootKeyInfo
}

func NewSamSetBootKeyInfo_EndEx1(instance *cim.WmiInstance) (newInstance *SamSetBootKeyInfo_End, err error) {
	tmp, err := NewSamSetBootKeyInfoEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetBootKeyInfo_End{
		SamSetBootKeyInfo: tmp,
	}
	return
}

func NewSamSetBootKeyInfo_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetBootKeyInfo_End, err error) {
	tmp, err := NewSamSetBootKeyInfoEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetBootKeyInfo_End{
		SamSetBootKeyInfo: tmp,
	}
	return
}
