// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_WdacBidTrace struct
type MSFT_WdacBidTrace struct {
	cim.WmiInstance

	//
	BidTraceAdapter string

	//
	IsDefined bool

	//
	IsEnabled bool

	//
	Mode uint32

	//
	PathOrFolder string

	//
	Platform string

	//
	ProcessId uint32
}

// SetBidTraceAdapter sets the value of BidTraceAdapter for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyBidTraceAdapter(value string) (err error) {
	return instance.SetProperty("BidTraceAdapter", value)
}

// GetBidTraceAdapter gets the value of BidTraceAdapter for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyBidTraceAdapter() (value string, err error) {
	retValue, err := instance.GetProperty("BidTraceAdapter")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsDefined sets the value of IsDefined for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyIsDefined(value bool) (err error) {
	return instance.SetProperty("IsDefined", value)
}

// GetIsDefined gets the value of IsDefined for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyIsDefined() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDefined")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", value)
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyMode(value uint32) (err error) {
	return instance.SetProperty("Mode", value)
}

// GetMode gets the value of Mode for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPathOrFolder sets the value of PathOrFolder for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyPathOrFolder(value string) (err error) {
	return instance.SetProperty("PathOrFolder", value)
}

// GetPathOrFolder gets the value of PathOrFolder for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyPathOrFolder() (value string, err error) {
	retValue, err := instance.GetProperty("PathOrFolder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", value)
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("Platform")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessId sets the value of ProcessId for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", value)
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
