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

// ads_ridmanager struct
type ads_ridmanager struct {
	*ds_top

	//
	DS_msDS_RIDPoolAllocationEnabled bool

	//
	DS_rIDAvailablePool int64
}

func Newads_ridmanagerEx1(instance *cim.WmiInstance) (newInstance *ads_ridmanager, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ridmanager{
		ds_top: tmp,
	}
	return
}

func Newads_ridmanagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ridmanager, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ridmanager{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_RIDPoolAllocationEnabled sets the value of DS_msDS_RIDPoolAllocationEnabled for the instance
func (instance *ads_ridmanager) SetPropertyDS_msDS_RIDPoolAllocationEnabled(value bool) (err error) {
	return instance.SetProperty("DS_msDS_RIDPoolAllocationEnabled", (value))
}

// GetDS_msDS_RIDPoolAllocationEnabled gets the value of DS_msDS_RIDPoolAllocationEnabled for the instance
func (instance *ads_ridmanager) GetPropertyDS_msDS_RIDPoolAllocationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RIDPoolAllocationEnabled")
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

// SetDS_rIDAvailablePool sets the value of DS_rIDAvailablePool for the instance
func (instance *ads_ridmanager) SetPropertyDS_rIDAvailablePool(value int64) (err error) {
	return instance.SetProperty("DS_rIDAvailablePool", (value))
}

// GetDS_rIDAvailablePool gets the value of DS_rIDAvailablePool for the instance
func (instance *ads_ridmanager) GetPropertyDS_rIDAvailablePool() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_rIDAvailablePool")
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
