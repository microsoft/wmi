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

// ads_msmqsettings struct
type ads_msmqsettings struct {
	*ds_top

	//
	DS_mSMQDependentClientService bool

	//
	DS_mSMQDsService bool

	//
	DS_mSMQMigrated bool

	//
	DS_mSMQNt4Flags int32

	//
	DS_mSMQOwnerID Uint8Array

	//
	DS_mSMQQMID Uint8Array

	//
	DS_mSMQRoutingService bool

	//
	DS_mSMQServices int32

	//
	DS_mSMQSiteName string

	//
	DS_mSMQSiteNameEx string
}

func Newads_msmqsettingsEx1(instance *cim.WmiInstance) (newInstance *ads_msmqsettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmqsettings{
		ds_top: tmp,
	}
	return
}

func Newads_msmqsettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmqsettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmqsettings{
		ds_top: tmp,
	}
	return
}

// SetDS_mSMQDependentClientService sets the value of DS_mSMQDependentClientService for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQDependentClientService(value bool) (err error) {
	return instance.SetProperty("DS_mSMQDependentClientService", (value))
}

// GetDS_mSMQDependentClientService gets the value of DS_mSMQDependentClientService for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQDependentClientService() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDependentClientService")
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

// SetDS_mSMQDsService sets the value of DS_mSMQDsService for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQDsService(value bool) (err error) {
	return instance.SetProperty("DS_mSMQDsService", (value))
}

// GetDS_mSMQDsService gets the value of DS_mSMQDsService for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQDsService() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDsService")
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

// SetDS_mSMQMigrated sets the value of DS_mSMQMigrated for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQMigrated(value bool) (err error) {
	return instance.SetProperty("DS_mSMQMigrated", (value))
}

// GetDS_mSMQMigrated gets the value of DS_mSMQMigrated for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQMigrated() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQMigrated")
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

// SetDS_mSMQNt4Flags sets the value of DS_mSMQNt4Flags for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQNt4Flags(value int32) (err error) {
	return instance.SetProperty("DS_mSMQNt4Flags", (value))
}

// GetDS_mSMQNt4Flags gets the value of DS_mSMQNt4Flags for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQNt4Flags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQNt4Flags")
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

// SetDS_mSMQOwnerID sets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQOwnerID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQOwnerID", (value))
}

// GetDS_mSMQOwnerID gets the value of DS_mSMQOwnerID for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQOwnerID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQOwnerID")
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

// SetDS_mSMQQMID sets the value of DS_mSMQQMID for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQQMID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQQMID", (value))
}

// GetDS_mSMQQMID gets the value of DS_mSMQQMID for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQQMID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQQMID")
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

// SetDS_mSMQRoutingService sets the value of DS_mSMQRoutingService for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQRoutingService(value bool) (err error) {
	return instance.SetProperty("DS_mSMQRoutingService", (value))
}

// GetDS_mSMQRoutingService gets the value of DS_mSMQRoutingService for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQRoutingService() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQRoutingService")
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

// SetDS_mSMQServices sets the value of DS_mSMQServices for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQServices(value int32) (err error) {
	return instance.SetProperty("DS_mSMQServices", (value))
}

// GetDS_mSMQServices gets the value of DS_mSMQServices for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQServices() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQServices")
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

// SetDS_mSMQSiteName sets the value of DS_mSMQSiteName for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQSiteName(value string) (err error) {
	return instance.SetProperty("DS_mSMQSiteName", (value))
}

// GetDS_mSMQSiteName gets the value of DS_mSMQSiteName for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQSiteName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSiteName")
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

// SetDS_mSMQSiteNameEx sets the value of DS_mSMQSiteNameEx for the instance
func (instance *ads_msmqsettings) SetPropertyDS_mSMQSiteNameEx(value string) (err error) {
	return instance.SetProperty("DS_mSMQSiteNameEx", (value))
}

// GetDS_mSMQSiteNameEx gets the value of DS_mSMQSiteNameEx for the instance
func (instance *ads_msmqsettings) GetPropertyDS_mSMQSiteNameEx() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSiteNameEx")
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
