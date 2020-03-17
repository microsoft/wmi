// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_VolumeModificationTemplate struct
type SDDC_VolumeModificationTemplate struct {
	cim.WmiInstance

	//
	MediaType []uint16

	//
	Name string

	//
	Resiliency []uint16

	//
	SizeAvailable []uint64
}

// SetMediaType sets the value of MediaType for the instance
func (instance *SDDC_VolumeModificationTemplate) SetPropertyMediaType(value []uint16) (err error) {
	return instance.SetProperty("MediaType", value)
}

// GetMediaType gets the value of MediaType for the instance
func (instance *SDDC_VolumeModificationTemplate) GetPropertyMediaType() (value []uint16, err error) {
	retValue, err := instance.GetProperty("MediaType")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *SDDC_VolumeModificationTemplate) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VolumeModificationTemplate) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResiliency sets the value of Resiliency for the instance
func (instance *SDDC_VolumeModificationTemplate) SetPropertyResiliency(value []uint16) (err error) {
	return instance.SetProperty("Resiliency", value)
}

// GetResiliency gets the value of Resiliency for the instance
func (instance *SDDC_VolumeModificationTemplate) GetPropertyResiliency() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Resiliency")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSizeAvailable sets the value of SizeAvailable for the instance
func (instance *SDDC_VolumeModificationTemplate) SetPropertySizeAvailable(value []uint64) (err error) {
	return instance.SetProperty("SizeAvailable", value)
}

// GetSizeAvailable gets the value of SizeAvailable for the instance
func (instance *SDDC_VolumeModificationTemplate) GetPropertySizeAvailable() (value []uint64, err error) {
	retValue, err := instance.GetProperty("SizeAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
