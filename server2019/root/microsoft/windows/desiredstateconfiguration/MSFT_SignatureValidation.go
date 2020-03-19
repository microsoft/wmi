// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SignatureValidation struct
type MSFT_SignatureValidation struct {
	*OMI_MetaConfigurationResource

	//
	SignedItemType []string

	//
	TrustedStorePath string
}

func NewMSFT_SignatureValidationEx1(instance *cim.WmiInstance) (newInstance *MSFT_SignatureValidation, err error) {
	tmp, err := NewOMI_MetaConfigurationResourceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_SignatureValidation{
		OMI_MetaConfigurationResource: tmp,
	}
	return
}

func NewMSFT_SignatureValidationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SignatureValidation, err error) {
	tmp, err := NewOMI_MetaConfigurationResourceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SignatureValidation{
		OMI_MetaConfigurationResource: tmp,
	}
	return
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
