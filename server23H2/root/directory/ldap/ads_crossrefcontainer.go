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

// ads_crossrefcontainer struct
type ads_crossrefcontainer struct {
	*ds_top

	//
	DS_msDS_Behavior_Version int32

	//
	DS_msDS_EnabledFeature []string

	//
	DS_msDS_ExecuteScriptPassword Uint8Array

	//
	DS_msDS_SPNSuffixes []string

	//
	DS_msDS_UpdateScript string

	//
	DS_uPNSuffixes []string
}

func Newads_crossrefcontainerEx1(instance *cim.WmiInstance) (newInstance *ads_crossrefcontainer, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_crossrefcontainer{
		ds_top: tmp,
	}
	return
}

func Newads_crossrefcontainerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_crossrefcontainer, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_crossrefcontainer{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_Behavior_Version sets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_msDS_Behavior_Version(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Behavior_Version", (value))
}

// GetDS_msDS_Behavior_Version gets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_msDS_Behavior_Version() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Behavior_Version")
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

// SetDS_msDS_EnabledFeature sets the value of DS_msDS_EnabledFeature for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_msDS_EnabledFeature(value []string) (err error) {
	return instance.SetProperty("DS_msDS_EnabledFeature", (value))
}

// GetDS_msDS_EnabledFeature gets the value of DS_msDS_EnabledFeature for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_msDS_EnabledFeature() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_EnabledFeature")
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

// SetDS_msDS_ExecuteScriptPassword sets the value of DS_msDS_ExecuteScriptPassword for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_msDS_ExecuteScriptPassword(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ExecuteScriptPassword", (value))
}

// GetDS_msDS_ExecuteScriptPassword gets the value of DS_msDS_ExecuteScriptPassword for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_msDS_ExecuteScriptPassword() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ExecuteScriptPassword")
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

// SetDS_msDS_SPNSuffixes sets the value of DS_msDS_SPNSuffixes for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_msDS_SPNSuffixes(value []string) (err error) {
	return instance.SetProperty("DS_msDS_SPNSuffixes", (value))
}

// GetDS_msDS_SPNSuffixes gets the value of DS_msDS_SPNSuffixes for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_msDS_SPNSuffixes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SPNSuffixes")
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

// SetDS_msDS_UpdateScript sets the value of DS_msDS_UpdateScript for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_msDS_UpdateScript(value string) (err error) {
	return instance.SetProperty("DS_msDS_UpdateScript", (value))
}

// GetDS_msDS_UpdateScript gets the value of DS_msDS_UpdateScript for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_msDS_UpdateScript() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UpdateScript")
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

// SetDS_uPNSuffixes sets the value of DS_uPNSuffixes for the instance
func (instance *ads_crossrefcontainer) SetPropertyDS_uPNSuffixes(value []string) (err error) {
	return instance.SetProperty("DS_uPNSuffixes", (value))
}

// GetDS_uPNSuffixes gets the value of DS_uPNSuffixes for the instance
func (instance *ads_crossrefcontainer) GetPropertyDS_uPNSuffixes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_uPNSuffixes")
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
