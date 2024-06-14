// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DACertificateContext struct
type DACertificateContext struct {
	*cim.WmiInstance

	//
	EncodedCertificate []uint8

	//
	PrivateKeyExists bool
}

func NewDACertificateContextEx1(instance *cim.WmiInstance) (newInstance *DACertificateContext, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DACertificateContext{
		WmiInstance: tmp,
	}
	return
}

func NewDACertificateContextEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DACertificateContext, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DACertificateContext{
		WmiInstance: tmp,
	}
	return
}

// SetEncodedCertificate sets the value of EncodedCertificate for the instance
func (instance *DACertificateContext) SetPropertyEncodedCertificate(value []uint8) (err error) {
	return instance.SetProperty("EncodedCertificate", (value))
}

// GetEncodedCertificate gets the value of EncodedCertificate for the instance
func (instance *DACertificateContext) GetPropertyEncodedCertificate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("EncodedCertificate")
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

// SetPrivateKeyExists sets the value of PrivateKeyExists for the instance
func (instance *DACertificateContext) SetPropertyPrivateKeyExists(value bool) (err error) {
	return instance.SetProperty("PrivateKeyExists", (value))
}

// GetPrivateKeyExists gets the value of PrivateKeyExists for the instance
func (instance *DACertificateContext) GetPropertyPrivateKeyExists() (value bool, err error) {
	retValue, err := instance.GetProperty("PrivateKeyExists")
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
