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

// ads_site struct
type ads_site struct {
	*ds_top

	//
	DS_gPLink string

	//
	DS_gPOptions int32

	//
	DS_location string

	//
	DS_managedBy string

	//
	DS_msDS_BridgeHeadServersUsed []DN_With_Binary

	//
	DS_mSMQInterval1 int32

	//
	DS_mSMQInterval2 int32

	//
	DS_mSMQNt4Stub []int32

	//
	DS_mSMQSiteForeign []bool

	//
	DS_mSMQSiteID Uint8Array

	//
	DS_notificationList string
}

func Newads_siteEx1(instance *cim.WmiInstance) (newInstance *ads_site, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_site{
		ds_top: tmp,
	}
	return
}

func Newads_siteEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_site, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_site{
		ds_top: tmp,
	}
	return
}

// SetDS_gPLink sets the value of DS_gPLink for the instance
func (instance *ads_site) SetPropertyDS_gPLink(value string) (err error) {
	return instance.SetProperty("DS_gPLink", (value))
}

// GetDS_gPLink gets the value of DS_gPLink for the instance
func (instance *ads_site) GetPropertyDS_gPLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPLink")
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

// SetDS_gPOptions sets the value of DS_gPOptions for the instance
func (instance *ads_site) SetPropertyDS_gPOptions(value int32) (err error) {
	return instance.SetProperty("DS_gPOptions", (value))
}

// GetDS_gPOptions gets the value of DS_gPOptions for the instance
func (instance *ads_site) GetPropertyDS_gPOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_gPOptions")
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

// SetDS_location sets the value of DS_location for the instance
func (instance *ads_site) SetPropertyDS_location(value string) (err error) {
	return instance.SetProperty("DS_location", (value))
}

// GetDS_location gets the value of DS_location for the instance
func (instance *ads_site) GetPropertyDS_location() (value string, err error) {
	retValue, err := instance.GetProperty("DS_location")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_site) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_site) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_msDS_BridgeHeadServersUsed sets the value of DS_msDS_BridgeHeadServersUsed for the instance
func (instance *ads_site) SetPropertyDS_msDS_BridgeHeadServersUsed(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_BridgeHeadServersUsed", (value))
}

// GetDS_msDS_BridgeHeadServersUsed gets the value of DS_msDS_BridgeHeadServersUsed for the instance
func (instance *ads_site) GetPropertyDS_msDS_BridgeHeadServersUsed() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msDS_BridgeHeadServersUsed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_mSMQInterval1 sets the value of DS_mSMQInterval1 for the instance
func (instance *ads_site) SetPropertyDS_mSMQInterval1(value int32) (err error) {
	return instance.SetProperty("DS_mSMQInterval1", (value))
}

// GetDS_mSMQInterval1 gets the value of DS_mSMQInterval1 for the instance
func (instance *ads_site) GetPropertyDS_mSMQInterval1() (value int32, err error) {
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
func (instance *ads_site) SetPropertyDS_mSMQInterval2(value int32) (err error) {
	return instance.SetProperty("DS_mSMQInterval2", (value))
}

// GetDS_mSMQInterval2 gets the value of DS_mSMQInterval2 for the instance
func (instance *ads_site) GetPropertyDS_mSMQInterval2() (value int32, err error) {
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

// SetDS_mSMQNt4Stub sets the value of DS_mSMQNt4Stub for the instance
func (instance *ads_site) SetPropertyDS_mSMQNt4Stub(value []int32) (err error) {
	return instance.SetProperty("DS_mSMQNt4Stub", (value))
}

// GetDS_mSMQNt4Stub gets the value of DS_mSMQNt4Stub for the instance
func (instance *ads_site) GetPropertyDS_mSMQNt4Stub() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_mSMQNt4Stub")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_mSMQSiteForeign sets the value of DS_mSMQSiteForeign for the instance
func (instance *ads_site) SetPropertyDS_mSMQSiteForeign(value []bool) (err error) {
	return instance.SetProperty("DS_mSMQSiteForeign", (value))
}

// GetDS_mSMQSiteForeign gets the value of DS_mSMQSiteForeign for the instance
func (instance *ads_site) GetPropertyDS_mSMQSiteForeign() (value []bool, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSiteForeign")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetDS_mSMQSiteID sets the value of DS_mSMQSiteID for the instance
func (instance *ads_site) SetPropertyDS_mSMQSiteID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSiteID", (value))
}

// GetDS_mSMQSiteID gets the value of DS_mSMQSiteID for the instance
func (instance *ads_site) GetPropertyDS_mSMQSiteID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSiteID")
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

// SetDS_notificationList sets the value of DS_notificationList for the instance
func (instance *ads_site) SetPropertyDS_notificationList(value string) (err error) {
	return instance.SetProperty("DS_notificationList", (value))
}

// GetDS_notificationList gets the value of DS_notificationList for the instance
func (instance *ads_site) GetPropertyDS_notificationList() (value string, err error) {
	retValue, err := instance.GetProperty("DS_notificationList")
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
