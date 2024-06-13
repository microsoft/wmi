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

// ads_msds_optionalfeature struct
type ads_msds_optionalfeature struct {
	*ds_top

	//
	DS_msDS_OptionalFeatureFlags int32

	//
	DS_msDS_OptionalFeatureGUID Uint8Array

	//
	DS_msDS_RequiredDomainBehaviorVersion int32

	//
	DS_msDS_RequiredForestBehaviorVersion int32
}

func Newads_msds_optionalfeatureEx1(instance *cim.WmiInstance) (newInstance *ads_msds_optionalfeature, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_optionalfeature{
		ds_top: tmp,
	}
	return
}

func Newads_msds_optionalfeatureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_optionalfeature, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_optionalfeature{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_OptionalFeatureFlags sets the value of DS_msDS_OptionalFeatureFlags for the instance
func (instance *ads_msds_optionalfeature) SetPropertyDS_msDS_OptionalFeatureFlags(value int32) (err error) {
	return instance.SetProperty("DS_msDS_OptionalFeatureFlags", (value))
}

// GetDS_msDS_OptionalFeatureFlags gets the value of DS_msDS_OptionalFeatureFlags for the instance
func (instance *ads_msds_optionalfeature) GetPropertyDS_msDS_OptionalFeatureFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OptionalFeatureFlags")
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

// SetDS_msDS_OptionalFeatureGUID sets the value of DS_msDS_OptionalFeatureGUID for the instance
func (instance *ads_msds_optionalfeature) SetPropertyDS_msDS_OptionalFeatureGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_OptionalFeatureGUID", (value))
}

// GetDS_msDS_OptionalFeatureGUID gets the value of DS_msDS_OptionalFeatureGUID for the instance
func (instance *ads_msds_optionalfeature) GetPropertyDS_msDS_OptionalFeatureGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OptionalFeatureGUID")
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

// SetDS_msDS_RequiredDomainBehaviorVersion sets the value of DS_msDS_RequiredDomainBehaviorVersion for the instance
func (instance *ads_msds_optionalfeature) SetPropertyDS_msDS_RequiredDomainBehaviorVersion(value int32) (err error) {
	return instance.SetProperty("DS_msDS_RequiredDomainBehaviorVersion", (value))
}

// GetDS_msDS_RequiredDomainBehaviorVersion gets the value of DS_msDS_RequiredDomainBehaviorVersion for the instance
func (instance *ads_msds_optionalfeature) GetPropertyDS_msDS_RequiredDomainBehaviorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RequiredDomainBehaviorVersion")
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

// SetDS_msDS_RequiredForestBehaviorVersion sets the value of DS_msDS_RequiredForestBehaviorVersion for the instance
func (instance *ads_msds_optionalfeature) SetPropertyDS_msDS_RequiredForestBehaviorVersion(value int32) (err error) {
	return instance.SetProperty("DS_msDS_RequiredForestBehaviorVersion", (value))
}

// GetDS_msDS_RequiredForestBehaviorVersion gets the value of DS_msDS_RequiredForestBehaviorVersion for the instance
func (instance *ads_msds_optionalfeature) GetPropertyDS_msDS_RequiredForestBehaviorVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RequiredForestBehaviorVersion")
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
