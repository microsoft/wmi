// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_SideLoader struct
type MDM_SideLoader struct {
	*cim.WmiInstance

	//
	key uint32

	//
	ProductKeyHash string
}

func NewMDM_SideLoaderEx1(instance *cim.WmiInstance) (newInstance *MDM_SideLoader, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_SideLoader{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_SideLoaderEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_SideLoader, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_SideLoader{
		WmiInstance: tmp,
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MDM_SideLoader) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_SideLoader) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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

// SetProductKeyHash sets the value of ProductKeyHash for the instance
func (instance *MDM_SideLoader) SetPropertyProductKeyHash(value string) (err error) {
	return instance.SetProperty("ProductKeyHash", (value))
}

// GetProductKeyHash gets the value of ProductKeyHash for the instance
func (instance *MDM_SideLoader) GetPropertyProductKeyHash() (value string, err error) {
	retValue, err := instance.GetProperty("ProductKeyHash")
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

//

// <param name="ProductKey" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_SideLoader) ActivateKey( /* IN */ ProductKey string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ActivateKey", ProductKey)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CertificateBlob" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_SideLoader) AddCertificate( /* IN */ CertificateBlob string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddCertificate", CertificateBlob)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_SideLoader) UnActivateLOB() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UnActivateLOB")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
