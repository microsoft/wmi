// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PhysicalLink struct
type CIM_PhysicalLink struct {
	CIM_PhysicalElement

	//
	Length float64

	//
	MaxLength float64

	//
	MediaType uint16

	//
	Wired bool
}

// SetLength sets the value of Length for the instance
func (instance *CIM_PhysicalLink) SetPropertyLength(value float64) (err error) {
	return instance.SetProperty("Length", value)
}

// GetLength gets the value of Length for the instance
func (instance *CIM_PhysicalLink) GetPropertyLength() (value float64, err error) {
	retValue, err := instance.GetProperty("Length")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxLength sets the value of MaxLength for the instance
func (instance *CIM_PhysicalLink) SetPropertyMaxLength(value float64) (err error) {
	return instance.SetProperty("MaxLength", value)
}

// GetMaxLength gets the value of MaxLength for the instance
func (instance *CIM_PhysicalLink) GetPropertyMaxLength() (value float64, err error) {
	retValue, err := instance.GetProperty("MaxLength")
	if err != nil {
		return
	}
	value, ok := retValue.(float64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaType sets the value of MediaType for the instance
func (instance *CIM_PhysicalLink) SetPropertyMediaType(value uint16) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *CIM_PhysicalLink) GetPropertyMediaType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWired sets the value of Wired for the instance
func (instance *CIM_PhysicalLink) SetPropertyWired(value bool) (err error) {
	return instance.SetProperty("Wired", value)
}

// GetWired gets the value of Wired for the instance
func (instance *CIM_PhysicalLink) GetPropertyWired() (value bool, err error) {
	retValue, err := instance.GetProperty("Wired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
