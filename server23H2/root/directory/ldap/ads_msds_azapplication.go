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

// ads_msds_azapplication struct
type ads_msds_azapplication struct {
	*ds_top

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzApplicationName string

	//
	DS_msDS_AzApplicationVersion string

	//
	DS_msDS_AzClassId string

	//
	DS_msDS_AzGenerateAudits bool

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzObjectGuid Uint8Array
}

func Newads_msds_azapplicationEx1(instance *cim.WmiInstance) (newInstance *ads_msds_azapplication, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azapplication{
		ds_top: tmp,
	}
	return
}

func Newads_msds_azapplicationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_azapplication, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azapplication{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationData")
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

// SetDS_msDS_AzApplicationName sets the value of DS_msDS_AzApplicationName for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzApplicationName(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationName", (value))
}

// GetDS_msDS_AzApplicationName gets the value of DS_msDS_AzApplicationName for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzApplicationName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationName")
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

// SetDS_msDS_AzApplicationVersion sets the value of DS_msDS_AzApplicationVersion for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzApplicationVersion(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationVersion", (value))
}

// GetDS_msDS_AzApplicationVersion gets the value of DS_msDS_AzApplicationVersion for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzApplicationVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationVersion")
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

// SetDS_msDS_AzClassId sets the value of DS_msDS_AzClassId for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzClassId(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzClassId", (value))
}

// GetDS_msDS_AzClassId gets the value of DS_msDS_AzClassId for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzClassId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzClassId")
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

// SetDS_msDS_AzGenerateAudits sets the value of DS_msDS_AzGenerateAudits for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzGenerateAudits(value bool) (err error) {
	return instance.SetProperty("DS_msDS_AzGenerateAudits", (value))
}

// GetDS_msDS_AzGenerateAudits gets the value of DS_msDS_AzGenerateAudits for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzGenerateAudits() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzGenerateAudits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDS_msDS_AzGenericData sets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzGenericData")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azapplication) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azapplication) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzObjectGuid")
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
