// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ds_ipsecbase struct
type ds_ipsecbase struct {
	*ds_top

	//
	DS_ipsecData Uint8Array

	//
	DS_ipsecDataType int32

	//
	DS_ipsecID string

	//
	DS_ipsecName string

	//
	DS_ipsecOwnersReference []string
}

func Newds_ipsecbaseEx1(instance *cim.WmiInstance) (newInstance *ds_ipsecbase, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecbase{
		ds_top: tmp,
	}
	return
}

func Newds_ipsecbaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_ipsecbase, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_ipsecbase{
		ds_top: tmp,
	}
	return
}

// SetDS_ipsecData sets the value of DS_ipsecData for the instance
func (instance *ds_ipsecbase) SetPropertyDS_ipsecData(value Uint8Array) (err error) {
	return instance.SetProperty("DS_ipsecData", (value))
}

// GetDS_ipsecData gets the value of DS_ipsecData for the instance
func (instance *ds_ipsecbase) GetPropertyDS_ipsecData() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_ipsecData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_ipsecDataType sets the value of DS_ipsecDataType for the instance
func (instance *ds_ipsecbase) SetPropertyDS_ipsecDataType(value int32) (err error) {
	return instance.SetProperty("DS_ipsecDataType", (value))
}

// GetDS_ipsecDataType gets the value of DS_ipsecDataType for the instance
func (instance *ds_ipsecbase) GetPropertyDS_ipsecDataType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_ipsecDataType")
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

	value = int32(valuetmp)

	return
}

// SetDS_ipsecID sets the value of DS_ipsecID for the instance
func (instance *ds_ipsecbase) SetPropertyDS_ipsecID(value string) (err error) {
	return instance.SetProperty("DS_ipsecID", (value))
}

// GetDS_ipsecID gets the value of DS_ipsecID for the instance
func (instance *ds_ipsecbase) GetPropertyDS_ipsecID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecID")
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

// SetDS_ipsecName sets the value of DS_ipsecName for the instance
func (instance *ds_ipsecbase) SetPropertyDS_ipsecName(value string) (err error) {
	return instance.SetProperty("DS_ipsecName", (value))
}

// GetDS_ipsecName gets the value of DS_ipsecName for the instance
func (instance *ds_ipsecbase) GetPropertyDS_ipsecName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecName")
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

// SetDS_ipsecOwnersReference sets the value of DS_ipsecOwnersReference for the instance
func (instance *ds_ipsecbase) SetPropertyDS_ipsecOwnersReference(value []string) (err error) {
	return instance.SetProperty("DS_ipsecOwnersReference", (value))
}

// GetDS_ipsecOwnersReference gets the value of DS_ipsecOwnersReference for the instance
func (instance *ds_ipsecbase) GetPropertyDS_ipsecOwnersReference() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecOwnersReference")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
