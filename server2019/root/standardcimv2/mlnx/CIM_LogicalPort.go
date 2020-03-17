// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_LogicalPort struct
type CIM_LogicalPort struct {
	CIM_LogicalDevice

	//
	MaxSpeed uint64

	//
	OtherPortType string

	//
	PortType LogicalPort_PortType

	//
	RequestedSpeed uint64

	//
	Speed uint64

	//
	UsageRestriction LogicalPort_UsageRestriction
}

// SetMaxSpeed sets the value of MaxSpeed for the instance
func (instance *CIM_LogicalPort) SetPropertyMaxSpeed(value uint64) (err error) {
	return instance.SetProperty("MaxSpeed", value)
}

// GetMaxSpeed gets the value of MaxSpeed for the instance
func (instance *CIM_LogicalPort) GetPropertyMaxSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherPortType sets the value of OtherPortType for the instance
func (instance *CIM_LogicalPort) SetPropertyOtherPortType(value string) (err error) {
	return instance.SetProperty("OtherPortType", value)
}

// GetOtherPortType gets the value of OtherPortType for the instance
func (instance *CIM_LogicalPort) GetPropertyOtherPortType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherPortType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortType sets the value of PortType for the instance
func (instance *CIM_LogicalPort) SetPropertyPortType(value LogicalPort_PortType) (err error) {
	return instance.SetProperty("PortType", value)
}

// GetPortType gets the value of PortType for the instance
func (instance *CIM_LogicalPort) GetPropertyPortType() (value LogicalPort_PortType, err error) {
	retValue, err := instance.GetProperty("PortType")
	if err != nil {
		return
	}
	value, ok := retValue.(LogicalPort_PortType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequestedSpeed sets the value of RequestedSpeed for the instance
func (instance *CIM_LogicalPort) SetPropertyRequestedSpeed(value uint64) (err error) {
	return instance.SetProperty("RequestedSpeed", value)
}

// GetRequestedSpeed gets the value of RequestedSpeed for the instance
func (instance *CIM_LogicalPort) GetPropertyRequestedSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestedSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpeed sets the value of Speed for the instance
func (instance *CIM_LogicalPort) SetPropertySpeed(value uint64) (err error) {
	return instance.SetProperty("Speed", value)
}

// GetSpeed gets the value of Speed for the instance
func (instance *CIM_LogicalPort) GetPropertySpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("Speed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsageRestriction sets the value of UsageRestriction for the instance
func (instance *CIM_LogicalPort) SetPropertyUsageRestriction(value LogicalPort_UsageRestriction) (err error) {
	return instance.SetProperty("UsageRestriction", value)
}

// GetUsageRestriction gets the value of UsageRestriction for the instance
func (instance *CIM_LogicalPort) GetPropertyUsageRestriction() (value LogicalPort_UsageRestriction, err error) {
	retValue, err := instance.GetProperty("UsageRestriction")
	if err != nil {
		return
	}
	value, ok := retValue.(LogicalPort_UsageRestriction)
	if !ok {
		// TODO: Set an error
	}
	return
}
