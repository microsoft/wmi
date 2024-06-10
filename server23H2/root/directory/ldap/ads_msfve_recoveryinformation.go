// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msfve_recoveryinformation struct
type ads_msfve_recoveryinformation struct {
	*ds_top

	//
	DS_msFVE_KeyPackage Uint8Array

	//
	DS_msFVE_RecoveryGuid Uint8Array

	//
	DS_msFVE_RecoveryPassword string

	//
	DS_msFVE_VolumeGuid Uint8Array
}

func Newads_msfve_recoveryinformationEx1(instance *cim.WmiInstance) (newInstance *ads_msfve_recoveryinformation, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msfve_recoveryinformation{
		ds_top: tmp,
	}
	return
}

func Newads_msfve_recoveryinformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msfve_recoveryinformation, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msfve_recoveryinformation{
		ds_top: tmp,
	}
	return
}

// SetDS_msFVE_KeyPackage sets the value of DS_msFVE_KeyPackage for the instance
func (instance *ads_msfve_recoveryinformation) SetPropertyDS_msFVE_KeyPackage(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msFVE_KeyPackage", (value))
}

// GetDS_msFVE_KeyPackage gets the value of DS_msFVE_KeyPackage for the instance
func (instance *ads_msfve_recoveryinformation) GetPropertyDS_msFVE_KeyPackage() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msFVE_KeyPackage")
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

// SetDS_msFVE_RecoveryGuid sets the value of DS_msFVE_RecoveryGuid for the instance
func (instance *ads_msfve_recoveryinformation) SetPropertyDS_msFVE_RecoveryGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msFVE_RecoveryGuid", (value))
}

// GetDS_msFVE_RecoveryGuid gets the value of DS_msFVE_RecoveryGuid for the instance
func (instance *ads_msfve_recoveryinformation) GetPropertyDS_msFVE_RecoveryGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msFVE_RecoveryGuid")
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

// SetDS_msFVE_RecoveryPassword sets the value of DS_msFVE_RecoveryPassword for the instance
func (instance *ads_msfve_recoveryinformation) SetPropertyDS_msFVE_RecoveryPassword(value string) (err error) {
	return instance.SetProperty("DS_msFVE_RecoveryPassword", (value))
}

// GetDS_msFVE_RecoveryPassword gets the value of DS_msFVE_RecoveryPassword for the instance
func (instance *ads_msfve_recoveryinformation) GetPropertyDS_msFVE_RecoveryPassword() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msFVE_RecoveryPassword")
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

// SetDS_msFVE_VolumeGuid sets the value of DS_msFVE_VolumeGuid for the instance
func (instance *ads_msfve_recoveryinformation) SetPropertyDS_msFVE_VolumeGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msFVE_VolumeGuid", (value))
}

// GetDS_msFVE_VolumeGuid gets the value of DS_msFVE_VolumeGuid for the instance
func (instance *ads_msfve_recoveryinformation) GetPropertyDS_msFVE_VolumeGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msFVE_VolumeGuid")
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
