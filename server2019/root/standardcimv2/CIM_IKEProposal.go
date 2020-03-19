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

// CIM_IKEProposal struct
type CIM_IKEProposal struct {
	*CIM_SAProposal

	//
	AuthenticationMethod uint16

	//
	CipherAlgorithm uint16

	//
	GroupId uint16

	//
	HashAlgorithm uint16

	//
	MaxLifetimeKilobytes uint64

	//
	MaxLifetimeSeconds uint64

	//
	OtherAuthenticationMethod string

	//
	OtherCipherAlgorithm string

	//
	OtherHashAlgorithm string

	//
	VendorID string
}

func NewCIM_IKEProposalEx1(instance *cim.WmiInstance) (newInstance *CIM_IKEProposal, err error) {
	tmp, err := NewCIM_SAProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_IKEProposal{
		CIM_SAProposal: tmp,
	}
	return
}

func NewCIM_IKEProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_IKEProposal, err error) {
	tmp, err := NewCIM_SAProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_IKEProposal{
		CIM_SAProposal: tmp,
	}
	return
}

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *CIM_IKEProposal) SetPropertyAuthenticationMethod(value uint16) (err error) {
	return instance.SetProperty("AuthenticationMethod", value)
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *CIM_IKEProposal) GetPropertyAuthenticationMethod() (value uint16, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCipherAlgorithm sets the value of CipherAlgorithm for the instance
func (instance *CIM_IKEProposal) SetPropertyCipherAlgorithm(value uint16) (err error) {
	return instance.SetProperty("CipherAlgorithm", value)
}

// GetCipherAlgorithm gets the value of CipherAlgorithm for the instance
func (instance *CIM_IKEProposal) GetPropertyCipherAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("CipherAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupId sets the value of GroupId for the instance
func (instance *CIM_IKEProposal) SetPropertyGroupId(value uint16) (err error) {
	return instance.SetProperty("GroupId", value)
}

// GetGroupId gets the value of GroupId for the instance
func (instance *CIM_IKEProposal) GetPropertyGroupId() (value uint16, err error) {
	retValue, err := instance.GetProperty("GroupId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHashAlgorithm sets the value of HashAlgorithm for the instance
func (instance *CIM_IKEProposal) SetPropertyHashAlgorithm(value uint16) (err error) {
	return instance.SetProperty("HashAlgorithm", value)
}

// GetHashAlgorithm gets the value of HashAlgorithm for the instance
func (instance *CIM_IKEProposal) GetPropertyHashAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("HashAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLifetimeKilobytes sets the value of MaxLifetimeKilobytes for the instance
func (instance *CIM_IKEProposal) SetPropertyMaxLifetimeKilobytes(value uint64) (err error) {
	return instance.SetProperty("MaxLifetimeKilobytes", value)
}

// GetMaxLifetimeKilobytes gets the value of MaxLifetimeKilobytes for the instance
func (instance *CIM_IKEProposal) GetPropertyMaxLifetimeKilobytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeKilobytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLifetimeSeconds sets the value of MaxLifetimeSeconds for the instance
func (instance *CIM_IKEProposal) SetPropertyMaxLifetimeSeconds(value uint64) (err error) {
	return instance.SetProperty("MaxLifetimeSeconds", value)
}

// GetMaxLifetimeSeconds gets the value of MaxLifetimeSeconds for the instance
func (instance *CIM_IKEProposal) GetPropertyMaxLifetimeSeconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxLifetimeSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherAuthenticationMethod sets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_IKEProposal) SetPropertyOtherAuthenticationMethod(value string) (err error) {
	return instance.SetProperty("OtherAuthenticationMethod", value)
}

// GetOtherAuthenticationMethod gets the value of OtherAuthenticationMethod for the instance
func (instance *CIM_IKEProposal) GetPropertyOtherAuthenticationMethod() (value string, err error) {
	retValue, err := instance.GetProperty("OtherAuthenticationMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherCipherAlgorithm sets the value of OtherCipherAlgorithm for the instance
func (instance *CIM_IKEProposal) SetPropertyOtherCipherAlgorithm(value string) (err error) {
	return instance.SetProperty("OtherCipherAlgorithm", value)
}

// GetOtherCipherAlgorithm gets the value of OtherCipherAlgorithm for the instance
func (instance *CIM_IKEProposal) GetPropertyOtherCipherAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCipherAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherHashAlgorithm sets the value of OtherHashAlgorithm for the instance
func (instance *CIM_IKEProposal) SetPropertyOtherHashAlgorithm(value string) (err error) {
	return instance.SetProperty("OtherHashAlgorithm", value)
}

// GetOtherHashAlgorithm gets the value of OtherHashAlgorithm for the instance
func (instance *CIM_IKEProposal) GetPropertyOtherHashAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("OtherHashAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorID sets the value of VendorID for the instance
func (instance *CIM_IKEProposal) SetPropertyVendorID(value string) (err error) {
	return instance.SetProperty("VendorID", value)
}

// GetVendorID gets the value of VendorID for the instance
func (instance *CIM_IKEProposal) GetPropertyVendorID() (value string, err error) {
	retValue, err := instance.GetProperty("VendorID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
