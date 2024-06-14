// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.RSOP.NS5B3DCE03_EB70_4EA2_8019_E0415C3C4C7B.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEProxySettings struct
type RSOP_IEProxySettings struct {
	*cim.WmiInstance

	//
	enableProxy bool

	//
	ftpProxyServer string

	//
	gopherProxyServer string

	//
	httpProxyServer string

	//
	proxyOverride string

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	secureProxyServer string

	//
	socksProxyServer string

	//
	useSameProxy bool
}

func NewRSOP_IEProxySettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEProxySettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEProxySettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEProxySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEProxySettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEProxySettings{
		WmiInstance: tmp,
	}
	return
}

// SetenableProxy sets the value of enableProxy for the instance
func (instance *RSOP_IEProxySettings) SetPropertyenableProxy(value bool) (err error) {
	return instance.SetProperty("enableProxy", (value))
}

// GetenableProxy gets the value of enableProxy for the instance
func (instance *RSOP_IEProxySettings) GetPropertyenableProxy() (value bool, err error) {
	retValue, err := instance.GetProperty("enableProxy")
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

// SetftpProxyServer sets the value of ftpProxyServer for the instance
func (instance *RSOP_IEProxySettings) SetPropertyftpProxyServer(value string) (err error) {
	return instance.SetProperty("ftpProxyServer", (value))
}

// GetftpProxyServer gets the value of ftpProxyServer for the instance
func (instance *RSOP_IEProxySettings) GetPropertyftpProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("ftpProxyServer")
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

// SetgopherProxyServer sets the value of gopherProxyServer for the instance
func (instance *RSOP_IEProxySettings) SetPropertygopherProxyServer(value string) (err error) {
	return instance.SetProperty("gopherProxyServer", (value))
}

// GetgopherProxyServer gets the value of gopherProxyServer for the instance
func (instance *RSOP_IEProxySettings) GetPropertygopherProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("gopherProxyServer")
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

// SethttpProxyServer sets the value of httpProxyServer for the instance
func (instance *RSOP_IEProxySettings) SetPropertyhttpProxyServer(value string) (err error) {
	return instance.SetProperty("httpProxyServer", (value))
}

// GethttpProxyServer gets the value of httpProxyServer for the instance
func (instance *RSOP_IEProxySettings) GetPropertyhttpProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("httpProxyServer")
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

// SetproxyOverride sets the value of proxyOverride for the instance
func (instance *RSOP_IEProxySettings) SetPropertyproxyOverride(value string) (err error) {
	return instance.SetProperty("proxyOverride", (value))
}

// GetproxyOverride gets the value of proxyOverride for the instance
func (instance *RSOP_IEProxySettings) GetPropertyproxyOverride() (value string, err error) {
	retValue, err := instance.GetProperty("proxyOverride")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEProxySettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEProxySettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEProxySettings) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEProxySettings) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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

// SetsecureProxyServer sets the value of secureProxyServer for the instance
func (instance *RSOP_IEProxySettings) SetPropertysecureProxyServer(value string) (err error) {
	return instance.SetProperty("secureProxyServer", (value))
}

// GetsecureProxyServer gets the value of secureProxyServer for the instance
func (instance *RSOP_IEProxySettings) GetPropertysecureProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("secureProxyServer")
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

// SetsocksProxyServer sets the value of socksProxyServer for the instance
func (instance *RSOP_IEProxySettings) SetPropertysocksProxyServer(value string) (err error) {
	return instance.SetProperty("socksProxyServer", (value))
}

// GetsocksProxyServer gets the value of socksProxyServer for the instance
func (instance *RSOP_IEProxySettings) GetPropertysocksProxyServer() (value string, err error) {
	retValue, err := instance.GetProperty("socksProxyServer")
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

// SetuseSameProxy sets the value of useSameProxy for the instance
func (instance *RSOP_IEProxySettings) SetPropertyuseSameProxy(value bool) (err error) {
	return instance.SetProperty("useSameProxy", (value))
}

// GetuseSameProxy gets the value of useSameProxy for the instance
func (instance *RSOP_IEProxySettings) GetPropertyuseSameProxy() (value bool, err error) {
	retValue, err := instance.GetProperty("useSameProxy")
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
