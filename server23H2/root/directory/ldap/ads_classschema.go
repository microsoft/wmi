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

// ads_classschema struct
type ads_classschema struct {
	*ds_top

	//
	DS_auxiliaryClass []string

	//
	DS_classDisplayName []string

	//
	DS_defaultHidingValue bool

	//
	DS_defaultObjectCategory string

	//
	DS_defaultSecurityDescriptor string

	//
	DS_governsID string

	//
	DS_isDefunct bool

	//
	DS_lDAPDisplayName string

	//
	DS_mayContain []string

	//
	DS_msDS_IntId int32

	//
	DS_msDs_Schema_Extensions []Uint8Array

	//
	DS_mustContain []string

	//
	DS_objectClassCategory int32

	//
	DS_possSuperiors []string

	//
	DS_rDNAttID string

	//
	DS_schemaFlagsEx int32

	//
	DS_schemaIDGUID Uint8Array

	//
	DS_subClassOf string

	//
	DS_systemAuxiliaryClass []string

	//
	DS_systemMayContain []string

	//
	DS_systemMustContain []string

	//
	DS_systemOnly bool

	//
	DS_systemPossSuperiors []string
}

func Newads_classschemaEx1(instance *cim.WmiInstance) (newInstance *ads_classschema, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_classschema{
		ds_top: tmp,
	}
	return
}

func Newads_classschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_classschema, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_classschema{
		ds_top: tmp,
	}
	return
}

// SetDS_auxiliaryClass sets the value of DS_auxiliaryClass for the instance
func (instance *ads_classschema) SetPropertyDS_auxiliaryClass(value []string) (err error) {
	return instance.SetProperty("DS_auxiliaryClass", (value))
}

// GetDS_auxiliaryClass gets the value of DS_auxiliaryClass for the instance
func (instance *ads_classschema) GetPropertyDS_auxiliaryClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_auxiliaryClass")
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

// SetDS_classDisplayName sets the value of DS_classDisplayName for the instance
func (instance *ads_classschema) SetPropertyDS_classDisplayName(value []string) (err error) {
	return instance.SetProperty("DS_classDisplayName", (value))
}

// GetDS_classDisplayName gets the value of DS_classDisplayName for the instance
func (instance *ads_classschema) GetPropertyDS_classDisplayName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_classDisplayName")
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

// SetDS_defaultHidingValue sets the value of DS_defaultHidingValue for the instance
func (instance *ads_classschema) SetPropertyDS_defaultHidingValue(value bool) (err error) {
	return instance.SetProperty("DS_defaultHidingValue", (value))
}

// GetDS_defaultHidingValue gets the value of DS_defaultHidingValue for the instance
func (instance *ads_classschema) GetPropertyDS_defaultHidingValue() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_defaultHidingValue")
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

// SetDS_defaultObjectCategory sets the value of DS_defaultObjectCategory for the instance
func (instance *ads_classschema) SetPropertyDS_defaultObjectCategory(value string) (err error) {
	return instance.SetProperty("DS_defaultObjectCategory", (value))
}

// GetDS_defaultObjectCategory gets the value of DS_defaultObjectCategory for the instance
func (instance *ads_classschema) GetPropertyDS_defaultObjectCategory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_defaultObjectCategory")
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

// SetDS_defaultSecurityDescriptor sets the value of DS_defaultSecurityDescriptor for the instance
func (instance *ads_classschema) SetPropertyDS_defaultSecurityDescriptor(value string) (err error) {
	return instance.SetProperty("DS_defaultSecurityDescriptor", (value))
}

// GetDS_defaultSecurityDescriptor gets the value of DS_defaultSecurityDescriptor for the instance
func (instance *ads_classschema) GetPropertyDS_defaultSecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("DS_defaultSecurityDescriptor")
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

// SetDS_governsID sets the value of DS_governsID for the instance
func (instance *ads_classschema) SetPropertyDS_governsID(value string) (err error) {
	return instance.SetProperty("DS_governsID", (value))
}

// GetDS_governsID gets the value of DS_governsID for the instance
func (instance *ads_classschema) GetPropertyDS_governsID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_governsID")
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

// SetDS_isDefunct sets the value of DS_isDefunct for the instance
func (instance *ads_classschema) SetPropertyDS_isDefunct(value bool) (err error) {
	return instance.SetProperty("DS_isDefunct", (value))
}

// GetDS_isDefunct gets the value of DS_isDefunct for the instance
func (instance *ads_classschema) GetPropertyDS_isDefunct() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isDefunct")
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

// SetDS_lDAPDisplayName sets the value of DS_lDAPDisplayName for the instance
func (instance *ads_classschema) SetPropertyDS_lDAPDisplayName(value string) (err error) {
	return instance.SetProperty("DS_lDAPDisplayName", (value))
}

// GetDS_lDAPDisplayName gets the value of DS_lDAPDisplayName for the instance
func (instance *ads_classschema) GetPropertyDS_lDAPDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_lDAPDisplayName")
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

// SetDS_mayContain sets the value of DS_mayContain for the instance
func (instance *ads_classschema) SetPropertyDS_mayContain(value []string) (err error) {
	return instance.SetProperty("DS_mayContain", (value))
}

// GetDS_mayContain gets the value of DS_mayContain for the instance
func (instance *ads_classschema) GetPropertyDS_mayContain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mayContain")
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

// SetDS_msDS_IntId sets the value of DS_msDS_IntId for the instance
func (instance *ads_classschema) SetPropertyDS_msDS_IntId(value int32) (err error) {
	return instance.SetProperty("DS_msDS_IntId", (value))
}

// GetDS_msDS_IntId gets the value of DS_msDS_IntId for the instance
func (instance *ads_classschema) GetPropertyDS_msDS_IntId() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IntId")
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

