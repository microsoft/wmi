// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_ForwardingService struct
type CIM_ForwardingService struct {
	CIM_NetworkService

	// This defines the type of protocol that is being forwarded when the value of the ProtocolType attribute is 1 (i.e., "Other"). This provides for future extensibility.
	OtherProtocolType string

	// This defines the type of protocol that is being forwarded.
	ProtocolType ForwardingService_ProtocolType
}

// SetOtherProtocolType sets the value of OtherProtocolType for the instance
func (instance *CIM_ForwardingService) SetPropertyOtherProtocolType(value string) (err error) {
	return instance.SetProperty("OtherProtocolType", value)
}

// GetOtherProtocolType gets the value of OtherProtocolType for the instance
func (instance *CIM_ForwardingService) GetPropertyOtherProtocolType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherProtocolType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolType sets the value of ProtocolType for the instance
func (instance *CIM_ForwardingService) SetPropertyProtocolType(value ForwardingService_ProtocolType) (err error) {
	return instance.SetProperty("ProtocolType", value)
}

// GetProtocolType gets the value of ProtocolType for the instance
func (instance *CIM_ForwardingService) GetPropertyProtocolType() (value ForwardingService_ProtocolType, err error) {
	retValue, err := instance.GetProperty("ProtocolType")
	if err != nil {
		return
	}
	value, ok := retValue.(ForwardingService_ProtocolType)
	if !ok {
		// TODO: Set an error
	}
	return
}
