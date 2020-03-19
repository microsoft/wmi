// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetConSecRuleEMAuthSet struct
type MSFT_NetConSecRuleEMAuthSet struct {
	*MSFT_NetSARuleEMAuth
}

func NewMSFT_NetConSecRuleEMAuthSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetConSecRuleEMAuthSet, err error) {
	tmp, err := NewMSFT_NetSARuleEMAuthEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleEMAuthSet{
		MSFT_NetSARuleEMAuth: tmp,
	}
	return
}

func NewMSFT_NetConSecRuleEMAuthSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetConSecRuleEMAuthSet, err error) {
	tmp, err := NewMSFT_NetSARuleEMAuthEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetConSecRuleEMAuthSet{
		MSFT_NetSARuleEMAuth: tmp,
	}
	return
}
