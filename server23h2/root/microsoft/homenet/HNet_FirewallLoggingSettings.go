// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HNet_FirewallLoggingSettings struct
type HNet_FirewallLoggingSettings struct {
	*cim.WmiInstance

	//
	Id string

	//
	LogConnections bool

	//
	LogDroppedPackets bool

	//
	MaxFileSize uint32

	//
	Path string
}

func NewHNet_FirewallLoggingSettingsEx1(instance *cim.WmiInstance) (newInstance *HNet_FirewallLoggingSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_FirewallLoggingSettings{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_FirewallLoggingSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_FirewallLoggingSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_FirewallLoggingSettings{
		WmiInstance: tmp,
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *HNet_FirewallLoggingSettings) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *HNet_FirewallLoggingSettings) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetLogConnections sets the value of LogConnections for the instance
func (instance *HNet_FirewallLoggingSettings) SetPropertyLogConnections(value bool) (err error) {
	return instance.SetProperty("LogConnections", (value))
}

// GetLogConnections gets the value of LogConnections for the instance
func (instance *HNet_FirewallLoggingSettings) GetPropertyLogConnections() (value bool, err error) {
	retValue, err := instance.GetProperty("LogConnections")
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

// SetLogDroppedPackets sets the value of LogDroppedPackets for the instance
func (instance *HNet_FirewallLoggingSettings) SetPropertyLogDroppedPackets(value bool) (err error) {
	return instance.SetProperty("LogDroppedPackets", (value))
}

// GetLogDroppedPackets gets the value of LogDroppedPackets for the instance
func (instance *HNet_FirewallLoggingSettings) GetPropertyLogDroppedPackets() (value bool, err error) {
	retValue, err := instance.GetProperty("LogDroppedPackets")
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

// SetMaxFileSize sets the value of MaxFileSize for the instance
func (instance *HNet_FirewallLoggingSettings) SetPropertyMaxFileSize(value uint32) (err error) {
	return instance.SetProperty("MaxFileSize", (value))
}

// GetMaxFileSize gets the value of MaxFileSize for the instance
func (instance *HNet_FirewallLoggingSettings) GetPropertyMaxFileSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFileSize")
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

// SetPath sets the value of Path for the instance
func (instance *HNet_FirewallLoggingSettings) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *HNet_FirewallLoggingSettings) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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
