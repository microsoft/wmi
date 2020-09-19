// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Hgs
//////////////////////////////////////////////
package hgs

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_HgsClientConfiguration struct
type MSFT_HgsClientConfiguration struct {
	*cim.WmiInstance

	//
	AttestationOperationMode uint16

	//
	AttestationServerUrl string

	//
	AttestationStatus uint16

	//
	AttestationSubstatus uint64

	//
	FallbackAttestationServerUrl []string

	//
	FallbackKeyProtectionServerUrl []string

	//
	IsHostGuarded bool

	//
	KeyProtectionServerUrl string

	//
	LastAttestationServerUrl string

	//
	LastKeyProtectionServerUrl string

	//
	Mode uint16
}

func NewMSFT_HgsClientConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_HgsClientConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsClientConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_HgsClientConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_HgsClientConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_HgsClientConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAttestationOperationMode sets the value of AttestationOperationMode for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyAttestationOperationMode(value uint16) (err error) {
	return instance.SetProperty("AttestationOperationMode", (value))
}

// GetAttestationOperationMode gets the value of AttestationOperationMode for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyAttestationOperationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttestationOperationMode")
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

// SetAttestationServerUrl sets the value of AttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyAttestationServerUrl(value string) (err error) {
	return instance.SetProperty("AttestationServerUrl", (value))
}

// GetAttestationServerUrl gets the value of AttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyAttestationServerUrl() (value string, err error) {
	retValue, err := instance.GetProperty("AttestationServerUrl")
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

// SetAttestationStatus sets the value of AttestationStatus for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyAttestationStatus(value uint16) (err error) {
	return instance.SetProperty("AttestationStatus", (value))
}

// GetAttestationStatus gets the value of AttestationStatus for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyAttestationStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttestationStatus")
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

// SetAttestationSubstatus sets the value of AttestationSubstatus for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyAttestationSubstatus(value uint64) (err error) {
	return instance.SetProperty("AttestationSubstatus", (value))
}

// GetAttestationSubstatus gets the value of AttestationSubstatus for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyAttestationSubstatus() (value uint64, err error) {
	retValue, err := instance.GetProperty("AttestationSubstatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFallbackAttestationServerUrl sets the value of FallbackAttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyFallbackAttestationServerUrl(value []string) (err error) {
	return instance.SetProperty("FallbackAttestationServerUrl", (value))
}

// GetFallbackAttestationServerUrl gets the value of FallbackAttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyFallbackAttestationServerUrl() (value []string, err error) {
	retValue, err := instance.GetProperty("FallbackAttestationServerUrl")
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

// SetFallbackKeyProtectionServerUrl sets the value of FallbackKeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyFallbackKeyProtectionServerUrl(value []string) (err error) {
	return instance.SetProperty("FallbackKeyProtectionServerUrl", (value))
}

// GetFallbackKeyProtectionServerUrl gets the value of FallbackKeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyFallbackKeyProtectionServerUrl() (value []string, err error) {
	retValue, err := instance.GetProperty("FallbackKeyProtectionServerUrl")
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

// SetIsHostGuarded sets the value of IsHostGuarded for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyIsHostGuarded(value bool) (err error) {
	return instance.SetProperty("IsHostGuarded", (value))
}

// GetIsHostGuarded gets the value of IsHostGuarded for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyIsHostGuarded() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHostGuarded")
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

// SetKeyProtectionServerUrl sets the value of KeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyKeyProtectionServerUrl(value string) (err error) {
	return instance.SetProperty("KeyProtectionServerUrl", (value))
}

// GetKeyProtectionServerUrl gets the value of KeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyKeyProtectionServerUrl() (value string, err error) {
	retValue, err := instance.GetProperty("KeyProtectionServerUrl")
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

// SetLastAttestationServerUrl sets the value of LastAttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyLastAttestationServerUrl(value string) (err error) {
	return instance.SetProperty("LastAttestationServerUrl", (value))
}

// GetLastAttestationServerUrl gets the value of LastAttestationServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyLastAttestationServerUrl() (value string, err error) {
	retValue, err := instance.GetProperty("LastAttestationServerUrl")
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

// SetLastKeyProtectionServerUrl sets the value of LastKeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyLastKeyProtectionServerUrl(value string) (err error) {
	return instance.SetProperty("LastKeyProtectionServerUrl", (value))
}

// GetLastKeyProtectionServerUrl gets the value of LastKeyProtectionServerUrl for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyLastKeyProtectionServerUrl() (value string, err error) {
	retValue, err := instance.GetProperty("LastKeyProtectionServerUrl")
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

// SetMode sets the value of Mode for the instance
func (instance *MSFT_HgsClientConfiguration) SetPropertyMode(value uint16) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *MSFT_HgsClientConfiguration) GetPropertyMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("Mode")
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

//

// <param name="cmdletOutput" type="MSFT_HgsClientConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsClientConfiguration) Get( /* OUT */ cmdletOutput MSFT_HgsClientConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AttestationServerUrl" type="string "></param>

// <param name="AttestationOperationMode" type="uint16 "></param>
// <param name="AttestationStatus" type="uint16 "></param>
// <param name="AttestationSubstatus" type="uint64 "></param>
// <param name="IsHostGuarded" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsClientConfiguration) IsHostTrusted( /* IN */ AttestationServerUrl string,
	/* OUT */ IsHostGuarded bool,
	/* OUT */ AttestationOperationMode uint16,
	/* OUT */ AttestationStatus uint16,
	/* OUT */ AttestationSubstatus uint64) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsHostTrusted", AttestationServerUrl)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EnableLocalMode" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_HgsClientConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsClientConfiguration) SetByChangeToLocalMode( /* IN */ EnableLocalMode bool,
	/* OUT */ cmdletOutput MSFT_HgsClientConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetByChangeToLocalMode", EnableLocalMode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AttestationServerUrl" type="string "></param>
// <param name="FallbackAttestationServerUrl" type="string []"></param>
// <param name="FallbackKeyProtectionServerUrl" type="string []"></param>
// <param name="KeyProtectionServerUrl" type="string "></param>

// <param name="cmdletOutput" type="MSFT_HgsClientConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_HgsClientConfiguration) SetBySecureHostingServiceMode( /* IN */ KeyProtectionServerUrl string,
	/* IN */ AttestationServerUrl string,
	/* OUT */ cmdletOutput MSFT_HgsClientConfiguration,
	/* OPTIONAL IN */ FallbackKeyProtectionServerUrl []string,
	/* OPTIONAL IN */ FallbackAttestationServerUrl []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetBySecureHostingServiceMode", KeyProtectionServerUrl, AttestationServerUrl, FallbackKeyProtectionServerUrl, FallbackAttestationServerUrl)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
