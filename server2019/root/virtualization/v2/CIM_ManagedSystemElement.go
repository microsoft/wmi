// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_ManagedSystemElement struct
type CIM_ManagedSystemElement struct {
	CIM_ManagedElement

	// CommunicationStatus indicates the ability of the instrumentation to communicate with the underlying ManagedElement. CommunicationStatus consists of one of the following values: Unknown, None, Communication OK, Lost Communication, or No Contact.
	///A Null return indicates the implementation (provider) does not implement this property.
	///"Unknown" indicates the implementation is in general capable of returning this property, but is unable to do so at this time.
	///"Not Available" indicates that the implementation (provider) is capable of returning a value for this property, but not ever for this particular piece of hardware/software or the property is intentionally not used because it adds no meaningful information (as in the case of a property that is intended to add additional info to another property).
	///"Communication OK " indicates communication is established with the element, but does not convey any quality of service.
	///"No Contact" indicates that the monitoring system has knowledge of this element, but has never been able to establish communications with it.
	///"Lost Communication" indicates that the Managed Element is known to exist and has been contacted successfully in the past, but is currently unreachable.
	CommunicationStatus ManagedSystemElement_CommunicationStatus

	// DetailedStatus compliments PrimaryStatus with additional status detail. It consists of one of the following values: Not Available, No Additional Information, Stressed, Predictive Failure, Error, Non-Recoverable Error, SupportingEntityInError. Detailed status is used to expand upon the PrimaryStatus of the element.
	///A Null return indicates the implementation (provider) does not implement this property.
	///"Not Available" indicates that the implementation (provider) is capable of returning a value for this property, but not ever for this particular piece of hardware/software or the property is intentionally not used because it adds no meaningful information (as in the case of a property that is intended to add additional info to another property).
	///"No Additional Information" indicates that the element is functioning normally as indicated by PrimaryStatus = "OK".
	///"Stressed" indicates that the element is functioning, but needs attention. Examples of "Stressed" states are overload, overheated, and so on.
	///"Predictive Failure" indicates that an element is functioning normally but a failure is predicted in the near future.
	///"Non-Recoverable Error " indicates that this element is in an error condition that requires human intervention.
	///"Supporting Entity in Error" indicates that this element might be "OK" but that another element, on which it is dependent, is in error. An example is a network service or endpoint that cannot function due to lower-layer networking problems.
	DetailedStatus ManagedSystemElement_DetailedStatus

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

	// OperatingStatus provides a current status value for the operational condition of the element and can be used for providing more detail with respect to the value of EnabledState. It can also provide the transitional states when an element is transitioning from one state to another, such as when an element is transitioning between EnabledState and RequestedState, as well as other transitional conditions.
	///OperatingStatus consists of one of the following values: Unknown, Not Available, In Service, Starting, Stopping, Stopped, Aborted, Dormant, Completed, Migrating, Emmigrating, Immigrating, Snapshotting. Shutting Down, In Test
	///A Null return indicates the implementation (provider) does not implement this property.
	///"Unknown" indicates the implementation is in general capable of returning this property, but is unable to do so at this time.
	///"None" indicates that the implementation (provider) is capable of returning a value for this property, but not ever for this particular piece of hardware/software or the property is intentionally not used because it adds no meaningful information (as in the case of a property that is intended to add additional info to another property).
	///"Servicing" describes an element being configured, maintained, cleaned, or otherwise administered.
	///"Starting" describes an element being initialized.
	///"Stopping" describes an element being brought to an orderly stop.
	///"Stopped" and "Aborted" are similar, although the former implies a clean and orderly stop, while the latter implies an abrupt stop where the state and configuration of the element might need to be updated.
	///"Dormant" indicates that the element is inactive or quiesced.
	///"Completed" indicates that the element has completed its operation. This value should be combined with either OK, Error, or Degraded in the PrimaryStatus so that a client can tell if the complete operation Completed with OK (passed), Completed with Error (failed), or Completed with Degraded (the operation finished, but it did not complete OK or did not report an error).
	///"Migrating" element is being moved between host elements.
	///"Immigrating" element is being moved to new host element.
	///"Emigrating" element is being moved away from host element.
	///"Shutting Down" describes an element being brought to an abrupt stop.
	///"In Test" element is performing test functions.
	///"Transitioning" describes an element that is between states, that is, it is not fully available in either its previous state or its next state. This value should be used if other values indicating a transition to a specific state are not applicable.
	///"In Service" describes an element that is in service and operational.
	OperatingStatus ManagedSystemElement_OperatingStatus

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

	// PrimaryStatus provides a high level status value, intended to align with Red-Yellow-Green type representation of status. It should be used in conjunction with DetailedStatus to provide high level and detailed health status of the ManagedElement and its subcomponents.
	///PrimaryStatus consists of one of the following values: Unknown, OK, Degraded or Error. "Unknown" indicates the implementation is in general capable of returning this property, but is unable to do so at this time.
	///"OK" indicates the ManagedElement is functioning normally.
	///"Degraded" indicates the ManagedElement is functioning below normal.
	///"Error" indicates the ManagedElement is in an Error condition.
	PrimaryStatus ManagedSystemElement_PrimaryStatus

	// A string indicating the current status of the object. Various operational and non-operational statuses are defined. This property is deprecated in lieu of OperationalStatus, which includes the same semantics in its enumeration. This change is made for 3 reasons:
	///1) Status is more correctly defined as an array. This definition overcomes the limitation of describing status using a single value, when it is really a multi-valued property (for example, an element might be OK AND Stopped.
	///2) A MaxLen of 10 is too restrictive and leads to unclear enumerated values.
	///3) The change to a uint16 data type was discussed when CIM V2.0 was defined. However, existing V1.0 implementations used the string property and did not want to modify their code. Therefore, Status was grandfathered into the Schema. Use of the deprecated qualifier allows the maintenance of the existing property, but also permits an improved definition using OperationalStatus.
	Status string

	// Strings describing the various OperationalStatus array values. For example, if "Stopping" is the value assigned to OperationalStatus, then this property may contain an explanation as to why an object is being stopped. Note that entries in this array are correlated with those at the same array index in OperationalStatus.
	StatusDescriptions []string
}

