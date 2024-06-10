// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAClientSettings struct
type DAClientSettings struct {
	*cim.WmiInstance

	//
	Downlevel string

	//
	ForceTunnel string

	//
	ForceTunnelingNrptEntry DnsClientNrptRule

	//
	OnlyRemoteComputers string
}

func NewDAClientSettingsEx1(instance *cim.WmiInstance) (newInstance *DAClientSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAClientSettings{
		WmiInstance: tmp,
	}
	return
}

func NewDAClientSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAClientSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAClientSettings{
		WmiInstance: tmp,
	}
	return
}

// SetDownlevel sets the value of Downlevel for the instance
func (instance *DAClientSettings) SetPropertyDownlevel(value string) (err error) {
	return instance.SetProperty("Downlevel", (value))
}

// GetDownlevel gets the value of Downlevel for the instance
func (instance *DAClientSettings) GetPropertyDownlevel() (value string, err error) {
	retValue, err := instance.GetProperty("Downlevel")
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

// SetForceTunnel sets the value of ForceTunnel for the instance
func (instance *DAClientSettings) SetPropertyForceTunnel(value string) (err error) {
	return instance.SetProperty("ForceTunnel", (value))
}

// GetForceTunnel gets the value of ForceTunnel for the instance
func (instance *DAClientSettings) GetPropertyForceTunnel() (value string, err error) {
	retValue, err := instance.GetProperty("ForceTunnel")
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

// SetForceTunnelingNrptEntry sets the value of ForceTunnelingNrptEntry for the instance
func (instance *DAClientSettings) SetPropertyForceTunnelingNrptEntry(value DnsClientNrptRule) (err error) {
	return instance.SetProperty("ForceTunnelingNrptEntry", (value))
}

// GetForceTunnelingNrptEntry gets the value of ForceTunnelingNrptEntry for the instance
func (instance *DAClientSettings) GetPropertyForceTunnelingNrptEntry() (value DnsClientNrptRule, err error) {
	retValue, err := instance.GetProperty("ForceTunnelingNrptEntry")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DnsClientNrptRule)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DnsClientNrptRule is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DnsClientNrptRule(valuetmp)

	return
}

// SetOnlyRemoteComputers sets the value of OnlyRemoteComputers for the instance
func (instance *DAClientSettings) SetPropertyOnlyRemoteComputers(value string) (err error) {
	return instance.SetProperty("OnlyRemoteComputers", (value))
}

// GetOnlyRemoteComputers gets the value of OnlyRemoteComputers for the instance
func (instance *DAClientSettings) GetPropertyOnlyRemoteComputers() (value string, err error) {
	retValue, err := instance.GetProperty("OnlyRemoteComputers")
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
