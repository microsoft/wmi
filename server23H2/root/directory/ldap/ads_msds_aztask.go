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

// ads_msds_aztask struct
type ads_msds_aztask struct {
	*ds_top

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzBizRule string

	//
	DS_msDS_AzBizRuleLanguage string

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzLastImportedBizRulePath string

	//
	DS_msDS_AzObjectGuid Uint8Array

	//
	DS_msDS_AzTaskIsRoleDefinition bool

	//
	DS_msDS_OperationsForAzTask []string

	//
	DS_msDS_TasksForAzTask []string
}

func Newads_msds_aztaskEx1(instance *cim.WmiInstance) (newInstance *ads_msds_aztask, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_aztask{
		ds_top: tmp,
	}
	return
}

func Newads_msds_aztaskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_aztask, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_aztask{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationData")
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

// SetDS_msDS_AzBizRule sets the value of DS_msDS_AzBizRule for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzBizRule(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzBizRule", (value))
}

// GetDS_msDS_AzBizRule gets the value of DS_msDS_AzBizRule for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzBizRule() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzBizRule")
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

// SetDS_msDS_AzBizRuleLanguage sets the value of DS_msDS_AzBizRuleLanguage for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzBizRuleLanguage(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzBizRuleLanguage", (value))
}

// GetDS_msDS_AzBizRuleLanguage gets the value of DS_msDS_AzBizRuleLanguage for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzBizRuleLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzBizRuleLanguage")
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

// SetDS_msDS_AzGenericData sets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzGenericData")
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

// SetDS_msDS_AzLastImportedBizRulePath sets the value of DS_msDS_AzLastImportedBizRulePath for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzLastImportedBizRulePath(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzLastImportedBizRulePath", (value))
}

// GetDS_msDS_AzLastImportedBizRulePath gets the value of DS_msDS_AzLastImportedBizRulePath for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzLastImportedBizRulePath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzLastImportedBizRulePath")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzObjectGuid")
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

// SetDS_msDS_AzTaskIsRoleDefinition sets the value of DS_msDS_AzTaskIsRoleDefinition for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_AzTaskIsRoleDefinition(value bool) (err error) {
	return instance.SetProperty("DS_msDS_AzTaskIsRoleDefinition", (value))
}

// GetDS_msDS_AzTaskIsRoleDefinition gets the value of DS_msDS_AzTaskIsRoleDefinition for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_AzTaskIsRoleDefinition() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzTaskIsRoleDefinition")
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

// SetDS_msDS_OperationsForAzTask sets the value of DS_msDS_OperationsForAzTask for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_OperationsForAzTask(value []string) (err error) {
	return instance.SetProperty("DS_msDS_OperationsForAzTask", (value))
}

// GetDS_msDS_OperationsForAzTask gets the value of DS_msDS_OperationsForAzTask for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_OperationsForAzTask() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OperationsForAzTask")
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

// SetDS_msDS_TasksForAzTask sets the value of DS_msDS_TasksForAzTask for the instance
func (instance *ads_msds_aztask) SetPropertyDS_msDS_TasksForAzTask(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TasksForAzTask", (value))
}

// GetDS_msDS_TasksForAzTask gets the value of DS_msDS_TasksForAzTask for the instance
func (instance *ads_msds_aztask) GetPropertyDS_msDS_TasksForAzTask() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TasksForAzTask")
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
