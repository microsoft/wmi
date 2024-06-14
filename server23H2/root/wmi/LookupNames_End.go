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

// LookupNames_End struct
type LookupNames_End struct {
	*LookupNames
}

func NewLookupNames_EndEx1(instance *cim.WmiInstance) (newInstance *LookupNames_End, err error) {
	tmp, err := NewLookupNamesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LookupNames_End{
		LookupNames: tmp,
	}
	return
}

func NewLookupNames_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LookupNames_End, err error) {
	tmp, err := NewLookupNamesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LookupNames_End{
		LookupNames: tmp,
	}
	return
}