// SetCommunicationStatus sets the value of CommunicationStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyCommunicationStatus(value ManagedSystemElement_CommunicationStatus) (err error) {
	return instance.SetProperty("CommunicationStatus", value)
}

// GetCommunicationStatus gets the value of CommunicationStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyCommunicationStatus() (value ManagedSystemElement_CommunicationStatus, err error) {
	retValue, err := instance.GetProperty("CommunicationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(ManagedSystemElement_CommunicationStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDetailedStatus sets the value of DetailedStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyDetailedStatus(value ManagedSystemElement_DetailedStatus) (err error) {
	return instance.SetProperty("DetailedStatus", value)
}

// GetDetailedStatus gets the value of DetailedStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyDetailedStatus() (value ManagedSystemElement_DetailedStatus, err error) {
	retValue, err := instance.GetProperty("DetailedStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(ManagedSystemElement_DetailedStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthState sets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyHealthState(value ManagedSystemElement_HealthState) (err error) {
	return instance.SetProperty("HealthState", value)
}

// GetHealthState gets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyHealthState() (value ManagedSystemElement_HealthState, err error) {
	retValue, err := instance.GetProperty("HealthState")
	if err != nil {
		return
	}
	value, ok := retValue.(ManagedSystemElement_HealthState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallDate sets the value of InstallDate for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", value)
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperatingStatus sets the value of OperatingStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyOperatingStatus(value ManagedSystemElement_OperatingStatus) (err error) {
	return instance.SetProperty("OperatingStatus", value)
}

// GetOperatingStatus gets the value of OperatingStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyOperatingStatus() (value ManagedSystemElement_OperatingStatus, err error) {
	retValue, err := instance.GetProperty("OperatingStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(ManagedSystemElement_OperatingStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyOperationalStatus(value []ManagedSystemElement_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyOperationalStatus() (value []ManagedSystemElement_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]ManagedSystemElement_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryStatus sets the value of PrimaryStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyPrimaryStatus(value ManagedSystemElement_PrimaryStatus) (err error) {
	return instance.SetProperty("PrimaryStatus", value)
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyPrimaryStatus() (value ManagedSystemElement_PrimaryStatus, err error) {
	retValue, err := instance.GetProperty("PrimaryStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(ManagedSystemElement_PrimaryStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusDescriptions sets the value of StatusDescriptions for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyStatusDescriptions(value []string) (err error) {
	return instance.SetProperty("StatusDescriptions", value)
}

// GetStatusDescriptions gets the value of StatusDescriptions for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyStatusDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("StatusDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
