// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Lenovo_GetMaxCertificateLength struct
type Lenovo_GetMaxCertificateLength struct {
	*Lenovo_CertificateStorageElement

	//
	Active bool

	//
	InstanceName string

	// Get maximum certificate length
	MaxCertificateLength uint32
}

func NewLenovo_GetMaxCertificateLengthEx1(instance *cim.WmiInstance) (newInstance *Lenovo_GetMaxCertificateLength, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_GetMaxCertificateLength{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

func NewLenovo_GetMaxCertificateLengthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_GetMaxCertificateLength, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_GetMaxCertificateLength{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_GetMaxCertificateLength) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_GetMaxCertificateLength) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Lenovo_GetMaxCertificateLength) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_GetMaxCertificateLength) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetMaxCertificateLength sets the value of MaxCertificateLength for the instance
func (instance *Lenovo_GetMaxCertificateLength) SetPropertyMaxCertificateLength(value uint32) (err error) {
	return instance.SetProperty("MaxCertificateLength", (value))
}

// GetMaxCertificateLength gets the value of MaxCertificateLength for the instance
func (instance *Lenovo_GetMaxCertificateLength) GetPropertyMaxCertificateLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCertificateLength")
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
