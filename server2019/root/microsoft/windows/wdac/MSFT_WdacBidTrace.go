// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Wdac
//
// ////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_WdacBidTrace struct
type MSFT_WdacBidTrace struct {
	*cim.WmiInstance

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

func NewMSFT_WdacBidTraceEx1(instance *cim.WmiInstance) (newInstance *MSFT_WdacBidTrace, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_WdacBidTrace{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_WdacBidTraceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WdacBidTrace, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WdacBidTrace{
		WmiInstance: tmp,
	}
	return
}

// SetBidTraceAdapter sets the value of BidTraceAdapter for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyBidTraceAdapter(value string) (err error) {
	return instance.SetProperty("BidTraceAdapter", (value))
}

// GetBidTraceAdapter gets the value of BidTraceAdapter for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyBidTraceAdapter() (value string, err error) {
	retValue, err := instance.GetProperty("BidTraceAdapter")
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

// SetIsDefined sets the value of IsDefined for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyIsDefined(value bool) (err error) {
	return instance.SetProperty("IsDefined", (value))
}

// GetIsDefined gets the value of IsDefined for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyIsDefined() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDefined")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", (value))
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMode sets the value of Mode for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyMode(value uint32) (err error) {
	return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Mode")
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

// SetPathOrFolder sets the value of PathOrFolder for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyPathOrFolder(value string) (err error) {
	return instance.SetProperty("PathOrFolder", (value))
}

// GetPathOrFolder gets the value of PathOrFolder for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyPathOrFolder() (value string, err error) {
	retValue, err := instance.GetProperty("PathOrFolder")
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

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", (value))
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("Platform")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *MSFT_WdacBidTrace) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *MSFT_WdacBidTrace) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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
