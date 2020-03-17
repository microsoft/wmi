// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.EventTracingManagement
//////////////////////////////////////////////
package eventtracingmanagement

// CIM_ManagedSystemElement struct
type CIM_ManagedSystemElement struct {
	CIM_ManagedElement

	//
	CommunicationStatus uint16

	//
	DetailedStatus uint16

	//
	HealthState uint16

	//
	InstallDate string

	//
	Name string

	//
	OperatingStatus uint16

	//
	OperationalStatus []uint16

	//
	PrimaryStatus uint16

	//
	Status string

	//
	StatusDescriptions []string
}

// SetCommunicationStatus sets the value of CommunicationStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyCommunicationStatus(value uint16) (err error) {
	return instance.SetProperty("CommunicationStatus", value)
}

// GetCommunicationStatus gets the value of CommunicationStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyCommunicationStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("CommunicationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDetailedStatus sets the value of DetailedStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyDetailedStatus(value uint16) (err error) {
	return instance.SetProperty("DetailedStatus", value)
}

// GetDetailedStatus gets the value of DetailedStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyDetailedStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("DetailedStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthState sets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyHealthState(value uint16) (err error) {
	return instance.SetProperty("HealthState", value)
}

// GetHealthState gets the value of HealthState for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyHealthState() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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
func (instance *CIM_ManagedSystemElement) SetPropertyOperatingStatus(value uint16) (err error) {
	return instance.SetProperty("OperatingStatus", value)
}

// GetOperatingStatus gets the value of OperatingStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyOperatingStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("OperatingStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryStatus sets the value of PrimaryStatus for the instance
func (instance *CIM_ManagedSystemElement) SetPropertyPrimaryStatus(value uint16) (err error) {
	return instance.SetProperty("PrimaryStatus", value)
}

// GetPrimaryStatus gets the value of PrimaryStatus for the instance
func (instance *CIM_ManagedSystemElement) GetPropertyPrimaryStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("PrimaryStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
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
