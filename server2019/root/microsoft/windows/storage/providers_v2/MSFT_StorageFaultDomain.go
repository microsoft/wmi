// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_StorageFaultDomain struct
type MSFT_StorageFaultDomain struct {
	MSFT_StorageObject

	// A user settable description of the fault domain object.
	Description string

	// A user-friendly string representing the name of the fault domain object.
	FriendlyName string

	//
	HealthStatus StorageFaultDomain_HealthStatus

	// This field represents the name of the company responsible for the hardware backing the fault domain oject. For physical disk it must match the disk's SCSI inquiry data.
	Manufacturer string

	// This field represents the model number of the hardware. For physical disk it must match the disk's SCSI inquiry data.
	Model string

	//
	OperationalDetails []string

	//
	OperationalStatus []StorageFaultDomain_OperationalStatus

	// This field is a free-form string indicating where the hardware is located.
	PhysicalLocation string

	// This field represents the serial number of the hardware. For physical disk it must match the disk's SCSI inquiry data.
	SerialNumber string
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyHealthStatus(value StorageFaultDomain_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyHealthStatus() (value StorageFaultDomain_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageFaultDomain_HealthStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", value)
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModel sets the value of Model for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyModel() (value string, err error) {
	retValue, err := instance.GetProperty("Model")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalDetails sets the value of OperationalDetails for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyOperationalDetails(value []string) (err error) {
	return instance.SetProperty("OperationalDetails", value)
}

// GetOperationalDetails gets the value of OperationalDetails for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyOperationalDetails() (value []string, err error) {
	retValue, err := instance.GetProperty("OperationalDetails")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyOperationalStatus(value []StorageFaultDomain_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyOperationalStatus() (value []StorageFaultDomain_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]StorageFaultDomain_OperationalStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhysicalLocation sets the value of PhysicalLocation for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertyPhysicalLocation(value string) (err error) {
	return instance.SetProperty("PhysicalLocation", value)
}

// GetPhysicalLocation gets the value of PhysicalLocation for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertyPhysicalLocation() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *MSFT_StorageFaultDomain) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", value)
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *MSFT_StorageFaultDomain) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
