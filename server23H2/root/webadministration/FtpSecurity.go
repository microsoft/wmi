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

// FtpSecurity struct
type FtpSecurity struct {
	*EmbeddedObject

	//
	Authentication FtpAuthentication

	//
	CommandFiltering FtpCommandFiltering

	//
	CustomAuthorization FtpCustomAuthorization

	//
	DataChannelSecurity FtpDataChannelSecurity

	//
	Ssl FtpSsl

	//
	SslClientCertificates FtpSslClientCertificates
}

func NewFtpSecurityEx1(instance *cim.WmiInstance) (newInstance *FtpSecurity, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpSecurity{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpSecurityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpSecurity, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpSecurity{
		EmbeddedObject: tmp,
	}
	return
}

// SetAuthentication sets the value of Authentication for the instance
func (instance *FtpSecurity) SetPropertyAuthentication(value FtpAuthentication) (err error) {
	return instance.SetProperty("Authentication", (value))
}

// GetAuthentication gets the value of Authentication for the instance
func (instance *FtpSecurity) GetPropertyAuthentication() (value FtpAuthentication, err error) {
	retValue, err := instance.GetProperty("Authentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpAuthentication(valuetmp)

	return
}

// SetCommandFiltering sets the value of CommandFiltering for the instance
func (instance *FtpSecurity) SetPropertyCommandFiltering(value FtpCommandFiltering) (err error) {
	return instance.SetProperty("CommandFiltering", (value))
}

// GetCommandFiltering gets the value of CommandFiltering for the instance
func (instance *FtpSecurity) GetPropertyCommandFiltering() (value FtpCommandFiltering, err error) {
	retValue, err := instance.GetProperty("CommandFiltering")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCommandFiltering)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCommandFiltering is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCommandFiltering(valuetmp)

	return
}

// SetCustomAuthorization sets the value of CustomAuthorization for the instance
func (instance *FtpSecurity) SetPropertyCustomAuthorization(value FtpCustomAuthorization) (err error) {
	return instance.SetProperty("CustomAuthorization", (value))
}

// GetCustomAuthorization gets the value of CustomAuthorization for the instance
func (instance *FtpSecurity) GetPropertyCustomAuthorization() (value FtpCustomAuthorization, err error) {
	retValue, err := instance.GetProperty("CustomAuthorization")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCustomAuthorization)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCustomAuthorization is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCustomAuthorization(valuetmp)

	return
}

// SetDataChannelSecurity sets the value of DataChannelSecurity for the instance
func (instance *FtpSecurity) SetPropertyDataChannelSecurity(value FtpDataChannelSecurity) (err error) {
	return instance.SetProperty("DataChannelSecurity", (value))
}

// GetDataChannelSecurity gets the value of DataChannelSecurity for the instance
func (instance *FtpSecurity) GetPropertyDataChannelSecurity() (value FtpDataChannelSecurity, err error) {
	retValue, err := instance.GetProperty("DataChannelSecurity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpDataChannelSecurity)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpDataChannelSecurity is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpDataChannelSecurity(valuetmp)

	return
}

// SetSsl sets the value of Ssl for the instance
func (instance *FtpSecurity) SetPropertySsl(value FtpSsl) (err error) {
	return instance.SetProperty("Ssl", (value))
}

// GetSsl gets the value of Ssl for the instance
func (instance *FtpSecurity) GetPropertySsl() (value FtpSsl, err error) {
	retValue, err := instance.GetProperty("Ssl")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpSsl)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpSsl is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpSsl(valuetmp)

	return
}

// SetSslClientCertificates sets the value of SslClientCertificates for the instance
func (instance *FtpSecurity) SetPropertySslClientCertificates(value FtpSslClientCertificates) (err error) {
	return instance.SetProperty("SslClientCertificates", (value))
}

// GetSslClientCertificates gets the value of SslClientCertificates for the instance
func (instance *FtpSecurity) GetPropertySslClientCertificates() (value FtpSslClientCertificates, err error) {
	retValue, err := instance.GetProperty("SslClientCertificates")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpSslClientCertificates)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpSslClientCertificates is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpSslClientCertificates(valuetmp)

	return
}
