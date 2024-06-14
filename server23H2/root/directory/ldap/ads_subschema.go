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

// ads_subschema struct
type ads_subschema struct {
	*ds_top

	//
	DS_attributeTypes []string

	//
	DS_dITContentRules []string

	//
	DS_extendedAttributeInfo []string

	//
	DS_extendedClassInfo []string

	//
	DS_objectClasses []string
}

func Newads_subschemaEx1(instance *cim.WmiInstance) (newInstance *ads_subschema, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_subschema{
		ds_top: tmp,
	}
	return
}

func Newads_subschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_subschema, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_subschema{
		ds_top: tmp,
	}
	return
}

// SetDS_attributeTypes sets the value of DS_attributeTypes for the instance
func (instance *ads_subschema) SetPropertyDS_attributeTypes(value []string) (err error) {
	return instance.SetProperty("DS_attributeTypes", (value))
}

// GetDS_attributeTypes gets the value of DS_attributeTypes for the instance
func (instance *ads_subschema) GetPropertyDS_attributeTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_attributeTypes")
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

// SetDS_dITContentRules sets the value of DS_dITContentRules for the instance
func (instance *ads_subschema) SetPropertyDS_dITContentRules(value []string) (err error) {
	return instance.SetProperty("DS_dITContentRules", (value))
}

// GetDS_dITContentRules gets the value of DS_dITContentRules for the instance
func (instance *ads_subschema) GetPropertyDS_dITContentRules() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dITContentRules")
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

// SetDS_extendedAttributeInfo sets the value of DS_extendedAttributeInfo for the instance
func (instance *ads_subschema) SetPropertyDS_extendedAttributeInfo(value []string) (err error) {
	return instance.SetProperty("DS_extendedAttributeInfo", (value))
}

// GetDS_extendedAttributeInfo gets the value of DS_extendedAttributeInfo for the instance
func (instance *ads_subschema) GetPropertyDS_extendedAttributeInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_extendedAttributeInfo")
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

// SetDS_extendedClassInfo sets the value of DS_extendedClassInfo for the instance
func (instance *ads_subschema) SetPropertyDS_extendedClassInfo(value []string) (err error) {
	return instance.SetProperty("DS_extendedClassInfo", (value))
}

// GetDS_extendedClassInfo gets the value of DS_extendedClassInfo for the instance
func (instance *ads_subschema) GetPropertyDS_extendedClassInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_extendedClassInfo")
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

// SetDS_objectClasses sets the value of DS_objectClasses for the instance
func (instance *ads_subschema) SetPropertyDS_objectClasses(value []string) (err error) {
	return instance.SetProperty("DS_objectClasses", (value))
}

// GetDS_objectClasses gets the value of DS_objectClasses for the instance
func (instance *ads_subschema) GetPropertyDS_objectClasses() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_objectClasses")
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
