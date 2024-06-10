// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIKECertAuthProposal struct
type MSFT_NetIKECertAuthProposal struct {
	*MSFT_NetIKEAuthProposal

	//
	CertName string

	//
	CertNameType uint16

	//
	EKUs []string

	//
	ExcludeCAName bool

	//
	FollowRenewal bool

	//
	MapToAccount bool

	//
	SelectionCriteria bool

	//
	SigningAlgorithm uint16

	//
	Thumbprint string

	//
	TrustedCA string

	//
	TrustedCAType uint16

	//
	ValidationCriteria bool
}

func NewMSFT_NetIKECertAuthProposalEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIKECertAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECertAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

func NewMSFT_NetIKECertAuthProposalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIKECertAuthProposal, err error) {
	tmp, err := NewMSFT_NetIKEAuthProposalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIKECertAuthProposal{
		MSFT_NetIKEAuthProposal: tmp,
	}
	return
}

// SetCertName sets the value of CertName for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyCertName(value string) (err error) {
	return instance.SetProperty("CertName", (value))
}

// GetCertName gets the value of CertName for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyCertName() (value string, err error) {
	retValue, err := instance.GetProperty("CertName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetCertNameType sets the value of CertNameType for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyCertNameType(value uint16) (err error) {
	return instance.SetProperty("CertNameType", (value))
}

// GetCertNameType gets the value of CertNameType for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyCertNameType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CertNameType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetEKUs sets the value of EKUs for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyEKUs(value []string) (err error) {
	return instance.SetProperty("EKUs", (value))
}

// GetEKUs gets the value of EKUs for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyEKUs() (value []string, err error) {
	retValue, err := instance.GetProperty("EKUs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetExcludeCAName sets the value of ExcludeCAName for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyExcludeCAName(value bool) (err error) {
	return instance.SetProperty("ExcludeCAName", (value))
}

// GetExcludeCAName gets the value of ExcludeCAName for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyExcludeCAName() (value bool, err error) {
	retValue, err := instance.GetProperty("ExcludeCAName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetFollowRenewal sets the value of FollowRenewal for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyFollowRenewal(value bool) (err error) {
	return instance.SetProperty("FollowRenewal", (value))
}

// GetFollowRenewal gets the value of FollowRenewal for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyFollowRenewal() (value bool, err error) {
	retValue, err := instance.GetProperty("FollowRenewal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMapToAccount sets the value of MapToAccount for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyMapToAccount(value bool) (err error) {
	return instance.SetProperty("MapToAccount", (value))
}

// GetMapToAccount gets the value of MapToAccount for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyMapToAccount() (value bool, err error) {
	retValue, err := instance.GetProperty("MapToAccount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetSelectionCriteria sets the value of SelectionCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertySelectionCriteria(value bool) (err error) {
	return instance.SetProperty("SelectionCriteria", (value))
}

// GetSelectionCriteria gets the value of SelectionCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertySelectionCriteria() (value bool, err error) {
	retValue, err := instance.GetProperty("SelectionCriteria")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetSigningAlgorithm sets the value of SigningAlgorithm for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertySigningAlgorithm(value uint16) (err error) {
	return instance.SetProperty("SigningAlgorithm", (value))
}

// GetSigningAlgorithm gets the value of SigningAlgorithm for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertySigningAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("SigningAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", (value))
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTrustedCA sets the value of TrustedCA for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyTrustedCA(value string) (err error) {
	return instance.SetProperty("TrustedCA", (value))
}

// GetTrustedCA gets the value of TrustedCA for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyTrustedCA() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedCA")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTrustedCAType sets the value of TrustedCAType for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyTrustedCAType(value uint16) (err error) {
	return instance.SetProperty("TrustedCAType", (value))
}

// GetTrustedCAType gets the value of TrustedCAType for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyTrustedCAType() (value uint16, err error) {
	retValue, err := instance.GetProperty("TrustedCAType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetValidationCriteria sets the value of ValidationCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyValidationCriteria(value bool) (err error) {
	return instance.SetProperty("ValidationCriteria", (value))
}

// GetValidationCriteria gets the value of ValidationCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyValidationCriteria() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidationCriteria")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
