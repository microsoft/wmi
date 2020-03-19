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

// MSFT_NetAdapterAdvancedPropertySettingData struct
type MSFT_NetAdapterAdvancedPropertySettingData struct {
	*MSFT_NetAdapterSettingData

	//
	DefaultDisplayValue string

	//
	DefaultRegistryValue string

	//
	DisplayName string

	//
	DisplayParameterType uint32

	//
	DisplayValue string

	//
	NumericParameterBaseValue string

	//
	NumericParameterMaxValue string

	//
	NumericParameterMinValue string

	//
	NumericParameterStepValue string

	//
	Optional bool

	//
	RegistryDataType uint32

	//
	RegistryKeyword string

	//
	RegistryValue []string

	//
	ValidDisplayValues []string

	//
	ValidRegistryValues []string
}

func NewMSFT_NetAdapterAdvancedPropertySettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterAdvancedPropertySettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterAdvancedPropertySettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterAdvancedPropertySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterAdvancedPropertySettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterAdvancedPropertySettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetDefaultDisplayValue sets the value of DefaultDisplayValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyDefaultDisplayValue(value string) (err error) {
	return instance.SetProperty("DefaultDisplayValue", value)
}

// GetDefaultDisplayValue gets the value of DefaultDisplayValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyDefaultDisplayValue() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultDisplayValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultRegistryValue sets the value of DefaultRegistryValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyDefaultRegistryValue(value string) (err error) {
	return instance.SetProperty("DefaultRegistryValue", value)
}

// GetDefaultRegistryValue gets the value of DefaultRegistryValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyDefaultRegistryValue() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultRegistryValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayParameterType sets the value of DisplayParameterType for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyDisplayParameterType(value uint32) (err error) {
	return instance.SetProperty("DisplayParameterType", value)
}

// GetDisplayParameterType gets the value of DisplayParameterType for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyDisplayParameterType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisplayParameterType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayValue sets the value of DisplayValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyDisplayValue(value string) (err error) {
	return instance.SetProperty("DisplayValue", value)
}

// GetDisplayValue gets the value of DisplayValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyDisplayValue() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericParameterBaseValue sets the value of NumericParameterBaseValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyNumericParameterBaseValue(value string) (err error) {
	return instance.SetProperty("NumericParameterBaseValue", value)
}

// GetNumericParameterBaseValue gets the value of NumericParameterBaseValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyNumericParameterBaseValue() (value string, err error) {
	retValue, err := instance.GetProperty("NumericParameterBaseValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericParameterMaxValue sets the value of NumericParameterMaxValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyNumericParameterMaxValue(value string) (err error) {
	return instance.SetProperty("NumericParameterMaxValue", value)
}

// GetNumericParameterMaxValue gets the value of NumericParameterMaxValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyNumericParameterMaxValue() (value string, err error) {
	retValue, err := instance.GetProperty("NumericParameterMaxValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericParameterMinValue sets the value of NumericParameterMinValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyNumericParameterMinValue(value string) (err error) {
	return instance.SetProperty("NumericParameterMinValue", value)
}

// GetNumericParameterMinValue gets the value of NumericParameterMinValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyNumericParameterMinValue() (value string, err error) {
	retValue, err := instance.GetProperty("NumericParameterMinValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumericParameterStepValue sets the value of NumericParameterStepValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyNumericParameterStepValue(value string) (err error) {
	return instance.SetProperty("NumericParameterStepValue", value)
}

// GetNumericParameterStepValue gets the value of NumericParameterStepValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyNumericParameterStepValue() (value string, err error) {
	retValue, err := instance.GetProperty("NumericParameterStepValue")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOptional sets the value of Optional for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyOptional(value bool) (err error) {
	return instance.SetProperty("Optional", value)
}

// GetOptional gets the value of Optional for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyOptional() (value bool, err error) {
	retValue, err := instance.GetProperty("Optional")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryDataType sets the value of RegistryDataType for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyRegistryDataType(value uint32) (err error) {
	return instance.SetProperty("RegistryDataType", value)
}

// GetRegistryDataType gets the value of RegistryDataType for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyRegistryDataType() (value uint32, err error) {
	retValue, err := instance.GetProperty("RegistryDataType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryKeyword sets the value of RegistryKeyword for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyRegistryKeyword(value string) (err error) {
	return instance.SetProperty("RegistryKeyword", value)
}

// GetRegistryKeyword gets the value of RegistryKeyword for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyRegistryKeyword() (value string, err error) {
	retValue, err := instance.GetProperty("RegistryKeyword")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryValue sets the value of RegistryValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyRegistryValue(value []string) (err error) {
	return instance.SetProperty("RegistryValue", value)
}

// GetRegistryValue gets the value of RegistryValue for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyRegistryValue() (value []string, err error) {
	retValue, err := instance.GetProperty("RegistryValue")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidDisplayValues sets the value of ValidDisplayValues for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyValidDisplayValues(value []string) (err error) {
	return instance.SetProperty("ValidDisplayValues", value)
}

// GetValidDisplayValues gets the value of ValidDisplayValues for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyValidDisplayValues() (value []string, err error) {
	retValue, err := instance.GetProperty("ValidDisplayValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetValidRegistryValues sets the value of ValidRegistryValues for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) SetPropertyValidRegistryValues(value []string) (err error) {
	return instance.SetProperty("ValidRegistryValues", value)
}

// GetValidRegistryValues gets the value of ValidRegistryValues for the instance
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) GetPropertyValidRegistryValues() (value []string, err error) {
	retValue, err := instance.GetProperty("ValidRegistryValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CmdletOutput" type="MSFT_NetAdapterAdvancedPropertySettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterAdvancedPropertySettingData) Reset( /* OUT */ CmdletOutput MSFT_NetAdapterAdvancedPropertySettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
