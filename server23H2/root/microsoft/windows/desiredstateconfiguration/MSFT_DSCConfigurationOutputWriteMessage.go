// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.DesiredStateConfiguration
//
// ////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DSCConfigurationOutputWriteMessage struct
type MSFT_DSCConfigurationOutputWriteMessage struct {
	*MSFT_DSCConfigurationOutput

	//
	Channel uint32

	//
	Message string
}

func NewMSFT_DSCConfigurationOutputWriteMessageEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutputWriteMessage, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteMessage{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationOutputWriteMessageEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationOutputWriteMessage, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputWriteMessage{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

// SetChannel sets the value of Channel for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) SetPropertyChannel(value uint32) (err error) {
	return instance.SetProperty("Channel", (value))
}

// GetChannel gets the value of Channel for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) GetPropertyChannel() (value uint32, err error) {
	retValue, err := instance.GetProperty("Channel")
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

// SetMessage sets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) SetPropertyMessage(value string) (err error) {
	return instance.SetProperty("Message", (value))
}

// GetMessage gets the value of Message for the instance
func (instance *MSFT_DSCConfigurationOutputWriteMessage) GetPropertyMessage() (value string, err error) {
	retValue, err := instance.GetProperty("Message")
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
