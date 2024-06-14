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

// ads_attributeschema struct
type ads_attributeschema struct {
	*ds_top

	//
	DS_attributeID string

	//
	DS_attributeSecurityGUID Uint8Array

	//
	DS_attributeSyntax string

	//
	DS_classDisplayName []string

	//
	DS_extendedCharsAllowed bool

	//
	DS_isDefunct bool

	//
	DS_isEphemeral bool

	//
	DS_isMemberOfPartialAttributeSet bool

	//
	DS_isSingleValued bool

	//
	DS_lDAPDisplayName string

	//
	DS_linkID int32

	//
	DS_mAPIID int32

	//
	DS_msDS_IntId int32

	//
	DS_msDs_Schema_Extensions []Uint8Array

	//
	DS_oMObjectClass Uint8Array

	//
	DS_oMSyntax int32

	//
	DS_rangeLower int32

	//
	DS_rangeUpper int32

	//
	DS_schemaFlagsEx int32

	//
	DS_schemaIDGUID Uint8Array

	//
	DS_searchFlags int32

	//
	DS_systemOnly bool
}

func Newads_attributeschemaEx1(instance *cim.WmiInstance) (newInstance *ads_attributeschema, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_attributeschema{
		ds_top: tmp,
	}
	return
}

func Newads_attributeschemaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_attributeschema, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_attributeschema{
		ds_top: tmp,
	}
	return
}

// SetDS_attributeID sets the value of DS_attributeID for the instance
func (instance *ads_attributeschema) SetPropertyDS_attributeID(value string) (err error) {
	return instance.SetProperty("DS_attributeID", (value))
}

// GetDS_attributeID gets the value of DS_attributeID for the instance
func (instance *ads_attributeschema) GetPropertyDS_attributeID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_attributeID")
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

// SetDS_attributeSecurityGUID sets the value of DS_attributeSecurityGUID for the instance
func (instance *ads_attributeschema) SetPropertyDS_attributeSecurityGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_attributeSecurityGUID", (value))
}

// GetDS_attributeSecurityGUID gets the value of DS_attributeSecurityGUID for the instance
func (instance *ads_attributeschema) GetPropertyDS_attributeSecurityGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_attributeSecurityGUID")
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

// SetDS_attributeSyntax sets the value of DS_attributeSyntax for the instance
func (instance *ads_attributeschema) SetPropertyDS_attributeSyntax(value string) (err error) {
	return instance.SetProperty("DS_attributeSyntax", (value))
}

