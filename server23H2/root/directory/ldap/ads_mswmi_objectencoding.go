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

// ads_mswmi_objectencoding struct
type ads_mswmi_objectencoding struct {
	*ds_top

	//
	DS_msWMI_Class string

	//
	DS_msWMI_Genus int32

	//
	DS_msWMI_ID string

	//
	DS_msWMI_intFlags1 int32

	//
	DS_msWMI_intFlags2 int32

	//
	DS_msWMI_intFlags3 int32

	//
	DS_msWMI_intFlags4 int32

	//
	DS_msWMI_Parm1 string

	//
	DS_msWMI_Parm2 string

	//
	DS_msWMI_Parm3 string

	//
	DS_msWMI_Parm4 string

	//
	DS_msWMI_ScopeGuid string

	//
	DS_msWMI_TargetObject []Uint8Array
}

func Newads_mswmi_objectencodingEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_objectencoding, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_objectencoding{
		ds_top: tmp,
	}
	return
}

func Newads_mswmi_objectencodingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_objectencoding, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_objectencoding{
		ds_top: tmp,
	}
	return
}

// SetDS_msWMI_Class sets the value of DS_msWMI_Class for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Class(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Class", (value))
}

// GetDS_msWMI_Class gets the value of DS_msWMI_Class for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Class() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Class")
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

// SetDS_msWMI_Genus sets the value of DS_msWMI_Genus for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Genus(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_Genus", (value))
}

// GetDS_msWMI_Genus gets the value of DS_msWMI_Genus for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Genus() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Genus")
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

// SetDS_msWMI_ID sets the value of DS_msWMI_ID for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_ID(value string) (err error) {
	return instance.SetProperty("DS_msWMI_ID", (value))
}

// GetDS_msWMI_ID gets the value of DS_msWMI_ID for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_ID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_ID")
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

// SetDS_msWMI_intFlags1 sets the value of DS_msWMI_intFlags1 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_intFlags1(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags1", (value))
}

// GetDS_msWMI_intFlags1 gets the value of DS_msWMI_intFlags1 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_intFlags1() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_intFlags1")
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

// SetDS_msWMI_intFlags2 sets the value of DS_msWMI_intFlags2 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_intFlags2(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags2", (value))
}

// GetDS_msWMI_intFlags2 gets the value of DS_msWMI_intFlags2 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_intFlags2() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_intFlags2")
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

// SetDS_msWMI_intFlags3 sets the value of DS_msWMI_intFlags3 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_intFlags3(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags3", (value))
}

// GetDS_msWMI_intFlags3 gets the value of DS_msWMI_intFlags3 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_intFlags3() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_intFlags3")
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

// SetDS_msWMI_intFlags4 sets the value of DS_msWMI_intFlags4 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_intFlags4(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags4", (value))
}

// GetDS_msWMI_intFlags4 gets the value of DS_msWMI_intFlags4 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_intFlags4() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_intFlags4")
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

// SetDS_msWMI_Parm1 sets the value of DS_msWMI_Parm1 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Parm1(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm1", (value))
}

// GetDS_msWMI_Parm1 gets the value of DS_msWMI_Parm1 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Parm1() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Parm1")
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

// SetDS_msWMI_Parm2 sets the value of DS_msWMI_Parm2 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Parm2(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm2", (value))
}

// GetDS_msWMI_Parm2 gets the value of DS_msWMI_Parm2 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Parm2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Parm2")
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

// SetDS_msWMI_Parm3 sets the value of DS_msWMI_Parm3 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Parm3(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm3", (value))
}

// GetDS_msWMI_Parm3 gets the value of DS_msWMI_Parm3 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Parm3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Parm3")
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

// SetDS_msWMI_Parm4 sets the value of DS_msWMI_Parm4 for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_Parm4(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm4", (value))
}

// GetDS_msWMI_Parm4 gets the value of DS_msWMI_Parm4 for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_Parm4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Parm4")
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

// SetDS_msWMI_ScopeGuid sets the value of DS_msWMI_ScopeGuid for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_ScopeGuid(value string) (err error) {
	return instance.SetProperty("DS_msWMI_ScopeGuid", (value))
}

// GetDS_msWMI_ScopeGuid gets the value of DS_msWMI_ScopeGuid for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_ScopeGuid() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_ScopeGuid")
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

// SetDS_msWMI_TargetObject sets the value of DS_msWMI_TargetObject for the instance
func (instance *ads_mswmi_objectencoding) SetPropertyDS_msWMI_TargetObject(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msWMI_TargetObject", (value))
}

// GetDS_msWMI_TargetObject gets the value of DS_msWMI_TargetObject for the instance
func (instance *ads_mswmi_objectencoding) GetPropertyDS_msWMI_TargetObject() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetObject")
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
