// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LogSection struct
type LogSection struct {
	*ConfigurationSection

	//
	CentralBinaryLogFile CentralBinaryLogFile

	//
	CentralLogFileMode int32

	//
	CentralW3CLogFile CentralW3CLogFile

	//
	LogInUTF8 bool
}

func NewLogSectionEx1(instance *cim.WmiInstance) (newInstance *LogSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LogSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewLogSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LogSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LogSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetCentralBinaryLogFile sets the value of CentralBinaryLogFile for the instance
func (instance *LogSection) SetPropertyCentralBinaryLogFile(value CentralBinaryLogFile) (err error) {
	return instance.SetProperty("CentralBinaryLogFile", (value))
}

// GetCentralBinaryLogFile gets the value of CentralBinaryLogFile for the instance
func (instance *LogSection) GetPropertyCentralBinaryLogFile() (value CentralBinaryLogFile, err error) {
	retValue, err := instance.GetProperty("CentralBinaryLogFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CentralBinaryLogFile)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CentralBinaryLogFile is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CentralBinaryLogFile(valuetmp)

	return
}

// SetCentralLogFileMode sets the value of CentralLogFileMode for the instance
func (instance *LogSection) SetPropertyCentralLogFileMode(value int32) (err error) {
	return instance.SetProperty("CentralLogFileMode", (value))
}

// GetCentralLogFileMode gets the value of CentralLogFileMode for the instance
func (instance *LogSection) GetPropertyCentralLogFileMode() (value int32, err error) {
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

// SetCentralW3CLogFile sets the value of CentralW3CLogFile for the instance
func (instance *LogSection) SetPropertyCentralW3CLogFile(value CentralW3CLogFile) (err error) {
	return instance.SetProperty("CentralW3CLogFile", (value))
}

// GetCentralW3CLogFile gets the value of CentralW3CLogFile for the instance
func (instance *LogSection) GetPropertyCentralW3CLogFile() (value CentralW3CLogFile, err error) {
	retValue, err := instance.GetProperty("CentralW3CLogFile")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CentralW3CLogFile)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CentralW3CLogFile is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CentralW3CLogFile(valuetmp)

	return
}

// SetLogInUTF8 sets the value of LogInUTF8 for the instance
func (instance *LogSection) SetPropertyLogInUTF8(value bool) (err error) {
	return instance.SetProperty("LogInUTF8", (value))
}

// GetLogInUTF8 gets the value of LogInUTF8 for the instance
func (instance *LogSection) GetPropertyLogInUTF8() (value bool, err error) {
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
