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

// FtpAuthentication struct
type FtpAuthentication struct {
	*EmbeddedObject

	//
	AnonymousAuthentication FtpAnonymousAuthentication

	//
	BasicAuthentication FtpBasicAuthentication

	//
	ClientCertAuthentication FtpClientCertAuthentication

	//
	CustomAuthentication FtpCustomAuthentication
}

func NewFtpAuthenticationEx1(instance *cim.WmiInstance) (newInstance *FtpAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

func NewFtpAuthenticationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpAuthentication, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpAuthentication{
		EmbeddedObject: tmp,
	}
	return
}

// SetAnonymousAuthentication sets the value of AnonymousAuthentication for the instance
func (instance *FtpAuthentication) SetPropertyAnonymousAuthentication(value FtpAnonymousAuthentication) (err error) {
	return instance.SetProperty("AnonymousAuthentication", (value))
}

// GetAnonymousAuthentication gets the value of AnonymousAuthentication for the instance
func (instance *FtpAuthentication) GetPropertyAnonymousAuthentication() (value FtpAnonymousAuthentication, err error) {
	retValue, err := instance.GetProperty("AnonymousAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpAnonymousAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpAnonymousAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpAnonymousAuthentication(valuetmp)

	return
}

// SetBasicAuthentication sets the value of BasicAuthentication for the instance
func (instance *FtpAuthentication) SetPropertyBasicAuthentication(value FtpBasicAuthentication) (err error) {
	return instance.SetProperty("BasicAuthentication", (value))
}

// GetBasicAuthentication gets the value of BasicAuthentication for the instance
func (instance *FtpAuthentication) GetPropertyBasicAuthentication() (value FtpBasicAuthentication, err error) {
	retValue, err := instance.GetProperty("BasicAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpBasicAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpBasicAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpBasicAuthentication(valuetmp)

	return
}

// SetClientCertAuthentication sets the value of ClientCertAuthentication for the instance
func (instance *FtpAuthentication) SetPropertyClientCertAuthentication(value FtpClientCertAuthentication) (err error) {
	return instance.SetProperty("ClientCertAuthentication", (value))
}

// GetClientCertAuthentication gets the value of ClientCertAuthentication for the instance
func (instance *FtpAuthentication) GetPropertyClientCertAuthentication() (value FtpClientCertAuthentication, err error) {
	retValue, err := instance.GetProperty("ClientCertAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpClientCertAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpClientCertAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpClientCertAuthentication(valuetmp)

	return
}

// SetCustomAuthentication sets the value of CustomAuthentication for the instance
func (instance *FtpAuthentication) SetPropertyCustomAuthentication(value FtpCustomAuthentication) (err error) {
	return instance.SetProperty("CustomAuthentication", (value))
}

// GetCustomAuthentication gets the value of CustomAuthentication for the instance
func (instance *FtpAuthentication) GetPropertyCustomAuthentication() (value FtpCustomAuthentication, err error) {
	retValue, err := instance.GetProperty("CustomAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCustomAuthentication)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCustomAuthentication is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCustomAuthentication(valuetmp)

	return
}
