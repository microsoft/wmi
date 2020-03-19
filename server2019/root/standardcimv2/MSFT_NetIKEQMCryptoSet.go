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

// MSFT_NetIKEQMCryptoSet struct
type MSFT_NetIKEQMCryptoSet struct {
	*MSFT_NetIKECryptoSet

	//
	PfsGroupID uint16
}

func NewMSFT_NetIKEQMCryptoSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKEQMCryptoSet, err error) {
	tmp, err := NewMSFT_NetIKECryptoSetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEQMCryptoSet{
		MSFT_NetIKECryptoSet: tmp,
	}
	return
}

func NewMSFT_NetIKEQMCryptoSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKEQMCryptoSet, err error) {
	tmp, err := NewMSFT_NetIKECryptoSetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKEQMCryptoSet{
		MSFT_NetIKECryptoSet: tmp,
	}
	return
}

// SetPfsGroupID sets the value of PfsGroupID for the instance
func (instance *MSFT_NetIKEQMCryptoSet) SetPropertyPfsGroupID(value uint16) (err error) {
	return instance.SetProperty("PfsGroupID", value)
}

// GetPfsGroupID gets the value of PfsGroupID for the instance
func (instance *MSFT_NetIKEQMCryptoSet) GetPropertyPfsGroupID() (value uint16, err error) {
	retValue, err := instance.GetProperty("PfsGroupID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="NewName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEQMCryptoSet) Rename( /* IN */ NewName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewGPOSession" type="string "></param>
// <param name="NewID" type="string "></param>
// <param name="NewName" type="string "></param>
// <param name="NewPolicyStore" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIKEQMCryptoSet) CloneObject( /* IN */ NewName string,
	/* IN */ NewID string,
	/* IN */ NewPolicyStore string,
	/* IN */ NewGPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloneObject", NewName, NewID, NewPolicyStore, NewGPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
