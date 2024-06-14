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

// ds_msds_claimtypepropertybase struct
type ds_msds_claimtypepropertybase struct {
	*ds_top

	//
	DS_Enabled bool

	//
	DS_msDS_ClaimPossibleValues string

	//
	DS_msDS_ClaimSharesPossibleValuesWith string
}

func Newds_msds_claimtypepropertybaseEx1(instance *cim.WmiInstance) (newInstance *ds_msds_claimtypepropertybase, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_claimtypepropertybase{
		ds_top: tmp,
	}
	return
}

func Newds_msds_claimtypepropertybaseEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_claimtypepropertybase, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_claimtypepropertybase{
		ds_top: tmp,
	}
	return
}

// SetDS_Enabled sets the value of DS_Enabled for the instance
func (instance *ds_msds_claimtypepropertybase) SetPropertyDS_Enabled(value bool) (err error) {
	return instance.SetProperty("DS_Enabled", (value))
}

// GetDS_Enabled gets the value of DS_Enabled for the instance
func (instance *ds_msds_claimtypepropertybase) GetPropertyDS_Enabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_Enabled")
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

// SetDS_msDS_ClaimPossibleValues sets the value of DS_msDS_ClaimPossibleValues for the instance
func (instance *ds_msds_claimtypepropertybase) SetPropertyDS_msDS_ClaimPossibleValues(value string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimPossibleValues", (value))
}

// GetDS_msDS_ClaimPossibleValues gets the value of DS_msDS_ClaimPossibleValues for the instance
func (instance *ds_msds_claimtypepropertybase) GetPropertyDS_msDS_ClaimPossibleValues() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimPossibleValues")
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

// SetDS_msDS_ClaimSharesPossibleValuesWith sets the value of DS_msDS_ClaimSharesPossibleValuesWith for the instance
func (instance *ds_msds_claimtypepropertybase) SetPropertyDS_msDS_ClaimSharesPossibleValuesWith(value string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimSharesPossibleValuesWith", (value))
}

// GetDS_msDS_ClaimSharesPossibleValuesWith gets the value of DS_msDS_ClaimSharesPossibleValuesWith for the instance
func (instance *ds_msds_claimtypepropertybase) GetPropertyDS_msDS_ClaimSharesPossibleValuesWith() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimSharesPossibleValuesWith")
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