// SetDS_msDs_Schema_Extensions sets the value of DS_msDs_Schema_Extensions for the instance
func (instance *ads_classschema) SetPropertyDS_msDs_Schema_Extensions(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDs_Schema_Extensions", (value))
}

// GetDS_msDs_Schema_Extensions gets the value of DS_msDs_Schema_Extensions for the instance
func (instance *ads_classschema) GetPropertyDS_msDs_Schema_Extensions() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDs_Schema_Extensions")
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

// SetDS_mustContain sets the value of DS_mustContain for the instance
func (instance *ads_classschema) SetPropertyDS_mustContain(value []string) (err error) {
	return instance.SetProperty("DS_mustContain", (value))
}

// GetDS_mustContain gets the value of DS_mustContain for the instance
func (instance *ads_classschema) GetPropertyDS_mustContain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_mustContain")
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

// SetDS_objectClassCategory sets the value of DS_objectClassCategory for the instance
func (instance *ads_classschema) SetPropertyDS_objectClassCategory(value int32) (err error) {
	return instance.SetProperty("DS_objectClassCategory", (value))
}

// GetDS_objectClassCategory gets the value of DS_objectClassCategory for the instance
func (instance *ads_classschema) GetPropertyDS_objectClassCategory() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_objectClassCategory")
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

// SetDS_possSuperiors sets the value of DS_possSuperiors for the instance
func (instance *ads_classschema) SetPropertyDS_possSuperiors(value []string) (err error) {
	return instance.SetProperty("DS_possSuperiors", (value))
}

// GetDS_possSuperiors gets the value of DS_possSuperiors for the instance
func (instance *ads_classschema) GetPropertyDS_possSuperiors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_possSuperiors")
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

// SetDS_rDNAttID sets the value of DS_rDNAttID for the instance
func (instance *ads_classschema) SetPropertyDS_rDNAttID(value string) (err error) {
	return instance.SetProperty("DS_rDNAttID", (value))
}

// GetDS_rDNAttID gets the value of DS_rDNAttID for the instance
func (instance *ads_classschema) GetPropertyDS_rDNAttID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rDNAttID")
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

// SetDS_schemaFlagsEx sets the value of DS_schemaFlagsEx for the instance
func (instance *ads_classschema) SetPropertyDS_schemaFlagsEx(value int32) (err error) {
	return instance.SetProperty("DS_schemaFlagsEx", (value))
}

// GetDS_schemaFlagsEx gets the value of DS_schemaFlagsEx for the instance
func (instance *ads_classschema) GetPropertyDS_schemaFlagsEx() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_schemaFlagsEx")
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

// SetDS_schemaIDGUID sets the value of DS_schemaIDGUID for the instance
func (instance *ads_classschema) SetPropertyDS_schemaIDGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schemaIDGUID", (value))
}

// GetDS_schemaIDGUID gets the value of DS_schemaIDGUID for the instance
func (instance *ads_classschema) GetPropertyDS_schemaIDGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schemaIDGUID")
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

// SetDS_subClassOf sets the value of DS_subClassOf for the instance
func (instance *ads_classschema) SetPropertyDS_subClassOf(value string) (err error) {
	return instance.SetProperty("DS_subClassOf", (value))
}

// GetDS_subClassOf gets the value of DS_subClassOf for the instance
func (instance *ads_classschema) GetPropertyDS_subClassOf() (value string, err error) {
	retValue, err := instance.GetProperty("DS_subClassOf")
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

// SetDS_systemAuxiliaryClass sets the value of DS_systemAuxiliaryClass for the instance
func (instance *ads_classschema) SetPropertyDS_systemAuxiliaryClass(value []string) (err error) {
	return instance.SetProperty("DS_systemAuxiliaryClass", (value))
}

// GetDS_systemAuxiliaryClass gets the value of DS_systemAuxiliaryClass for the instance
func (instance *ads_classschema) GetPropertyDS_systemAuxiliaryClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_systemAuxiliaryClass")
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

// SetDS_systemMayContain sets the value of DS_systemMayContain for the instance
func (instance *ads_classschema) SetPropertyDS_systemMayContain(value []string) (err error) {
	return instance.SetProperty("DS_systemMayContain", (value))
}

// GetDS_systemMayContain gets the value of DS_systemMayContain for the instance
func (instance *ads_classschema) GetPropertyDS_systemMayContain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_systemMayContain")
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

// SetDS_systemMustContain sets the value of DS_systemMustContain for the instance
func (instance *ads_classschema) SetPropertyDS_systemMustContain(value []string) (err error) {
	return instance.SetProperty("DS_systemMustContain", (value))
}

// GetDS_systemMustContain gets the value of DS_systemMustContain for the instance
func (instance *ads_classschema) GetPropertyDS_systemMustContain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_systemMustContain")
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

// SetDS_systemOnly sets the value of DS_systemOnly for the instance
func (instance *ads_classschema) SetPropertyDS_systemOnly(value bool) (err error) {
	return instance.SetProperty("DS_systemOnly", (value))
}

// GetDS_systemOnly gets the value of DS_systemOnly for the instance
func (instance *ads_classschema) GetPropertyDS_systemOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_systemOnly")
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

// SetDS_systemPossSuperiors sets the value of DS_systemPossSuperiors for the instance
func (instance *ads_classschema) SetPropertyDS_systemPossSuperiors(value []string) (err error) {
	return instance.SetProperty("DS_systemPossSuperiors", (value))
}

// GetDS_systemPossSuperiors gets the value of DS_systemPossSuperiors for the instance
func (instance *ads_classschema) GetPropertyDS_systemPossSuperiors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_systemPossSuperiors")
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
