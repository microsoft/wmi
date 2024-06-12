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

// ds_msds_cloudextensions struct
type ds_msds_cloudextensions struct {
	*ds_top

	//
	DS_msDS_cloudExtensionAttribute1 string

	//
	DS_msDS_cloudExtensionAttribute10 string

	//
	DS_msDS_cloudExtensionAttribute11 string

	//
	DS_msDS_cloudExtensionAttribute12 string

	//
	DS_msDS_cloudExtensionAttribute13 string

	//
	DS_msDS_cloudExtensionAttribute14 string

	//
	DS_msDS_cloudExtensionAttribute15 string

	//
	DS_msDS_cloudExtensionAttribute16 string

	//
	DS_msDS_cloudExtensionAttribute17 string

	//
	DS_msDS_cloudExtensionAttribute18 string

	//
	DS_msDS_cloudExtensionAttribute19 string

	//
	DS_msDS_cloudExtensionAttribute2 string

	//
	DS_msDS_cloudExtensionAttribute20 string

	//
	DS_msDS_cloudExtensionAttribute3 string

	//
	DS_msDS_cloudExtensionAttribute4 string

	//
	DS_msDS_cloudExtensionAttribute5 string

	//
	DS_msDS_cloudExtensionAttribute6 string

	//
	DS_msDS_cloudExtensionAttribute7 string

	//
	DS_msDS_cloudExtensionAttribute8 string

	//
	DS_msDS_cloudExtensionAttribute9 string
}

func Newds_msds_cloudextensionsEx1(instance *cim.WmiInstance) (newInstance *ds_msds_cloudextensions, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_msds_cloudextensions{
		ds_top: tmp,
	}
	return
}

func Newds_msds_cloudextensionsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_msds_cloudextensions, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_msds_cloudextensions{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_cloudExtensionAttribute1 sets the value of DS_msDS_cloudExtensionAttribute1 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute1(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute1", (value))
}

// GetDS_msDS_cloudExtensionAttribute1 gets the value of DS_msDS_cloudExtensionAttribute1 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute1() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute1")
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

// SetDS_msDS_cloudExtensionAttribute10 sets the value of DS_msDS_cloudExtensionAttribute10 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute10(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute10", (value))
}

// GetDS_msDS_cloudExtensionAttribute10 gets the value of DS_msDS_cloudExtensionAttribute10 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute10() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute10")
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

// SetDS_msDS_cloudExtensionAttribute11 sets the value of DS_msDS_cloudExtensionAttribute11 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute11(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute11", (value))
}

// GetDS_msDS_cloudExtensionAttribute11 gets the value of DS_msDS_cloudExtensionAttribute11 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute11() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute11")
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

// SetDS_msDS_cloudExtensionAttribute12 sets the value of DS_msDS_cloudExtensionAttribute12 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute12(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute12", (value))
}

// GetDS_msDS_cloudExtensionAttribute12 gets the value of DS_msDS_cloudExtensionAttribute12 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute12() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute12")
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

// SetDS_msDS_cloudExtensionAttribute13 sets the value of DS_msDS_cloudExtensionAttribute13 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute13(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute13", (value))
}

// GetDS_msDS_cloudExtensionAttribute13 gets the value of DS_msDS_cloudExtensionAttribute13 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute13() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute13")
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

// SetDS_msDS_cloudExtensionAttribute14 sets the value of DS_msDS_cloudExtensionAttribute14 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute14(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute14", (value))
}

// GetDS_msDS_cloudExtensionAttribute14 gets the value of DS_msDS_cloudExtensionAttribute14 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute14() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute14")
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

// SetDS_msDS_cloudExtensionAttribute15 sets the value of DS_msDS_cloudExtensionAttribute15 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute15(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute15", (value))
}

// GetDS_msDS_cloudExtensionAttribute15 gets the value of DS_msDS_cloudExtensionAttribute15 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute15() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute15")
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

// SetDS_msDS_cloudExtensionAttribute16 sets the value of DS_msDS_cloudExtensionAttribute16 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute16(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute16", (value))
}

