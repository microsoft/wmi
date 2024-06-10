// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAOtpAuthentication struct
type DAOtpAuthentication struct {
	*cim.WmiInstance

	//
	CAServers []string

	//
	CertificateTemplateName string

	//
	OtpStatus string

	//
	RadiusServer []string

	//
	SigningCertificateTemplateName string

	//
	UserSecurityGroupName string
}

func NewDAOtpAuthenticationEx1(instance *cim.WmiInstance) (newInstance *DAOtpAuthentication, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAOtpAuthentication{
		WmiInstance: tmp,
	}
	return
}

func NewDAOtpAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAOtpAuthentication, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAOtpAuthentication{
		WmiInstance: tmp,
	}
	return
}

// SetCAServers sets the value of CAServers for the instance
func (instance *DAOtpAuthentication) SetPropertyCAServers(value []string) (err error) {
	return instance.SetProperty("CAServers", (value))
}

// GetCAServers gets the value of CAServers for the instance
func (instance *DAOtpAuthentication) GetPropertyCAServers() (value []string, err error) {
	retValue, err := instance.GetProperty("CAServers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetCertificateTemplateName sets the value of CertificateTemplateName for the instance
func (instance *DAOtpAuthentication) SetPropertyCertificateTemplateName(value string) (err error) {
	return instance.SetProperty("CertificateTemplateName", (value))
}

// GetCertificateTemplateName gets the value of CertificateTemplateName for the instance
func (instance *DAOtpAuthentication) GetPropertyCertificateTemplateName() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateTemplateName")
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

// SetOtpStatus sets the value of OtpStatus for the instance
func (instance *DAOtpAuthentication) SetPropertyOtpStatus(value string) (err error) {
	return instance.SetProperty("OtpStatus", (value))
}

// GetOtpStatus gets the value of OtpStatus for the instance
func (instance *DAOtpAuthentication) GetPropertyOtpStatus() (value string, err error) {
	retValue, err := instance.GetProperty("OtpStatus")
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

// SetRadiusServer sets the value of RadiusServer for the instance
func (instance *DAOtpAuthentication) SetPropertyRadiusServer(value []string) (err error) {
	return instance.SetProperty("RadiusServer", (value))
}

// GetRadiusServer gets the value of RadiusServer for the instance
func (instance *DAOtpAuthentication) GetPropertyRadiusServer() (value []string, err error) {
	retValue, err := instance.GetProperty("RadiusServer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSigningCertificateTemplateName sets the value of SigningCertificateTemplateName for the instance
func (instance *DAOtpAuthentication) SetPropertySigningCertificateTemplateName(value string) (err error) {
	return instance.SetProperty("SigningCertificateTemplateName", (value))
}

// GetSigningCertificateTemplateName gets the value of SigningCertificateTemplateName for the instance
func (instance *DAOtpAuthentication) GetPropertySigningCertificateTemplateName() (value string, err error) {
	retValue, err := instance.GetProperty("SigningCertificateTemplateName")
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

// SetUserSecurityGroupName sets the value of UserSecurityGroupName for the instance
func (instance *DAOtpAuthentication) SetPropertyUserSecurityGroupName(value string) (err error) {
	return instance.SetProperty("UserSecurityGroupName", (value))
}

// GetUserSecurityGroupName gets the value of UserSecurityGroupName for the instance
func (instance *DAOtpAuthentication) GetPropertyUserSecurityGroupName() (value string, err error) {
	retValue, err := instance.GetProperty("UserSecurityGroupName")
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
