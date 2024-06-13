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

// ds_dynamicobject struct
type ds_dynamicobject struct {
	*ds_top

	//
	DS_entryTTL int32

	//
	DS_msDS_Entry_Time_To_Die string
}

func Newds_dynamicobjectEx1(instance *cim.WmiInstance) (newInstance *ds_dynamicobject, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_dynamicobject{
		ds_top: tmp,
	}
	return
}

func Newds_dynamicobjectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_dynamicobject, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_dynamicobject{
		ds_top: tmp,
	}
	return
}

// SetDS_entryTTL sets the value of DS_entryTTL for the instance
func (instance *ds_dynamicobject) SetPropertyDS_entryTTL(value int32) (err error) {
	return instance.SetProperty("DS_entryTTL", (value))
}

// GetDS_entryTTL gets the value of DS_entryTTL for the instance
func (instance *ds_dynamicobject) GetPropertyDS_entryTTL() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_entryTTL")
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

// SetDS_msDS_Entry_Time_To_Die sets the value of DS_msDS_Entry_Time_To_Die for the instance
func (instance *ds_dynamicobject) SetPropertyDS_msDS_Entry_Time_To_Die(value string) (err error) {
	return instance.SetProperty("DS_msDS_Entry_Time_To_Die", (value))
}

// GetDS_msDS_Entry_Time_To_Die gets the value of DS_msDS_Entry_Time_To_Die for the instance
func (instance *ds_dynamicobject) GetPropertyDS_msDS_Entry_Time_To_Die() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Entry_Time_To_Die")
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