// GetDS_msDS_cloudExtensionAttribute16 gets the value of DS_msDS_cloudExtensionAttribute16 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute16() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute16")
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

// SetDS_msDS_cloudExtensionAttribute17 sets the value of DS_msDS_cloudExtensionAttribute17 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute17(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute17", (value))
}

// GetDS_msDS_cloudExtensionAttribute17 gets the value of DS_msDS_cloudExtensionAttribute17 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute17() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute17")
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

// SetDS_msDS_cloudExtensionAttribute18 sets the value of DS_msDS_cloudExtensionAttribute18 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute18(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute18", (value))
}

// GetDS_msDS_cloudExtensionAttribute18 gets the value of DS_msDS_cloudExtensionAttribute18 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute18() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute18")
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

// SetDS_msDS_cloudExtensionAttribute19 sets the value of DS_msDS_cloudExtensionAttribute19 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute19(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute19", (value))
}

// GetDS_msDS_cloudExtensionAttribute19 gets the value of DS_msDS_cloudExtensionAttribute19 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute19() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute19")
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

// SetDS_msDS_cloudExtensionAttribute2 sets the value of DS_msDS_cloudExtensionAttribute2 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute2(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute2", (value))
}

// GetDS_msDS_cloudExtensionAttribute2 gets the value of DS_msDS_cloudExtensionAttribute2 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute2")
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

// SetDS_msDS_cloudExtensionAttribute20 sets the value of DS_msDS_cloudExtensionAttribute20 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute20(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute20", (value))
}

// GetDS_msDS_cloudExtensionAttribute20 gets the value of DS_msDS_cloudExtensionAttribute20 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute20() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute20")
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

// SetDS_msDS_cloudExtensionAttribute3 sets the value of DS_msDS_cloudExtensionAttribute3 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute3(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute3", (value))
}

// GetDS_msDS_cloudExtensionAttribute3 gets the value of DS_msDS_cloudExtensionAttribute3 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute3")
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

// SetDS_msDS_cloudExtensionAttribute4 sets the value of DS_msDS_cloudExtensionAttribute4 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute4(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute4", (value))
}

// GetDS_msDS_cloudExtensionAttribute4 gets the value of DS_msDS_cloudExtensionAttribute4 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute4")
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

// SetDS_msDS_cloudExtensionAttribute5 sets the value of DS_msDS_cloudExtensionAttribute5 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute5(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute5", (value))
}

// GetDS_msDS_cloudExtensionAttribute5 gets the value of DS_msDS_cloudExtensionAttribute5 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute5() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute5")
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

// SetDS_msDS_cloudExtensionAttribute6 sets the value of DS_msDS_cloudExtensionAttribute6 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute6(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute6", (value))
}

// GetDS_msDS_cloudExtensionAttribute6 gets the value of DS_msDS_cloudExtensionAttribute6 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute6() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute6")
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

// SetDS_msDS_cloudExtensionAttribute7 sets the value of DS_msDS_cloudExtensionAttribute7 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute7(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute7", (value))
}

// GetDS_msDS_cloudExtensionAttribute7 gets the value of DS_msDS_cloudExtensionAttribute7 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute7() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute7")
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

// SetDS_msDS_cloudExtensionAttribute8 sets the value of DS_msDS_cloudExtensionAttribute8 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute8(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute8", (value))
}

// GetDS_msDS_cloudExtensionAttribute8 gets the value of DS_msDS_cloudExtensionAttribute8 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute8() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute8")
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

// SetDS_msDS_cloudExtensionAttribute9 sets the value of DS_msDS_cloudExtensionAttribute9 for the instance
func (instance *ds_msds_cloudextensions) SetPropertyDS_msDS_cloudExtensionAttribute9(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute9", (value))
}

// GetDS_msDS_cloudExtensionAttribute9 gets the value of DS_msDS_cloudExtensionAttribute9 for the instance
func (instance *ds_msds_cloudextensions) GetPropertyDS_msDS_cloudExtensionAttribute9() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute9")
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
