// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_VideoBIOSFeature struct
type CIM_VideoBIOSFeature struct {
	CIM_SoftwareFeature

	//
	CharacteristicDescriptions []string

	//
	Characteristics []uint16
}

// SetCharacteristicDescriptions sets the value of CharacteristicDescriptions for the instance
func (instance *CIM_VideoBIOSFeature) SetPropertyCharacteristicDescriptions(value []string) (err error) {
	return instance.SetProperty("CharacteristicDescriptions", value)
}

// GetCharacteristicDescriptions gets the value of CharacteristicDescriptions for the instance
func (instance *CIM_VideoBIOSFeature) GetPropertyCharacteristicDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CharacteristicDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *CIM_VideoBIOSFeature) SetPropertyCharacteristics(value []uint16) (err error) {
	return instance.SetProperty("Characteristics", value)
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_VideoBIOSFeature) GetPropertyCharacteristics() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Characteristics")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
