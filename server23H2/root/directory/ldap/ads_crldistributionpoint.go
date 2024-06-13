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

// ads_crldistributionpoint struct
type ads_crldistributionpoint struct {
	*ds_top

	//
	DS_authorityRevocationList []Uint8Array

	//
	DS_certificateAuthorityObject string

	//
	DS_certificateRevocationList Uint8Array

	//
	DS_cRLPartitionedRevocationList Uint8Array

	//
	DS_deltaRevocationList []Uint8Array
}

func Newads_crldistributionpointEx1(instance *cim.WmiInstance) (newInstance *ads_crldistributionpoint, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_crldistributionpoint{
		ds_top: tmp,
	}
	return
}

func Newads_crldistributionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_crldistributionpoint, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_crldistributionpoint{
		ds_top: tmp,
	}
	return
}

// SetDS_authorityRevocationList sets the value of DS_authorityRevocationList for the instance
func (instance *ads_crldistributionpoint) SetPropertyDS_authorityRevocationList(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_authorityRevocationList", (value))
}

// GetDS_authorityRevocationList gets the value of DS_authorityRevocationList for the instance
func (instance *ads_crldistributionpoint) GetPropertyDS_authorityRevocationList() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_authorityRevocationList")
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

// SetDS_certificateAuthorityObject sets the value of DS_certificateAuthorityObject for the instance
func (instance *ads_crldistributionpoint) SetPropertyDS_certificateAuthorityObject(value string) (err error) {
	return instance.SetProperty("DS_certificateAuthorityObject", (value))
}

// GetDS_certificateAuthorityObject gets the value of DS_certificateAuthorityObject for the instance
func (instance *ads_crldistributionpoint) GetPropertyDS_certificateAuthorityObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_certificateAuthorityObject")
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

// SetDS_certificateRevocationList sets the value of DS_certificateRevocationList for the instance
func (instance *ads_crldistributionpoint) SetPropertyDS_certificateRevocationList(value Uint8Array) (err error) {
	return instance.SetProperty("DS_certificateRevocationList", (value))
}

// GetDS_certificateRevocationList gets the value of DS_certificateRevocationList for the instance
func (instance *ads_crldistributionpoint) GetPropertyDS_certificateRevocationList() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_certificateRevocationList")
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

// SetDS_cRLPartitionedRevocationList sets the value of DS_cRLPartitionedRevocationList for the instance
func (instance *ads_crldistributionpoint) SetPropertyDS_cRLPartitionedRevocationList(value Uint8Array) (err error) {
	return instance.SetProperty("DS_cRLPartitionedRevocationList", (value))
}

// GetDS_cRLPartitionedRevocationList gets the value of DS_cRLPartitionedRevocationList for the instance
func (instance *ads_crldistributionpoint) GetPropertyDS_cRLPartitionedRevocationList() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_cRLPartitionedRevocationList")
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

// SetDS_deltaRevocationList sets the value of DS_deltaRevocationList for the instance
func (instance *ads_crldistributionpoint) SetPropertyDS_deltaRevocationList(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_deltaRevocationList", (value))
}

// GetDS_deltaRevocationList gets the value of DS_deltaRevocationList for the instance
func (instance *ads_crldistributionpoint) GetPropertyDS_deltaRevocationList() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_deltaRevocationList")
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
