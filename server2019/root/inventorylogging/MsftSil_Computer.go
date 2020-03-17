// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// MsftSil_Computer struct
type MsftSil_Computer struct {
	MsftSil_Data

	//
	ChassisSerialNumber string

	//
	CollectedDateTime string

	//
	Model string

	//
	Name string

	//
	NumberOfCores uint32

	//
	NumberOfLogicalProcessors uint32

	//
	NumberOfProcessors uint32

	//
	OSName string

	//
	OSSku uint32

	//
	OSSuite uint32

	//
	OSSuiteMask uint32

	//
	OSVersion string

	//
	ProcessorFamily uint32

	//
	ProcessorManufacturer string

	//
	ProcessorName string

	//
	SystemManufacturer string
}

// SetChassisSerialNumber sets the value of ChassisSerialNumber for the instance
func (instance *MsftSil_Computer) SetPropertyChassisSerialNumber(value string) (err error) {
	return instance.SetProperty("ChassisSerialNumber", value)
}

// GetChassisSerialNumber gets the value of ChassisSerialNumber for the instance
func (instance *MsftSil_Computer) GetPropertyChassisSerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("ChassisSerialNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCollectedDateTime sets the value of CollectedDateTime for the instance
func (instance *MsftSil_Computer) SetPropertyCollectedDateTime(value string) (err error) {
	return instance.SetProperty("CollectedDateTime", value)
}

// GetCollectedDateTime gets the value of CollectedDateTime for the instance
func (instance *MsftSil_Computer) GetPropertyCollectedDateTime() (value string, err error) {
	retValue, err := instance.GetProperty("CollectedDateTime")
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
func (instance *MsftSil_Computer) SetPropertyModel(value string) (err error) {
	return instance.SetProperty("Model", value)
}

// GetModel gets the value of Model for the instance
func (instance *MsftSil_Computer) GetPropertyModel() (value string, err error) {
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

// SetName sets the value of Name for the instance
func (instance *MsftSil_Computer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MsftSil_Computer) GetPropertyName() (value string, err error) {
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

// SetNumberOfCores sets the value of NumberOfCores for the instance
func (instance *MsftSil_Computer) SetPropertyNumberOfCores(value uint32) (err error) {
	return instance.SetProperty("NumberOfCores", value)
}

// GetNumberOfCores gets the value of NumberOfCores for the instance
func (instance *MsftSil_Computer) GetPropertyNumberOfCores() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfCores")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfLogicalProcessors sets the value of NumberOfLogicalProcessors for the instance
func (instance *MsftSil_Computer) SetPropertyNumberOfLogicalProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfLogicalProcessors", value)
}

// GetNumberOfLogicalProcessors gets the value of NumberOfLogicalProcessors for the instance
func (instance *MsftSil_Computer) GetPropertyNumberOfLogicalProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfLogicalProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfProcessors sets the value of NumberOfProcessors for the instance
func (instance *MsftSil_Computer) SetPropertyNumberOfProcessors(value uint32) (err error) {
	return instance.SetProperty("NumberOfProcessors", value)
}

// GetNumberOfProcessors gets the value of NumberOfProcessors for the instance
func (instance *MsftSil_Computer) GetPropertyNumberOfProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfProcessors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSName sets the value of OSName for the instance
func (instance *MsftSil_Computer) SetPropertyOSName(value string) (err error) {
	return instance.SetProperty("OSName", value)
}

// GetOSName gets the value of OSName for the instance
func (instance *MsftSil_Computer) GetPropertyOSName() (value string, err error) {
	retValue, err := instance.GetProperty("OSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSSku sets the value of OSSku for the instance
func (instance *MsftSil_Computer) SetPropertyOSSku(value uint32) (err error) {
	return instance.SetProperty("OSSku", value)
}

// GetOSSku gets the value of OSSku for the instance
func (instance *MsftSil_Computer) GetPropertyOSSku() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSSku")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSSuite sets the value of OSSuite for the instance
func (instance *MsftSil_Computer) SetPropertyOSSuite(value uint32) (err error) {
	return instance.SetProperty("OSSuite", value)
}

// GetOSSuite gets the value of OSSuite for the instance
func (instance *MsftSil_Computer) GetPropertyOSSuite() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSSuite")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSSuiteMask sets the value of OSSuiteMask for the instance
func (instance *MsftSil_Computer) SetPropertyOSSuiteMask(value uint32) (err error) {
	return instance.SetProperty("OSSuiteMask", value)
}

// GetOSSuiteMask gets the value of OSSuiteMask for the instance
func (instance *MsftSil_Computer) GetPropertyOSSuiteMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("OSSuiteMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOSVersion sets the value of OSVersion for the instance
func (instance *MsftSil_Computer) SetPropertyOSVersion(value string) (err error) {
	return instance.SetProperty("OSVersion", value)
}

// GetOSVersion gets the value of OSVersion for the instance
func (instance *MsftSil_Computer) GetPropertyOSVersion() (value string, err error) {
	retValue, err := instance.GetProperty("OSVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorFamily sets the value of ProcessorFamily for the instance
func (instance *MsftSil_Computer) SetPropertyProcessorFamily(value uint32) (err error) {
	return instance.SetProperty("ProcessorFamily", value)
}

// GetProcessorFamily gets the value of ProcessorFamily for the instance
func (instance *MsftSil_Computer) GetPropertyProcessorFamily() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessorFamily")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorManufacturer sets the value of ProcessorManufacturer for the instance
func (instance *MsftSil_Computer) SetPropertyProcessorManufacturer(value string) (err error) {
	return instance.SetProperty("ProcessorManufacturer", value)
}

// GetProcessorManufacturer gets the value of ProcessorManufacturer for the instance
func (instance *MsftSil_Computer) GetPropertyProcessorManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessorManufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorName sets the value of ProcessorName for the instance
func (instance *MsftSil_Computer) SetPropertyProcessorName(value string) (err error) {
	return instance.SetProperty("ProcessorName", value)
}

// GetProcessorName gets the value of ProcessorName for the instance
func (instance *MsftSil_Computer) GetPropertyProcessorName() (value string, err error) {
	retValue, err := instance.GetProperty("ProcessorName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemManufacturer sets the value of SystemManufacturer for the instance
func (instance *MsftSil_Computer) SetPropertySystemManufacturer(value string) (err error) {
	return instance.SetProperty("SystemManufacturer", value)
}

// GetSystemManufacturer gets the value of SystemManufacturer for the instance
func (instance *MsftSil_Computer) GetPropertySystemManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("SystemManufacturer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
