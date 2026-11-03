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

// AMT_ProvisioningCertificateHash struct
type AMT_ProvisioningCertificateHash struct {
	*CIM_Credential

	// Indicates whether this certificate hash is in use
	Enabled bool

	// Hash value (length depending on hash type)
	HashData []uint8

	// The hash type.
	HashType ProvisioningCertificateHash_HashType

	// Indicates whether this entry is a factory default.
	IsDefault bool
}

func NewAMT_ProvisioningCertificateHashEx1(instance *cim.WmiInstance) (newInstance *AMT_ProvisioningCertificateHash, err error) {
	tmp, err := NewCIM_CredentialEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AMT_ProvisioningCertificateHash{
		CIM_Credential: tmp,
	}
	return
}

func NewAMT_ProvisioningCertificateHashEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AMT_ProvisioningCertificateHash, err error) {
	tmp, err := NewCIM_CredentialEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AMT_ProvisioningCertificateHash{
		CIM_Credential: tmp,
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *AMT_ProvisioningCertificateHash) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *AMT_ProvisioningCertificateHash) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetHashData sets the value of HashData for the instance
func (instance *AMT_ProvisioningCertificateHash) SetPropertyHashData(value []uint8) (err error) {
	return instance.SetProperty("HashData", (value))
}

// GetHashData gets the value of HashData for the instance
func (instance *AMT_ProvisioningCertificateHash) GetPropertyHashData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("HashData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetHashType sets the value of HashType for the instance
func (instance *AMT_ProvisioningCertificateHash) SetPropertyHashType(value ProvisioningCertificateHash_HashType) (err error) {
	return instance.SetProperty("HashType", (value))
}

// GetHashType gets the value of HashType for the instance
func (instance *AMT_ProvisioningCertificateHash) GetPropertyHashType() (value ProvisioningCertificateHash_HashType, err error) {
	retValue, err := instance.GetProperty("HashType")
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

	value = ProvisioningCertificateHash_HashType(valuetmp)

	return
}

// SetIsDefault sets the value of IsDefault for the instance
func (instance *AMT_ProvisioningCertificateHash) SetPropertyIsDefault(value bool) (err error) {
	return instance.SetProperty("IsDefault", (value))
}

// GetIsDefault gets the value of IsDefault for the instance
func (instance *AMT_ProvisioningCertificateHash) GetPropertyIsDefault() (value bool, err error) {
	retValue, err := instance.GetProperty("IsDefault")
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
