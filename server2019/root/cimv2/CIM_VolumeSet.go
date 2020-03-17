// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_VolumeSet struct
type CIM_VolumeSet struct {
	CIM_StorageExtent

	//
	PSExtentInterleaveDepth uint64

	//
	PSExtentStripeLength uint64
}

// SetPSExtentInterleaveDepth sets the value of PSExtentInterleaveDepth for the instance
func (instance *CIM_VolumeSet) SetPropertyPSExtentInterleaveDepth(value uint64) (err error) {
	return instance.SetProperty("PSExtentInterleaveDepth", value)
}

// GetPSExtentInterleaveDepth gets the value of PSExtentInterleaveDepth for the instance
func (instance *CIM_VolumeSet) GetPropertyPSExtentInterleaveDepth() (value uint64, err error) {
	retValue, err := instance.GetProperty("PSExtentInterleaveDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPSExtentStripeLength sets the value of PSExtentStripeLength for the instance
func (instance *CIM_VolumeSet) SetPropertyPSExtentStripeLength(value uint64) (err error) {
	return instance.SetProperty("PSExtentStripeLength", value)
}

// GetPSExtentStripeLength gets the value of PSExtentStripeLength for the instance
func (instance *CIM_VolumeSet) GetPropertyPSExtentStripeLength() (value uint64, err error) {
	retValue, err := instance.GetProperty("PSExtentStripeLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
