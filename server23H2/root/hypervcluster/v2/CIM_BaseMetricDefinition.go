// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_BaseMetricDefinition struct
type CIM_BaseMetricDefinition struct {
	*CIM_ManagedElement

	// Defines one or more strings that can be used to refine (break down) queries against the BaseMetricValues along a certain dimension. An example is a transaction name, allowing the break down of the total value for all transactions into a set of values, one for each transaction name. Other examples might be application system or user group name. The strings are free format and should be meaningful to the end users of the metric data. The strings indicate which break down dimensions are supported for this metric definition, by the underlying instrumentation.
	BreakdownDimensions []string

	// An enumerated value that describes the characteristics of the metric, for purposes of performing calculations. The property can take one of the following values:
	///1="Non-calculable" -> a string. Arithmetic makes no sense.
	///2="Summable" -> It is reasonable to sum this value over many instances of e.g., UnitOfWork, such as the number of files processed in a backup job. For example, if each backup job is a UnitOfWork, and each job backs up 27,000 files on average, then it makes sense to say that 100 backup jobs processed 2,700,000 files.
	///3="Non-summable" -> It does not make sense to sum this value over many instances of UnitOfWork. An example would be a metric that measures the queue length when a job arrives at a server. If each job is a UnitOfWork, and the average queue length when each job arrives is 33, it does not make sense to say that the queue length for 100 jobs is 3300. It does make sense to say that the mean is 33.
	Calculable BaseMetricDefinition_Calculable

	// ChangeType indicates how the metric value changes, in the form of typical combinations of finer grain attributes such as direction change, minimum and maximum values, and wrapping semantics.
	///0="Unknown": The metric designer did not qualify the ChangeType.
	///2="N/A": If the "IsContinuous" property is "false", ChangeType does not make sense and MUST be is set to "N/A".
	///3="Counter": The metric is a counter metric. These have non-negative integer values which increase monotonically until reaching the maximum representable number and then wrap around and start increasing from 0. Such counters, also known as rollover counters, can be used for instance to count the number of network errors or the number of transactions processed. The only way for a client application to keep track of wrap arounds is to retrieve the value of the counter in appropriately short intervals.
	///4="Gauge": The metric is a gauge metric. These have integer or float values that can increase and decrease arbitrarily. A gauge MUST NOT wrap when reaching the minimum or maximum representable number, instead, the value "sticks" at that number. Minimum or maximum values inside of the representable value range at which the metric value "sticks", may or may not be defined.
	///Vendors may extend this property in the vendor reserved range.
	ChangeType BaseMetricDefinition_ChangeType

	// The data type of the metric. For example, "boolean" (value=1) or "datetime" (=3) may be specified. These types represent the datatypes defined for CIM.
	DataType BaseMetricDefinition_DataType

	// GatheringType indicates how the metric values are gathered by the underlying instrumentation. This allows the client application to choose the right metric for the purpose.
	///0="Unknown": Indicates that the GatheringType is not known.
	///2="OnChange": Indicates that the CIM metric values get updated immediately when the values inside of the measured resource change. The values of OnChange metrics truly reflect the current situation within the resource at any time. An example is the number of logged on users that gets updated immediately as users log on and off.
	///3="Periodic": Indicates that the CIM metric values get updated periodically. For instance, to a client application, a metric value applying to the current time will appear constant during each gathering interval, and then jumps to the new value at the end of each gathering interval.
	///4="OnRequest": Indicates that the CIM metric value is determined each time a client application reads it. The values of OnRequest metrics truly return the current situation within the resource if somebody asks for it. However, they do not change "unobserved", and therefore subscribing for value changes of OnRequest metrics is NOT RECOMMENDED.
	GatheringType BaseMetricDefinition_GatheringType

	// A string that uniquely identifies the metric definition. The use of OSF UUID/GUIDs is recommended.
	Id string

	// IsContinuous indicates whether or not the metric value is continuous or scalar. Performance metrics are an example of a linear metric. Examples of non-linear metrics include error codes or operational states. Continuous metrics can be compared using the "greater than" relation.
	IsContinuous bool

	// The name of the metric. This name does not have to be unique, but should be descriptive and may contain blanks.
	Name string

	// Identifies the specific units of a value. The value of this property shall be a legal value of the Programmatic Units qualifier as defined in Appendix C.1 of DSP0004 V2.4 or later.
	ProgrammaticUnits string

	// TimeScope indicates the time scope to which the metric value applies.
	///0="Unknown" indicates the time scope was not qualified by the metric designer, or is unknown to the provider.
	///2="Point" indicates that the metric applies to a point in time. On the corresponding BaseMetricValue instances, TimeStamp specifies the point in time and Duration is always 0.
	///3="Interval" indicates that the metric applies to a time interval. On the corresponding BaseMetricValue instances, TimeStamp specifies the end of the time interval and Duration specifies its duration.
	///4="StartupInterval" indicates that the metric applies to a time interval that began at the startup of the measured resource (i.e. the ManagedElement associated by MetricDefForMe). On the corresponding BaseMetricValue instances, TimeStamp specifies the end of the time interval. If Duration is 0, this indicates that the startup time of the measured resource is unknown. Else, Duration specifies the duration between startup of the resource and TimeStamp.
	TimeScope BaseMetricDefinition_TimeScope

	// Identifies the specific units of a value. Examples are Bytes, Packets, Jobs, Files, Milliseconds, and Amps.
	Units string
}

