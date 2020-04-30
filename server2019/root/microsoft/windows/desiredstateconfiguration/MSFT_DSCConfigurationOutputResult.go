// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DSCConfigurationOutputResult struct
type MSFT_DSCConfigurationOutputResult struct {
	*MSFT_DSCConfigurationOutput

	//
	Object interface{}

	//
	Result uint32
}

func NewMSFT_DSCConfigurationOutputResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutputResult, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputResult{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationOutputResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationOutputResult, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputResult{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

// SetObject sets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputResult) SetPropertyObject(value interface{}) (err error) {
	return instance.SetProperty("Object", value)
}

// GetObject gets the value of Object for the instance
func (instance *MSFT_DSCConfigurationOutputResult) GetPropertyObject() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Object")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResult sets the value of Result for the instance
func (instance *MSFT_DSCConfigurationOutputResult) SetPropertyResult(value uint32) (err error) {
	return instance.SetProperty("Result", value)
}

// GetResult gets the value of Result for the instance
func (instance *MSFT_DSCConfigurationOutputResult) GetPropertyResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("Result")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}