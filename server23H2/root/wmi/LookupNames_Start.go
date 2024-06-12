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

// LookupNames_Start struct
type LookupNames_Start struct {
	*LookupNames
}

func NewLookupNames_StartEx1(instance *cim.WmiInstance) (newInstance *LookupNames_Start, err error) {
	tmp, err := NewLookupNamesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LookupNames_Start{
		LookupNames: tmp,
	}
	return
}

func NewLookupNames_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LookupNames_Start, err error) {
	tmp, err := NewLookupNamesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LookupNames_Start{
		LookupNames: tmp,
	}
	return
}
