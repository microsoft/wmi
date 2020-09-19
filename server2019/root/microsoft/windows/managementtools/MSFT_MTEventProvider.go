// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTEventProvider struct
type MSFT_MTEventProvider struct {
	*CIM_ManagedElement

	//
	DisplayName string

	//
	DisplayPath string

	//
	ExportedChannelsCount uint32

	//
	Name string
}

func NewMSFT_MTEventProviderEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTEventProvider, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTEventProvider{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTEventProviderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTEventProvider, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTEventProvider{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_MTEventProvider) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_MTEventProvider) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetDisplayPath sets the value of DisplayPath for the instance
func (instance *MSFT_MTEventProvider) SetPropertyDisplayPath(value string) (err error) {
	return instance.SetProperty("DisplayPath", (value))
}

// GetDisplayPath gets the value of DisplayPath for the instance
func (instance *MSFT_MTEventProvider) GetPropertyDisplayPath() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayPath")
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

// SetExportedChannelsCount sets the value of ExportedChannelsCount for the instance
func (instance *MSFT_MTEventProvider) SetPropertyExportedChannelsCount(value uint32) (err error) {
	return instance.SetProperty("ExportedChannelsCount", (value))
}

// GetExportedChannelsCount gets the value of ExportedChannelsCount for the instance
func (instance *MSFT_MTEventProvider) GetPropertyExportedChannelsCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExportedChannelsCount")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_MTEventProvider) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTEventProvider) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

//

// <param name="Result" type="MSFT_MTEventChannel []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTEventProvider) GetChannels( /* OUT */ Result []MSFT_MTEventChannel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetChannels")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EventProviders" type="MSFT_MTEventProvider []"></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="WindowsEventChannels" type="MSFT_MTEventChannel []"></param>
func (instance *MSFT_MTEventProvider) GetProvidersAndWindowsEventChannels( /* OUT */ EventProviders []MSFT_MTEventProvider,
	/* OUT */ WindowsEventChannels []MSFT_MTEventChannel) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetProvidersAndWindowsEventChannels")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
