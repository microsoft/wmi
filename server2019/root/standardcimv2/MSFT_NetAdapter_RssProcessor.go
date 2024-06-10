// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.StandardCimv2
//
// ////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapter_RssProcessor struct
type MSFT_NetAdapter_RssProcessor struct {
	*cim.WmiInstance

	//
	PreferenceIndex uint16

	//
	ProcessorGroup uint16

	//
	ProcessorNumber uint8
}

func NewMSFT_NetAdapter_RssProcessorEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapter_RssProcessor, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RssProcessor{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_NetAdapter_RssProcessorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapter_RssProcessor, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapter_RssProcessor{
		WmiInstance: tmp,
	}
	return
}

// SetPreferenceIndex sets the value of PreferenceIndex for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyPreferenceIndex(value uint16) (err error) {
	return instance.SetProperty("PreferenceIndex", (value))
}

// GetPreferenceIndex gets the value of PreferenceIndex for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyPreferenceIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("PreferenceIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetProcessorGroup sets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("ProcessorGroup", (value))
}

// GetProcessorGroup gets the value of ProcessorGroup for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetProcessorNumber sets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_RssProcessor) SetPropertyProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("ProcessorNumber", (value))
}

// GetProcessorNumber gets the value of ProcessorNumber for the instance
func (instance *MSFT_NetAdapter_RssProcessor) GetPropertyProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("ProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
