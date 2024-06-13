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

// ads_controlaccessright struct
type ads_controlaccessright struct {
	*ds_top

	//
	DS_appliesTo []string

	//
	DS_localizationDisplayId int32

	//
	DS_rightsGuid string

	//
	DS_validAccesses int32
}

func Newads_controlaccessrightEx1(instance *cim.WmiInstance) (newInstance *ads_controlaccessright, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_controlaccessright{
		ds_top: tmp,
	}
	return
}

func Newads_controlaccessrightEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_controlaccessright, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_controlaccessright{
		ds_top: tmp,
	}
	return
}

// SetDS_appliesTo sets the value of DS_appliesTo for the instance
func (instance *ads_controlaccessright) SetPropertyDS_appliesTo(value []string) (err error) {
	return instance.SetProperty("DS_appliesTo", (value))
}

// GetDS_appliesTo gets the value of DS_appliesTo for the instance
func (instance *ads_controlaccessright) GetPropertyDS_appliesTo() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_appliesTo")
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

// SetDS_localizationDisplayId sets the value of DS_localizationDisplayId for the instance
func (instance *ads_controlaccessright) SetPropertyDS_localizationDisplayId(value int32) (err error) {
	return instance.SetProperty("DS_localizationDisplayId", (value))
}

// GetDS_localizationDisplayId gets the value of DS_localizationDisplayId for the instance
func (instance *ads_controlaccessright) GetPropertyDS_localizationDisplayId() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_localizationDisplayId")
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

// SetDS_rightsGuid sets the value of DS_rightsGuid for the instance
func (instance *ads_controlaccessright) SetPropertyDS_rightsGuid(value string) (err error) {
	return instance.SetProperty("DS_rightsGuid", (value))
}

// GetDS_rightsGuid gets the value of DS_rightsGuid for the instance
func (instance *ads_controlaccessright) GetPropertyDS_rightsGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rightsGuid")
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

// SetDS_validAccesses sets the value of DS_validAccesses for the instance
func (instance *ads_controlaccessright) SetPropertyDS_validAccesses(value int32) (err error) {
	return instance.SetProperty("DS_validAccesses", (value))
}

// GetDS_validAccesses gets the value of DS_validAccesses for the instance
func (instance *ads_controlaccessright) GetPropertyDS_validAccesses() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_validAccesses")
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
