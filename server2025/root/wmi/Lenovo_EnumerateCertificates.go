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

// Lenovo_EnumerateCertificates struct
type Lenovo_EnumerateCertificates struct {
	*Lenovo_CertificateStorageElement

	//
	Active bool

	// Enumerate certificates
	CertificateDataSets []Lenovo_CertificateDataSets

	//
	InstanceName string
}

func NewLenovo_EnumerateCertificatesEx1(instance *cim.WmiInstance) (newInstance *Lenovo_EnumerateCertificates, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_EnumerateCertificates{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

func NewLenovo_EnumerateCertificatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_EnumerateCertificates, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_EnumerateCertificates{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_EnumerateCertificates) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_EnumerateCertificates) GetPropertyActive() (value bool, err error) {
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

// SetCertificateDataSets sets the value of CertificateDataSets for the instance
func (instance *Lenovo_EnumerateCertificates) SetPropertyCertificateDataSets(value []Lenovo_CertificateDataSets) (err error) {
	return instance.SetProperty("CertificateDataSets", (value))
}

// GetCertificateDataSets gets the value of CertificateDataSets for the instance
func (instance *Lenovo_EnumerateCertificates) GetPropertyCertificateDataSets() (value []Lenovo_CertificateDataSets, err error) {
	retValue, err := instance.GetProperty("CertificateDataSets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Lenovo_CertificateDataSets)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Lenovo_CertificateDataSets is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Lenovo_CertificateDataSets(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *Lenovo_EnumerateCertificates) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_EnumerateCertificates) GetPropertyInstanceName() (value string, err error) {
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
