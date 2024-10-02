// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DSCConfigurationOutputWriteProgress struct
type MSFT_DSCConfigurationOutputWriteProgress struct {
	*MSFT_DSCConfigurationOutput

	//
	Activity string

	//
	CurrentOperation string

	//
	PercentComplete uint32

	//
	SecondsRemaining uint32

	//
	StatusDescription string
}

func NewMSFT_DSCConfigurationOutputWriteProgressEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutputWriteProgress, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteProgress{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationOutputWriteProgressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationOutputWriteProgress, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteProgress{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

// SetActivity sets the value of Activity for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyActivity(value string) (err error) {
	return instance.SetProperty("Activity", (value))
}

// GetActivity gets the value of Activity for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyActivity() (value string, err error) {
	retValue, err := instance.GetProperty("Activity")
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

// SetCurrentOperation sets the value of CurrentOperation for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyCurrentOperation(value string) (err error) {
	return instance.SetProperty("CurrentOperation", (value))
}

// GetCurrentOperation gets the value of CurrentOperation for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyCurrentOperation() (value string, err error) {
	retValue, err := instance.GetProperty("CurrentOperation")
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

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyPercentComplete(value uint32) (err error) {
	return instance.SetProperty("PercentComplete", (value))
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyPercentComplete() (value uint32, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
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

// SetSecondsRemaining sets the value of SecondsRemaining for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertySecondsRemaining(value uint32) (err error) {
	return instance.SetProperty("SecondsRemaining", (value))
}

// GetSecondsRemaining gets the value of SecondsRemaining for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertySecondsRemaining() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondsRemaining")
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

// SetStatusDescription sets the value of StatusDescription for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) SetPropertyStatusDescription(value string) (err error) {
	return instance.SetProperty("StatusDescription", (value))
}

// GetStatusDescription gets the value of StatusDescription for the instance
func (instance *MSFT_DSCConfigurationOutputWriteProgress) GetPropertyStatusDescription() (value string, err error) {
	retValue, err := instance.GetProperty("StatusDescription")
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
