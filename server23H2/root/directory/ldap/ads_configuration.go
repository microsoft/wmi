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

// ads_configuration struct
type ads_configuration struct {
	*ds_top

	//
	DS_gPLink string

	//
	DS_gPOptions int32

	//
	DS_msDS_USNLastSyncSuccess int64
}

func Newads_configurationEx1(instance *cim.WmiInstance) (newInstance *ads_configuration, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_configuration{
		ds_top: tmp,
	}
	return
}

func Newads_configurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_configuration, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_configuration{
		ds_top: tmp,
	}
	return
}

// SetDS_gPLink sets the value of DS_gPLink for the instance
func (instance *ads_configuration) SetPropertyDS_gPLink(value string) (err error) {
	return instance.SetProperty("DS_gPLink", (value))
}

// GetDS_gPLink gets the value of DS_gPLink for the instance
func (instance *ads_configuration) GetPropertyDS_gPLink() (value string, err error) {
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
func (instance *ads_configuration) SetPropertyDS_gPOptions(value int32) (err error) {
	return instance.SetProperty("DS_gPOptions", (value))
}

// GetDS_gPOptions gets the value of DS_gPOptions for the instance
func (instance *ads_configuration) GetPropertyDS_gPOptions() (value int32, err error) {
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

// SetDS_msDS_USNLastSyncSuccess sets the value of DS_msDS_USNLastSyncSuccess for the instance
func (instance *ads_configuration) SetPropertyDS_msDS_USNLastSyncSuccess(value int64) (err error) {
	return instance.SetProperty("DS_msDS_USNLastSyncSuccess", (value))
}

// GetDS_msDS_USNLastSyncSuccess gets the value of DS_msDS_USNLastSyncSuccess for the instance
func (instance *ads_configuration) GetPropertyDS_msDS_USNLastSyncSuccess() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_USNLastSyncSuccess")
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
