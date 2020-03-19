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

// MSFT_NetIKEKerbAuthProposal struct
type MSFT_NetIKEKerbAuthProposal struct {
	*MSFT_NetIKEAuthProposal

	//
	KerbProxy string
}

func NewMSFT_NetIKEKerbAuthProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEKerbAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEKerbAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

func NewMSFT_NetIKEKerbAuthProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEKerbAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEKerbAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

// SetKerbProxy sets the value of KerbProxy for the instance
func (instance *MSFT_NetIKEKerbAuthProposal) SetPropertyKerbProxy(value string) (err error) {
	return instance.SetProperty("KerbProxy", value)
}

// GetKerbProxy gets the value of KerbProxy for the instance
func (instance *MSFT_NetIKEKerbAuthProposal) GetPropertyKerbProxy() (value string, err error) {
	retValue, err := instance.GetProperty("KerbProxy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
