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

// FtpMessages struct
type FtpMessages struct {
	*EmbeddedObject

	//
	AllowLocalDetailedErrors bool

	//
	BannerMessage string

	//
	ExitMessage string

	//
	ExpandVariables bool

	//
	GreetingMessage string

	//
	MaxClientsMessage string

	//
	SuppressDefaultBanner bool
}

func NewFtpMessagesEx1(instance *cim.WmiInstance) (newInstance *FtpMessages, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpMessages{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpMessagesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpMessages, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpMessages{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowLocalDetailedErrors sets the value of AllowLocalDetailedErrors for the instance
func (instance *FtpMessages) SetPropertyAllowLocalDetailedErrors(value bool) (err error) {
	return instance.SetProperty("AllowLocalDetailedErrors", (value))
}

// GetAllowLocalDetailedErrors gets the value of AllowLocalDetailedErrors for the instance
func (instance *FtpMessages) GetPropertyAllowLocalDetailedErrors() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowLocalDetailedErrors")
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

// SetBannerMessage sets the value of BannerMessage for the instance
func (instance *FtpMessages) SetPropertyBannerMessage(value string) (err error) {
	return instance.SetProperty("BannerMessage", (value))
}

// GetBannerMessage gets the value of BannerMessage for the instance
func (instance *FtpMessages) GetPropertyBannerMessage() (value string, err error) {
	retValue, err := instance.GetProperty("BannerMessage")
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

// SetExitMessage sets the value of ExitMessage for the instance
func (instance *FtpMessages) SetPropertyExitMessage(value string) (err error) {
	return instance.SetProperty("ExitMessage", (value))
}

// GetExitMessage gets the value of ExitMessage for the instance
func (instance *FtpMessages) GetPropertyExitMessage() (value string, err error) {
	retValue, err := instance.GetProperty("ExitMessage")
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

// SetExpandVariables sets the value of ExpandVariables for the instance
func (instance *FtpMessages) SetPropertyExpandVariables(value bool) (err error) {
	return instance.SetProperty("ExpandVariables", (value))
}

// GetExpandVariables gets the value of ExpandVariables for the instance
func (instance *FtpMessages) GetPropertyExpandVariables() (value bool, err error) {
	retValue, err := instance.GetProperty("ExpandVariables")
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

// SetGreetingMessage sets the value of GreetingMessage for the instance
func (instance *FtpMessages) SetPropertyGreetingMessage(value string) (err error) {
	return instance.SetProperty("GreetingMessage", (value))
}

// GetGreetingMessage gets the value of GreetingMessage for the instance
func (instance *FtpMessages) GetPropertyGreetingMessage() (value string, err error) {
	retValue, err := instance.GetProperty("GreetingMessage")
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

// SetMaxClientsMessage sets the value of MaxClientsMessage for the instance
func (instance *FtpMessages) SetPropertyMaxClientsMessage(value string) (err error) {
	return instance.SetProperty("MaxClientsMessage", (value))
}

// GetMaxClientsMessage gets the value of MaxClientsMessage for the instance
func (instance *FtpMessages) GetPropertyMaxClientsMessage() (value string, err error) {
	retValue, err := instance.GetProperty("MaxClientsMessage")
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

// SetSuppressDefaultBanner sets the value of SuppressDefaultBanner for the instance
func (instance *FtpMessages) SetPropertySuppressDefaultBanner(value bool) (err error) {
	return instance.SetProperty("SuppressDefaultBanner", (value))
}

// GetSuppressDefaultBanner gets the value of SuppressDefaultBanner for the instance
func (instance *FtpMessages) GetPropertySuppressDefaultBanner() (value bool, err error) {
	retValue, err := instance.GetProperty("SuppressDefaultBanner")
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
