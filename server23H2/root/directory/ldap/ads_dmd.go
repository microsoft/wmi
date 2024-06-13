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

// ads_dmd struct
type ads_dmd struct {
	*ds_top

	//
	DS_dmdName string

	//
	DS_msDS_IntId int32

	//
	DS_msDs_Schema_Extensions []Uint8Array

	//
	DS_msDS_USNLastSyncSuccess int64

	//
	DS_prefixMap Uint8Array

	//
	DS_schemaInfo []Uint8Array

	//
	DS_schemaUpdate string
}

func Newads_dmdEx1(instance *cim.WmiInstance) (newInstance *ads_dmd, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dmd{
		ds_top: tmp,
	}
	return
}

func Newads_dmdEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dmd, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dmd{
		ds_top: tmp,
	}
	return
}

// SetDS_dmdName sets the value of DS_dmdName for the instance
func (instance *ads_dmd) SetPropertyDS_dmdName(value string) (err error) {
	return instance.SetProperty("DS_dmdName", (value))
}

// GetDS_dmdName gets the value of DS_dmdName for the instance
func (instance *ads_dmd) GetPropertyDS_dmdName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dmdName")
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

// SetDS_msDS_IntId sets the value of DS_msDS_IntId for the instance
func (instance *ads_dmd) SetPropertyDS_msDS_IntId(value int32) (err error) {
	return instance.SetProperty("DS_msDS_IntId", (value))
}

// GetDS_msDS_IntId gets the value of DS_msDS_IntId for the instance
func (instance *ads_dmd) GetPropertyDS_msDS_IntId() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IntId")
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

// SetDS_msDs_Schema_Extensions sets the value of DS_msDs_Schema_Extensions for the instance
func (instance *ads_dmd) SetPropertyDS_msDs_Schema_Extensions(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDs_Schema_Extensions", (value))
}

// GetDS_msDs_Schema_Extensions gets the value of DS_msDs_Schema_Extensions for the instance
func (instance *ads_dmd) GetPropertyDS_msDs_Schema_Extensions() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDs_Schema_Extensions")
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

// SetDS_msDS_USNLastSyncSuccess sets the value of DS_msDS_USNLastSyncSuccess for the instance
func (instance *ads_dmd) SetPropertyDS_msDS_USNLastSyncSuccess(value int64) (err error) {
	return instance.SetProperty("DS_msDS_USNLastSyncSuccess", (value))
}

// GetDS_msDS_USNLastSyncSuccess gets the value of DS_msDS_USNLastSyncSuccess for the instance
func (instance *ads_dmd) GetPropertyDS_msDS_USNLastSyncSuccess() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_USNLastSyncSuccess")
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

// SetDS_prefixMap sets the value of DS_prefixMap for the instance
func (instance *ads_dmd) SetPropertyDS_prefixMap(value Uint8Array) (err error) {
	return instance.SetProperty("DS_prefixMap", (value))
}

// GetDS_prefixMap gets the value of DS_prefixMap for the instance
func (instance *ads_dmd) GetPropertyDS_prefixMap() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_prefixMap")
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

// SetDS_schemaInfo sets the value of DS_schemaInfo for the instance
func (instance *ads_dmd) SetPropertyDS_schemaInfo(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_schemaInfo", (value))
}

// GetDS_schemaInfo gets the value of DS_schemaInfo for the instance
func (instance *ads_dmd) GetPropertyDS_schemaInfo() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schemaInfo")
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

// SetDS_schemaUpdate sets the value of DS_schemaUpdate for the instance
func (instance *ads_dmd) SetPropertyDS_schemaUpdate(value string) (err error) {
	return instance.SetProperty("DS_schemaUpdate", (value))
}

// GetDS_schemaUpdate gets the value of DS_schemaUpdate for the instance
func (instance *ads_dmd) GetPropertyDS_schemaUpdate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_schemaUpdate")
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
