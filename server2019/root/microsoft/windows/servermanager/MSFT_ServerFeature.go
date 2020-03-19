// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ServerFeature struct
type MSFT_ServerFeature struct {
	*cim.WmiInstance

	//
	BpaModels []string

	//
	ConfigurationStatus uint8

	//
	DisplayName string

	//
	EventQuery string

	//
	Id int32

	//
	ParentName string

	//
	Services []string

	//
	State uint8

	//
	Type uint8

	//
	UniqueName string
}

func NewMSFT_ServerFeatureEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerFeature, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerFeature{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerFeatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerFeature, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerFeature{
		WmiInstance: tmp,
	}
	return
}

// SetBpaModels sets the value of BpaModels for the instance
func (instance *MSFT_ServerFeature) SetPropertyBpaModels(value []string) (err error) {
	return instance.SetProperty("BpaModels", value)
}

// GetBpaModels gets the value of BpaModels for the instance
func (instance *MSFT_ServerFeature) GetPropertyBpaModels() (value []string, err error) {
	retValue, err := instance.GetProperty("BpaModels")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationStatus sets the value of ConfigurationStatus for the instance
func (instance *MSFT_ServerFeature) SetPropertyConfigurationStatus(value uint8) (err error) {
	return instance.SetProperty("ConfigurationStatus", value)
}

// GetConfigurationStatus gets the value of ConfigurationStatus for the instance
func (instance *MSFT_ServerFeature) GetPropertyConfigurationStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("ConfigurationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_ServerFeature) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", value)
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_ServerFeature) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventQuery sets the value of EventQuery for the instance
func (instance *MSFT_ServerFeature) SetPropertyEventQuery(value string) (err error) {
	return instance.SetProperty("EventQuery", value)
}

// GetEventQuery gets the value of EventQuery for the instance
func (instance *MSFT_ServerFeature) GetPropertyEventQuery() (value string, err error) {
	retValue, err := instance.GetProperty("EventQuery")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSFT_ServerFeature) SetPropertyId(value int32) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ServerFeature) GetPropertyId() (value int32, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentName sets the value of ParentName for the instance
func (instance *MSFT_ServerFeature) SetPropertyParentName(value string) (err error) {
	return instance.SetProperty("ParentName", value)
}

// GetParentName gets the value of ParentName for the instance
func (instance *MSFT_ServerFeature) GetPropertyParentName() (value string, err error) {
	retValue, err := instance.GetProperty("ParentName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServices sets the value of Services for the instance
func (instance *MSFT_ServerFeature) SetPropertyServices(value []string) (err error) {
	return instance.SetProperty("Services", value)
}

// GetServices gets the value of Services for the instance
func (instance *MSFT_ServerFeature) GetPropertyServices() (value []string, err error) {
	retValue, err := instance.GetProperty("Services")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_ServerFeature) SetPropertyState(value uint8) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_ServerFeature) GetPropertyState() (value uint8, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_ServerFeature) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_ServerFeature) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueName sets the value of UniqueName for the instance
func (instance *MSFT_ServerFeature) SetPropertyUniqueName(value string) (err error) {
	return instance.SetProperty("UniqueName", value)
}

// GetUniqueName gets the value of UniqueName for the instance
func (instance *MSFT_ServerFeature) GetPropertyUniqueName() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
