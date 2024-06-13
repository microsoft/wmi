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

// OpenPolicy_End struct
type OpenPolicy_End struct {
	*OpenPolicy
}

func NewOpenPolicy_EndEx1(instance *cim.WmiInstance) (newInstance *OpenPolicy_End, err error) {
	tmp, err := NewOpenPolicyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &OpenPolicy_End{
		OpenPolicy: tmp,
	}
	return
}

func NewOpenPolicy_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *OpenPolicy_End, err error) {
	tmp, err := NewOpenPolicyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &OpenPolicy_End{
		OpenPolicy: tmp,
	}
	return
}
