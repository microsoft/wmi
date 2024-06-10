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

// FtpLogSection struct
type FtpLogSection struct {
	*ConfigurationSection

	//
	CentralLogFile CentralLogFileSettings

	//
	CentralLogFileMode int32

	//
	LogInUTF8 bool
}

func NewFtpLogSectionEx1(instance *cim.WmiInstance) (newInstance *FtpLogSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpLogSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewFtpLogSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpLogSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpLogSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCentralLogFile sets the value of CentralLogFile for the instance
func (instance *FtpLogSection) SetPropertyCentralLogFile(value CentralLogFileSettings) (err error) {
	return instance.SetProperty("CentralLogFile", (value))
}

// GetCentralLogFile gets the value of CentralLogFile for the instance
func (instance *FtpLogSection) GetPropertyCentralLogFile() (value CentralLogFileSettings, err error) {
	retValue, err := instance.GetProperty("CentralLogFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CentralLogFileSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CentralLogFileSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CentralLogFileSettings(valuetmp)

	return
}

// SetCentralLogFileMode sets the value of CentralLogFileMode for the instance
func (instance *FtpLogSection) SetPropertyCentralLogFileMode(value int32) (err error) {
	return instance.SetProperty("CentralLogFileMode", (value))
}

// GetCentralLogFileMode gets the value of CentralLogFileMode for the instance
func (instance *FtpLogSection) GetPropertyCentralLogFileMode() (value int32, err error) {
	retValue, err := instance.GetProperty("CentralLogFileMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetLogInUTF8 sets the value of LogInUTF8 for the instance
func (instance *FtpLogSection) SetPropertyLogInUTF8(value bool) (err error) {
	return instance.SetProperty("LogInUTF8", (value))
}

// GetLogInUTF8 gets the value of LogInUTF8 for the instance
func (instance *FtpLogSection) GetPropertyLogInUTF8() (value bool, err error) {
	retValue, err := instance.GetProperty("LogInUTF8")
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
