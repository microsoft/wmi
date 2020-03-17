// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_SideLoader struct
type MDM_SideLoader struct {
	cim.WmiInstance

	//
	key uint32

	//
	ProductKeyHash string
}

// Setkey sets the value of key for the instance
func (instance *MDM_SideLoader) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MDM_SideLoader) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProductKeyHash sets the value of ProductKeyHash for the instance
func (instance *MDM_SideLoader) SetPropertyProductKeyHash(value string) (err error) {
	return instance.SetProperty("ProductKeyHash", value)
}

// GetProductKeyHash gets the value of ProductKeyHash for the instance
func (instance *MDM_SideLoader) GetPropertyProductKeyHash() (value string, err error) {
	retValue, err := instance.GetProperty("ProductKeyHash")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
