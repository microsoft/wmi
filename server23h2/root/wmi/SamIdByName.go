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

// SamIdByName struct
type SamIdByName struct {
	*MSSAMTrace
}

func NewSamIdByNameEx1(instance *cim.WmiInstance) (newInstance *SamIdByName, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamIdByName{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamIdByNameEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamIdByName, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamIdByName{
		MSSAMTrace: tmp,
	}
	return
}
