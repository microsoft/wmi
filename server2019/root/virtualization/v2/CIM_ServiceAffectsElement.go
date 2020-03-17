// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_ServiceAffectsElement struct
type CIM_ServiceAffectsElement struct {
	cim.WmiInstance

	// The Managed Element that is affected by the Service.
	AffectedElement CIM_ManagedElement

	// The Service that is affecting the ManagedElement.
	AffectingElement CIM_Service

	// An enumeration that describes the effect on the ManagedElement. This array corresponds to the OtherElementEffectsDescriptions array, where the latter provides details that are related to the high-level effects enumerated by this property. Additional detail is required if the ElementEffects array contains the value 1 (Other). The values are defined as follows:
	///- Exclusive Use (2): No other Service may have this association to the element.
	///- Performance Impact (3): Deprecated in favor of "Consumes", "Enhances Performance", or "Degrades Performance". Execution of the Service may enhance or degrade the performance of the element. This may be as a side-effect of execution or as an intended consequence of methods provided by the Service.
	///- Element Integrity (4): Deprecated in favor of "Consumes", "Enhances Integrity", or "Degrades Integrity". Execution of the Service may enhance or degrade the integrity of the element. This may be as a side-effect of execution or as an intended consequence of methods provided by the Service.
	///- Manages (5): The Service manages the element.
	///- Consumes (6): Execution of the Service consumes some or all of the associated element as a consequence of running the Service. For example, the Service may consume CPU cycles, which may affect performance, or Storage which may affect both performance and integrity. (For instance, the lack of free storage can degrade integrity by reducing the ability to save state. ) "Consumes" may be used alone or in conjunction with other values, in particular, "Degrades Performance" and "Degrades Integrity".
	///"Manages" and not "Consumes" should be used to reflect allocation services that may be provided by a Service.
	///- Enhances Integrity (7): The Service may enhance integrity of the associated element.
	///- Degrades Integrity (8): The Service may degrade integrity of the associated element.
	///- Enhances Performance (9): The Service may enhance performance of the associated element.
	///- Degrades Performance (10): The Service may degrade performance of the associated element.
	ElementEffects []ServiceAffectsElement_ElementEffects

	// Provides details for the effect at the corresponding array position in ElementEffects. This information is required if ElementEffects contains the value 1 (Other).
	OtherElementEffectsDescriptions []string
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
