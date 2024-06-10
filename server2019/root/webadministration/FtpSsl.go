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

// FtpSsl struct
type FtpSsl struct {
	*EmbeddedObject

	//
	ControlChannelPolicy int32

	//
	DataChannelPolicy int32

	//
	ServerCertHash string

	//
	ServerCertStoreName string

	//
	Ssl128 bool
}

func NewFtpSslEx1(instance *cim.WmiInstance) (newInstance *FtpSsl, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpSsl{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpSslEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpSsl, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpSsl{
		EmbeddedObject: tmp,
	}
	return
}

// SetControlChannelPolicy sets the value of ControlChannelPolicy for the instance
func (instance *FtpSsl) SetPropertyControlChannelPolicy(value int32) (err error) {
	return instance.SetProperty("ControlChannelPolicy", (value))
}

// GetControlChannelPolicy gets the value of ControlChannelPolicy for the instance
func (instance *FtpSsl) GetPropertyControlChannelPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("ControlChannelPolicy")
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

// SetDataChannelPolicy sets the value of DataChannelPolicy for the instance
func (instance *FtpSsl) SetPropertyDataChannelPolicy(value int32) (err error) {
	return instance.SetProperty("DataChannelPolicy", (value))
}

// GetDataChannelPolicy gets the value of DataChannelPolicy for the instance
func (instance *FtpSsl) GetPropertyDataChannelPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("DataChannelPolicy")
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

// SetServerCertHash sets the value of ServerCertHash for the instance
func (instance *FtpSsl) SetPropertyServerCertHash(value string) (err error) {
	return instance.SetProperty("ServerCertHash", (value))
}

// GetServerCertHash gets the value of ServerCertHash for the instance
func (instance *FtpSsl) GetPropertyServerCertHash() (value string, err error) {
	retValue, err := instance.GetProperty("ServerCertHash")
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

// SetServerCertStoreName sets the value of ServerCertStoreName for the instance
func (instance *FtpSsl) SetPropertyServerCertStoreName(value string) (err error) {
	return instance.SetProperty("ServerCertStoreName", (value))
}

// GetServerCertStoreName gets the value of ServerCertStoreName for the instance
func (instance *FtpSsl) GetPropertyServerCertStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerCertStoreName")
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

// SetSsl128 sets the value of Ssl128 for the instance
func (instance *FtpSsl) SetPropertySsl128(value bool) (err error) {
	return instance.SetProperty("Ssl128", (value))
}

// GetSsl128 gets the value of Ssl128 for the instance
func (instance *FtpSsl) GetPropertySsl128() (value bool, err error) {
	retValue, err := instance.GetProperty("Ssl128")
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
