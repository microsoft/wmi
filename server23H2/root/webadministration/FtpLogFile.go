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

// FtpLogFile struct
type FtpLogFile struct {
	*EmbeddedObject

	//
	Directory string

	//
	Enabled bool

	//
	LocalTimeRollover bool

	//
	LogExtFileFlags int32

	//
	Period int32

	//
	SelectiveLogging int32

	//
	TruncateSize string
}

func NewFtpLogFileEx1(instance *cim.WmiInstance) (newInstance *FtpLogFile, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpLogFile{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpLogFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpLogFile, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpLogFile{
		EmbeddedObject: tmp,
	}
	return
}

// SetDirectory sets the value of Directory for the instance
func (instance *FtpLogFile) SetPropertyDirectory(value string) (err error) {
	return instance.SetProperty("Directory", (value))
}

// GetDirectory gets the value of Directory for the instance
func (instance *FtpLogFile) GetPropertyDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("Directory")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *FtpLogFile) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *FtpLogFile) GetPropertyEnabled() (value bool, err error) {
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

// SetLocalTimeRollover sets the value of LocalTimeRollover for the instance
func (instance *FtpLogFile) SetPropertyLocalTimeRollover(value bool) (err error) {
	return instance.SetProperty("LocalTimeRollover", (value))
}

// GetLocalTimeRollover gets the value of LocalTimeRollover for the instance
func (instance *FtpLogFile) GetPropertyLocalTimeRollover() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalTimeRollover")
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

// SetLogExtFileFlags sets the value of LogExtFileFlags for the instance
func (instance *FtpLogFile) SetPropertyLogExtFileFlags(value int32) (err error) {
	return instance.SetProperty("LogExtFileFlags", (value))
}

// GetLogExtFileFlags gets the value of LogExtFileFlags for the instance
func (instance *FtpLogFile) GetPropertyLogExtFileFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("LogExtFileFlags")
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

// SetPeriod sets the value of Period for the instance
func (instance *FtpLogFile) SetPropertyPeriod(value int32) (err error) {
	return instance.SetProperty("Period", (value))
}

// GetPeriod gets the value of Period for the instance
func (instance *FtpLogFile) GetPropertyPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("Period")
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

// SetSelectiveLogging sets the value of SelectiveLogging for the instance
func (instance *FtpLogFile) SetPropertySelectiveLogging(value int32) (err error) {
	return instance.SetProperty("SelectiveLogging", (value))
}

// GetSelectiveLogging gets the value of SelectiveLogging for the instance
func (instance *FtpLogFile) GetPropertySelectiveLogging() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectiveLogging")
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

// SetTruncateSize sets the value of TruncateSize for the instance
func (instance *FtpLogFile) SetPropertyTruncateSize(value string) (err error) {
	return instance.SetProperty("TruncateSize", (value))
}

// GetTruncateSize gets the value of TruncateSize for the instance
func (instance *FtpLogFile) GetPropertyTruncateSize() (value string, err error) {
	retValue, err := instance.GetProperty("TruncateSize")
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