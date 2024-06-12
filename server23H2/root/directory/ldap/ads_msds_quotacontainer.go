// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msds_quotacontainer struct
type ads_msds_quotacontainer struct {
	*ds_top

	//
	DS_msDS_DefaultQuota int32

	//
	DS_msDS_QuotaEffective int32

	//
	DS_msDS_QuotaUsed int32

	//
	DS_msDS_TombstoneQuotaFactor int32

	//
	DS_msDS_TopQuotaUsage []string
}

func Newads_msds_quotacontainerEx1(instance *cim.WmiInstance) (newInstance *ads_msds_quotacontainer, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_quotacontainer{
		ds_top: tmp,
	}
	return
}

func Newads_msds_quotacontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_quotacontainer, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_quotacontainer{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_DefaultQuota sets the value of DS_msDS_DefaultQuota for the instance
func (instance *ads_msds_quotacontainer) SetPropertyDS_msDS_DefaultQuota(value int32) (err error) {
	return instance.SetProperty("DS_msDS_DefaultQuota", (value))
}

// GetDS_msDS_DefaultQuota gets the value of DS_msDS_DefaultQuota for the instance
func (instance *ads_msds_quotacontainer) GetPropertyDS_msDS_DefaultQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_DefaultQuota")
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

// SetDS_msDS_QuotaEffective sets the value of DS_msDS_QuotaEffective for the instance
func (instance *ads_msds_quotacontainer) SetPropertyDS_msDS_QuotaEffective(value int32) (err error) {
	return instance.SetProperty("DS_msDS_QuotaEffective", (value))
}

// GetDS_msDS_QuotaEffective gets the value of DS_msDS_QuotaEffective for the instance
func (instance *ads_msds_quotacontainer) GetPropertyDS_msDS_QuotaEffective() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_QuotaEffective")
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

// SetDS_msDS_QuotaUsed sets the value of DS_msDS_QuotaUsed for the instance
func (instance *ads_msds_quotacontainer) SetPropertyDS_msDS_QuotaUsed(value int32) (err error) {
	return instance.SetProperty("DS_msDS_QuotaUsed", (value))
}

// GetDS_msDS_QuotaUsed gets the value of DS_msDS_QuotaUsed for the instance
func (instance *ads_msds_quotacontainer) GetPropertyDS_msDS_QuotaUsed() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_QuotaUsed")
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

// SetDS_msDS_TombstoneQuotaFactor sets the value of DS_msDS_TombstoneQuotaFactor for the instance
func (instance *ads_msds_quotacontainer) SetPropertyDS_msDS_TombstoneQuotaFactor(value int32) (err error) {
	return instance.SetProperty("DS_msDS_TombstoneQuotaFactor", (value))
}

// GetDS_msDS_TombstoneQuotaFactor gets the value of DS_msDS_TombstoneQuotaFactor for the instance
func (instance *ads_msds_quotacontainer) GetPropertyDS_msDS_TombstoneQuotaFactor() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TombstoneQuotaFactor")
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

// SetDS_msDS_TopQuotaUsage sets the value of DS_msDS_TopQuotaUsage for the instance
func (instance *ads_msds_quotacontainer) SetPropertyDS_msDS_TopQuotaUsage(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TopQuotaUsage", (value))
}

// GetDS_msDS_TopQuotaUsage gets the value of DS_msDS_TopQuotaUsage for the instance
func (instance *ads_msds_quotacontainer) GetPropertyDS_msDS_TopQuotaUsage() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TopQuotaUsage")
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
