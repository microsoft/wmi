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

// CreateTrustedDomEx_End struct
type CreateTrustedDomEx_End struct {
	*CreateTrustedDomEx
}

func NewCreateTrustedDomEx_EndEx1(instance *cim.WmiInstance) (newInstance *CreateTrustedDomEx_End, err error) {
	tmp, err := NewCreateTrustedDomExEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx_End{
		CreateTrustedDomEx: tmp,
	}
	return
}

func NewCreateTrustedDomEx_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CreateTrustedDomEx_End, err error) {
	tmp, err := NewCreateTrustedDomExEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx_End{
		CreateTrustedDomEx: tmp,
	}
	return
}
