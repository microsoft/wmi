// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KerbAcceptSecurityContext_End struct
type KerbAcceptSecurityContext_End struct {
	*KerbAcceptSecurityContext

	// Credentials Source
	CredSource string

	// Domain Name
	DomainName string

	// Status
	Status uint32

	// Target
	Target string

	// User Name
	UserName string
}

func NewKerbAcceptSecurityContext_EndEx1(instance *cim.WmiInstance) (newInstance *KerbAcceptSecurityContext_End, err error) {
	tmp, err := NewKerbAcceptSecurityContextEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbAcceptSecurityContext_End{
		KerbAcceptSecurityContext: tmp,
	}
	return
}

func NewKerbAcceptSecurityContext_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbAcceptSecurityContext_End, err error) {
	tmp, err := NewKerbAcceptSecurityContextEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbAcceptSecurityContext_End{
		KerbAcceptSecurityContext: tmp,
	}
	return
}

// SetCredSource sets the value of CredSource for the instance
func (instance *KerbAcceptSecurityContext_End) SetPropertyCredSource(value string) (err error) {
	return instance.SetProperty("CredSource", (value))
}

// GetCredSource gets the value of CredSource for the instance
func (instance *KerbAcceptSecurityContext_End) GetPropertyCredSource() (value string, err error) {
	retValue, err := instance.GetProperty("CredSource")
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

// SetDomainName sets the value of DomainName for the instance
func (instance *KerbAcceptSecurityContext_End) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *KerbAcceptSecurityContext_End) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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

// SetStatus sets the value of Status for the instance
func (instance *KerbAcceptSecurityContext_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KerbAcceptSecurityContext_End) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetTarget sets the value of Target for the instance
func (instance *KerbAcceptSecurityContext_End) SetPropertyTarget(value string) (err error) {
	return instance.SetProperty("Target", (value))
}

// GetTarget gets the value of Target for the instance
func (instance *KerbAcceptSecurityContext_End) GetPropertyTarget() (value string, err error) {
	retValue, err := instance.GetProperty("Target")
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

// SetUserName sets the value of UserName for the instance
func (instance *KerbAcceptSecurityContext_End) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *KerbAcceptSecurityContext_End) GetPropertyUserName() (value string, err error) {
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
