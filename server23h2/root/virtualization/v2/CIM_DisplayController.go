// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DisplayController struct
type CIM_DisplayController struct {
	*CIM_Controller

	// An array of integers indicating the graphics and 3D capabilities of the DisplayController.
	AcceleratorCapabilities []DisplayController_AcceleratorCapabilities

	// An array of free-form strings providing more detailed explanations for any of the video Accelerator features indicated in the Capabilities array. Note, each entry of this array is related to the entry in the Capabilities array that is located at the same index.
	CapabilityDescriptions []string

	// Maximum amount of memory supported in bytes.
	MaxMemorySupported uint32

	// Number of video pages supported given the current resolutions and available memory.
	NumberOfVideoPages uint32

	// A string describing the video architecture type when the instance's VideoArchitecture property is 1 ("Other").
	OtherVideoArchitecture string

	// A string describing the video memory type when the instance's VideoMemoryType property is 1 ("Other").
	OtherVideoMemoryType string

	// An integer enumeration indicating the display controllers video architecture used to generate the video signal. Usually, a dedicated video processor generates the video signal in accordance with the specified architecture.It is an indicator of the maximum resolution capability of the display controller.
	VideoArchitecture DisplayController_VideoArchitecture

	// An integer enumeration indicating the type of video memory.
	VideoMemoryType DisplayController_VideoMemoryType

	// A free-form string describing the video processor/Controller.
	VideoProcessor string
}

func NewCIM_DisplayControllerEx1(instance *cim.WmiInstance) (newInstance *CIM_DisplayController, err error) {
	tmp, err := NewCIM_ControllerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DisplayController{
		CIM_Controller: tmp,
	}
	return
}

func NewCIM_DisplayControllerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DisplayController, err error) {
	tmp, err := NewCIM_ControllerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DisplayController{
		CIM_Controller: tmp,
	}
	return
}

// SetAcceleratorCapabilities sets the value of AcceleratorCapabilities for the instance
func (instance *CIM_DisplayController) SetPropertyAcceleratorCapabilities(value []DisplayController_AcceleratorCapabilities) (err error) {
	return instance.SetProperty("AcceleratorCapabilities", (value))
}

// GetAcceleratorCapabilities gets the value of AcceleratorCapabilities for the instance
func (instance *CIM_DisplayController) GetPropertyAcceleratorCapabilities() (value []DisplayController_AcceleratorCapabilities, err error) {
	retValue, err := instance.GetProperty("AcceleratorCapabilities")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DisplayController_AcceleratorCapabilities(valuetmp))
	}

	return
}

// SetCapabilityDescriptions sets the value of CapabilityDescriptions for the instance
func (instance *CIM_DisplayController) SetPropertyCapabilityDescriptions(value []string) (err error) {
	return instance.SetProperty("CapabilityDescriptions", (value))
}

// GetCapabilityDescriptions gets the value of CapabilityDescriptions for the instance
func (instance *CIM_DisplayController) GetPropertyCapabilityDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("CapabilityDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetMaxMemorySupported sets the value of MaxMemorySupported for the instance
func (instance *CIM_DisplayController) SetPropertyMaxMemorySupported(value uint32) (err error) {
	return instance.SetProperty("MaxMemorySupported", (value))
}

// GetMaxMemorySupported gets the value of MaxMemorySupported for the instance
func (instance *CIM_DisplayController) GetPropertyMaxMemorySupported() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMemorySupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNumberOfVideoPages sets the value of NumberOfVideoPages for the instance
func (instance *CIM_DisplayController) SetPropertyNumberOfVideoPages(value uint32) (err error) {
	return instance.SetProperty("NumberOfVideoPages", (value))
}

// GetNumberOfVideoPages gets the value of NumberOfVideoPages for the instance
func (instance *CIM_DisplayController) GetPropertyNumberOfVideoPages() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfVideoPages")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetOtherVideoArchitecture sets the value of OtherVideoArchitecture for the instance
func (instance *CIM_DisplayController) SetPropertyOtherVideoArchitecture(value string) (err error) {
	return instance.SetProperty("OtherVideoArchitecture", (value))
}

// GetOtherVideoArchitecture gets the value of OtherVideoArchitecture for the instance
func (instance *CIM_DisplayController) GetPropertyOtherVideoArchitecture() (value string, err error) {
	retValue, err := instance.GetProperty("OtherVideoArchitecture")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOtherVideoMemoryType sets the value of OtherVideoMemoryType for the instance
func (instance *CIM_DisplayController) SetPropertyOtherVideoMemoryType(value string) (err error) {
	return instance.SetProperty("OtherVideoMemoryType", (value))
}

// GetOtherVideoMemoryType gets the value of OtherVideoMemoryType for the instance
func (instance *CIM_DisplayController) GetPropertyOtherVideoMemoryType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherVideoMemoryType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVideoArchitecture sets the value of VideoArchitecture for the instance
func (instance *CIM_DisplayController) SetPropertyVideoArchitecture(value DisplayController_VideoArchitecture) (err error) {
	return instance.SetProperty("VideoArchitecture", (value))
}

// GetVideoArchitecture gets the value of VideoArchitecture for the instance
func (instance *CIM_DisplayController) GetPropertyVideoArchitecture() (value DisplayController_VideoArchitecture, err error) {
	retValue, err := instance.GetProperty("VideoArchitecture")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DisplayController_VideoArchitecture(valuetmp)

	return
}

// SetVideoMemoryType sets the value of VideoMemoryType for the instance
func (instance *CIM_DisplayController) SetPropertyVideoMemoryType(value DisplayController_VideoMemoryType) (err error) {
	return instance.SetProperty("VideoMemoryType", (value))
}

// GetVideoMemoryType gets the value of VideoMemoryType for the instance
func (instance *CIM_DisplayController) GetPropertyVideoMemoryType() (value DisplayController_VideoMemoryType, err error) {
	retValue, err := instance.GetProperty("VideoMemoryType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DisplayController_VideoMemoryType(valuetmp)

	return
}

// SetVideoProcessor sets the value of VideoProcessor for the instance
func (instance *CIM_DisplayController) SetPropertyVideoProcessor(value string) (err error) {
	return instance.SetProperty("VideoProcessor", (value))
}

// GetVideoProcessor gets the value of VideoProcessor for the instance
func (instance *CIM_DisplayController) GetPropertyVideoProcessor() (value string, err error) {
	retValue, err := instance.GetProperty("VideoProcessor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}
