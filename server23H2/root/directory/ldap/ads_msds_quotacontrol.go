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

// ads_msds_quotacontrol struct
type ads_msds_quotacontrol struct {
	*ds_top

	//
	DS_msDS_QuotaAmount int32

	//
	DS_msDS_QuotaTrustee Uint8Array
}

func Newads_msds_quotacontrolEx1(instance *cim.WmiInstance) (newInstance *ads_msds_quotacontrol, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_quotacontrol{
		ds_top: tmp,
	}
	return
}

func Newads_msds_quotacontrolEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_quotacontrol, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_quotacontrol{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_QuotaAmount sets the value of DS_msDS_QuotaAmount for the instance
func (instance *ads_msds_quotacontrol) SetPropertyDS_msDS_QuotaAmount(value int32) (err error) {
	return instance.SetProperty("DS_msDS_QuotaAmount", (value))
}

// GetDS_msDS_QuotaAmount gets the value of DS_msDS_QuotaAmount for the instance
func (instance *ads_msds_quotacontrol) GetPropertyDS_msDS_QuotaAmount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_QuotaAmount")
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

// SetDS_msDS_QuotaTrustee sets the value of DS_msDS_QuotaTrustee for the instance
func (instance *ads_msds_quotacontrol) SetPropertyDS_msDS_QuotaTrustee(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_QuotaTrustee", (value))
}

// GetDS_msDS_QuotaTrustee gets the value of DS_msDS_QuotaTrustee for the instance
func (instance *ads_msds_quotacontrol) GetPropertyDS_msDS_QuotaTrustee() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_QuotaTrustee")
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
