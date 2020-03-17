// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PortConnector struct
type Win32_PortConnector struct {
	CIM_PhysicalConnector

	//
	ExternalReferenceDesignator string

	//
	InternalReferenceDesignator string

	//
	PortType uint16
}

// SetExternalReferenceDesignator sets the value of ExternalReferenceDesignator for the instance
func (instance *Win32_PortConnector) SetPropertyExternalReferenceDesignator(value string) (err error) {
	return instance.SetProperty("ExternalReferenceDesignator", value)
}

// GetExternalReferenceDesignator gets the value of ExternalReferenceDesignator for the instance
func (instance *Win32_PortConnector) GetPropertyExternalReferenceDesignator() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalReferenceDesignator")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalReferenceDesignator sets the value of InternalReferenceDesignator for the instance
func (instance *Win32_PortConnector) SetPropertyInternalReferenceDesignator(value string) (err error) {
	return instance.SetProperty("InternalReferenceDesignator", value)
}

// GetInternalReferenceDesignator gets the value of InternalReferenceDesignator for the instance
func (instance *Win32_PortConnector) GetPropertyInternalReferenceDesignator() (value string, err error) {
	retValue, err := instance.GetProperty("InternalReferenceDesignator")
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
func (instance *Win32_PortConnector) SetPropertyPortType(value uint16) (err error) {
	return instance.SetProperty("PortType", value)
}

// GetPortType gets the value of PortType for the instance
func (instance *Win32_PortConnector) GetPropertyPortType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
