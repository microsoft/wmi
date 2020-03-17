// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_RssProcessor struct
type MSFT_NetAdapter_RssProcessor struct {
	cim.WmiInstance

	//
	PreferenceIndex uint16

	//
	ProcessorGroup uint16

	//
	ProcessorNumber uint8
}

// SetPreferenceIndex sets the value of PreferenceIndex for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyPreferenceIndex(value uint16) (err error) {
	return instance.SetProperty("PreferenceIndex", value)
}

// GetPreferenceIndex gets the value of PreferenceIndex for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyPreferenceIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("PreferenceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorGroup sets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", value)
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("ProcessorNumber", value)
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("ProcessorNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
