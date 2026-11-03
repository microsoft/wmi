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

// Lenovo_CertificateDataSets struct
type Lenovo_CertificateDataSets struct {
	*Lenovo_CertificateStorageElement

	//
	DataSize uint32

	//
	Guid []uint8

	//
	Type uint32
}

func NewLenovo_CertificateDataSetsEx1(instance *cim.WmiInstance) (newInstance *Lenovo_CertificateDataSets, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_CertificateDataSets{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

func NewLenovo_CertificateDataSetsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_CertificateDataSets, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_CertificateDataSets{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

// SetDataSize sets the value of DataSize for the instance
func (instance *Lenovo_CertificateDataSets) SetPropertyDataSize(value uint32) (err error) {
	return instance.SetProperty("DataSize", (value))
}

// GetDataSize gets the value of DataSize for the instance
func (instance *Lenovo_CertificateDataSets) GetPropertyDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataSize")
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

// SetGuid sets the value of Guid for the instance
func (instance *Lenovo_CertificateDataSets) SetPropertyGuid(value []uint8) (err error) {
	return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *Lenovo_CertificateDataSets) GetPropertyGuid() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Guid")
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

// SetType sets the value of Type for the instance
func (instance *Lenovo_CertificateDataSets) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *Lenovo_CertificateDataSets) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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
