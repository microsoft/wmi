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

// ads_mswmi_policytemplate struct
type ads_mswmi_policytemplate struct {
	*ds_top

	//
	DS_msWMI_Author string

	//
	DS_msWMI_ChangeDate string

	//
	DS_msWMI_CreationDate string

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
	DS_msWMI_Name string

	//
	DS_msWMI_NormalizedClass []string

	//
	DS_msWMI_Parm1 string

	//
	DS_msWMI_Parm2 string

	//
	DS_msWMI_Parm3 string

	//
	DS_msWMI_Parm4 string

	//
	DS_msWMI_SourceOrganization string

	//
	DS_msWMI_TargetClass string

	//
	DS_msWMI_TargetNameSpace string

	//
	DS_msWMI_TargetPath string

	//
	DS_msWMI_TargetType string
}

func Newads_mswmi_policytemplateEx1(instance *cim.WmiInstance) (newInstance *ads_mswmi_policytemplate, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_policytemplate{
		ds_top: tmp,
	}
	return
}

func Newads_mswmi_policytemplateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_mswmi_policytemplate, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_mswmi_policytemplate{
		ds_top: tmp,
	}
	return
}

// SetDS_msWMI_Author sets the value of DS_msWMI_Author for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Author(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Author", (value))
}

// GetDS_msWMI_Author gets the value of DS_msWMI_Author for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Author() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Author")
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

// SetDS_msWMI_ChangeDate sets the value of DS_msWMI_ChangeDate for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_ChangeDate(value string) (err error) {
	return instance.SetProperty("DS_msWMI_ChangeDate", (value))
}

// GetDS_msWMI_ChangeDate gets the value of DS_msWMI_ChangeDate for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_ChangeDate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_ChangeDate")
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

// SetDS_msWMI_CreationDate sets the value of DS_msWMI_CreationDate for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_CreationDate(value string) (err error) {
	return instance.SetProperty("DS_msWMI_CreationDate", (value))
}

// GetDS_msWMI_CreationDate gets the value of DS_msWMI_CreationDate for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_CreationDate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_CreationDate")
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

// SetDS_msWMI_ID sets the value of DS_msWMI_ID for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_ID(value string) (err error) {
	return instance.SetProperty("DS_msWMI_ID", (value))
}

// GetDS_msWMI_ID gets the value of DS_msWMI_ID for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_ID() (value string, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_intFlags1(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags1", (value))
}

// GetDS_msWMI_intFlags1 gets the value of DS_msWMI_intFlags1 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_intFlags1() (value int32, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_intFlags2(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags2", (value))
}

// GetDS_msWMI_intFlags2 gets the value of DS_msWMI_intFlags2 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_intFlags2() (value int32, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_intFlags3(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags3", (value))
}

// GetDS_msWMI_intFlags3 gets the value of DS_msWMI_intFlags3 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_intFlags3() (value int32, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_intFlags4(value int32) (err error) {
	return instance.SetProperty("DS_msWMI_intFlags4", (value))
}

// GetDS_msWMI_intFlags4 gets the value of DS_msWMI_intFlags4 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_intFlags4() (value int32, err error) {
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

// SetDS_msWMI_Name sets the value of DS_msWMI_Name for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Name(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Name", (value))
}

// GetDS_msWMI_Name gets the value of DS_msWMI_Name for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_Name")
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

// SetDS_msWMI_NormalizedClass sets the value of DS_msWMI_NormalizedClass for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_NormalizedClass(value []string) (err error) {
	return instance.SetProperty("DS_msWMI_NormalizedClass", (value))
}

// GetDS_msWMI_NormalizedClass gets the value of DS_msWMI_NormalizedClass for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_NormalizedClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_NormalizedClass")
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

// SetDS_msWMI_Parm1 sets the value of DS_msWMI_Parm1 for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Parm1(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm1", (value))
}

// GetDS_msWMI_Parm1 gets the value of DS_msWMI_Parm1 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Parm1() (value string, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Parm2(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm2", (value))
}

// GetDS_msWMI_Parm2 gets the value of DS_msWMI_Parm2 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Parm2() (value string, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Parm3(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm3", (value))
}

// GetDS_msWMI_Parm3 gets the value of DS_msWMI_Parm3 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Parm3() (value string, err error) {
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_Parm4(value string) (err error) {
	return instance.SetProperty("DS_msWMI_Parm4", (value))
}

// GetDS_msWMI_Parm4 gets the value of DS_msWMI_Parm4 for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_Parm4() (value string, err error) {
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

// SetDS_msWMI_SourceOrganization sets the value of DS_msWMI_SourceOrganization for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_SourceOrganization(value string) (err error) {
	return instance.SetProperty("DS_msWMI_SourceOrganization", (value))
}

// GetDS_msWMI_SourceOrganization gets the value of DS_msWMI_SourceOrganization for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_SourceOrganization() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_SourceOrganization")
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

// SetDS_msWMI_TargetClass sets the value of DS_msWMI_TargetClass for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_TargetClass(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetClass", (value))
}

// GetDS_msWMI_TargetClass gets the value of DS_msWMI_TargetClass for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_TargetClass() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetClass")
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
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_TargetNameSpace(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetNameSpace", (value))
}

// GetDS_msWMI_TargetNameSpace gets the value of DS_msWMI_TargetNameSpace for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_TargetNameSpace() (value string, err error) {
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

// SetDS_msWMI_TargetPath sets the value of DS_msWMI_TargetPath for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_TargetPath(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetPath", (value))
}

// GetDS_msWMI_TargetPath gets the value of DS_msWMI_TargetPath for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_TargetPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetPath")
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

// SetDS_msWMI_TargetType sets the value of DS_msWMI_TargetType for the instance
func (instance *ads_mswmi_policytemplate) SetPropertyDS_msWMI_TargetType(value string) (err error) {
	return instance.SetProperty("DS_msWMI_TargetType", (value))
}

// GetDS_msWMI_TargetType gets the value of DS_msWMI_TargetType for the instance
func (instance *ads_mswmi_policytemplate) GetPropertyDS_msWMI_TargetType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msWMI_TargetType")
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
