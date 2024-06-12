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

// CreateTrustedDomEx_Start struct
type CreateTrustedDomEx_Start struct {
	*CreateTrustedDomEx
}

func NewCreateTrustedDomEx_StartEx1(instance *cim.WmiInstance) (newInstance *CreateTrustedDomEx_Start, err error) {
	tmp, err := NewCreateTrustedDomExEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx_Start{
		CreateTrustedDomEx: tmp,
	}
	return
}

func NewCreateTrustedDomEx_StartEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CreateTrustedDomEx_Start, err error) {
	tmp, err := NewCreateTrustedDomExEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CreateTrustedDomEx_Start{
		CreateTrustedDomEx: tmp,
	}
	return
}
