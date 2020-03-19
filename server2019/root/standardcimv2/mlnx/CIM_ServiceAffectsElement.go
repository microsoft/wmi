// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ServiceAffectsElement struct
type CIM_ServiceAffectsElement struct {
	*cim.WmiInstance

	//
	AffectedElement CIM_ManagedElement

	//
	AffectingElement CIM_Service

	//
	ElementEffects []ServiceAffectsElement_ElementEffects

	//
	OtherElementEffectsDescriptions []string
}

func NewCIM_ServiceAffectsElementEx1(instance *cim.WmiInstance) (newInstance *CIM_ServiceAffectsElement, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAffectsElement{
		WmiInstance: tmp,
	}
	return
}

func NewCIM_ServiceAffectsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ServiceAffectsElement, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ServiceAffectsElement{
		WmiInstance: tmp,
	}
	return
}

// SetAffectedElement sets the value of AffectedElement for the instance
func (instance *CIM_ServiceAffectsElement) SetPropertyAffectedElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("AffectedElement", value)
}

// GetAffectedElement gets the value of AffectedElement for the instance
func (instance *CIM_ServiceAffectsElement) GetPropertyAffectedElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("AffectedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAffectingElement sets the value of AffectingElement for the instance
func (instance *CIM_ServiceAffectsElement) SetPropertyAffectingElement(value CIM_Service) (err error) {
	return instance.SetProperty("AffectingElement", value)
}

// GetAffectingElement gets the value of AffectingElement for the instance
func (instance *CIM_ServiceAffectsElement) GetPropertyAffectingElement() (value CIM_Service, err error) {
	retValue, err := instance.GetProperty("AffectingElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Service)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetElementEffects sets the value of ElementEffects for the instance
func (instance *CIM_ServiceAffectsElement) SetPropertyElementEffects(value []ServiceAffectsElement_ElementEffects) (err error) {
	return instance.SetProperty("ElementEffects", value)
}

// GetElementEffects gets the value of ElementEffects for the instance
func (instance *CIM_ServiceAffectsElement) GetPropertyElementEffects() (value []ServiceAffectsElement_ElementEffects, err error) {
	retValue, err := instance.GetProperty("ElementEffects")
	if err != nil {
		return
	}
	value, ok := retValue.([]ServiceAffectsElement_ElementEffects)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherElementEffectsDescriptions sets the value of OtherElementEffectsDescriptions for the instance
func (instance *CIM_ServiceAffectsElement) SetPropertyOtherElementEffectsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherElementEffectsDescriptions", value)
}

// GetOtherElementEffectsDescriptions gets the value of OtherElementEffectsDescriptions for the instance
func (instance *CIM_ServiceAffectsElement) GetPropertyOtherElementEffectsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherElementEffectsDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
