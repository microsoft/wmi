// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_BaseMetricValue struct
type CIM_BaseMetricValue struct {
	*CIM_ManagedElement

	// If present, specifies one BreakdownDimension from the BreakdownDimensions array defined in the associated CIM_ BaseMetricDefinition. This is the dimension along which this set of metric values is broken down. For a description of the concept, see the class CIM_BaseMetricDefinition.
	BreakdownDimension string

	// Defines a value of the BreakdownDimension property defined for this metric value instance. For instance, if the BreakdownDimension is "TransactionName", this property could name the actual transaction to which this particular metric value applies.
	BreakdownValue string

	// Property that represents the time duration over which this metric value is valid. This property should not exist for timestamps that apply only to a point in time but should be defined for values that are considered valid for a certain time period (ex. sampling). If the "Duration" property exists and is nonNull, the TimeStamp is to be considered the end of the interval.
	Duration string

	// A descriptive name for the element to which the metric value belongs (i.e., the measured element). This property is required by behavior if there is no association defined to a ManagedElement, but may be used in other cases to provide supplemental information. This allows metrics to be captured independently of any ManagedElement. An example is where a metric value belongs to a combination of elements, such as the input and output ports of the traffic in a switch. If there are multiple ManagedElements associated with the metric value, then usually there is one that naturally belongs to the metric value and that one should be used to create the supplemental information. The property is not meant to be used as a foreign key to search on the measured element. Instead, the association to the ManagedElement should be used.
	MeasuredElementName string

	// The key of the BaseMetricDefinition instance for this CIM_BaseMetricValue instance value.
	MetricDefinitionId string

	// The value of the metric represented as a string. Its original data type is specified in CIM_BaseMetricDefinition.
	MetricValue string

	// Identifies the time when the value of a metric instance is computed. Note that this is different from the time when the instance is created. For a given CIM_BaseMetricValue instance, the TimeStamp changes whenever a new measurement snapshot is taken if Volatile is true. A managmenet application may establish a time series of metric data by retrieving the instances of CIM_BaseMetricValue and sorting them according to their TimeStamp.
	TimeStamp string

	// If true, Volatile indicates that the value for the next point in time may use the same object and just change its properties (such as the value or timestamp). If false, the existing objects remain unchanged and a new object is created for the new point in time.
	Volatile bool
}

func NewCIM_BaseMetricValueEx1(instance *cim.WmiInstance) (newInstance *CIM_BaseMetricValue, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BaseMetricValue{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_BaseMetricValueEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BaseMetricValue, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BaseMetricValue{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetBreakdownDimension sets the value of BreakdownDimension for the instance
func (instance *CIM_BaseMetricValue) SetPropertyBreakdownDimension(value string) (err error) {
	return instance.SetProperty("BreakdownDimension", (value))
}

// GetBreakdownDimension gets the value of BreakdownDimension for the instance
func (instance *CIM_BaseMetricValue) GetPropertyBreakdownDimension() (value string, err error) {
	retValue, err := instance.GetProperty("BreakdownDimension")
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

// SetBreakdownValue sets the value of BreakdownValue for the instance
func (instance *CIM_BaseMetricValue) SetPropertyBreakdownValue(value string) (err error) {
	return instance.SetProperty("BreakdownValue", (value))
}

// GetBreakdownValue gets the value of BreakdownValue for the instance
func (instance *CIM_BaseMetricValue) GetPropertyBreakdownValue() (value string, err error) {
	retValue, err := instance.GetProperty("BreakdownValue")
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

// SetDuration sets the value of Duration for the instance
func (instance *CIM_BaseMetricValue) SetPropertyDuration(value string) (err error) {
	return instance.SetProperty("Duration", (value))
}

// GetDuration gets the value of Duration for the instance
func (instance *CIM_BaseMetricValue) GetPropertyDuration() (value string, err error) {
	retValue, err := instance.GetProperty("Duration")
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

// SetMeasuredElementName sets the value of MeasuredElementName for the instance
func (instance *CIM_BaseMetricValue) SetPropertyMeasuredElementName(value string) (err error) {
	return instance.SetProperty("MeasuredElementName", (value))
}

// GetMeasuredElementName gets the value of MeasuredElementName for the instance
func (instance *CIM_BaseMetricValue) GetPropertyMeasuredElementName() (value string, err error) {
	retValue, err := instance.GetProperty("MeasuredElementName")
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

// SetMetricDefinitionId sets the value of MetricDefinitionId for the instance
func (instance *CIM_BaseMetricValue) SetPropertyMetricDefinitionId(value string) (err error) {
	return instance.SetProperty("MetricDefinitionId", (value))
}

// GetMetricDefinitionId gets the value of MetricDefinitionId for the instance
func (instance *CIM_BaseMetricValue) GetPropertyMetricDefinitionId() (value string, err error) {
	retValue, err := instance.GetProperty("MetricDefinitionId")
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

// SetMetricValue sets the value of MetricValue for the instance
func (instance *CIM_BaseMetricValue) SetPropertyMetricValue(value string) (err error) {
	return instance.SetProperty("MetricValue", (value))
}

// GetMetricValue gets the value of MetricValue for the instance
func (instance *CIM_BaseMetricValue) GetPropertyMetricValue() (value string, err error) {
	retValue, err := instance.GetProperty("MetricValue")
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

// SetTimeStamp sets the value of TimeStamp for the instance
func (instance *CIM_BaseMetricValue) SetPropertyTimeStamp(value string) (err error) {
	return instance.SetProperty("TimeStamp", (value))
}

// GetTimeStamp gets the value of TimeStamp for the instance
func (instance *CIM_BaseMetricValue) GetPropertyTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("TimeStamp")
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

// Setvolatile sets the value of volatile for the instance
func (instance *CIM_BaseMetricValue) SetPropertyvolatile(value bool) (err error) {
	return instance.SetProperty("volatile", (value))
}

// Getvolatile gets the value of volatile for the instance
func (instance *CIM_BaseMetricValue) GetPropertyvolatile() (value bool, err error) {
	retValue, err := instance.GetProperty("volatile")
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
