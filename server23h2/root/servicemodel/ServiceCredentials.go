// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ServiceCredentials struct
type ServiceCredentials struct {
	*Behavior

	// The client certificate authentication and provisioning settings for this service.
	ClientCertificate string

	// The current issued token authentication settings for this service.
	IssuedTokenAuthentication string

	// The current credential authentication and provisioning settings to be used by peer transport endpoints.
	Peer string

	// Specifies the current secure conversation settings.
	SecureConversationAuthentication string

	// The certificate associated with this service.
	ServiceCertificate string

	// The username/password settings for this service.
	UserNameAuthentication string

	// The Windows authentication settings for this service.
	WindowsAuthentication string
}

func NewServiceCredentialsEx1(instance *cim.WmiInstance) (newInstance *ServiceCredentials, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ServiceCredentials{
		Behavior: tmp,
	}
	return
}

func NewServiceCredentialsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ServiceCredentials, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ServiceCredentials{
		Behavior: tmp,
	}
	return
}

// SetClientCertificate sets the value of ClientCertificate for the instance
func (instance *ServiceCredentials) SetPropertyClientCertificate(value string) (err error) {
	return instance.SetProperty("ClientCertificate", (value))
}

// GetClientCertificate gets the value of ClientCertificate for the instance
func (instance *ServiceCredentials) GetPropertyClientCertificate() (value string, err error) {
	retValue, err := instance.GetProperty("ClientCertificate")
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

// SetIssuedTokenAuthentication sets the value of IssuedTokenAuthentication for the instance
func (instance *ServiceCredentials) SetPropertyIssuedTokenAuthentication(value string) (err error) {
	return instance.SetProperty("IssuedTokenAuthentication", (value))
}

// GetIssuedTokenAuthentication gets the value of IssuedTokenAuthentication for the instance
func (instance *ServiceCredentials) GetPropertyIssuedTokenAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("IssuedTokenAuthentication")
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

// SetPeer sets the value of Peer for the instance
func (instance *ServiceCredentials) SetPropertyPeer(value string) (err error) {
	return instance.SetProperty("Peer", (value))
}

// GetPeer gets the value of Peer for the instance
func (instance *ServiceCredentials) GetPropertyPeer() (value string, err error) {
	retValue, err := instance.GetProperty("Peer")
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

// SetSecureConversationAuthentication sets the value of SecureConversationAuthentication for the instance
func (instance *ServiceCredentials) SetPropertySecureConversationAuthentication(value string) (err error) {
	return instance.SetProperty("SecureConversationAuthentication", (value))
}

// GetSecureConversationAuthentication gets the value of SecureConversationAuthentication for the instance
func (instance *ServiceCredentials) GetPropertySecureConversationAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("SecureConversationAuthentication")
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

// SetServiceCertificate sets the value of ServiceCertificate for the instance
func (instance *ServiceCredentials) SetPropertyServiceCertificate(value string) (err error) {
	return instance.SetProperty("ServiceCertificate", (value))
}

// GetServiceCertificate gets the value of ServiceCertificate for the instance
func (instance *ServiceCredentials) GetPropertyServiceCertificate() (value string, err error) {
	retValue, err := instance.GetProperty("ServiceCertificate")
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

// SetUserNameAuthentication sets the value of UserNameAuthentication for the instance
func (instance *ServiceCredentials) SetPropertyUserNameAuthentication(value string) (err error) {
	return instance.SetProperty("UserNameAuthentication", (value))
}

// GetUserNameAuthentication gets the value of UserNameAuthentication for the instance
func (instance *ServiceCredentials) GetPropertyUserNameAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("UserNameAuthentication")
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

// SetWindowsAuthentication sets the value of WindowsAuthentication for the instance
func (instance *ServiceCredentials) SetPropertyWindowsAuthentication(value string) (err error) {
	return instance.SetProperty("WindowsAuthentication", (value))
}

// GetWindowsAuthentication gets the value of WindowsAuthentication for the instance
func (instance *ServiceCredentials) GetPropertyWindowsAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("WindowsAuthentication")
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
