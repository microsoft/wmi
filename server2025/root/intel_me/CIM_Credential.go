// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.Intel_ME
//////////////////////////////////////////////
package intel_me

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Credential struct
type CIM_Credential struct {
	*CIM_ManagedElement

	// The date and time when the credential expires (and is not appropriate for use for authentication/ authorization). If this information is not applicable, Use a value of '99991231235959.999999+999', (1 microsecond before midnight, December 31, 9999 CE, at coordinated universal time + 999 minutes).
	///On CreateInstance, if this property is unspecified, or set to NULL, then a value of '99991231235959.999999+999' is assumed. Note that this property does not define how the expiration is set - but that there IS an expiration. The property may be set to either a specific date/time or an interval (calculated from the Issued datetime). For example, for Certificate Authority-signed public key, the expiration is determined by the CA. Another example is a voice mail password that expires 60 days after it is set/issued.
	Expires string

	// The date and time when the credential was issued. Use a value of all '00000101000000.000000+000', (midnight, January 1, 1 BCE at coordinated universal time +0 minutes), if this information is not applicable. On CreateInstance, if this property is unspecified, or set to NULL, then current time is assumed.
	Issued string
}

func NewCIM_CredentialEx1(instance *cim.WmiInstance) (newInstance *CIM_Credential, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Credential{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_CredentialEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Credential, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Credential{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetExpires sets the value of Expires for the instance
func (instance *CIM_Credential) SetPropertyExpires(value string) (err error) {
	return instance.SetProperty("Expires", (value))
}

// GetExpires gets the value of Expires for the instance
func (instance *CIM_Credential) GetPropertyExpires() (value string, err error) {
	retValue, err := instance.GetProperty("Expires")
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

// SetIssued sets the value of Issued for the instance
func (instance *CIM_Credential) SetPropertyIssued(value string) (err error) {
	return instance.SetProperty("Issued", (value))
}

// GetIssued gets the value of Issued for the instance
func (instance *CIM_Credential) GetPropertyIssued() (value string, err error) {
	retValue, err := instance.GetProperty("Issued")
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
