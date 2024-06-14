// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.WebApplicationProxy
//////////////////////////////////////////////
package webapplicationproxy

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AppProxyGlobalConfiguration struct
type AppProxyGlobalConfiguration struct {
	*cim.WmiInstance

	//
	ADFSSignOutUrl string

	//
	ADFSTokenAcceptanceDurationSec uint32

	//
	ADFSTokenSigningCertificatePublicKey string

	//
	ADFSUrl string

	//
	ADFSWebApplicationProxyRelyingPartyUri string

	//
	ConfigurationChangesPollingIntervalSec uint32

	//
	ConfigurationVersion string

	//
	ConnectedServersName []string

	//
	OAuthAuthenticationURL string

	//
	UserIdleTimeoutAction string

	//
	UserIdleTimeoutSec uint32
}

func NewAppProxyGlobalConfigurationEx1(instance *cim.WmiInstance) (newInstance *AppProxyGlobalConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AppProxyGlobalConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewAppProxyGlobalConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AppProxyGlobalConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AppProxyGlobalConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetADFSSignOutUrl sets the value of ADFSSignOutUrl for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyADFSSignOutUrl(value string) (err error) {
	return instance.SetProperty("ADFSSignOutUrl", (value))
}

// GetADFSSignOutUrl gets the value of ADFSSignOutUrl for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyADFSSignOutUrl() (value string, err error) {
	retValue, err := instance.GetProperty("ADFSSignOutUrl")
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

// SetADFSTokenAcceptanceDurationSec sets the value of ADFSTokenAcceptanceDurationSec for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyADFSTokenAcceptanceDurationSec(value uint32) (err error) {
	return instance.SetProperty("ADFSTokenAcceptanceDurationSec", (value))
}

// GetADFSTokenAcceptanceDurationSec gets the value of ADFSTokenAcceptanceDurationSec for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyADFSTokenAcceptanceDurationSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ADFSTokenAcceptanceDurationSec")
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

// SetADFSTokenSigningCertificatePublicKey sets the value of ADFSTokenSigningCertificatePublicKey for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyADFSTokenSigningCertificatePublicKey(value string) (err error) {
	return instance.SetProperty("ADFSTokenSigningCertificatePublicKey", (value))
}

// GetADFSTokenSigningCertificatePublicKey gets the value of ADFSTokenSigningCertificatePublicKey for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyADFSTokenSigningCertificatePublicKey() (value string, err error) {
	retValue, err := instance.GetProperty("ADFSTokenSigningCertificatePublicKey")
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

// SetADFSUrl sets the value of ADFSUrl for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyADFSUrl(value string) (err error) {
	return instance.SetProperty("ADFSUrl", (value))
}

// GetADFSUrl gets the value of ADFSUrl for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyADFSUrl() (value string, err error) {
	retValue, err := instance.GetProperty("ADFSUrl")
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

// SetADFSWebApplicationProxyRelyingPartyUri sets the value of ADFSWebApplicationProxyRelyingPartyUri for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyADFSWebApplicationProxyRelyingPartyUri(value string) (err error) {
	return instance.SetProperty("ADFSWebApplicationProxyRelyingPartyUri", (value))
}

// GetADFSWebApplicationProxyRelyingPartyUri gets the value of ADFSWebApplicationProxyRelyingPartyUri for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyADFSWebApplicationProxyRelyingPartyUri() (value string, err error) {
	retValue, err := instance.GetProperty("ADFSWebApplicationProxyRelyingPartyUri")
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

// SetConfigurationChangesPollingIntervalSec sets the value of ConfigurationChangesPollingIntervalSec for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyConfigurationChangesPollingIntervalSec(value uint32) (err error) {
	return instance.SetProperty("ConfigurationChangesPollingIntervalSec", (value))
}

// GetConfigurationChangesPollingIntervalSec gets the value of ConfigurationChangesPollingIntervalSec for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyConfigurationChangesPollingIntervalSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfigurationChangesPollingIntervalSec")
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

// SetConfigurationVersion sets the value of ConfigurationVersion for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyConfigurationVersion(value string) (err error) {
	return instance.SetProperty("ConfigurationVersion", (value))
}

// GetConfigurationVersion gets the value of ConfigurationVersion for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyConfigurationVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationVersion")
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

// SetConnectedServersName sets the value of ConnectedServersName for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyConnectedServersName(value []string) (err error) {
	return instance.SetProperty("ConnectedServersName", (value))
}

// GetConnectedServersName gets the value of ConnectedServersName for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyConnectedServersName() (value []string, err error) {
	retValue, err := instance.GetProperty("ConnectedServersName")
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

// SetOAuthAuthenticationURL sets the value of OAuthAuthenticationURL for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyOAuthAuthenticationURL(value string) (err error) {
	return instance.SetProperty("OAuthAuthenticationURL", (value))
}

// GetOAuthAuthenticationURL gets the value of OAuthAuthenticationURL for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyOAuthAuthenticationURL() (value string, err error) {
	retValue, err := instance.GetProperty("OAuthAuthenticationURL")
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

// SetUserIdleTimeoutAction sets the value of UserIdleTimeoutAction for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyUserIdleTimeoutAction(value string) (err error) {
	return instance.SetProperty("UserIdleTimeoutAction", (value))
}

// GetUserIdleTimeoutAction gets the value of UserIdleTimeoutAction for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyUserIdleTimeoutAction() (value string, err error) {
	retValue, err := instance.GetProperty("UserIdleTimeoutAction")
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

// SetUserIdleTimeoutSec sets the value of UserIdleTimeoutSec for the instance
func (instance *AppProxyGlobalConfiguration) SetPropertyUserIdleTimeoutSec(value uint32) (err error) {
	return instance.SetProperty("UserIdleTimeoutSec", (value))
}

// GetUserIdleTimeoutSec gets the value of UserIdleTimeoutSec for the instance
func (instance *AppProxyGlobalConfiguration) GetPropertyUserIdleTimeoutSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserIdleTimeoutSec")
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
