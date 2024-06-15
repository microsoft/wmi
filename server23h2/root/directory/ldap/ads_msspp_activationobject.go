// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msspp_activationobject struct
type ads_msspp_activationobject struct {
	*ds_top

	//
	DS_msSPP_ConfigLicense Uint8Array

	//
	DS_msSPP_ConfirmationId string

	//
	DS_msSPP_CSVLKPartialProductKey string

	//
	DS_msSPP_CSVLKPid string

	//
	DS_msSPP_CSVLKSkuId Uint8Array

	//
	DS_msSPP_InstallationId string

	//
	DS_msSPP_IssuanceLicense Uint8Array

	//
	DS_msSPP_KMSIds []Uint8Array

	//
	DS_msSPP_OnlineLicense Uint8Array

	//
	DS_msSPP_PhoneLicense Uint8Array
}

func Newads_msspp_activationobjectEx1(instance *cim.WmiInstance) (newInstance *ads_msspp_activationobject, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msspp_activationobject{
		ds_top: tmp,
	}
	return
}

func Newads_msspp_activationobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msspp_activationobject, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msspp_activationobject{
		ds_top: tmp,
	}
	return
}

// SetDS_msSPP_ConfigLicense sets the value of DS_msSPP_ConfigLicense for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_ConfigLicense(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_ConfigLicense", (value))
}

// GetDS_msSPP_ConfigLicense gets the value of DS_msSPP_ConfigLicense for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_ConfigLicense() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_ConfigLicense")
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

// SetDS_msSPP_ConfirmationId sets the value of DS_msSPP_ConfirmationId for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_ConfirmationId(value string) (err error) {
	return instance.SetProperty("DS_msSPP_ConfirmationId", (value))
}

// GetDS_msSPP_ConfirmationId gets the value of DS_msSPP_ConfirmationId for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_ConfirmationId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_ConfirmationId")
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

// SetDS_msSPP_CSVLKPartialProductKey sets the value of DS_msSPP_CSVLKPartialProductKey for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_CSVLKPartialProductKey(value string) (err error) {
	return instance.SetProperty("DS_msSPP_CSVLKPartialProductKey", (value))
}

// GetDS_msSPP_CSVLKPartialProductKey gets the value of DS_msSPP_CSVLKPartialProductKey for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_CSVLKPartialProductKey() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_CSVLKPartialProductKey")
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

// SetDS_msSPP_CSVLKPid sets the value of DS_msSPP_CSVLKPid for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_CSVLKPid(value string) (err error) {
	return instance.SetProperty("DS_msSPP_CSVLKPid", (value))
}

// GetDS_msSPP_CSVLKPid gets the value of DS_msSPP_CSVLKPid for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_CSVLKPid() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_CSVLKPid")
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

// SetDS_msSPP_CSVLKSkuId sets the value of DS_msSPP_CSVLKSkuId for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_CSVLKSkuId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_CSVLKSkuId", (value))
}

// GetDS_msSPP_CSVLKSkuId gets the value of DS_msSPP_CSVLKSkuId for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_CSVLKSkuId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_CSVLKSkuId")
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

// SetDS_msSPP_InstallationId sets the value of DS_msSPP_InstallationId for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_InstallationId(value string) (err error) {
	return instance.SetProperty("DS_msSPP_InstallationId", (value))
}

// GetDS_msSPP_InstallationId gets the value of DS_msSPP_InstallationId for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_InstallationId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_InstallationId")
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

// SetDS_msSPP_IssuanceLicense sets the value of DS_msSPP_IssuanceLicense for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_IssuanceLicense(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_IssuanceLicense", (value))
}

// GetDS_msSPP_IssuanceLicense gets the value of DS_msSPP_IssuanceLicense for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_IssuanceLicense() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_IssuanceLicense")
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

// SetDS_msSPP_KMSIds sets the value of DS_msSPP_KMSIds for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_KMSIds(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_KMSIds", (value))
}

// GetDS_msSPP_KMSIds gets the value of DS_msSPP_KMSIds for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_KMSIds() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_KMSIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_msSPP_OnlineLicense sets the value of DS_msSPP_OnlineLicense for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_OnlineLicense(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_OnlineLicense", (value))
}

// GetDS_msSPP_OnlineLicense gets the value of DS_msSPP_OnlineLicense for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_OnlineLicense() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_OnlineLicense")
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

// SetDS_msSPP_PhoneLicense sets the value of DS_msSPP_PhoneLicense for the instance
func (instance *ads_msspp_activationobject) SetPropertyDS_msSPP_PhoneLicense(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msSPP_PhoneLicense", (value))
}

// GetDS_msSPP_PhoneLicense gets the value of DS_msSPP_PhoneLicense for the instance
func (instance *ads_msspp_activationobject) GetPropertyDS_msSPP_PhoneLicense() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msSPP_PhoneLicense")
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
