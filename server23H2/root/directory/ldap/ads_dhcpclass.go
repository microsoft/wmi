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

// ads_dhcpclass struct
type ads_dhcpclass struct {
	*ds_top

	//
	DS_dhcpClasses []Uint8Array

	//
	DS_dhcpFlags int64

	//
	DS_dhcpIdentification string

	//
	DS_dhcpMask []string

	//
	DS_dhcpMaxKey int64

	//
	DS_dhcpObjDescription string

	//
	DS_dhcpObjName string

	//
	DS_dhcpOptions []Uint8Array

	//
	DS_dhcpProperties []Uint8Array

	//
	DS_dhcpRanges []string

	//
	DS_dhcpReservations []string

	//
	DS_dhcpServers []string

	//
	DS_dhcpSites []string

	//
	DS_dhcpState []string

	//
	DS_dhcpSubnets []string

	//
	DS_dhcpType int32

	//
	DS_dhcpUniqueKey int64

	//
	DS_dhcpUpdateTime int64

	//
	DS_mscopeId string

	//
	DS_networkAddress []string

	//
	DS_optionDescription []string

	//
	DS_optionsLocation []string

	//
	DS_superScopeDescription []string

	//
	DS_superScopes []string
}

func Newads_dhcpclassEx1(instance *cim.WmiInstance) (newInstance *ads_dhcpclass, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dhcpclass{
		ds_top: tmp,
	}
	return
}

func Newads_dhcpclassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dhcpclass, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dhcpclass{
		ds_top: tmp,
	}
	return
}

// SetDS_dhcpClasses sets the value of DS_dhcpClasses for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpClasses(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dhcpClasses", (value))
}

// GetDS_dhcpClasses gets the value of DS_dhcpClasses for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpClasses() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dhcpClasses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_dhcpFlags sets the value of DS_dhcpFlags for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpFlags(value int64) (err error) {
	return instance.SetProperty("DS_dhcpFlags", (value))
}

// GetDS_dhcpFlags gets the value of DS_dhcpFlags for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpFlags() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_dhcpFlags")
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

// SetDS_dhcpIdentification sets the value of DS_dhcpIdentification for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpIdentification(value string) (err error) {
	return instance.SetProperty("DS_dhcpIdentification", (value))
}

// GetDS_dhcpIdentification gets the value of DS_dhcpIdentification for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpIdentification() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpIdentification")
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

// SetDS_dhcpMask sets the value of DS_dhcpMask for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpMask(value []string) (err error) {
	return instance.SetProperty("DS_dhcpMask", (value))
}

// GetDS_dhcpMask gets the value of DS_dhcpMask for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpMask() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpMask")
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

// SetDS_dhcpMaxKey sets the value of DS_dhcpMaxKey for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpMaxKey(value int64) (err error) {
	return instance.SetProperty("DS_dhcpMaxKey", (value))
}

// GetDS_dhcpMaxKey gets the value of DS_dhcpMaxKey for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpMaxKey() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_dhcpMaxKey")
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

// SetDS_dhcpObjDescription sets the value of DS_dhcpObjDescription for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpObjDescription(value string) (err error) {
	return instance.SetProperty("DS_dhcpObjDescription", (value))
}

// GetDS_dhcpObjDescription gets the value of DS_dhcpObjDescription for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpObjDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpObjDescription")
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

// SetDS_dhcpObjName sets the value of DS_dhcpObjName for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpObjName(value string) (err error) {
	return instance.SetProperty("DS_dhcpObjName", (value))
}

// GetDS_dhcpObjName gets the value of DS_dhcpObjName for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpObjName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpObjName")
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

// SetDS_dhcpOptions sets the value of DS_dhcpOptions for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpOptions(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dhcpOptions", (value))
}

// GetDS_dhcpOptions gets the value of DS_dhcpOptions for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpOptions() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dhcpOptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_dhcpProperties sets the value of DS_dhcpProperties for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpProperties(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_dhcpProperties", (value))
}

// GetDS_dhcpProperties gets the value of DS_dhcpProperties for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpProperties() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dhcpProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_dhcpRanges sets the value of DS_dhcpRanges for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpRanges(value []string) (err error) {
	return instance.SetProperty("DS_dhcpRanges", (value))
}

// GetDS_dhcpRanges gets the value of DS_dhcpRanges for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpRanges() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpRanges")
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

// SetDS_dhcpReservations sets the value of DS_dhcpReservations for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpReservations(value []string) (err error) {
	return instance.SetProperty("DS_dhcpReservations", (value))
}

// GetDS_dhcpReservations gets the value of DS_dhcpReservations for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpReservations() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpReservations")
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

