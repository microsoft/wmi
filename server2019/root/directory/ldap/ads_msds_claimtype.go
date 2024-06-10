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

// ads_msds_claimtype struct
type ads_msds_claimtype struct {
	*ds_msds_claimtypepropertybase

	//
	DS_msDS_ClaimAttributeSource string

	//
	DS_msDS_ClaimIsSingleValued bool

	//
	DS_msDS_ClaimIsValueSpaceRestricted bool

	//
	DS_msDS_ClaimSource string

	//
	DS_msDS_ClaimSourceType string

	//
	DS_msDS_ClaimTypeAppliesToClass []string

	//
	DS_msDS_ClaimValueType int64
}

func Newads_msds_claimtypeEx1(instance *cim.WmiInstance) (newInstance *ads_msds_claimtype, err error) {
	tmp, err := Newds_msds_claimtypepropertybaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_claimtype{
		ds_msds_claimtypepropertybase: tmp,
	}
	return
}

func Newads_msds_claimtypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_claimtype, err error) {
	tmp, err := Newds_msds_claimtypepropertybaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_claimtype{
		ds_msds_claimtypepropertybase: tmp,
	}
	return
}

// SetDS_msDS_ClaimAttributeSource sets the value of DS_msDS_ClaimAttributeSource for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimAttributeSource(value string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimAttributeSource", (value))
}

// GetDS_msDS_ClaimAttributeSource gets the value of DS_msDS_ClaimAttributeSource for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimAttributeSource() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimAttributeSource")
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

// SetDS_msDS_ClaimIsSingleValued sets the value of DS_msDS_ClaimIsSingleValued for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimIsSingleValued(value bool) (err error) {
	return instance.SetProperty("DS_msDS_ClaimIsSingleValued", (value))
}

// GetDS_msDS_ClaimIsSingleValued gets the value of DS_msDS_ClaimIsSingleValued for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimIsSingleValued() (value bool, err error) {
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
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimIsValueSpaceRestricted(value bool) (err error) {
	return instance.SetProperty("DS_msDS_ClaimIsValueSpaceRestricted", (value))
}

// GetDS_msDS_ClaimIsValueSpaceRestricted gets the value of DS_msDS_ClaimIsValueSpaceRestricted for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimIsValueSpaceRestricted() (value bool, err error) {
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

// SetDS_msDS_ClaimSource sets the value of DS_msDS_ClaimSource for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimSource(value string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimSource", (value))
}

// GetDS_msDS_ClaimSource gets the value of DS_msDS_ClaimSource for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimSource() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimSource")
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

// SetDS_msDS_ClaimSourceType sets the value of DS_msDS_ClaimSourceType for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimSourceType(value string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimSourceType", (value))
}

// GetDS_msDS_ClaimSourceType gets the value of DS_msDS_ClaimSourceType for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimSourceType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimSourceType")
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

// SetDS_msDS_ClaimTypeAppliesToClass sets the value of DS_msDS_ClaimTypeAppliesToClass for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimTypeAppliesToClass(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimTypeAppliesToClass", (value))
}

// GetDS_msDS_ClaimTypeAppliesToClass gets the value of DS_msDS_ClaimTypeAppliesToClass for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimTypeAppliesToClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimTypeAppliesToClass")
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

// SetDS_msDS_ClaimValueType sets the value of DS_msDS_ClaimValueType for the instance
func (instance *ads_msds_claimtype) SetPropertyDS_msDS_ClaimValueType(value int64) (err error) {
	return instance.SetProperty("DS_msDS_ClaimValueType", (value))
}

// GetDS_msDS_ClaimValueType gets the value of DS_msDS_ClaimValueType for the instance
func (instance *ads_msds_claimtype) GetPropertyDS_msDS_ClaimValueType() (value int64, err error) {
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
