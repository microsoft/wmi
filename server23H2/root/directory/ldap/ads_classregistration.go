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

// ads_classregistration struct
type ads_classregistration struct {
	*ds_leaf

	//
	DS_cOMCLSID string

	//
	DS_cOMInterfaceID []string

	//
	DS_cOMOtherProgId []string

	//
	DS_cOMProgID []string

	//
	DS_cOMTreatAsClassId string

	//
	DS_implementedCategories []Uint8Array

	//
	DS_managedBy string

	//
	DS_requiredCategories []Uint8Array
}

func Newads_classregistrationEx1(instance *cim.WmiInstance) (newInstance *ads_classregistration, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_classregistration{
		ds_leaf: tmp,
	}
	return
}

func Newads_classregistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_classregistration, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_classregistration{
		ds_leaf: tmp,
	}
	return
}

// SetDS_cOMCLSID sets the value of DS_cOMCLSID for the instance
func (instance *ads_classregistration) SetPropertyDS_cOMCLSID(value string) (err error) {
	return instance.SetProperty("DS_cOMCLSID", (value))
}

// GetDS_cOMCLSID gets the value of DS_cOMCLSID for the instance
func (instance *ads_classregistration) GetPropertyDS_cOMCLSID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cOMCLSID")
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

// SetDS_cOMInterfaceID sets the value of DS_cOMInterfaceID for the instance
func (instance *ads_classregistration) SetPropertyDS_cOMInterfaceID(value []string) (err error) {
	return instance.SetProperty("DS_cOMInterfaceID", (value))
}

// GetDS_cOMInterfaceID gets the value of DS_cOMInterfaceID for the instance
func (instance *ads_classregistration) GetPropertyDS_cOMInterfaceID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMInterfaceID")
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

// SetDS_cOMOtherProgId sets the value of DS_cOMOtherProgId for the instance
func (instance *ads_classregistration) SetPropertyDS_cOMOtherProgId(value []string) (err error) {
	return instance.SetProperty("DS_cOMOtherProgId", (value))
}

// GetDS_cOMOtherProgId gets the value of DS_cOMOtherProgId for the instance
func (instance *ads_classregistration) GetPropertyDS_cOMOtherProgId() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMOtherProgId")
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

// SetDS_cOMProgID sets the value of DS_cOMProgID for the instance
func (instance *ads_classregistration) SetPropertyDS_cOMProgID(value []string) (err error) {
	return instance.SetProperty("DS_cOMProgID", (value))
}

// GetDS_cOMProgID gets the value of DS_cOMProgID for the instance
func (instance *ads_classregistration) GetPropertyDS_cOMProgID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMProgID")
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

// SetDS_cOMTreatAsClassId sets the value of DS_cOMTreatAsClassId for the instance
func (instance *ads_classregistration) SetPropertyDS_cOMTreatAsClassId(value string) (err error) {
	return instance.SetProperty("DS_cOMTreatAsClassId", (value))
}

// GetDS_cOMTreatAsClassId gets the value of DS_cOMTreatAsClassId for the instance
func (instance *ads_classregistration) GetPropertyDS_cOMTreatAsClassId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cOMTreatAsClassId")
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

// SetDS_implementedCategories sets the value of DS_implementedCategories for the instance
func (instance *ads_classregistration) SetPropertyDS_implementedCategories(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_implementedCategories", (value))
}

// GetDS_implementedCategories gets the value of DS_implementedCategories for the instance
func (instance *ads_classregistration) GetPropertyDS_implementedCategories() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_implementedCategories")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_classregistration) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_classregistration) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_requiredCategories sets the value of DS_requiredCategories for the instance
func (instance *ads_classregistration) SetPropertyDS_requiredCategories(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_requiredCategories", (value))
}

// GetDS_requiredCategories gets the value of DS_requiredCategories for the instance
func (instance *ads_classregistration) GetPropertyDS_requiredCategories() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_requiredCategories")
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
