// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_VideoHead struct
type CIM_VideoHead struct {
	*CIM_LogicalDevice

	// The number of bits used to display each pixel.
	CurrentBitsPerPixel uint32

	// Current number of horizontal pixels.
	CurrentHorizontalResolution uint32

	// Number of colors supported at the current resolutions.
	CurrentNumberOfColors uint64

	// If in character mode, number of columns for this DisplayController. Otherwise, enter 0.
	CurrentNumberOfColumns uint32

	// If in character mode, number of rows for this Video Controller. Otherwise, enter 0.
	CurrentNumberOfRows uint32

	// Current refresh rate in Hertz.
	CurrentRefreshRate uint32

	// Current scan mode.
	CurrentScanMode VideoHead_CurrentScanMode

	// Current number of vertical pixels.
	CurrentVerticalResolution uint32

	// Maximum refresh rate of the DisplayController in Hertz.
	MaxRefreshRate uint32

	// Minimum refresh rate of the Video Controller in Hertz.
	MinRefreshRate uint32

	// A string describing the current scan mode when the instance's CurrentScanMode property is 1 ("Other").
	OtherCurrentScanMode string
}

func NewCIM_VideoHeadEx1(instance *cim.WmiInstance) (newInstance *CIM_VideoHead, err error) {
	tmp, err := NewCIM_LogicalDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VideoHead{
		CIM_LogicalDevice: tmp,
	}
	return
}

func NewCIM_VideoHeadEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VideoHead, err error) {
	tmp, err := NewCIM_LogicalDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VideoHead{
		CIM_LogicalDevice: tmp,
	}
	return
}

// SetCurrentBitsPerPixel sets the value of CurrentBitsPerPixel for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentBitsPerPixel(value uint32) (err error) {
	return instance.SetProperty("CurrentBitsPerPixel", value)
}

// GetCurrentBitsPerPixel gets the value of CurrentBitsPerPixel for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentBitsPerPixel() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentBitsPerPixel")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentHorizontalResolution sets the value of CurrentHorizontalResolution for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentHorizontalResolution(value uint32) (err error) {
	return instance.SetProperty("CurrentHorizontalResolution", value)
}

// GetCurrentHorizontalResolution gets the value of CurrentHorizontalResolution for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentHorizontalResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentHorizontalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentNumberOfColors sets the value of CurrentNumberOfColors for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentNumberOfColors(value uint64) (err error) {
	return instance.SetProperty("CurrentNumberOfColors", value)
}

// GetCurrentNumberOfColors gets the value of CurrentNumberOfColors for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentNumberOfColors() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentNumberOfColors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentNumberOfColumns sets the value of CurrentNumberOfColumns for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentNumberOfColumns(value uint32) (err error) {
	return instance.SetProperty("CurrentNumberOfColumns", value)
}

// GetCurrentNumberOfColumns gets the value of CurrentNumberOfColumns for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentNumberOfColumns() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentNumberOfColumns")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentNumberOfRows sets the value of CurrentNumberOfRows for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentNumberOfRows(value uint32) (err error) {
	return instance.SetProperty("CurrentNumberOfRows", value)
}

// GetCurrentNumberOfRows gets the value of CurrentNumberOfRows for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentNumberOfRows() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentNumberOfRows")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentRefreshRate sets the value of CurrentRefreshRate for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentRefreshRate(value uint32) (err error) {
	return instance.SetProperty("CurrentRefreshRate", value)
}

// GetCurrentRefreshRate gets the value of CurrentRefreshRate for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentRefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentScanMode sets the value of CurrentScanMode for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentScanMode(value VideoHead_CurrentScanMode) (err error) {
	return instance.SetProperty("CurrentScanMode", value)
}

// GetCurrentScanMode gets the value of CurrentScanMode for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentScanMode() (value VideoHead_CurrentScanMode, err error) {
	retValue, err := instance.GetProperty("CurrentScanMode")
	if err != nil {
		return
	}
	value, ok := retValue.(VideoHead_CurrentScanMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentVerticalResolution sets the value of CurrentVerticalResolution for the instance
func (instance *CIM_VideoHead) SetPropertyCurrentVerticalResolution(value uint32) (err error) {
	return instance.SetProperty("CurrentVerticalResolution", value)
}

// GetCurrentVerticalResolution gets the value of CurrentVerticalResolution for the instance
func (instance *CIM_VideoHead) GetPropertyCurrentVerticalResolution() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentVerticalResolution")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxRefreshRate sets the value of MaxRefreshRate for the instance
func (instance *CIM_VideoHead) SetPropertyMaxRefreshRate(value uint32) (err error) {
	return instance.SetProperty("MaxRefreshRate", value)
}

// GetMaxRefreshRate gets the value of MaxRefreshRate for the instance
func (instance *CIM_VideoHead) GetPropertyMaxRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxRefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinRefreshRate sets the value of MinRefreshRate for the instance
func (instance *CIM_VideoHead) SetPropertyMinRefreshRate(value uint32) (err error) {
	return instance.SetProperty("MinRefreshRate", value)
}

// GetMinRefreshRate gets the value of MinRefreshRate for the instance
func (instance *CIM_VideoHead) GetPropertyMinRefreshRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinRefreshRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherCurrentScanMode sets the value of OtherCurrentScanMode for the instance
func (instance *CIM_VideoHead) SetPropertyOtherCurrentScanMode(value string) (err error) {
	return instance.SetProperty("OtherCurrentScanMode", value)
}

// GetOtherCurrentScanMode gets the value of OtherCurrentScanMode for the instance
func (instance *CIM_VideoHead) GetPropertyOtherCurrentScanMode() (value string, err error) {
	retValue, err := instance.GetProperty("OtherCurrentScanMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
