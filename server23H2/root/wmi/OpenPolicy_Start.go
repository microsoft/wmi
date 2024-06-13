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

// OpenPolicy_Start struct
type OpenPolicy_Start struct {
	*OpenPolicy
}

func NewOpenPolicy_StartEx1(instance *cim.WmiInstance) (newInstance *OpenPolicy_Start, err error) {
	tmp, err := NewOpenPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OpenPolicy_Start{
		OpenPolicy: tmp,
	}
	return
}

func NewOpenPolicy_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OpenPolicy_Start, err error) {
	tmp, err := NewOpenPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OpenPolicy_Start{
		OpenPolicy: tmp,
	}
	return
}
