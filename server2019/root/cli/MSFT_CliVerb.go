// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CliVerb struct
type MSFT_CliVerb struct {
	*cim.WmiInstance

	//
	Derivation string

	//
	Description string

	//
	Name string

	//
	Parameters []MSFT_CliParam

	//
	Qualifiers []MSFT_CliQualifier

	//
	Usage string

	//
	VerbType uint32
}

func NewMSFT_CliVerbEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliVerb, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliVerb{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliVerbEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliVerb, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliVerb{
		WmiInstance: tmp,
	}
	return
}

// SetDerivation sets the value of Derivation for the instance
func (instance *MSFT_CliVerb) SetPropertyDerivation(value string) (err error) {
	return instance.SetProperty("Derivation", (value))
}

// GetDerivation gets the value of Derivation for the instance
func (instance *MSFT_CliVerb) GetPropertyDerivation() (value string, err error) {
	retValue, err := instance.GetProperty("Derivation")
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliVerb) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliVerb) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_CliVerb) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliVerb) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetParameters sets the value of Parameters for the instance
func (instance *MSFT_CliVerb) SetPropertyParameters(value []MSFT_CliParam) (err error) {
	return instance.SetProperty("Parameters", (value))
}

// GetParameters gets the value of Parameters for the instance
func (instance *MSFT_CliVerb) GetPropertyParameters() (value []MSFT_CliParam, err error) {
	retValue, err := instance.GetProperty("Parameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_CliParam)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_CliParam is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_CliParam(valuetmp))
	}

	return
}

// SetQualifiers sets the value of Qualifiers for the instance
func (instance *MSFT_CliVerb) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) {
	return instance.SetProperty("Qualifiers", (value))
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliVerb) GetPropertyQualifiers() (value []MSFT_CliQualifier, err error) {
	retValue, err := instance.GetProperty("Qualifiers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_CliQualifier)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_CliQualifier is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_CliQualifier(valuetmp))
	}

	return
}

// SetUsage sets the value of Usage for the instance
func (instance *MSFT_CliVerb) SetPropertyUsage(value string) (err error) {
	return instance.SetProperty("Usage", (value))
}

// GetUsage gets the value of Usage for the instance
func (instance *MSFT_CliVerb) GetPropertyUsage() (value string, err error) {
	retValue, err := instance.GetProperty("Usage")
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

// SetVerbType sets the value of VerbType for the instance
func (instance *MSFT_CliVerb) SetPropertyVerbType(value uint32) (err error) {
	return instance.SetProperty("VerbType", (value))
}

// GetVerbType gets the value of VerbType for the instance
func (instance *MSFT_CliVerb) GetPropertyVerbType() (value uint32, err error) {
	retValue, err := instance.GetProperty("VerbType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
