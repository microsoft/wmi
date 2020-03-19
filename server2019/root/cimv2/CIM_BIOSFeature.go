// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_BIOSFeature struct
type CIM_BIOSFeature struct {
	*CIM_SoftwareFeature

	//
	CharacteristicDescriptions []string

	//
	Characteristics []uint16
}

func NewCIM_BIOSFeatureEx1(instance *cim.WmiInstance) (newInstance *CIM_BIOSFeature, err error) {
	tmp, err := NewCIM_SoftwareFeatureEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSFeature{
		CIM_SoftwareFeature: tmp,
	}
	return
}

func NewCIM_BIOSFeatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BIOSFeature, err error) {
	tmp, err := NewCIM_SoftwareFeatureEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BIOSFeature{
		CIM_SoftwareFeature: tmp,
	}
	return
}

// SetCharacteristicDescriptions sets the value of CharacteristicDescriptions for the instance
func (instance *CIM_BIOSFeature) SetPropertyCharacteristicDescriptions(value []string) (err error) {
	return instance.SetProperty("CharacteristicDescriptions", value)
}

// GetCharacteristicDescriptions gets the value of CharacteristicDescriptions for the instance
func (instance *CIM_BIOSFeature) GetPropertyCharacteristicDescriptions() (value []string, err error) {
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
func (instance *CIM_BIOSFeature) SetPropertyCharacteristics(value []uint16) (err error) {
	return instance.SetProperty("Characteristics", value)
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *CIM_BIOSFeature) GetPropertyCharacteristics() (value []uint16, err error) {
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
