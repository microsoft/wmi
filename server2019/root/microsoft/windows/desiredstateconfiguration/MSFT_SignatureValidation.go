// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
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
	return instance.SetProperty("SignedItemType", (value))
}

// GetSignedItemType gets the value of SignedItemType for the instance
func (instance *MSFT_SignatureValidation) GetPropertySignedItemType() (value []string, err error) {
	retValue, err := instance.GetProperty("SignedItemType")
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

// SetTrustedStorePath sets the value of TrustedStorePath for the instance
func (instance *MSFT_SignatureValidation) SetPropertyTrustedStorePath(value string) (err error) {
	return instance.SetProperty("TrustedStorePath", (value))
}

// GetTrustedStorePath gets the value of TrustedStorePath for the instance
func (instance *MSFT_SignatureValidation) GetPropertyTrustedStorePath() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedStorePath")
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
