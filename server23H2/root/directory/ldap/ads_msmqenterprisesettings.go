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

// ads_msmqenterprisesettings struct
type ads_msmqenterprisesettings struct {
	*ds_top

	//
	DS_mSMQCSPName string

	//
	DS_mSMQInterval1 int32

	//
	DS_mSMQInterval2 int32

	//
	DS_mSMQLongLived int32

	//
	DS_mSMQNameStyle bool

	//
	DS_mSMQVersion int32
}

func Newads_msmqenterprisesettingsEx1(instance *cim.WmiInstance) (newInstance *ads_msmqenterprisesettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msmqenterprisesettings{
		ds_top: tmp,
	}
	return
}

func Newads_msmqenterprisesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msmqenterprisesettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msmqenterprisesettings{
		ds_top: tmp,
	}
	return
}

// SetDS_mSMQCSPName sets the value of DS_mSMQCSPName for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQCSPName(value string) (err error) {
	return instance.SetProperty("DS_mSMQCSPName", (value))
}

// GetDS_mSMQCSPName gets the value of DS_mSMQCSPName for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQCSPName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mSMQCSPName")
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

// SetDS_mSMQInterval1 sets the value of DS_mSMQInterval1 for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQInterval1(value int32) (err error) {
	return instance.SetProperty("DS_mSMQInterval1", (value))
}

// GetDS_mSMQInterval1 gets the value of DS_mSMQInterval1 for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQInterval1() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQInterval1")
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

// SetDS_mSMQInterval2 sets the value of DS_mSMQInterval2 for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQInterval2(value int32) (err error) {
	return instance.SetProperty("DS_mSMQInterval2", (value))
}

// GetDS_mSMQInterval2 gets the value of DS_mSMQInterval2 for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQInterval2() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQInterval2")
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

// SetDS_mSMQLongLived sets the value of DS_mSMQLongLived for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQLongLived(value int32) (err error) {
	return instance.SetProperty("DS_mSMQLongLived", (value))
}

// GetDS_mSMQLongLived gets the value of DS_mSMQLongLived for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQLongLived() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQLongLived")
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

// SetDS_mSMQNameStyle sets the value of DS_mSMQNameStyle for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQNameStyle(value bool) (err error) {
	return instance.SetProperty("DS_mSMQNameStyle", (value))
}

// GetDS_mSMQNameStyle gets the value of DS_mSMQNameStyle for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQNameStyle() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQNameStyle")
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

// SetDS_mSMQVersion sets the value of DS_mSMQVersion for the instance
func (instance *ads_msmqenterprisesettings) SetPropertyDS_mSMQVersion(value int32) (err error) {
	return instance.SetProperty("DS_mSMQVersion", (value))
}

// GetDS_mSMQVersion gets the value of DS_mSMQVersion for the instance
func (instance *ads_msmqenterprisesettings) GetPropertyDS_mSMQVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQVersion")
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
