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

// Close_End struct
type Close_End struct {
	*Close
}

func NewClose_EndEx1(instance *cim.WmiInstance) (newInstance *Close_End, err error) {
	tmp, err := NewCloseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Close_End{
		Close: tmp,
	}
	return
}

func NewClose_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Close_End, err error) {
	tmp, err := NewCloseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Close_End{
		Close: tmp,
	}
	return
}
