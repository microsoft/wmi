// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MDM_Certificate struct
type MDM_Certificate struct {
	*cim.WmiInstance

	//
	Blob string

	//
	IsInstalled bool

	//
	StoreLocation uint8

	//
	StoreName string

	//
	Thumbprint string
}

func NewMDM_CertificateEx1(instance *cim.WmiInstance) (newInstance *MDM_Certificate, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Certificate{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_CertificateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Certificate, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Certificate{
		WmiInstance: tmp,
	}
	return
}

// SetBlob sets the value of Blob for the instance
func (instance *MDM_Certificate) SetPropertyBlob(value string) (err error) {
	return instance.SetProperty("Blob", (value))
}

// GetBlob gets the value of Blob for the instance
func (instance *MDM_Certificate) GetPropertyBlob() (value string, err error) {
	retValue, err := instance.GetProperty("Blob")
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

// SetIsInstalled sets the value of IsInstalled for the instance
func (instance *MDM_Certificate) SetPropertyIsInstalled(value bool) (err error) {
	return instance.SetProperty("IsInstalled", (value))
}

// GetIsInstalled gets the value of IsInstalled for the instance
func (instance *MDM_Certificate) GetPropertyIsInstalled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsInstalled")
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

// SetStoreLocation sets the value of StoreLocation for the instance
func (instance *MDM_Certificate) SetPropertyStoreLocation(value uint8) (err error) {
	return instance.SetProperty("StoreLocation", (value))
}

// GetStoreLocation gets the value of StoreLocation for the instance
func (instance *MDM_Certificate) GetPropertyStoreLocation() (value uint8, err error) {
	retValue, err := instance.GetProperty("StoreLocation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetStoreName sets the value of StoreName for the instance
func (instance *MDM_Certificate) SetPropertyStoreName(value string) (err error) {
	return instance.SetProperty("StoreName", (value))
}

// GetStoreName gets the value of StoreName for the instance
func (instance *MDM_Certificate) GetPropertyStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("StoreName")
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

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MDM_Certificate) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", (value))
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MDM_Certificate) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
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
