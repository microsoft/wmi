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

// ads_mscom_partitionset struct
type ads_mscom_partitionset struct {
	*ds_top

	//
	DS_msCOM_DefaultPartitionLink string

	//
	DS_msCOM_ObjectId Uint8Array

	//
	DS_msCOM_PartitionLink []string
}

func Newads_mscom_partitionsetEx1(instance *cim.WmiInstance) (newInstance *ads_mscom_partitionset, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mscom_partitionset{
		ds_top: tmp,
	}
	return
}

func Newads_mscom_partitionsetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mscom_partitionset, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mscom_partitionset{
		ds_top: tmp,
	}
	return
}

// SetDS_msCOM_DefaultPartitionLink sets the value of DS_msCOM_DefaultPartitionLink for the instance
func (instance *ads_mscom_partitionset) SetPropertyDS_msCOM_DefaultPartitionLink(value string) (err error) {
	return instance.SetProperty("DS_msCOM_DefaultPartitionLink", (value))
}

// GetDS_msCOM_DefaultPartitionLink gets the value of DS_msCOM_DefaultPartitionLink for the instance
func (instance *ads_mscom_partitionset) GetPropertyDS_msCOM_DefaultPartitionLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_DefaultPartitionLink")
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

// SetDS_msCOM_ObjectId sets the value of DS_msCOM_ObjectId for the instance
func (instance *ads_mscom_partitionset) SetPropertyDS_msCOM_ObjectId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msCOM_ObjectId", (value))
}

// GetDS_msCOM_ObjectId gets the value of DS_msCOM_ObjectId for the instance
func (instance *ads_mscom_partitionset) GetPropertyDS_msCOM_ObjectId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_ObjectId")
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

// SetDS_msCOM_PartitionLink sets the value of DS_msCOM_PartitionLink for the instance
func (instance *ads_mscom_partitionset) SetPropertyDS_msCOM_PartitionLink(value []string) (err error) {
	return instance.SetProperty("DS_msCOM_PartitionLink", (value))
}

// GetDS_msCOM_PartitionLink gets the value of DS_msCOM_PartitionLink for the instance
func (instance *ads_mscom_partitionset) GetPropertyDS_msCOM_PartitionLink() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_PartitionLink")
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
