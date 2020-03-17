// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_TransparentBridgingService struct
type CIM_TransparentBridgingService struct {
	CIM_ForwardingService

	// The timeout period in seconds for aging out dynamically learned forwarding information. 802.1D-1990 recommends a default of 300 seconds.
	AgingTime uint32

	// Filtering Database Identifier used by VLAN-aware switches that have more than one filtering database.
	FID uint32
}

// SetAgingTime sets the value of AgingTime for the instance
func (instance *CIM_TransparentBridgingService) SetPropertyAgingTime(value uint32) (err error) {
	return instance.SetProperty("AgingTime", value)
}

// GetAgingTime gets the value of AgingTime for the instance
func (instance *CIM_TransparentBridgingService) GetPropertyAgingTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("AgingTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFID sets the value of FID for the instance
func (instance *CIM_TransparentBridgingService) SetPropertyFID(value uint32) (err error) {
	return instance.SetProperty("FID", value)
}

// GetFID gets the value of FID for the instance
func (instance *CIM_TransparentBridgingService) GetPropertyFID() (value uint32, err error) {
	retValue, err := instance.GetProperty("FID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