// SetDS_dhcpServers sets the value of DS_dhcpServers for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpServers(value []string) (err error) {
	return instance.SetProperty("DS_dhcpServers", (value))
}

// GetDS_dhcpServers gets the value of DS_dhcpServers for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpServers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpServers")
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

// SetDS_dhcpSites sets the value of DS_dhcpSites for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpSites(value []string) (err error) {
	return instance.SetProperty("DS_dhcpSites", (value))
}

// GetDS_dhcpSites gets the value of DS_dhcpSites for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpSites() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpSites")
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

// SetDS_dhcpState sets the value of DS_dhcpState for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpState(value []string) (err error) {
	return instance.SetProperty("DS_dhcpState", (value))
}

// GetDS_dhcpState gets the value of DS_dhcpState for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpState() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpState")
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

// SetDS_dhcpSubnets sets the value of DS_dhcpSubnets for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpSubnets(value []string) (err error) {
	return instance.SetProperty("DS_dhcpSubnets", (value))
}

// GetDS_dhcpSubnets gets the value of DS_dhcpSubnets for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpSubnets() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dhcpSubnets")
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

// SetDS_dhcpType sets the value of DS_dhcpType for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpType(value int32) (err error) {
	return instance.SetProperty("DS_dhcpType", (value))
}

// GetDS_dhcpType gets the value of DS_dhcpType for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_dhcpType")
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

// SetDS_dhcpUniqueKey sets the value of DS_dhcpUniqueKey for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpUniqueKey(value int64) (err error) {
	return instance.SetProperty("DS_dhcpUniqueKey", (value))
}

// GetDS_dhcpUniqueKey gets the value of DS_dhcpUniqueKey for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpUniqueKey() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_dhcpUniqueKey")
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

// SetDS_dhcpUpdateTime sets the value of DS_dhcpUpdateTime for the instance
func (instance *ads_dhcpclass) SetPropertyDS_dhcpUpdateTime(value int64) (err error) {
	return instance.SetProperty("DS_dhcpUpdateTime", (value))
}

// GetDS_dhcpUpdateTime gets the value of DS_dhcpUpdateTime for the instance
func (instance *ads_dhcpclass) GetPropertyDS_dhcpUpdateTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_dhcpUpdateTime")
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

// SetDS_mscopeId sets the value of DS_mscopeId for the instance
func (instance *ads_dhcpclass) SetPropertyDS_mscopeId(value string) (err error) {
	return instance.SetProperty("DS_mscopeId", (value))
}

// GetDS_mscopeId gets the value of DS_mscopeId for the instance
func (instance *ads_dhcpclass) GetPropertyDS_mscopeId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mscopeId")
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

// SetDS_networkAddress sets the value of DS_networkAddress for the instance
func (instance *ads_dhcpclass) SetPropertyDS_networkAddress(value []string) (err error) {
	return instance.SetProperty("DS_networkAddress", (value))
}

// GetDS_networkAddress gets the value of DS_networkAddress for the instance
func (instance *ads_dhcpclass) GetPropertyDS_networkAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_networkAddress")
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

// SetDS_optionDescription sets the value of DS_optionDescription for the instance
func (instance *ads_dhcpclass) SetPropertyDS_optionDescription(value []string) (err error) {
	return instance.SetProperty("DS_optionDescription", (value))
}

// GetDS_optionDescription gets the value of DS_optionDescription for the instance
func (instance *ads_dhcpclass) GetPropertyDS_optionDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_optionDescription")
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

// SetDS_optionsLocation sets the value of DS_optionsLocation for the instance
func (instance *ads_dhcpclass) SetPropertyDS_optionsLocation(value []string) (err error) {
	return instance.SetProperty("DS_optionsLocation", (value))
}

// GetDS_optionsLocation gets the value of DS_optionsLocation for the instance
func (instance *ads_dhcpclass) GetPropertyDS_optionsLocation() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_optionsLocation")
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

// SetDS_superScopeDescription sets the value of DS_superScopeDescription for the instance
func (instance *ads_dhcpclass) SetPropertyDS_superScopeDescription(value []string) (err error) {
	return instance.SetProperty("DS_superScopeDescription", (value))
}

// GetDS_superScopeDescription gets the value of DS_superScopeDescription for the instance
func (instance *ads_dhcpclass) GetPropertyDS_superScopeDescription() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_superScopeDescription")
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

// SetDS_superScopes sets the value of DS_superScopes for the instance
func (instance *ads_dhcpclass) SetPropertyDS_superScopes(value []string) (err error) {
	return instance.SetProperty("DS_superScopes", (value))
}

// GetDS_superScopes gets the value of DS_superScopes for the instance
func (instance *ads_dhcpclass) GetPropertyDS_superScopes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_superScopes")
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
