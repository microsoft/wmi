// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_ManagedSystemElement struct
type CIM_ManagedSystemElement struct {
	*CIM_ManagedElement

	// Indicates the current health of the element. This attribute expresses the health of this element but not necessarily that of its subcomponents. The possible values are 0 to 30, where 5 means the element is entirely healthy and 30 means the element is completely non-functional. The following continuum is defined:
	///"Non-recoverable Error" (30) - The element has completely failed, and recovery is not possible. All functionality provided by this element has been lost.
	///"Critical Failure" (25) - The element is non-functional and recovery might not be possible.
	///"Major Failure" (20) - The element is failing. It is possible that some or all of the functionality of this component is degraded or not working.
	///"Minor Failure" (15) - All functionality is available but some might be degraded.
	///"Degraded/Warning" (10) - The element is in working order and all functionality is provided. However, the element is not working to the best of its abilities. For example, the element might not be operating at optimal performance or it might be reporting recoverable errors.
	///"OK" (5) - The element is fully functional and is operating within normal operational parameters and without error.
	///"Unknown" (0) - The implementation cannot report on HealthState at this time.
	///DMTF has reserved the unused portion of the continuum for additional HealthStates in the future.
	HealthState ManagedSystemElement_HealthState

	// A datetime value that indicates when the object was installed. Lack of a value does not indicate that the object is not installed.
	InstallDate string

	// The Name property defines the label by which the object is known. When subclassed, the Name property can be overridden to be a Key property.
	Name string

	// Indicates the current statuses of the element. Various operational statuses are defined. Many of the enumeration's values are self-explanatory. However, a few are not and are described here in more detail.
	///"Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
	///"Predictive Failure" indicates that an element is functioning nominally but predicting a failure in the near future.
	///"In Service" describes an element being configured, maintained, cleaned, or otherwise administered.
	///"No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
	///"Lost Communication" indicates that the ManagedSystem Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
	///"Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
	///"Dormant" indicates that the element is inactive or quiesced.
	///"Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
	///"Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
	///"Power Mode" indicates that the element has additional power model information contained in the Associated PowerManagementService association.
	///OperationalStatus replaces the Status property on ManagedSystemElement to provide a consistent approach to enumerations, to address implementation needs for an array property, and to provide a migration path from today's environment to the future. This change was not made earlier because it required the deprecated qualifier. Due to the widespread use of the existing Status property in management applications, it is strongly recommended that providers or instrumentation provide both the Status and OperationalStatus properties. Further, the first value of OperationalStatus should contain the primary status for the element. When instrumented, Status (because it is single-valued) should also provide the primary status of the element.
	OperationalStatus []ManagedSystemElement_OperationalStatus

	// A string indicating the current status of the object. Various operational and non-operational statuses are defined. This property is deprecated in lieu of OperationalStatus, which includes the same semantics in its enumeration. This change is made for 3 reasons:
	///1) Status is more correctly defined as an array. This definition overcomes the limitation of describing status using a single value, when it is really a multi-valued property (for example, an element might be OK AND Stopped.
	///2) A MaxLen of 10 is too restrictive and leads to unclear enumerated values.
	///3) The change to a uint16 data type was discussed when CIM V2.0 was defined. However, existing V1.0 implementations used the string property and did not want to modify their code. Therefore, Status was grandfathered into the Schema. Use of the deprecated qualifier allows the maintenance of the existing property, but also permits an improved definition using OperationalStatus.
	Status string

	// Strings describing the various OperationalStatus array values. For example, if "Stopping" is the value assigned to OperationalStatus, then this property may contain an explanation as to why an object is being stopped. Note that entries in this array are correlated with those at the same array index in OperationalStatus.
	StatusDescriptions []string
}

func NewCIM_ManagedSystemElementEx1(instance *cim.WmiInstance) (newInstance *CIM_ManagedSystemElement, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_ManagedSystemElement{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_ManagedSystemElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_ManagedSystemElement, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_ManagedSystemElement{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetHealthState sets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyHealthState(value ManagedSystemElement_HealthState) (err error) {
	return instance.SetProperty("HealthState", (value))
}

// GetHealthState gets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyHealthState() (value ManagedSystemElement_HealthState, err error) {
	retValue, err := instance.GetProperty("HealthState")
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

	value = ManagedSystemElement_HealthState(valuetmp)

	return
}

// SetInstallDate sets the value of InstallDate for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", (value))
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
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
func (instance *CIM_ManagedSystemElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyName() (value string, err error) {
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyOperationalStatus(value []ManagedSystemElement_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyOperationalStatus() (value []ManagedSystemElement_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
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
		value = append(value, ManagedSystemElement_OperationalStatus(valuetmp))
	}

	return
}

// SetStatus sets the value of Status for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStatusDescriptions sets the value of StatusDescriptions for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyStatusDescriptions(value []string) (err error) {
	return instance.SetProperty("StatusDescriptions", (value))
}

// GetStatusDescriptions gets the value of StatusDescriptions for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyStatusDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("StatusDescriptions")
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
