// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Processor struct
type CIM_Processor struct {
	CIM_LogicalDevice

	//
	AddressWidth uint16

	//
	CurrentClockSpeed uint32

	//
	DataWidth uint16

	//
	Family uint16

	//
	LoadPercentage uint16

	//
	MaxClockSpeed uint32

	//
	OtherFamilyDescription string

	//
	Role string

	//
	Stepping string

	//
	UniqueId string

	//
	UpgradeMethod uint16
}

// SetAddressWidth sets the value of AddressWidth for the instance
func (instance *CIM_Processor) SetPropertyAddressWidth(value uint16) (err error) {
	return instance.SetProperty("AddressWidth", value)
}

// GetAddressWidth gets the value of AddressWidth for the instance
func (instance *CIM_Processor) GetPropertyAddressWidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentClockSpeed sets the value of CurrentClockSpeed for the instance
func (instance *CIM_Processor) SetPropertyCurrentClockSpeed(value uint32) (err error) {
	return instance.SetProperty("CurrentClockSpeed", value)
}

// GetCurrentClockSpeed gets the value of CurrentClockSpeed for the instance
func (instance *CIM_Processor) GetPropertyCurrentClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentClockSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDataWidth sets the value of DataWidth for the instance
func (instance *CIM_Processor) SetPropertyDataWidth(value uint16) (err error) {
	return instance.SetProperty("DataWidth", value)
}

// GetDataWidth gets the value of DataWidth for the instance
func (instance *CIM_Processor) GetPropertyDataWidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("DataWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFamily sets the value of Family for the instance
func (instance *CIM_Processor) SetPropertyFamily(value uint16) (err error) {
	return instance.SetProperty("Family", value)
}

// GetFamily gets the value of Family for the instance
func (instance *CIM_Processor) GetPropertyFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("Family")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoadPercentage sets the value of LoadPercentage for the instance
func (instance *CIM_Processor) SetPropertyLoadPercentage(value uint16) (err error) {
	return instance.SetProperty("LoadPercentage", value)
}

// GetLoadPercentage gets the value of LoadPercentage for the instance
func (instance *CIM_Processor) GetPropertyLoadPercentage() (value uint16, err error) {
	retValue, err := instance.GetProperty("LoadPercentage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxClockSpeed sets the value of MaxClockSpeed for the instance
func (instance *CIM_Processor) SetPropertyMaxClockSpeed(value uint32) (err error) {
	return instance.SetProperty("MaxClockSpeed", value)
}

// GetMaxClockSpeed gets the value of MaxClockSpeed for the instance
func (instance *CIM_Processor) GetPropertyMaxClockSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxClockSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherFamilyDescription sets the value of OtherFamilyDescription for the instance
func (instance *CIM_Processor) SetPropertyOtherFamilyDescription(value string) (err error) {
	return instance.SetProperty("OtherFamilyDescription", value)
}

// GetOtherFamilyDescription gets the value of OtherFamilyDescription for the instance
func (instance *CIM_Processor) GetPropertyOtherFamilyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherFamilyDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRole sets the value of Role for the instance
func (instance *CIM_Processor) SetPropertyRole(value string) (err error) {
	return instance.SetProperty("Role", value)
}

// GetRole gets the value of Role for the instance
func (instance *CIM_Processor) GetPropertyRole() (value string, err error) {
	retValue, err := instance.GetProperty("Role")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStepping sets the value of Stepping for the instance
func (instance *CIM_Processor) SetPropertyStepping(value string) (err error) {
	return instance.SetProperty("Stepping", value)
}

// GetStepping gets the value of Stepping for the instance
func (instance *CIM_Processor) GetPropertyStepping() (value string, err error) {
	retValue, err := instance.GetProperty("Stepping")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueId sets the value of UniqueId for the instance
func (instance *CIM_Processor) SetPropertyUniqueId(value string) (err error) {
	return instance.SetProperty("UniqueId", value)
}

// GetUniqueId gets the value of UniqueId for the instance
func (instance *CIM_Processor) GetPropertyUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpgradeMethod sets the value of UpgradeMethod for the instance
func (instance *CIM_Processor) SetPropertyUpgradeMethod(value uint16) (err error) {
	return instance.SetProperty("UpgradeMethod", value)
}

// GetUpgradeMethod gets the value of UpgradeMethod for the instance
func (instance *CIM_Processor) GetPropertyUpgradeMethod() (value uint16, err error) {
	retValue, err := instance.GetProperty("UpgradeMethod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
