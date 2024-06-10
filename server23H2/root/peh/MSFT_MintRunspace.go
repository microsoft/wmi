// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.PEH
//
// ////////////////////////////////////////////
package peh

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MintRunspace struct
type MSFT_MintRunspace struct {
	*MSFT_Runspace
}

func NewMSFT_MintRunspaceEx1(instance *cim.WmiInstance) (newInstance *MSFT_MintRunspace, err error) {
	tmp, err := NewMSFT_RunspaceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintRunspace{
		MSFT_Runspace: tmp,
	}
	return
}

func NewMSFT_MintRunspaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MintRunspace, err error) {
	tmp, err := NewMSFT_RunspaceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MintRunspace{
		MSFT_Runspace: tmp,
	}
	return
}
