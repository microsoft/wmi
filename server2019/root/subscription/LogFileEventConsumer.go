// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.subscription
//////////////////////////////////////////////
package subscription

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LogFileEventConsumer struct
type LogFileEventConsumer struct {
	*__EventConsumer

	//
	Filename string

	//
	IsUnicode bool

	//
	MaximumFileSize uint64

	//
	Name string

	//
	Text string
}

func NewLogFileEventConsumerEx1(instance *cim.WmiInstance) (newInstance *LogFileEventConsumer, err error) {
	tmp, err := New__EventConsumerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &LogFileEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

func NewLogFileEventConsumerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LogFileEventConsumer, err error) {
	tmp, err := New__EventConsumerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LogFileEventConsumer{
		__EventConsumer: tmp,
	}
	return
}

// SetFilename sets the value of Filename for the instance
func (instance *LogFileEventConsumer) SetPropertyFilename(value string) (err error) {
	return instance.SetProperty("Filename", (value))
}

// GetFilename gets the value of Filename for the instance
func (instance *LogFileEventConsumer) GetPropertyFilename() (value string, err error) {
	retValue, err := instance.GetProperty("Filename")
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

// SetIsUnicode sets the value of IsUnicode for the instance
func (instance *LogFileEventConsumer) SetPropertyIsUnicode(value bool) (err error) {
	return instance.SetProperty("IsUnicode", (value))
}

// GetIsUnicode gets the value of IsUnicode for the instance
func (instance *LogFileEventConsumer) GetPropertyIsUnicode() (value bool, err error) {
	retValue, err := instance.GetProperty("IsUnicode")
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

// SetMaximumFileSize sets the value of MaximumFileSize for the instance
func (instance *LogFileEventConsumer) SetPropertyMaximumFileSize(value uint64) (err error) {
	return instance.SetProperty("MaximumFileSize", (value))
}

// GetMaximumFileSize gets the value of MaximumFileSize for the instance
func (instance *LogFileEventConsumer) GetPropertyMaximumFileSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaximumFileSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *LogFileEventConsumer) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *LogFileEventConsumer) GetPropertyName() (value string, err error) {
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

// SetText sets the value of Text for the instance
func (instance *LogFileEventConsumer) SetPropertyText(value string) (err error) {
	return instance.SetProperty("Text", (value))
}

// GetText gets the value of Text for the instance
func (instance *LogFileEventConsumer) GetPropertyText() (value string, err error) {
	retValue, err := instance.GetProperty("Text")
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
