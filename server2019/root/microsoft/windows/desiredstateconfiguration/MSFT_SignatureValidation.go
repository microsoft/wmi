// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

// MSFT_SignatureValidation struct
type MSFT_SignatureValidation struct {
	OMI_MetaConfigurationResource

	//
	SignedItemType []string

	//
	TrustedStorePath string
}

// SetSignedItemType sets the value of SignedItemType for the instance
func (instance *MSFT_SignatureValidation) SetPropertySignedItemType(value []string) (err error) {
	return instance.SetProperty("SignedItemType", value)
}

// GetSignedItemType gets the value of SignedItemType for the instance
func (instance *MSFT_SignatureValidation) GetPropertySignedItemType() (value []string, err error) {
	retValue, err := instance.GetProperty("SignedItemType")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTrustedStorePath sets the value of TrustedStorePath for the instance
func (instance *MSFT_SignatureValidation) SetPropertyTrustedStorePath(value string) (err error) {
	return instance.SetProperty("TrustedStorePath", value)
}

// GetTrustedStorePath gets the value of TrustedStorePath for the instance
func (instance *MSFT_SignatureValidation) GetPropertyTrustedStorePath() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedStorePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
