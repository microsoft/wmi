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

// MSFT_NetIKEPSKAuthProposal struct
type MSFT_NetIKEPSKAuthProposal struct {
	*MSFT_NetIKEAuthProposal

	//
	PreSharedKey string
}

func NewMSFT_NetIKEPSKAuthProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEPSKAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEPSKAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

func NewMSFT_NetIKEPSKAuthProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEPSKAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEPSKAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

// SetPreSharedKey sets the value of PreSharedKey for the instance
func (instance *MSFT_NetIKEPSKAuthProposal) SetPropertyPreSharedKey(value string) (err error) {
	return instance.SetProperty("PreSharedKey", value)
}

// GetPreSharedKey gets the value of PreSharedKey for the instance
func (instance *MSFT_NetIKEPSKAuthProposal) GetPropertyPreSharedKey() (value string, err error) {
	retValue, err := instance.GetProperty("PreSharedKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
