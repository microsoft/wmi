// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpDenyByFailureSettings struct
type FtpDenyByFailureSettings struct {
	*EmbeddedObject

	//
	Enabled bool

	//
	EntryExpiration string

	//
	LoggingOnlyMode bool

	//
	MaxFailure uint32
}

func NewFtpDenyByFailureSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpDenyByFailureSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpDenyByFailureSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpDenyByFailureSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpDenyByFailureSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpDenyByFailureSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *FtpDenyByFailureSettings) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *FtpDenyByFailureSettings) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetEntryExpiration sets the value of EntryExpiration for the instance
func (instance *FtpDenyByFailureSettings) SetPropertyEntryExpiration(value string) (err error) {
	return instance.SetProperty("EntryExpiration", (value))
}

// GetEntryExpiration gets the value of EntryExpiration for the instance
func (instance *FtpDenyByFailureSettings) GetPropertyEntryExpiration() (value string, err error) {
	retValue, err := instance.GetProperty("EntryExpiration")
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

// SetLoggingOnlyMode sets the value of LoggingOnlyMode for the instance
func (instance *FtpDenyByFailureSettings) SetPropertyLoggingOnlyMode(value bool) (err error) {
	return instance.SetProperty("LoggingOnlyMode", (value))
}

// GetLoggingOnlyMode gets the value of LoggingOnlyMode for the instance
func (instance *FtpDenyByFailureSettings) GetPropertyLoggingOnlyMode() (value bool, err error) {
	retValue, err := instance.GetProperty("LoggingOnlyMode")
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

// SetMaxFailure sets the value of MaxFailure for the instance
func (instance *FtpDenyByFailureSettings) SetPropertyMaxFailure(value uint32) (err error) {
	return instance.SetProperty("MaxFailure", (value))
}

// GetMaxFailure gets the value of MaxFailure for the instance
func (instance *FtpDenyByFailureSettings) GetPropertyMaxFailure() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFailure")
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
