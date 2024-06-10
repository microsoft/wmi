// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// QuerySecret_End struct
type QuerySecret_End struct {
	*QuerySecret
}

func NewQuerySecret_EndEx1(instance *cim.WmiInstance) (newInstance *QuerySecret_End, err error) {
	tmp, err := NewQuerySecretEx1(instance)

	if err != nil {
		return
	}
	newInstance = &QuerySecret_End{
		QuerySecret: tmp,
	}
	return
}

func NewQuerySecret_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *QuerySecret_End, err error) {
	tmp, err := NewQuerySecretEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &QuerySecret_End{
		QuerySecret: tmp,
	}
	return
}
