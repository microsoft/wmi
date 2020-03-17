// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_LogicalPort struct
type CIM_LogicalPort struct {
	CIM_LogicalDevice

	// The maximum bandwidth of the Port in Bits per Second.
	MaxSpeed uint64

	// Describes the type of module, when PortType is set to 1 ("Other").
	OtherPortType string

	// PortType is defined to force consistent naming of the 'type' property in subclasses and to guarantee unique enum values for all instances of NetworkPort. When set to 1 ("Other"), related property OtherPortType contains a string description of the type of port. A range of values, DMTF_Reserved, has been defined that allows subclasses to override and define their specific types of ports.
	PortType LogicalPort_PortType

	// The requested bandwidth of the Port in Bits per Second. The actual bandwidth is reported in LogicalPort.Speed.
	RequestedSpeed uint64

	// The bandwidth of the Port in Bits per Second.
	Speed uint64

	// In some circumstances, a LogicalPort might be identifiable as a front end or back end port. An example of this situation would be a storage array that might have back end ports to communicate with disk drives and front end ports to communicate with hosts. If there is no restriction on the use of the port, then the value should be set to 'not restricted'.
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
