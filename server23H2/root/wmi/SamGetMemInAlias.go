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

// SamGetMemInAlias struct
type SamGetMemInAlias struct {
	*MSSAMTrace
}

func NewSamGetMemInAliasEx1(instance *cim.WmiInstance) (newInstance *SamGetMemInAlias, err error) {
	tmp, err := NewMSSAMTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInAlias{
		MSSAMTrace: tmp,
	}
	return
}

func NewSamGetMemInAliasEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SamGetMemInAlias, err error) {
	tmp, err := NewMSSAMTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SamGetMemInAlias{
		MSSAMTrace: tmp,
	}
	return
}
