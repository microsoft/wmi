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

// FtpSslClientCertificates struct
type FtpSslClientCertificates struct {
	*EmbeddedObject

	//
	ClientCertificatePolicy int32

	//
	RevocationFreshnessTime string

	//
	RevocationURLRetrievalTimeout string

	//
	UseActiveDirectoryMapping bool

	//
	ValidationFlags int32
}

func NewFtpSslClientCertificatesEx1(instance *cim.WmiInstance) (newInstance *FtpSslClientCertificates, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpSslClientCertificates{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpSslClientCertificatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpSslClientCertificates, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpSslClientCertificates{
		EmbeddedObject: tmp,
	}
	return
}

// SetClientCertificatePolicy sets the value of ClientCertificatePolicy for the instance
func (instance *FtpSslClientCertificates) SetPropertyClientCertificatePolicy(value int32) (err error) {
	return instance.SetProperty("ClientCertificatePolicy", (value))
}

// GetClientCertificatePolicy gets the value of ClientCertificatePolicy for the instance
func (instance *FtpSslClientCertificates) GetPropertyClientCertificatePolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("ClientCertificatePolicy")
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

// SetRevocationFreshnessTime sets the value of RevocationFreshnessTime for the instance
func (instance *FtpSslClientCertificates) SetPropertyRevocationFreshnessTime(value string) (err error) {
	return instance.SetProperty("RevocationFreshnessTime", (value))
}

// GetRevocationFreshnessTime gets the value of RevocationFreshnessTime for the instance
func (instance *FtpSslClientCertificates) GetPropertyRevocationFreshnessTime() (value string, err error) {
	retValue, err := instance.GetProperty("RevocationFreshnessTime")
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

// SetRevocationURLRetrievalTimeout sets the value of RevocationURLRetrievalTimeout for the instance
func (instance *FtpSslClientCertificates) SetPropertyRevocationURLRetrievalTimeout(value string) (err error) {
	return instance.SetProperty("RevocationURLRetrievalTimeout", (value))
}

// GetRevocationURLRetrievalTimeout gets the value of RevocationURLRetrievalTimeout for the instance
func (instance *FtpSslClientCertificates) GetPropertyRevocationURLRetrievalTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("RevocationURLRetrievalTimeout")
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

// SetUseActiveDirectoryMapping sets the value of UseActiveDirectoryMapping for the instance
func (instance *FtpSslClientCertificates) SetPropertyUseActiveDirectoryMapping(value bool) (err error) {
	return instance.SetProperty("UseActiveDirectoryMapping", (value))
}

// GetUseActiveDirectoryMapping gets the value of UseActiveDirectoryMapping for the instance
func (instance *FtpSslClientCertificates) GetPropertyUseActiveDirectoryMapping() (value bool, err error) {
	retValue, err := instance.GetProperty("UseActiveDirectoryMapping")
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

// SetValidationFlags sets the value of ValidationFlags for the instance
func (instance *FtpSslClientCertificates) SetPropertyValidationFlags(value int32) (err error) {
	return instance.SetProperty("ValidationFlags", (value))
}

// GetValidationFlags gets the value of ValidationFlags for the instance
func (instance *FtpSslClientCertificates) GetPropertyValidationFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("ValidationFlags")
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
