// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PhysicalConnector struct
type CIM_PhysicalConnector struct {
	CIM_PhysicalElement

	//
	ConnectorPinout string

	//
	ConnectorType []uint16
}

// SetConnectorPinout sets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorPinout(value string) (err error) {
	return instance.SetProperty("ConnectorPinout", value)
}

// GetConnectorPinout gets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorPinout() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectorPinout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorType sets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorType(value []uint16) (err error) {
	return instance.SetProperty("ConnectorType", value)
}

// GetConnectorType gets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorType() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ConnectorType")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
