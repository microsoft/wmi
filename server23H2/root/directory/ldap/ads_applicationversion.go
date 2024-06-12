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

// ads_applicationversion struct
type ads_applicationversion struct {
	*ds_applicationsettings

	//
	DS_appSchemaVersion int32

	//
	DS_keywords []string

	//
	DS_managedBy string

	//
	DS_owner string

	//
	DS_vendor string

	//
	DS_versionNumber int32

	//
	DS_versionNumberHi int32

	//
	DS_versionNumberLo int32
}

func Newads_applicationversionEx1(instance *cim.WmiInstance) (newInstance *ads_applicationversion, err error) {
	tmp, err := Newds_applicationsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_applicationversion{
		ds_applicationsettings: tmp,
	}
	return
}

func Newads_applicationversionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_applicationversion, err error) {
	tmp, err := Newds_applicationsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_applicationversion{
		ds_applicationsettings: tmp,
	}
	return
}

// SetDS_appSchemaVersion sets the value of DS_appSchemaVersion for the instance
func (instance *ads_applicationversion) SetPropertyDS_appSchemaVersion(value int32) (err error) {
	return instance.SetProperty("DS_appSchemaVersion", (value))
}

// GetDS_appSchemaVersion gets the value of DS_appSchemaVersion for the instance
func (instance *ads_applicationversion) GetPropertyDS_appSchemaVersion() (value int32, err error) {
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

// SetDS_keywords sets the value of DS_keywords for the instance
func (instance *ads_applicationversion) SetPropertyDS_keywords(value []string) (err error) {
	return instance.SetProperty("DS_keywords", (value))
}

// GetDS_keywords gets the value of DS_keywords for the instance
func (instance *ads_applicationversion) GetPropertyDS_keywords() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_keywords")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_applicationversion) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_applicationversion) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_owner sets the value of DS_owner for the instance
func (instance *ads_applicationversion) SetPropertyDS_owner(value string) (err error) {
	return instance.SetProperty("DS_owner", (value))
}

// GetDS_owner gets the value of DS_owner for the instance
func (instance *ads_applicationversion) GetPropertyDS_owner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_owner")
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
func (instance *ads_applicationversion) SetPropertyDS_vendor(value string) (err error) {
	return instance.SetProperty("DS_vendor", (value))
}

// GetDS_vendor gets the value of DS_vendor for the instance
func (instance *ads_applicationversion) GetPropertyDS_vendor() (value string, err error) {
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
func (instance *ads_applicationversion) SetPropertyDS_versionNumber(value int32) (err error) {
	return instance.SetProperty("DS_versionNumber", (value))
}

// GetDS_versionNumber gets the value of DS_versionNumber for the instance
func (instance *ads_applicationversion) GetPropertyDS_versionNumber() (value int32, err error) {
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
func (instance *ads_applicationversion) SetPropertyDS_versionNumberHi(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberHi", (value))
}

// GetDS_versionNumberHi gets the value of DS_versionNumberHi for the instance
func (instance *ads_applicationversion) GetPropertyDS_versionNumberHi() (value int32, err error) {
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
func (instance *ads_applicationversion) SetPropertyDS_versionNumberLo(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberLo", (value))
}

// GetDS_versionNumberLo gets the value of DS_versionNumberLo for the instance
func (instance *ads_applicationversion) GetPropertyDS_versionNumberLo() (value int32, err error) {
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
