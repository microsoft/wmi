// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// NlServerAuth_End struct
type NlServerAuth_End struct {
	*NlServerAuth

	// Account
	Account string

	// Channel Type
	ChannelType uint32

	// Client
	Client string

	// Negotiated Flags
	NegotiatedFlags uint32

	// Status
	Status uint32
}

func NewNlServerAuth_EndEx1(instance *cim.WmiInstance) (newInstance *NlServerAuth_End, err error) {
	tmp, err := NewNlServerAuthEx1(instance)

	if err != nil {
		return
	}
	newInstance = &NlServerAuth_End{
		NlServerAuth: tmp,
	}
	return
}

func NewNlServerAuth_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *NlServerAuth_End, err error) {
	tmp, err := NewNlServerAuthEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &NlServerAuth_End{
		NlServerAuth: tmp,
	}
	return
}

// SetAccount sets the value of Account for the instance
func (instance *NlServerAuth_End) SetPropertyAccount(value string) (err error) {
	return instance.SetProperty("Account", (value))
}

// GetAccount gets the value of Account for the instance
func (instance *NlServerAuth_End) GetPropertyAccount() (value string, err error) {
	retValue, err := instance.GetProperty("Account")
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

// SetChannelType sets the value of ChannelType for the instance
func (instance *NlServerAuth_End) SetPropertyChannelType(value uint32) (err error) {
	return instance.SetProperty("ChannelType", (value))
}

// GetChannelType gets the value of ChannelType for the instance
func (instance *NlServerAuth_End) GetPropertyChannelType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ChannelType")
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

// SetClient sets the value of Client for the instance
func (instance *NlServerAuth_End) SetPropertyClient(value string) (err error) {
	return instance.SetProperty("Client", (value))
}

// GetClient gets the value of Client for the instance
func (instance *NlServerAuth_End) GetPropertyClient() (value string, err error) {
	retValue, err := instance.GetProperty("Client")
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

// SetNegotiatedFlags sets the value of NegotiatedFlags for the instance
func (instance *NlServerAuth_End) SetPropertyNegotiatedFlags(value uint32) (err error) {
	return instance.SetProperty("NegotiatedFlags", (value))
}

// GetNegotiatedFlags gets the value of NegotiatedFlags for the instance
func (instance *NlServerAuth_End) GetPropertyNegotiatedFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("NegotiatedFlags")
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

// SetStatus sets the value of Status for the instance
func (instance *NlServerAuth_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *NlServerAuth_End) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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
