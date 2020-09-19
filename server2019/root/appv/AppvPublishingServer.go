// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Appv
//////////////////////////////////////////////
package appv

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppvPublishingServer struct
type AppvPublishingServer struct {
	*cim.WmiInstance

	//
	GlobalRefreshEnabled bool

	//
	GlobalRefreshInterval uint32

	//
	GlobalRefreshIntervalUnit string

	//
	GlobalRefreshOnLogon bool

	//
	ID uint32

	//
	SetByGroupPolicy bool

	//
	Url string

	//
	UserRefreshEnabled bool

	//
	UserRefreshInterval uint32

	//
	UserRefreshIntervalUnit string

	//
	UserRefreshOnLogon bool
}

func NewAppvPublishingServerEx1(instance *cim.WmiInstance) (newInstance *AppvPublishingServer, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppvPublishingServer{
		WmiInstance: tmp,
	}
	return
}

func NewAppvPublishingServerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppvPublishingServer, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppvPublishingServer{
		WmiInstance: tmp,
	}
	return
}

// SetGlobalRefreshEnabled sets the value of GlobalRefreshEnabled for the instance
func (instance *AppvPublishingServer) SetPropertyGlobalRefreshEnabled(value bool) (err error) {
	return instance.SetProperty("GlobalRefreshEnabled", (value))
}

// GetGlobalRefreshEnabled gets the value of GlobalRefreshEnabled for the instance
func (instance *AppvPublishingServer) GetPropertyGlobalRefreshEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalRefreshEnabled")
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

// SetGlobalRefreshInterval sets the value of GlobalRefreshInterval for the instance
func (instance *AppvPublishingServer) SetPropertyGlobalRefreshInterval(value uint32) (err error) {
	return instance.SetProperty("GlobalRefreshInterval", (value))
}

// GetGlobalRefreshInterval gets the value of GlobalRefreshInterval for the instance
func (instance *AppvPublishingServer) GetPropertyGlobalRefreshInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("GlobalRefreshInterval")
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

// SetGlobalRefreshIntervalUnit sets the value of GlobalRefreshIntervalUnit for the instance
func (instance *AppvPublishingServer) SetPropertyGlobalRefreshIntervalUnit(value string) (err error) {
	return instance.SetProperty("GlobalRefreshIntervalUnit", (value))
}

// GetGlobalRefreshIntervalUnit gets the value of GlobalRefreshIntervalUnit for the instance
func (instance *AppvPublishingServer) GetPropertyGlobalRefreshIntervalUnit() (value string, err error) {
	retValue, err := instance.GetProperty("GlobalRefreshIntervalUnit")
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

// SetGlobalRefreshOnLogon sets the value of GlobalRefreshOnLogon for the instance
func (instance *AppvPublishingServer) SetPropertyGlobalRefreshOnLogon(value bool) (err error) {
	return instance.SetProperty("GlobalRefreshOnLogon", (value))
}

// GetGlobalRefreshOnLogon gets the value of GlobalRefreshOnLogon for the instance
func (instance *AppvPublishingServer) GetPropertyGlobalRefreshOnLogon() (value bool, err error) {
	retValue, err := instance.GetProperty("GlobalRefreshOnLogon")
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

// SetID sets the value of ID for the instance
func (instance *AppvPublishingServer) SetPropertyID(value uint32) (err error) {
	return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *AppvPublishingServer) GetPropertyID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ID")
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

// SetSetByGroupPolicy sets the value of SetByGroupPolicy for the instance
func (instance *AppvPublishingServer) SetPropertySetByGroupPolicy(value bool) (err error) {
	return instance.SetProperty("SetByGroupPolicy", (value))
}

// GetSetByGroupPolicy gets the value of SetByGroupPolicy for the instance
func (instance *AppvPublishingServer) GetPropertySetByGroupPolicy() (value bool, err error) {
	retValue, err := instance.GetProperty("SetByGroupPolicy")
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

// SetUrl sets the value of Url for the instance
func (instance *AppvPublishingServer) SetPropertyUrl(value string) (err error) {
	return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *AppvPublishingServer) GetPropertyUrl() (value string, err error) {
	retValue, err := instance.GetProperty("Url")
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

// SetUserRefreshEnabled sets the value of UserRefreshEnabled for the instance
func (instance *AppvPublishingServer) SetPropertyUserRefreshEnabled(value bool) (err error) {
	return instance.SetProperty("UserRefreshEnabled", (value))
}

// GetUserRefreshEnabled gets the value of UserRefreshEnabled for the instance
func (instance *AppvPublishingServer) GetPropertyUserRefreshEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("UserRefreshEnabled")
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

// SetUserRefreshInterval sets the value of UserRefreshInterval for the instance
func (instance *AppvPublishingServer) SetPropertyUserRefreshInterval(value uint32) (err error) {
	return instance.SetProperty("UserRefreshInterval", (value))
}

// GetUserRefreshInterval gets the value of UserRefreshInterval for the instance
func (instance *AppvPublishingServer) GetPropertyUserRefreshInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserRefreshInterval")
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

// SetUserRefreshIntervalUnit sets the value of UserRefreshIntervalUnit for the instance
func (instance *AppvPublishingServer) SetPropertyUserRefreshIntervalUnit(value string) (err error) {
	return instance.SetProperty("UserRefreshIntervalUnit", (value))
}

// GetUserRefreshIntervalUnit gets the value of UserRefreshIntervalUnit for the instance
func (instance *AppvPublishingServer) GetPropertyUserRefreshIntervalUnit() (value string, err error) {
	retValue, err := instance.GetProperty("UserRefreshIntervalUnit")
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

// SetUserRefreshOnLogon sets the value of UserRefreshOnLogon for the instance
func (instance *AppvPublishingServer) SetPropertyUserRefreshOnLogon(value bool) (err error) {
	return instance.SetProperty("UserRefreshOnLogon", (value))
}

// GetUserRefreshOnLogon gets the value of UserRefreshOnLogon for the instance
func (instance *AppvPublishingServer) GetPropertyUserRefreshOnLogon() (value bool, err error) {
	retValue, err := instance.GetProperty("UserRefreshOnLogon")
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
