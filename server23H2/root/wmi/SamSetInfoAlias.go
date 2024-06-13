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

// SamSetInfoAlias struct
type SamSetInfoAlias struct {
	*MSSAMTrace
}

func NewSamSetInfoAliasEx1(instance *cim.WmiInstance) (newInstance *SamSetInfoAlias, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoAlias{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamSetInfoAliasEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamSetInfoAlias, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamSetInfoAlias{
		MSSAMTrace: tmp,
	}
	return
}
