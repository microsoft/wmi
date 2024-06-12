// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msds_keycredential struct
type ads_msds_keycredential struct {
	*ds_top

	//
	DS_msDS_ComputerSID Uint8Array

	//
	DS_msDS_CustomKeyInformation Uint8Array

	//
	DS_msDS_DeviceDN string

	//
	DS_msDS_DeviceID Uint8Array

	//
	DS_msDS_KeyApproximateLastLogonTimeStamp int64

	//
	DS_msDS_KeyId Uint8Array

	//
	DS_msDS_KeyMaterial Uint8Array

	//
	DS_msDS_KeyPrincipal string

	//
	DS_msDS_KeyUsage string
}

func Newads_msds_keycredentialEx1(instance *cim.WmiInstance) (newInstance *ads_msds_keycredential, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_keycredential{
		ds_top: tmp,
	}
	return
}

func Newads_msds_keycredentialEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_keycredential, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_keycredential{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_ComputerSID sets the value of DS_msDS_ComputerSID for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_ComputerSID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ComputerSID", (value))
}

// GetDS_msDS_ComputerSID gets the value of DS_msDS_ComputerSID for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_ComputerSID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ComputerSID")
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

// SetDS_msDS_CustomKeyInformation sets the value of DS_msDS_CustomKeyInformation for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_CustomKeyInformation(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_CustomKeyInformation", (value))
}

// GetDS_msDS_CustomKeyInformation gets the value of DS_msDS_CustomKeyInformation for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_CustomKeyInformation() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_CustomKeyInformation")
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

// SetDS_msDS_DeviceDN sets the value of DS_msDS_DeviceDN for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_DeviceDN(value string) (err error) {
	return instance.SetProperty("DS_msDS_DeviceDN", (value))
}

// GetDS_msDS_DeviceDN gets the value of DS_msDS_DeviceDN for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_DeviceDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DeviceDN")
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

// SetDS_msDS_DeviceID sets the value of DS_msDS_DeviceID for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_DeviceID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_DeviceID", (value))
}

// GetDS_msDS_DeviceID gets the value of DS_msDS_DeviceID for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_DeviceID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DeviceID")
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

// SetDS_msDS_KeyApproximateLastLogonTimeStamp sets the value of DS_msDS_KeyApproximateLastLogonTimeStamp for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_KeyApproximateLastLogonTimeStamp(value int64) (err error) {
	return instance.SetProperty("DS_msDS_KeyApproximateLastLogonTimeStamp", (value))
}

// GetDS_msDS_KeyApproximateLastLogonTimeStamp gets the value of DS_msDS_KeyApproximateLastLogonTimeStamp for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_KeyApproximateLastLogonTimeStamp() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyApproximateLastLogonTimeStamp")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_KeyId sets the value of DS_msDS_KeyId for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_KeyId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_KeyId", (value))
}

// GetDS_msDS_KeyId gets the value of DS_msDS_KeyId for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_KeyId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyId")
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

// SetDS_msDS_KeyMaterial sets the value of DS_msDS_KeyMaterial for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_KeyMaterial(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_KeyMaterial", (value))
}

// GetDS_msDS_KeyMaterial gets the value of DS_msDS_KeyMaterial for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_KeyMaterial() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyMaterial")
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

// SetDS_msDS_KeyPrincipal sets the value of DS_msDS_KeyPrincipal for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_KeyPrincipal(value string) (err error) {
	return instance.SetProperty("DS_msDS_KeyPrincipal", (value))
}

// GetDS_msDS_KeyPrincipal gets the value of DS_msDS_KeyPrincipal for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_KeyPrincipal() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyPrincipal")
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

// SetDS_msDS_KeyUsage sets the value of DS_msDS_KeyUsage for the instance
func (instance *ads_msds_keycredential) SetPropertyDS_msDS_KeyUsage(value string) (err error) {
	return instance.SetProperty("DS_msDS_KeyUsage", (value))
}

// GetDS_msDS_KeyUsage gets the value of DS_msDS_KeyUsage for the instance
func (instance *ads_msds_keycredential) GetPropertyDS_msDS_KeyUsage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyUsage")
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
