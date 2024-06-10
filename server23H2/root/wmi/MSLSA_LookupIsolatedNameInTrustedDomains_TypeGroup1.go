// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1 struct
type MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1 struct {
	*MSLSA_LookupIsolatedNameInTrustedDomains

	// Client Network Address
	ClientNetworkAddress string

	// Isolated Name
	IsolatedName string
}

func NewMSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1Ex1(instance *cim.WmiInstance) (newInstance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1, err error) {
	tmp, err := NewMSLSA_LookupIsolatedNameInTrustedDomainsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1{
		MSLSA_LookupIsolatedNameInTrustedDomains: tmp,
	}
	return
}

func NewMSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1, err error) {
	tmp, err := NewMSLSA_LookupIsolatedNameInTrustedDomainsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1{
		MSLSA_LookupIsolatedNameInTrustedDomains: tmp,
	}
	return
}

// SetClientNetworkAddress sets the value of ClientNetworkAddress for the instance
func (instance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1) SetPropertyClientNetworkAddress(value string) (err error) {
	return instance.SetProperty("ClientNetworkAddress", (value))
}

// GetClientNetworkAddress gets the value of ClientNetworkAddress for the instance
func (instance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1) GetPropertyClientNetworkAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ClientNetworkAddress")
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

// SetIsolatedName sets the value of IsolatedName for the instance
func (instance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1) SetPropertyIsolatedName(value string) (err error) {
	return instance.SetProperty("IsolatedName", (value))
}

// GetIsolatedName gets the value of IsolatedName for the instance
func (instance *MSLSA_LookupIsolatedNameInTrustedDomains_TypeGroup1) GetPropertyIsolatedName() (value string, err error) {
	retValue, err := instance.GetProperty("IsolatedName")
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
