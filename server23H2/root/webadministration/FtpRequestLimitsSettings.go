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

// FtpRequestLimitsSettings struct
type FtpRequestLimitsSettings struct {
	*EmbeddedObject

	//
	MaxAllowedContentLength string

	//
	MaxUrl uint32
}

func NewFtpRequestLimitsSettingsEx1(instance *cim.WmiInstance) (newInstance *FtpRequestLimitsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpRequestLimitsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpRequestLimitsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpRequestLimitsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpRequestLimitsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetMaxAllowedContentLength sets the value of MaxAllowedContentLength for the instance
func (instance *FtpRequestLimitsSettings) SetPropertyMaxAllowedContentLength(value string) (err error) {
	return instance.SetProperty("MaxAllowedContentLength", (value))
}

// GetMaxAllowedContentLength gets the value of MaxAllowedContentLength for the instance
func (instance *FtpRequestLimitsSettings) GetPropertyMaxAllowedContentLength() (value string, err error) {
	retValue, err := instance.GetProperty("MaxAllowedContentLength")
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

// SetMaxUrl sets the value of MaxUrl for the instance
func (instance *FtpRequestLimitsSettings) SetPropertyMaxUrl(value uint32) (err error) {
	return instance.SetProperty("MaxUrl", (value))
}

// GetMaxUrl gets the value of MaxUrl for the instance
func (instance *FtpRequestLimitsSettings) GetPropertyMaxUrl() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxUrl")
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