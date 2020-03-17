// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIKECertAuthProposal struct
type MSFT_NetIKECertAuthProposal struct {
	MSFT_NetIKEAuthProposal

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

// SetCertName sets the value of CertName for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyCertName(value string) (err error) {
	return instance.SetProperty("CertName", value)
}

// GetCertName gets the value of CertName for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyCertName() (value string, err error) {
	retValue, err := instance.GetProperty("CertName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCertNameType sets the value of CertNameType for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyCertNameType(value uint16) (err error) {
	return instance.SetProperty("CertNameType", value)
}

// GetCertNameType gets the value of CertNameType for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyCertNameType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CertNameType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEKUs sets the value of EKUs for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyEKUs(value []string) (err error) {
	return instance.SetProperty("EKUs", value)
}

// GetEKUs gets the value of EKUs for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyEKUs() (value []string, err error) {
	retValue, err := instance.GetProperty("EKUs")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExcludeCAName sets the value of ExcludeCAName for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyExcludeCAName(value bool) (err error) {
	return instance.SetProperty("ExcludeCAName", value)
}

// GetExcludeCAName gets the value of ExcludeCAName for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyExcludeCAName() (value bool, err error) {
	retValue, err := instance.GetProperty("ExcludeCAName")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFollowRenewal sets the value of FollowRenewal for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyFollowRenewal(value bool) (err error) {
	return instance.SetProperty("FollowRenewal", value)
}

// GetFollowRenewal gets the value of FollowRenewal for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyFollowRenewal() (value bool, err error) {
	retValue, err := instance.GetProperty("FollowRenewal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMapToAccount sets the value of MapToAccount for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyMapToAccount(value bool) (err error) {
	return instance.SetProperty("MapToAccount", value)
}

// GetMapToAccount gets the value of MapToAccount for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyMapToAccount() (value bool, err error) {
	retValue, err := instance.GetProperty("MapToAccount")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSelectionCriteria sets the value of SelectionCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertySelectionCriteria(value bool) (err error) {
	return instance.SetProperty("SelectionCriteria", value)
}

// GetSelectionCriteria gets the value of SelectionCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertySelectionCriteria() (value bool, err error) {
	retValue, err := instance.GetProperty("SelectionCriteria")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSigningAlgorithm sets the value of SigningAlgorithm for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertySigningAlgorithm(value uint16) (err error) {
	return instance.SetProperty("SigningAlgorithm", value)
}

// GetSigningAlgorithm gets the value of SigningAlgorithm for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertySigningAlgorithm() (value uint16, err error) {
	retValue, err := instance.GetProperty("SigningAlgorithm")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", value)
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedCA sets the value of TrustedCA for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyTrustedCA(value string) (err error) {
	return instance.SetProperty("TrustedCA", value)
}

// GetTrustedCA gets the value of TrustedCA for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyTrustedCA() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedCA")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedCAType sets the value of TrustedCAType for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyTrustedCAType(value uint16) (err error) {
	return instance.SetProperty("TrustedCAType", value)
}

// GetTrustedCAType gets the value of TrustedCAType for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyTrustedCAType() (value uint16, err error) {
	retValue, err := instance.GetProperty("TrustedCAType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidationCriteria sets the value of ValidationCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) SetPropertyValidationCriteria(value bool) (err error) {
	return instance.SetProperty("ValidationCriteria", value)
}

// GetValidationCriteria gets the value of ValidationCriteria for the instance
func (instance *MSFT_NetIKECertAuthProposal) GetPropertyValidationCriteria() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidationCriteria")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
