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

// ads_msds_azadminmanager struct
type ads_msds_azadminmanager struct {
	*ds_top

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzDomainTimeout int32

	//
	DS_msDS_AzGenerateAudits bool

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzMajorVersion int32

	//
	DS_msDS_AzMinorVersion int32

	//
	DS_msDS_AzObjectGuid Uint8Array

	//
	DS_msDS_AzScriptEngineCacheMax int32

	//
	DS_msDS_AzScriptTimeout int32
}

func Newads_msds_azadminmanagerEx1(instance *cim.WmiInstance) (newInstance *ads_msds_azadminmanager, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azadminmanager{
		ds_top: tmp,
	}
	return
}

func Newads_msds_azadminmanagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_azadminmanager, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_azadminmanager{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
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

// SetDS_msDS_AzDomainTimeout sets the value of DS_msDS_AzDomainTimeout for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzDomainTimeout(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzDomainTimeout", (value))
}

// GetDS_msDS_AzDomainTimeout gets the value of DS_msDS_AzDomainTimeout for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzDomainTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzDomainTimeout")
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

// SetDS_msDS_AzGenerateAudits sets the value of DS_msDS_AzGenerateAudits for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzGenerateAudits(value bool) (err error) {
	return instance.SetProperty("DS_msDS_AzGenerateAudits", (value))
}

// GetDS_msDS_AzGenerateAudits gets the value of DS_msDS_AzGenerateAudits for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzGenerateAudits() (value bool, err error) {
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
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
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

// SetDS_msDS_AzMajorVersion sets the value of DS_msDS_AzMajorVersion for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzMajorVersion(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzMajorVersion", (value))
}

// GetDS_msDS_AzMajorVersion gets the value of DS_msDS_AzMajorVersion for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzMajorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzMajorVersion")
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

// SetDS_msDS_AzMinorVersion sets the value of DS_msDS_AzMinorVersion for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzMinorVersion(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzMinorVersion", (value))
}

// GetDS_msDS_AzMinorVersion gets the value of DS_msDS_AzMinorVersion for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzMinorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzMinorVersion")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
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

// SetDS_msDS_AzScriptEngineCacheMax sets the value of DS_msDS_AzScriptEngineCacheMax for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzScriptEngineCacheMax(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzScriptEngineCacheMax", (value))
}

// GetDS_msDS_AzScriptEngineCacheMax gets the value of DS_msDS_AzScriptEngineCacheMax for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzScriptEngineCacheMax() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzScriptEngineCacheMax")
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

// SetDS_msDS_AzScriptTimeout sets the value of DS_msDS_AzScriptTimeout for the instance
func (instance *ads_msds_azadminmanager) SetPropertyDS_msDS_AzScriptTimeout(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AzScriptTimeout", (value))
}

// GetDS_msDS_AzScriptTimeout gets the value of DS_msDS_AzScriptTimeout for the instance
func (instance *ads_msds_azadminmanager) GetPropertyDS_msDS_AzScriptTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzScriptTimeout")
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
