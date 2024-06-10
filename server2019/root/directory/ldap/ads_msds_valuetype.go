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

// ads_msds_valuetype struct
type ads_msds_valuetype struct {
	*ds_top

	//
	DS_msDS_ClaimIsSingleValued bool

	//
	DS_msDS_ClaimIsValueSpaceRestricted bool

	//
	DS_msDS_ClaimValueType int64

	//
	DS_msDS_IsPossibleValuesPresent bool
}

func Newads_msds_valuetypeEx1(instance *cim.WmiInstance) (newInstance *ads_msds_valuetype, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_valuetype{
		ds_top: tmp,
	}
	return
}

func Newads_msds_valuetypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_valuetype, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_valuetype{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_ClaimIsSingleValued sets the value of DS_msDS_ClaimIsSingleValued for the instance
func (instance *ads_msds_valuetype) SetPropertyDS_msDS_ClaimIsSingleValued(value bool) (err error) {
	return instance.SetProperty("DS_msDS_ClaimIsSingleValued", (value))
}

// GetDS_msDS_ClaimIsSingleValued gets the value of DS_msDS_ClaimIsSingleValued for the instance
func (instance *ads_msds_valuetype) GetPropertyDS_msDS_ClaimIsSingleValued() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimIsSingleValued")
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

// SetDS_msDS_ClaimIsValueSpaceRestricted sets the value of DS_msDS_ClaimIsValueSpaceRestricted for the instance
func (instance *ads_msds_valuetype) SetPropertyDS_msDS_ClaimIsValueSpaceRestricted(value bool) (err error) {
	return instance.SetProperty("DS_msDS_ClaimIsValueSpaceRestricted", (value))
}

// GetDS_msDS_ClaimIsValueSpaceRestricted gets the value of DS_msDS_ClaimIsValueSpaceRestricted for the instance
func (instance *ads_msds_valuetype) GetPropertyDS_msDS_ClaimIsValueSpaceRestricted() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimIsValueSpaceRestricted")
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

// SetDS_msDS_ClaimValueType sets the value of DS_msDS_ClaimValueType for the instance
func (instance *ads_msds_valuetype) SetPropertyDS_msDS_ClaimValueType(value int64) (err error) {
	return instance.SetProperty("DS_msDS_ClaimValueType", (value))
}

// GetDS_msDS_ClaimValueType gets the value of DS_msDS_ClaimValueType for the instance
func (instance *ads_msds_valuetype) GetPropertyDS_msDS_ClaimValueType() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimValueType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_IsPossibleValuesPresent sets the value of DS_msDS_IsPossibleValuesPresent for the instance
func (instance *ads_msds_valuetype) SetPropertyDS_msDS_IsPossibleValuesPresent(value bool) (err error) {
	return instance.SetProperty("DS_msDS_IsPossibleValuesPresent", (value))
}

// GetDS_msDS_IsPossibleValuesPresent gets the value of DS_msDS_IsPossibleValuesPresent for the instance
func (instance *ads_msds_valuetype) GetPropertyDS_msDS_IsPossibleValuesPresent() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsPossibleValuesPresent")
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
