// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_ntdssitesettings struct
type ads_ntdssitesettings struct {
	*ds_applicationsitesettings

	//
	DS_interSiteTopologyFailover int32

	//
	DS_interSiteTopologyGenerator string

	//
	DS_interSiteTopologyRenew int32

	//
	DS_managedBy string

	//
	DS_msDS_Preferred_GC_Site string

	//
	DS_options int32

	//
	DS_queryPolicyObject string

	//
	DS_schedule Uint8Array
}

func Newads_ntdssitesettingsEx1(instance *cim.WmiInstance) (newInstance *ads_ntdssitesettings, err error) {
	tmp, err := Newds_applicationsitesettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntdssitesettings{
		ds_applicationsitesettings: tmp,
	}
	return
}

func Newads_ntdssitesettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntdssitesettings, err error) {
	tmp, err := Newds_applicationsitesettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntdssitesettings{
		ds_applicationsitesettings: tmp,
	}
	return
}

// SetDS_interSiteTopologyFailover sets the value of DS_interSiteTopologyFailover for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_interSiteTopologyFailover(value int32) (err error) {
	return instance.SetProperty("DS_interSiteTopologyFailover", (value))
}

// GetDS_interSiteTopologyFailover gets the value of DS_interSiteTopologyFailover for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_interSiteTopologyFailover() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_interSiteTopologyFailover")
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

// SetDS_interSiteTopologyGenerator sets the value of DS_interSiteTopologyGenerator for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_interSiteTopologyGenerator(value string) (err error) {
	return instance.SetProperty("DS_interSiteTopologyGenerator", (value))
}

// GetDS_interSiteTopologyGenerator gets the value of DS_interSiteTopologyGenerator for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_interSiteTopologyGenerator() (value string, err error) {
	retValue, err := instance.GetProperty("DS_interSiteTopologyGenerator")
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

// SetDS_interSiteTopologyRenew sets the value of DS_interSiteTopologyRenew for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_interSiteTopologyRenew(value int32) (err error) {
	return instance.SetProperty("DS_interSiteTopologyRenew", (value))
}

// GetDS_interSiteTopologyRenew gets the value of DS_interSiteTopologyRenew for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_interSiteTopologyRenew() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_interSiteTopologyRenew")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_msDS_Preferred_GC_Site sets the value of DS_msDS_Preferred_GC_Site for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_msDS_Preferred_GC_Site(value string) (err error) {
	return instance.SetProperty("DS_msDS_Preferred_GC_Site", (value))
}

// GetDS_msDS_Preferred_GC_Site gets the value of DS_msDS_Preferred_GC_Site for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_msDS_Preferred_GC_Site() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Preferred_GC_Site")
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

// SetDS_options sets the value of DS_options for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_options(value int32) (err error) {
	return instance.SetProperty("DS_options", (value))
}

// GetDS_options gets the value of DS_options for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_options")
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

// SetDS_queryPolicyObject sets the value of DS_queryPolicyObject for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_queryPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_queryPolicyObject", (value))
}

// GetDS_queryPolicyObject gets the value of DS_queryPolicyObject for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_queryPolicyObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_queryPolicyObject")
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

// SetDS_schedule sets the value of DS_schedule for the instance
func (instance *ads_ntdssitesettings) SetPropertyDS_schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schedule", (value))
}

// GetDS_schedule gets the value of DS_schedule for the instance
func (instance *ads_ntdssitesettings) GetPropertyDS_schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schedule")
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
