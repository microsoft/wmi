// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_WebDownloadManager struct
type MSFT_WebDownloadManager struct {
	*OMI_ConfigurationDownloadManager

	//
	AllowUnsecureConnection bool

	//
	CertificateID string

	//
	ConfigurationNames []string

	//
	ProxyCredential MSFT_Credential

	//
	ProxyURL string

	//
	RegistrationKey string

	//
	ServerURL string
}

func NewMSFT_WebDownloadManagerEx1(instance *cim.WmiInstance) (newInstance *MSFT_WebDownloadManager, err error) {
	tmp, err := NewOMI_ConfigurationDownloadManagerEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_WebDownloadManager{
		OMI_ConfigurationDownloadManager: tmp,
	}
	return
}

func NewMSFT_WebDownloadManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_WebDownloadManager, err error) {
	tmp, err := NewOMI_ConfigurationDownloadManagerEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_WebDownloadManager{
		OMI_ConfigurationDownloadManager: tmp,
	}
	return
}

// SetAllowUnsecureConnection sets the value of AllowUnsecureConnection for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyAllowUnsecureConnection(value bool) (err error) {
	return instance.SetProperty("AllowUnsecureConnection", (value))
}

// GetAllowUnsecureConnection gets the value of AllowUnsecureConnection for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyAllowUnsecureConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnsecureConnection")
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

// SetCertificateID sets the value of CertificateID for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyCertificateID(value string) (err error) {
	return instance.SetProperty("CertificateID", (value))
}

// GetCertificateID gets the value of CertificateID for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyCertificateID() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateID")
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

// SetConfigurationNames sets the value of ConfigurationNames for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyConfigurationNames(value []string) (err error) {
	return instance.SetProperty("ConfigurationNames", (value))
}

// GetConfigurationNames gets the value of ConfigurationNames for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyConfigurationNames() (value []string, err error) {
	retValue, err := instance.GetProperty("ConfigurationNames")
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

// SetProxyCredential sets the value of ProxyCredential for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyProxyCredential(value MSFT_Credential) (err error) {
	return instance.SetProperty("ProxyCredential", (value))
}

// GetProxyCredential gets the value of ProxyCredential for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyProxyCredential() (value MSFT_Credential, err error) {
	retValue, err := instance.GetProperty("ProxyCredential")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_Credential)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_Credential is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_Credential(valuetmp)

	return
}

// SetProxyURL sets the value of ProxyURL for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyProxyURL(value string) (err error) {
	return instance.SetProperty("ProxyURL", (value))
}

// GetProxyURL gets the value of ProxyURL for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyProxyURL() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyURL")
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

// SetRegistrationKey sets the value of RegistrationKey for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyRegistrationKey(value string) (err error) {
	return instance.SetProperty("RegistrationKey", (value))
}

// GetRegistrationKey gets the value of RegistrationKey for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyRegistrationKey() (value string, err error) {
	retValue, err := instance.GetProperty("RegistrationKey")
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

// SetServerURL sets the value of ServerURL for the instance
func (instance *MSFT_WebDownloadManager) SetPropertyServerURL(value string) (err error) {
	return instance.SetProperty("ServerURL", (value))
}

// GetServerURL gets the value of ServerURL for the instance
func (instance *MSFT_WebDownloadManager) GetPropertyServerURL() (value string, err error) {
	retValue, err := instance.GetProperty("ServerURL")
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
