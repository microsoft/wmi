// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DtcNetworkSettings struct
type DtcNetworkSettings struct {
	*cim.WmiInstance

	//
	AuthenticationLevel string

	//
	InboundTransactionsEnabled bool

	//
	LUTransactionsEnabled bool

	//
	OutboundTransactionsEnabled bool

	//
	RemoteAdministrationAccessEnabled bool

	//
	RemoteClientAccessEnabled bool

	//
	XATransactionsEnabled bool
}

func NewDtcNetworkSettingsEx1(instance *cim.WmiInstance) (newInstance *DtcNetworkSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcNetworkSettings{
		WmiInstance: tmp,
	}
	return
}

func NewDtcNetworkSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcNetworkSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcNetworkSettings{
		WmiInstance: tmp,
	}
	return
}

// SetAuthenticationLevel sets the value of AuthenticationLevel for the instance
func (instance *DtcNetworkSettings) SetPropertyAuthenticationLevel(value string) (err error) {
	return instance.SetProperty("AuthenticationLevel", (value))
}

// GetAuthenticationLevel gets the value of AuthenticationLevel for the instance
func (instance *DtcNetworkSettings) GetPropertyAuthenticationLevel() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationLevel")
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

// SetInboundTransactionsEnabled sets the value of InboundTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyInboundTransactionsEnabled(value bool) (err error) {
	return instance.SetProperty("InboundTransactionsEnabled", (value))
}

// GetInboundTransactionsEnabled gets the value of InboundTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyInboundTransactionsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("InboundTransactionsEnabled")
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

// SetLUTransactionsEnabled sets the value of LUTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyLUTransactionsEnabled(value bool) (err error) {
	return instance.SetProperty("LUTransactionsEnabled", (value))
}

// GetLUTransactionsEnabled gets the value of LUTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyLUTransactionsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("LUTransactionsEnabled")
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

// SetOutboundTransactionsEnabled sets the value of OutboundTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyOutboundTransactionsEnabled(value bool) (err error) {
	return instance.SetProperty("OutboundTransactionsEnabled", (value))
}

// GetOutboundTransactionsEnabled gets the value of OutboundTransactionsEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyOutboundTransactionsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("OutboundTransactionsEnabled")
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

// SetRemoteAdministrationAccessEnabled sets the value of RemoteAdministrationAccessEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyRemoteAdministrationAccessEnabled(value bool) (err error) {
	return instance.SetProperty("RemoteAdministrationAccessEnabled", (value))
}

// GetRemoteAdministrationAccessEnabled gets the value of RemoteAdministrationAccessEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyRemoteAdministrationAccessEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoteAdministrationAccessEnabled")
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

// SetRemoteClientAccessEnabled sets the value of RemoteClientAccessEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyRemoteClientAccessEnabled(value bool) (err error) {
	return instance.SetProperty("RemoteClientAccessEnabled", (value))
}

// GetRemoteClientAccessEnabled gets the value of RemoteClientAccessEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyRemoteClientAccessEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("RemoteClientAccessEnabled")
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

// SetXATransactionsEnabled sets the value of XATransactionsEnabled for the instance
func (instance *DtcNetworkSettings) SetPropertyXATransactionsEnabled(value bool) (err error) {
	return instance.SetProperty("XATransactionsEnabled", (value))
}

// GetXATransactionsEnabled gets the value of XATransactionsEnabled for the instance
func (instance *DtcNetworkSettings) GetPropertyXATransactionsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("XATransactionsEnabled")
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
