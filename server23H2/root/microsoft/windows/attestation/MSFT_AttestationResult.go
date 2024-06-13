// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Attestation
//////////////////////////////////////////////
package attestation

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_AttestationResult struct
type MSFT_AttestationResult struct {
	*cim.WmiInstance

	//
	AttestationStatus uint16

	//
	AttestationSubstatus uint64

	//
	Data []uint8

	//
	Type uint16

	//
	Url string
}

func NewMSFT_AttestationResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_AttestationResult, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_AttestationResult{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_AttestationResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_AttestationResult, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_AttestationResult{
		WmiInstance: tmp,
	}
	return
}

// SetAttestationStatus sets the value of AttestationStatus for the instance
func (instance *MSFT_AttestationResult) SetPropertyAttestationStatus(value uint16) (err error) {
	return instance.SetProperty("AttestationStatus", (value))
}

// GetAttestationStatus gets the value of AttestationStatus for the instance
func (instance *MSFT_AttestationResult) GetPropertyAttestationStatus() (value uint16, err error) {
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
func (instance *MSFT_AttestationResult) SetPropertyAttestationSubstatus(value uint64) (err error) {
	return instance.SetProperty("AttestationSubstatus", (value))
}

// GetAttestationSubstatus gets the value of AttestationSubstatus for the instance
func (instance *MSFT_AttestationResult) GetPropertyAttestationSubstatus() (value uint64, err error) {
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

// SetData sets the value of Data for the instance
func (instance *MSFT_AttestationResult) SetPropertyData(value []uint8) (err error) {
	return instance.SetProperty("Data", (value))
}

// GetData gets the value of Data for the instance
func (instance *MSFT_AttestationResult) GetPropertyData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Data")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_AttestationResult) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_AttestationResult) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
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

// SetUrl sets the value of Url for the instance
func (instance *MSFT_AttestationResult) SetPropertyUrl(value string) (err error) {
	return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *MSFT_AttestationResult) GetPropertyUrl() (value string, err error) {
	retValue, err := instance.GetProperty("Url")
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
