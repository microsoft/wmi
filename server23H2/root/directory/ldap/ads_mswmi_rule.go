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

// ads_mswmi_rule struct
type ads_mswmi_rule struct {
	*ds_top

	//
	DS_msWMI_Query string

	//
	DS_msWMI_QueryLanguage string

	//
	DS_msWMI_TargetNameSpace string
}

func Newads_mswmi_ruleEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_rule, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_rule{
		ds_top: tmp,
	}
	return
}

func Newads_mswmi_ruleEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_rule, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_rule{
		ds_top: tmp,
	}
	return
}

// SetDS_msWMI_Query sets the value of DS_msWMI_Query for the instance
func (instance *ads_mswmi_rule) SetPropertyDS_msWMI_Query(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Query", (value))
}

// GetDS_msWMI_Query gets the value of DS_msWMI_Query for the instance
func (instance *ads_mswmi_rule) GetPropertyDS_msWMI_Query() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Query")
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

// SetDS_msWMI_QueryLanguage sets the value of DS_msWMI_QueryLanguage for the instance
func (instance *ads_mswmi_rule) SetPropertyDS_msWMI_QueryLanguage(value string) (err error) {
	return instance.SetProperty("DS_msWMI_QueryLanguage", (value))
}

// GetDS_msWMI_QueryLanguage gets the value of DS_msWMI_QueryLanguage for the instance
func (instance *ads_mswmi_rule) GetPropertyDS_msWMI_QueryLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_QueryLanguage")
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

// SetDS_msWMI_TargetNameSpace sets the value of DS_msWMI_TargetNameSpace for the instance
func (instance *ads_mswmi_rule) SetPropertyDS_msWMI_TargetNameSpace(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetNameSpace", (value))
}

// GetDS_msWMI_TargetNameSpace gets the value of DS_msWMI_TargetNameSpace for the instance
func (instance *ads_mswmi_rule) GetPropertyDS_msWMI_TargetNameSpace() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetNameSpace")
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
