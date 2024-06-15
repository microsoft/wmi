// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KerbInitSecurityContext_End struct
type KerbInitSecurityContext_End struct {
	*KerbInitSecurityContext

	// Credentials Source
	CredSource string

	// Domain Name
	DomainName string

	// Extended Error Code
	ExtError uint32

	// Extended Error klininfo
	klininfo uint32

	// Status
	Status uint32

	// Target
	Target string

	// User Name
	UserName string
}

func NewKerbInitSecurityContext_EndEx1(instance *cim.WmiInstance) (newInstance *KerbInitSecurityContext_End, err error) {
	tmp, err := NewKerbInitSecurityContextEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext_End{
		KerbInitSecurityContext: tmp,
	}
	return
}

func NewKerbInitSecurityContext_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbInitSecurityContext_End, err error) {
	tmp, err := NewKerbInitSecurityContextEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbInitSecurityContext_End{
		KerbInitSecurityContext: tmp,
	}
	return
}

// SetCredSource sets the value of CredSource for the instance
func (instance *KerbInitSecurityContext_End) SetPropertyCredSource(value string) (err error) {
	return instance.SetProperty("CredSource", (value))
}

// GetCredSource gets the value of CredSource for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyCredSource() (value string, err error) {
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
func (instance *KerbInitSecurityContext_End) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyDomainName() (value string, err error) {
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

// SetExtError sets the value of ExtError for the instance
func (instance *KerbInitSecurityContext_End) SetPropertyExtError(value uint32) (err error) {
	return instance.SetProperty("ExtError", (value))
}

// GetExtError gets the value of ExtError for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyExtError() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtError")
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

// Setklininfo sets the value of klininfo for the instance
func (instance *KerbInitSecurityContext_End) SetPropertyklininfo(value uint32) (err error) {
	return instance.SetProperty("klininfo", (value))
}

// Getklininfo gets the value of klininfo for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyklininfo() (value uint32, err error) {
	retValue, err := instance.GetProperty("klininfo")
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

// SetStatus sets the value of Status for the instance
func (instance *KerbInitSecurityContext_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyStatus() (value uint32, err error) {
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
func (instance *KerbInitSecurityContext_End) SetPropertyTarget(value string) (err error) {
	return instance.SetProperty("Target", (value))
}

// GetTarget gets the value of Target for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyTarget() (value string, err error) {
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
func (instance *KerbInitSecurityContext_End) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *KerbInitSecurityContext_End) GetPropertyUserName() (value string, err error) {
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
