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

// ads_serviceconnectionpoint struct
type ads_serviceconnectionpoint struct {
	*ds_connectionpoint

	//
	DS_appSchemaVersion int32

	//
	DS_serviceBindingInformation []string

	//
	DS_serviceClassName string

	//
	DS_serviceDNSName string

	//
	DS_serviceDNSNameType string

	//
	DS_vendor string

	//
	DS_versionNumber int32

	//
	DS_versionNumberHi int32

	//
	DS_versionNumberLo int32
}

func Newads_serviceconnectionpointEx1(instance *cim.WmiInstance) (newInstance *ads_serviceconnectionpoint, err error) {
	tmp, err := Newds_connectionpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_serviceconnectionpoint{
		ds_connectionpoint: tmp,
	}
	return
}

func Newads_serviceconnectionpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_serviceconnectionpoint, err error) {
	tmp, err := Newds_connectionpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_serviceconnectionpoint{
		ds_connectionpoint: tmp,
	}
	return
}

// SetDS_appSchemaVersion sets the value of DS_appSchemaVersion for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_appSchemaVersion(value int32) (err error) {
	return instance.SetProperty("DS_appSchemaVersion", (value))
}

// GetDS_appSchemaVersion gets the value of DS_appSchemaVersion for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_appSchemaVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_appSchemaVersion")
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

// SetDS_serviceBindingInformation sets the value of DS_serviceBindingInformation for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_serviceBindingInformation(value []string) (err error) {
	return instance.SetProperty("DS_serviceBindingInformation", (value))
}

// GetDS_serviceBindingInformation gets the value of DS_serviceBindingInformation for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_serviceBindingInformation() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_serviceBindingInformation")
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

// SetDS_serviceClassName sets the value of DS_serviceClassName for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_serviceClassName(value string) (err error) {
	return instance.SetProperty("DS_serviceClassName", (value))
}

// GetDS_serviceClassName gets the value of DS_serviceClassName for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_serviceClassName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serviceClassName")
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

// SetDS_serviceDNSName sets the value of DS_serviceDNSName for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_serviceDNSName(value string) (err error) {
	return instance.SetProperty("DS_serviceDNSName", (value))
}

// GetDS_serviceDNSName gets the value of DS_serviceDNSName for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_serviceDNSName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serviceDNSName")
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

// SetDS_serviceDNSNameType sets the value of DS_serviceDNSNameType for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_serviceDNSNameType(value string) (err error) {
	return instance.SetProperty("DS_serviceDNSNameType", (value))
}

// GetDS_serviceDNSNameType gets the value of DS_serviceDNSNameType for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_serviceDNSNameType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serviceDNSNameType")
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

// SetDS_vendor sets the value of DS_vendor for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_vendor(value string) (err error) {
	return instance.SetProperty("DS_vendor", (value))
}

// GetDS_vendor gets the value of DS_vendor for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_vendor() (value string, err error) {
	retValue, err := instance.GetProperty("DS_vendor")
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

// SetDS_versionNumber sets the value of DS_versionNumber for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_versionNumber(value int32) (err error) {
	return instance.SetProperty("DS_versionNumber", (value))
}

// GetDS_versionNumber gets the value of DS_versionNumber for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_versionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumber")
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

// SetDS_versionNumberHi sets the value of DS_versionNumberHi for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_versionNumberHi(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberHi", (value))
}

// GetDS_versionNumberHi gets the value of DS_versionNumberHi for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_versionNumberHi() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumberHi")
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

// SetDS_versionNumberLo sets the value of DS_versionNumberLo for the instance
func (instance *ads_serviceconnectionpoint) SetPropertyDS_versionNumberLo(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberLo", (value))
}

// GetDS_versionNumberLo gets the value of DS_versionNumberLo for the instance
func (instance *ads_serviceconnectionpoint) GetPropertyDS_versionNumberLo() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumberLo")
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