func NewCIM_BaseMetricDefinitionEx1(instance *cim.WmiInstance) (newInstance *CIM_BaseMetricDefinition, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_BaseMetricDefinition{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_BaseMetricDefinitionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_BaseMetricDefinition, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_BaseMetricDefinition{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetBreakdownDimensions sets the value of BreakdownDimensions for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyBreakdownDimensions(value []string) (err error) {
	return instance.SetProperty("BreakdownDimensions", (value))
}

// GetBreakdownDimensions gets the value of BreakdownDimensions for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyBreakdownDimensions() (value []string, err error) {
	retValue, err := instance.GetProperty("BreakdownDimensions")
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

// SetCalculable sets the value of Calculable for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyCalculable(value BaseMetricDefinition_Calculable) (err error) {
	return instance.SetProperty("Calculable", (value))
}

// GetCalculable gets the value of Calculable for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyCalculable() (value BaseMetricDefinition_Calculable, err error) {
	retValue, err := instance.GetProperty("Calculable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BaseMetricDefinition_Calculable(valuetmp)

	return
}

// SetChangeType sets the value of ChangeType for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyChangeType(value BaseMetricDefinition_ChangeType) (err error) {
	return instance.SetProperty("ChangeType", (value))
}

// GetChangeType gets the value of ChangeType for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyChangeType() (value BaseMetricDefinition_ChangeType, err error) {
	retValue, err := instance.GetProperty("ChangeType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BaseMetricDefinition_ChangeType(valuetmp)

	return
}

// SetDataType sets the value of DataType for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyDataType(value BaseMetricDefinition_DataType) (err error) {
	return instance.SetProperty("DataType", (value))
}

// GetDataType gets the value of DataType for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyDataType() (value BaseMetricDefinition_DataType, err error) {
	retValue, err := instance.GetProperty("DataType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BaseMetricDefinition_DataType(valuetmp)

	return
}

// SetGatheringType sets the value of GatheringType for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyGatheringType(value BaseMetricDefinition_GatheringType) (err error) {
	return instance.SetProperty("GatheringType", (value))
}

// GetGatheringType gets the value of GatheringType for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyGatheringType() (value BaseMetricDefinition_GatheringType, err error) {
	retValue, err := instance.GetProperty("GatheringType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BaseMetricDefinition_GatheringType(valuetmp)

	return
}

// SetId sets the value of Id for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetIsContinuous sets the value of IsContinuous for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyIsContinuous(value bool) (err error) {
	return instance.SetProperty("IsContinuous", (value))
}

// GetIsContinuous gets the value of IsContinuous for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyIsContinuous() (value bool, err error) {
	retValue, err := instance.GetProperty("IsContinuous")
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

// SetName sets the value of Name for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyName() (value string, err error) {
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

// SetProgrammaticUnits sets the value of ProgrammaticUnits for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyProgrammaticUnits(value string) (err error) {
	return instance.SetProperty("ProgrammaticUnits", (value))
}

// GetProgrammaticUnits gets the value of ProgrammaticUnits for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyProgrammaticUnits() (value string, err error) {
	retValue, err := instance.GetProperty("ProgrammaticUnits")
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

// SetTimeScope sets the value of TimeScope for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyTimeScope(value BaseMetricDefinition_TimeScope) (err error) {
	return instance.SetProperty("TimeScope", (value))
}

// GetTimeScope gets the value of TimeScope for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyTimeScope() (value BaseMetricDefinition_TimeScope, err error) {
	retValue, err := instance.GetProperty("TimeScope")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BaseMetricDefinition_TimeScope(valuetmp)

	return
}

// SetUnits sets the value of Units for the instance
func (instance *CIM_BaseMetricDefinition) SetPropertyUnits(value string) (err error) {
	return instance.SetProperty("Units", (value))
}

// GetUnits gets the value of Units for the instance
func (instance *CIM_BaseMetricDefinition) GetPropertyUnits() (value string, err error) {
	retValue, err := instance.GetProperty("Units")
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
