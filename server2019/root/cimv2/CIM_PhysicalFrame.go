// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PhysicalFrame struct
type CIM_PhysicalFrame struct {
	CIM_PhysicalPackage

	//
	AudibleAlarm bool

	//
	BreachDescription string

	//
	CableManagementStrategy string

	//
	LockPresent bool

	//
	SecurityBreach uint16

	//
	ServiceDescriptions []string

	//
	ServicePhilosophy []uint16

	//
	VisibleAlarm bool
}

// SetAudibleAlarm sets the value of AudibleAlarm for the instance
func (instance *CIM_PhysicalFrame) SetPropertyAudibleAlarm(value bool) (err error) {
	return instance.SetProperty("AudibleAlarm", value)
}

// GetAudibleAlarm gets the value of AudibleAlarm for the instance
func (instance *CIM_PhysicalFrame) GetPropertyAudibleAlarm() (value bool, err error) {
	retValue, err := instance.GetProperty("AudibleAlarm")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBreachDescription sets the value of BreachDescription for the instance
func (instance *CIM_PhysicalFrame) SetPropertyBreachDescription(value string) (err error) {
	return instance.SetProperty("BreachDescription", value)
}

// GetBreachDescription gets the value of BreachDescription for the instance
func (instance *CIM_PhysicalFrame) GetPropertyBreachDescription() (value string, err error) {
	retValue, err := instance.GetProperty("BreachDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCableManagementStrategy sets the value of CableManagementStrategy for the instance
func (instance *CIM_PhysicalFrame) SetPropertyCableManagementStrategy(value string) (err error) {
	return instance.SetProperty("CableManagementStrategy", value)
}

// GetCableManagementStrategy gets the value of CableManagementStrategy for the instance
func (instance *CIM_PhysicalFrame) GetPropertyCableManagementStrategy() (value string, err error) {
	retValue, err := instance.GetProperty("CableManagementStrategy")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLockPresent sets the value of LockPresent for the instance
func (instance *CIM_PhysicalFrame) SetPropertyLockPresent(value bool) (err error) {
	return instance.SetProperty("LockPresent", value)
}

// GetLockPresent gets the value of LockPresent for the instance
func (instance *CIM_PhysicalFrame) GetPropertyLockPresent() (value bool, err error) {
	retValue, err := instance.GetProperty("LockPresent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityBreach sets the value of SecurityBreach for the instance
func (instance *CIM_PhysicalFrame) SetPropertySecurityBreach(value uint16) (err error) {
	return instance.SetProperty("SecurityBreach", value)
}

// GetSecurityBreach gets the value of SecurityBreach for the instance
func (instance *CIM_PhysicalFrame) GetPropertySecurityBreach() (value uint16, err error) {
	retValue, err := instance.GetProperty("SecurityBreach")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServiceDescriptions sets the value of ServiceDescriptions for the instance
func (instance *CIM_PhysicalFrame) SetPropertyServiceDescriptions(value []string) (err error) {
	return instance.SetProperty("ServiceDescriptions", value)
}

// GetServiceDescriptions gets the value of ServiceDescriptions for the instance
func (instance *CIM_PhysicalFrame) GetPropertyServiceDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("ServiceDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServicePhilosophy sets the value of ServicePhilosophy for the instance
func (instance *CIM_PhysicalFrame) SetPropertyServicePhilosophy(value []uint16) (err error) {
	return instance.SetProperty("ServicePhilosophy", value)
}

// GetServicePhilosophy gets the value of ServicePhilosophy for the instance
func (instance *CIM_PhysicalFrame) GetPropertyServicePhilosophy() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ServicePhilosophy")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVisibleAlarm sets the value of VisibleAlarm for the instance
func (instance *CIM_PhysicalFrame) SetPropertyVisibleAlarm(value bool) (err error) {
	return instance.SetProperty("VisibleAlarm", value)
}

// GetVisibleAlarm gets the value of VisibleAlarm for the instance
func (instance *CIM_PhysicalFrame) GetPropertyVisibleAlarm() (value bool, err error) {
	retValue, err := instance.GetProperty("VisibleAlarm")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
