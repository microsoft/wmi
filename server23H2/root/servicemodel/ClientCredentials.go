// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClientCredentials struct
type ClientCredentials struct {
	*Behavior

	// The X.509 certificate the client uses to authenticate to the service.
	ClientCertificate string

	// The current Http Digest credential.
	HttpDigest string

	// The endpoint address and binding used to contact the locat security token service.
	IssuedToken string

	// The credentials peer node uses to authenticate itself to other nodes in the mesh.
	Peer string

	// The service's x.509 certificate.
	ServiceCertificate string

	// A Boolean value that specifies if the credential supports interactive negotiation.
	SupportInteractive bool

	// The username and password the client will use to authenticate itself to the service.
	UserName string

	// The windows credentials the client will use to authenticate itself to the service.
	Windows string
}

func NewClientCredentialsEx1(instance *cim.WmiInstance) (newInstance *ClientCredentials, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ClientCredentials{
		Behavior: tmp,
	}
	return
}

func NewClientCredentialsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClientCredentials, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClientCredentials{
		Behavior: tmp,
	}
	return
}

// SetClientCertificate sets the value of ClientCertificate for the instance
func (instance *ClientCredentials) SetPropertyClientCertificate(value string) (err error) {
	return instance.SetProperty("ClientCertificate", (value))
}

// GetClientCertificate gets the value of ClientCertificate for the instance
func (instance *ClientCredentials) GetPropertyClientCertificate() (value string, err error) {
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

// SetHttpDigest sets the value of HttpDigest for the instance
func (instance *ClientCredentials) SetPropertyHttpDigest(value string) (err error) {
	return instance.SetProperty("HttpDigest", (value))
}

// GetHttpDigest gets the value of HttpDigest for the instance
func (instance *ClientCredentials) GetPropertyHttpDigest() (value string, err error) {
	retValue, err := instance.GetProperty("HttpDigest")
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

// SetIssuedToken sets the value of IssuedToken for the instance
func (instance *ClientCredentials) SetPropertyIssuedToken(value string) (err error) {
	return instance.SetProperty("IssuedToken", (value))
}

// GetIssuedToken gets the value of IssuedToken for the instance
func (instance *ClientCredentials) GetPropertyIssuedToken() (value string, err error) {
	retValue, err := instance.GetProperty("IssuedToken")
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
func (instance *ClientCredentials) SetPropertyPeer(value string) (err error) {
	return instance.SetProperty("Peer", (value))
}

// GetPeer gets the value of Peer for the instance
func (instance *ClientCredentials) GetPropertyPeer() (value string, err error) {
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

// SetServiceCertificate sets the value of ServiceCertificate for the instance
func (instance *ClientCredentials) SetPropertyServiceCertificate(value string) (err error) {
	return instance.SetProperty("ServiceCertificate", (value))
}

// GetServiceCertificate gets the value of ServiceCertificate for the instance
func (instance *ClientCredentials) GetPropertyServiceCertificate() (value string, err error) {
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

// SetSupportInteractive sets the value of SupportInteractive for the instance
func (instance *ClientCredentials) SetPropertySupportInteractive(value bool) (err error) {
	return instance.SetProperty("SupportInteractive", (value))
}

// GetSupportInteractive gets the value of SupportInteractive for the instance
func (instance *ClientCredentials) GetPropertySupportInteractive() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportInteractive")
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

// SetUserName sets the value of UserName for the instance
func (instance *ClientCredentials) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *ClientCredentials) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

// SetWindows sets the value of Windows for the instance
func (instance *ClientCredentials) SetPropertyWindows(value string) (err error) {
	return instance.SetProperty("Windows", (value))
}

// GetWindows gets the value of Windows for the instance
func (instance *ClientCredentials) GetPropertyWindows() (value string, err error) {
	retValue, err := instance.GetProperty("Windows")
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
