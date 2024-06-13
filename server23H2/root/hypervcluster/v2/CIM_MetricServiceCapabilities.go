// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_MetricServiceCapabilities struct
type CIM_MetricServiceCapabilities struct {
	*CIM_EnabledLogicalElementCapabilities

	// ControllableManagedElements identifies the instances of CIM_ManagedElement that can be controlled by the associated CIM_MetricService instance. Each value shall be formatted as a WBEM URI defined according to DSP0207 identifying an instance of CIM_ManagedElement If a value corresponding to an instance of CIM_ManagedElement is included in the ControllableManagedElements property, the associated instance of CIM_MetricService shall support enabling and/or disabling at least one metric defined for the CIM_ManagedElement instance.
	ControllableManagedElements []string

	// ControllableMetrics identifies the instances of CIM_BaseMetricDefinition that can be controlled by the associated CIM_MetricService instance. Each string value shall be formatted as a WBEM URI defined as in accordance with DSP0207 that identifies an instance of CIM_BaseMetricDefinition. An instance of CIM_BaseMetricDefinition shall not be identified by a value of the ControllableMetrics property unless it is associated through CIM_ServiceAffectsElement to the associated instance of CIM_MetricService. If a value corresponding to an instance of CIM_BaseMetricDefinition is included in the ControllableMetrics property, the associated instance of CIM_MetricService shall support enabling and/or disabling at least one metric defined by the CIM_BaseMetricDefinition instance.
	ControllableMetrics []string

	// ManagedElementControlTypes identifies the type of control supported by the associated CIM_MetricService instance for the CIM_ManagedElement identified by the value at the same array index in the ControllableManagedElements property. A value of 2 "Discrete" shall indicate that individual metrics controlled by the associated instance of CIM_MetricService may be enabled and or disabled for the instance of CIM_ManagedElement identified at the corresponding array index of ControllableManagedElements.A value of 3 "Bulk" shall indicate that all metrics controlled by the associated instance of CIM_MetricService may be enabled and or disabled for the instance of CIM_ManagedElement identified at the corresponding array index of ControllableManagedElements. A value of 4 "Both" shall indicate that all metrics controlled by the associated instance of CIM_MetricService may be enabled and or disabled with a single operation or individually for the instance of CIM_ManagedElement identified by the value at the same array index of ControllableManagedElements.
	ManagedElementControlTypes []MetricServiceCapabilities_ManagedElementControlTypes

	// MetricControlTypes identifies the type of control supported by the associated CIM_MetricService instance for the CIM_BaseMetricDefinition identified by the value at the same array index in the ControllableMetrics property. A value of 2 "Discrete" shall indicate that individual metrics defined by the instance of CIM_BaseMetricDefinition identified at the corresponding array index of ControllableMetrics may be enabled and or disabled by the associated instance of CIM_MetricService.A value of 3 "Bulk" shall indicate that all metrics defined by the instance of CIM_BaseMetricDefinition identified by the value at the same array index of ControllableMetrics may be enabled and or disabled with a single operation. A value of 4 "Both" shall indicate that all metrics defined by the instance of CIM_BaseMetricDefinition identified by the value at the same array index of ControllableMetrics may be enabled and or disabled individually or as a single operation.
	MetricsControlTypes []MetricServiceCapabilities_MetricsControlTypes

	// Each enumeration corresponds to support for the like-named method of the MetricService.
	SupportedMethods []MetricServiceCapabilities_SupportedMethods
}

func NewCIM_MetricServiceCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *CIM_MetricServiceCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricServiceCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}

func NewCIM_MetricServiceCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_MetricServiceCapabilities, err error) {
	tmp, err := NewCIM_EnabledLogicalElementCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_MetricServiceCapabilities{
		CIM_EnabledLogicalElementCapabilities: tmp,
	}
	return
}

// SetControllableManagedElements sets the value of ControllableManagedElements for the instance
func (instance *CIM_MetricServiceCapabilities) SetPropertyControllableManagedElements(value []string) (err error) {
	return instance.SetProperty("ControllableManagedElements", (value))
}

// GetControllableManagedElements gets the value of ControllableManagedElements for the instance
func (instance *CIM_MetricServiceCapabilities) GetPropertyControllableManagedElements() (value []string, err error) {
	retValue, err := instance.GetProperty("ControllableManagedElements")
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

// SetControllableMetrics sets the value of ControllableMetrics for the instance
func (instance *CIM_MetricServiceCapabilities) SetPropertyControllableMetrics(value []string) (err error) {
	return instance.SetProperty("ControllableMetrics", (value))
}

// GetControllableMetrics gets the value of ControllableMetrics for the instance
func (instance *CIM_MetricServiceCapabilities) GetPropertyControllableMetrics() (value []string, err error) {
	retValue, err := instance.GetProperty("ControllableMetrics")
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

// SetManagedElementControlTypes sets the value of ManagedElementControlTypes for the instance
func (instance *CIM_MetricServiceCapabilities) SetPropertyManagedElementControlTypes(value []MetricServiceCapabilities_ManagedElementControlTypes) (err error) {
	return instance.SetProperty("ManagedElementControlTypes", (value))
}

// GetManagedElementControlTypes gets the value of ManagedElementControlTypes for the instance
func (instance *CIM_MetricServiceCapabilities) GetPropertyManagedElementControlTypes() (value []MetricServiceCapabilities_ManagedElementControlTypes, err error) {
	retValue, err := instance.GetProperty("ManagedElementControlTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MetricServiceCapabilities_ManagedElementControlTypes(valuetmp))
	}

	return
}

// SetMetricsControlTypes sets the value of MetricsControlTypes for the instance
func (instance *CIM_MetricServiceCapabilities) SetPropertyMetricsControlTypes(value []MetricServiceCapabilities_MetricsControlTypes) (err error) {
	return instance.SetProperty("MetricsControlTypes", (value))
}

// GetMetricsControlTypes gets the value of MetricsControlTypes for the instance
func (instance *CIM_MetricServiceCapabilities) GetPropertyMetricsControlTypes() (value []MetricServiceCapabilities_MetricsControlTypes, err error) {
	retValue, err := instance.GetProperty("MetricsControlTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MetricServiceCapabilities_MetricsControlTypes(valuetmp))
	}

	return
}

// SetSupportedMethods sets the value of SupportedMethods for the instance
func (instance *CIM_MetricServiceCapabilities) SetPropertySupportedMethods(value []MetricServiceCapabilities_SupportedMethods) (err error) {
	return instance.SetProperty("SupportedMethods", (value))
}

// GetSupportedMethods gets the value of SupportedMethods for the instance
func (instance *CIM_MetricServiceCapabilities) GetPropertySupportedMethods() (value []MetricServiceCapabilities_SupportedMethods, err error) {
	retValue, err := instance.GetProperty("SupportedMethods")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MetricServiceCapabilities_SupportedMethods(valuetmp))
	}

	return
}
