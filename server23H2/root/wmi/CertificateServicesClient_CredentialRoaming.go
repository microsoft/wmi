// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CertificateServicesClient_CredentialRoaming struct
type CertificateServicesClient_CredentialRoaming struct {
	*EventTrace

	// Enable Flags
	Flags CredentialRoaming_Flags
}

func NewCertificateServicesClient_CredentialRoamingEx1(instance *cim.WmiInstance) (newInstance *CertificateServicesClient_CredentialRoaming, err error) {
	tmp, err := NewEventTraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CertificateServicesClient_CredentialRoaming{
		EventTrace: tmp,
	}
	return
}

func NewCertificateServicesClient_CredentialRoamingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CertificateServicesClient_CredentialRoaming, err error) {
	tmp, err := NewEventTraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CertificateServicesClient_CredentialRoaming{
		EventTrace: tmp,
	}
	return
}

// SetFlags sets the value of Flags for the instance
func (instance *CertificateServicesClient_CredentialRoaming) SetPropertyFlags(value CredentialRoaming_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *CertificateServicesClient_CredentialRoaming) GetPropertyFlags() (value CredentialRoaming_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
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

	value = CredentialRoaming_Flags(valuetmp)

	return
}
