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

// Lenovo_SetBiosCertificate struct
type Lenovo_SetBiosCertificate struct {
	*Lenovo_CertificateStorageElement

	//
	Active bool

	//
	InstanceName string
}

func NewLenovo_SetBiosCertificateEx1(instance *cim.WmiInstance) (newInstance *Lenovo_SetBiosCertificate, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SetBiosCertificate{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

func NewLenovo_SetBiosCertificateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Lenovo_SetBiosCertificate, err error) {
	tmp, err := NewLenovo_CertificateStorageElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Lenovo_SetBiosCertificate{
		Lenovo_CertificateStorageElement: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *Lenovo_SetBiosCertificate) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *Lenovo_SetBiosCertificate) GetPropertyActive() (value bool, err error) {
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
func (instance *Lenovo_SetBiosCertificate) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *Lenovo_SetBiosCertificate) GetPropertyInstanceName() (value string, err error) {
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

// Set Bios Certificate

// <param name="Data" type="uint8 []"></param>
// <param name="DataSize" type="uint32 "></param>
// <param name="Guid" type="uint8 []"></param>
// <param name="Parameter" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="return" type="string "></param>
func (instance *Lenovo_SetBiosCertificate) SetCertificate( /* IN */ Guid []uint8,
	/* IN */ Type uint32,
	/* IN */ DataSize uint32,
	/* IN */ Data []uint8,
	/* IN */ Parameter string,
	/* OUT */ Return string) (err error) {
	_, err = instance.InvokeMethod("SetCertificate", Guid, Type, DataSize, Data, Parameter)
	if err != nil {
		return
	}
	return

}