// GetDS_attributeSyntax gets the value of DS_attributeSyntax for the instance
func (instance *ads_attributeschema) GetPropertyDS_attributeSyntax() (value string, err error) {
	retValue, err := instance.GetProperty("DS_attributeSyntax")
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

// SetDS_classDisplayName sets the value of DS_classDisplayName for the instance
func (instance *ads_attributeschema) SetPropertyDS_classDisplayName(value []string) (err error) {
	return instance.SetProperty("DS_classDisplayName", (value))
}

// GetDS_classDisplayName gets the value of DS_classDisplayName for the instance
func (instance *ads_attributeschema) GetPropertyDS_classDisplayName() (value []string, err error) {
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

// SetDS_extendedCharsAllowed sets the value of DS_extendedCharsAllowed for the instance
func (instance *ads_attributeschema) SetPropertyDS_extendedCharsAllowed(value bool) (err error) {
	return instance.SetProperty("DS_extendedCharsAllowed", (value))
}

// GetDS_extendedCharsAllowed gets the value of DS_extendedCharsAllowed for the instance
func (instance *ads_attributeschema) GetPropertyDS_extendedCharsAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_extendedCharsAllowed")
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

// SetDS_isDefunct sets the value of DS_isDefunct for the instance
func (instance *ads_attributeschema) SetPropertyDS_isDefunct(value bool) (err error) {
	return instance.SetProperty("DS_isDefunct", (value))
}

// GetDS_isDefunct gets the value of DS_isDefunct for the instance
func (instance *ads_attributeschema) GetPropertyDS_isDefunct() (value bool, err error) {
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

// SetDS_isEphemeral sets the value of DS_isEphemeral for the instance
func (instance *ads_attributeschema) SetPropertyDS_isEphemeral(value bool) (err error) {
	return instance.SetProperty("DS_isEphemeral", (value))
}

// GetDS_isEphemeral gets the value of DS_isEphemeral for the instance
func (instance *ads_attributeschema) GetPropertyDS_isEphemeral() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isEphemeral")
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

// SetDS_isMemberOfPartialAttributeSet sets the value of DS_isMemberOfPartialAttributeSet for the instance
func (instance *ads_attributeschema) SetPropertyDS_isMemberOfPartialAttributeSet(value bool) (err error) {
	return instance.SetProperty("DS_isMemberOfPartialAttributeSet", (value))
}

// GetDS_isMemberOfPartialAttributeSet gets the value of DS_isMemberOfPartialAttributeSet for the instance
func (instance *ads_attributeschema) GetPropertyDS_isMemberOfPartialAttributeSet() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isMemberOfPartialAttributeSet")
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

// SetDS_isSingleValued sets the value of DS_isSingleValued for the instance
func (instance *ads_attributeschema) SetPropertyDS_isSingleValued(value bool) (err error) {
	return instance.SetProperty("DS_isSingleValued", (value))
}

// GetDS_isSingleValued gets the value of DS_isSingleValued for the instance
func (instance *ads_attributeschema) GetPropertyDS_isSingleValued() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isSingleValued")
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
func (instance *ads_attributeschema) SetPropertyDS_lDAPDisplayName(value string) (err error) {
	return instance.SetProperty("DS_lDAPDisplayName", (value))
}

// GetDS_lDAPDisplayName gets the value of DS_lDAPDisplayName for the instance
func (instance *ads_attributeschema) GetPropertyDS_lDAPDisplayName() (value string, err error) {
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

// SetDS_linkID sets the value of DS_linkID for the instance
func (instance *ads_attributeschema) SetPropertyDS_linkID(value int32) (err error) {
	return instance.SetProperty("DS_linkID", (value))
}

// GetDS_linkID gets the value of DS_linkID for the instance
func (instance *ads_attributeschema) GetPropertyDS_linkID() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_linkID")
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

// SetDS_mAPIID sets the value of DS_mAPIID for the instance
func (instance *ads_attributeschema) SetPropertyDS_mAPIID(value int32) (err error) {
	return instance.SetProperty("DS_mAPIID", (value))
}

// GetDS_mAPIID gets the value of DS_mAPIID for the instance
func (instance *ads_attributeschema) GetPropertyDS_mAPIID() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mAPIID")
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

// SetDS_msDS_IntId sets the value of DS_msDS_IntId for the instance
func (instance *ads_attributeschema) SetPropertyDS_msDS_IntId(value int32) (err error) {
	return instance.SetProperty("DS_msDS_IntId", (value))
}

// GetDS_msDS_IntId gets the value of DS_msDS_IntId for the instance
func (instance *ads_attributeschema) GetPropertyDS_msDS_IntId() (value int32, err error) {
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
func (instance *ads_attributeschema) SetPropertyDS_msDs_Schema_Extensions(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDs_Schema_Extensions", (value))
}

// GetDS_msDs_Schema_Extensions gets the value of DS_msDs_Schema_Extensions for the instance
func (instance *ads_attributeschema) GetPropertyDS_msDs_Schema_Extensions() (value []Uint8Array, err error) {
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

// SetDS_oMObjectClass sets the value of DS_oMObjectClass for the instance
func (instance *ads_attributeschema) SetPropertyDS_oMObjectClass(value Uint8Array) (err error) {
	return instance.SetProperty("DS_oMObjectClass", (value))
}

// GetDS_oMObjectClass gets the value of DS_oMObjectClass for the instance
func (instance *ads_attributeschema) GetPropertyDS_oMObjectClass() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_oMObjectClass")
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

// SetDS_oMSyntax sets the value of DS_oMSyntax for the instance
func (instance *ads_attributeschema) SetPropertyDS_oMSyntax(value int32) (err error) {
	return instance.SetProperty("DS_oMSyntax", (value))
}

// GetDS_oMSyntax gets the value of DS_oMSyntax for the instance
func (instance *ads_attributeschema) GetPropertyDS_oMSyntax() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_oMSyntax")
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

// SetDS_rangeLower sets the value of DS_rangeLower for the instance
func (instance *ads_attributeschema) SetPropertyDS_rangeLower(value int32) (err error) {
	return instance.SetProperty("DS_rangeLower", (value))
}

// GetDS_rangeLower gets the value of DS_rangeLower for the instance
func (instance *ads_attributeschema) GetPropertyDS_rangeLower() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_rangeLower")
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

// SetDS_rangeUpper sets the value of DS_rangeUpper for the instance
func (instance *ads_attributeschema) SetPropertyDS_rangeUpper(value int32) (err error) {
	return instance.SetProperty("DS_rangeUpper", (value))
}

// GetDS_rangeUpper gets the value of DS_rangeUpper for the instance
func (instance *ads_attributeschema) GetPropertyDS_rangeUpper() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_rangeUpper")
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

// SetDS_schemaFlagsEx sets the value of DS_schemaFlagsEx for the instance
func (instance *ads_attributeschema) SetPropertyDS_schemaFlagsEx(value int32) (err error) {
	return instance.SetProperty("DS_schemaFlagsEx", (value))
}

// GetDS_schemaFlagsEx gets the value of DS_schemaFlagsEx for the instance
func (instance *ads_attributeschema) GetPropertyDS_schemaFlagsEx() (value int32, err error) {
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
func (instance *ads_attributeschema) SetPropertyDS_schemaIDGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schemaIDGUID", (value))
}

// GetDS_schemaIDGUID gets the value of DS_schemaIDGUID for the instance
func (instance *ads_attributeschema) GetPropertyDS_schemaIDGUID() (value Uint8Array, err error) {
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

// SetDS_searchFlags sets the value of DS_searchFlags for the instance
func (instance *ads_attributeschema) SetPropertyDS_searchFlags(value int32) (err error) {
	return instance.SetProperty("DS_searchFlags", (value))
}

// GetDS_searchFlags gets the value of DS_searchFlags for the instance
func (instance *ads_attributeschema) GetPropertyDS_searchFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_searchFlags")
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

// SetDS_systemOnly sets the value of DS_systemOnly for the instance
func (instance *ads_attributeschema) SetPropertyDS_systemOnly(value bool) (err error) {
	return instance.SetProperty("DS_systemOnly", (value))
}

// GetDS_systemOnly gets the value of DS_systemOnly for the instance
func (instance *ads_attributeschema) GetPropertyDS_systemOnly() (value bool, err error) {
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
