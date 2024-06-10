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

// ads_dsuisettings struct
type ads_dsuisettings struct {
	*ds_top

	//
	DS_dSUIAdminMaximum int32

	//
	DS_dSUIAdminNotification []string

	//
	DS_dSUIShellMaximum int32

	//
	DS_msDS_FilterContainers []string

	//
	DS_msDS_Non_Security_Group_Extra_Classes []string

	//
	DS_msDS_Security_Group_Extra_Classes []string
}

func Newads_dsuisettingsEx1(instance *cim.WmiInstance) (newInstance *ads_dsuisettings, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_dsuisettings{
		ds_top: tmp,
	}
	return
}

func Newads_dsuisettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_dsuisettings, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_dsuisettings{
		ds_top: tmp,
	}
	return
}

// SetDS_dSUIAdminMaximum sets the value of DS_dSUIAdminMaximum for the instance
func (instance *ads_dsuisettings) SetPropertyDS_dSUIAdminMaximum(value int32) (err error) {
	return instance.SetProperty("DS_dSUIAdminMaximum", (value))
}

// GetDS_dSUIAdminMaximum gets the value of DS_dSUIAdminMaximum for the instance
func (instance *ads_dsuisettings) GetPropertyDS_dSUIAdminMaximum() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_dSUIAdminMaximum")
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

// SetDS_dSUIAdminNotification sets the value of DS_dSUIAdminNotification for the instance
func (instance *ads_dsuisettings) SetPropertyDS_dSUIAdminNotification(value []string) (err error) {
	return instance.SetProperty("DS_dSUIAdminNotification", (value))
}

// GetDS_dSUIAdminNotification gets the value of DS_dSUIAdminNotification for the instance
func (instance *ads_dsuisettings) GetPropertyDS_dSUIAdminNotification() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dSUIAdminNotification")
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

// SetDS_dSUIShellMaximum sets the value of DS_dSUIShellMaximum for the instance
func (instance *ads_dsuisettings) SetPropertyDS_dSUIShellMaximum(value int32) (err error) {
	return instance.SetProperty("DS_dSUIShellMaximum", (value))
}

// GetDS_dSUIShellMaximum gets the value of DS_dSUIShellMaximum for the instance
func (instance *ads_dsuisettings) GetPropertyDS_dSUIShellMaximum() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_dSUIShellMaximum")
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

// SetDS_msDS_FilterContainers sets the value of DS_msDS_FilterContainers for the instance
func (instance *ads_dsuisettings) SetPropertyDS_msDS_FilterContainers(value []string) (err error) {
	return instance.SetProperty("DS_msDS_FilterContainers", (value))
}

// GetDS_msDS_FilterContainers gets the value of DS_msDS_FilterContainers for the instance
func (instance *ads_dsuisettings) GetPropertyDS_msDS_FilterContainers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_FilterContainers")
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

// SetDS_msDS_Non_Security_Group_Extra_Classes sets the value of DS_msDS_Non_Security_Group_Extra_Classes for the instance
func (instance *ads_dsuisettings) SetPropertyDS_msDS_Non_Security_Group_Extra_Classes(value []string) (err error) {
	return instance.SetProperty("DS_msDS_Non_Security_Group_Extra_Classes", (value))
}

// GetDS_msDS_Non_Security_Group_Extra_Classes gets the value of DS_msDS_Non_Security_Group_Extra_Classes for the instance
func (instance *ads_dsuisettings) GetPropertyDS_msDS_Non_Security_Group_Extra_Classes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Non_Security_Group_Extra_Classes")
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

// SetDS_msDS_Security_Group_Extra_Classes sets the value of DS_msDS_Security_Group_Extra_Classes for the instance
func (instance *ads_dsuisettings) SetPropertyDS_msDS_Security_Group_Extra_Classes(value []string) (err error) {
	return instance.SetProperty("DS_msDS_Security_Group_Extra_Classes", (value))
}

// GetDS_msDS_Security_Group_Extra_Classes gets the value of DS_msDS_Security_Group_Extra_Classes for the instance
func (instance *ads_dsuisettings) GetPropertyDS_msDS_Security_Group_Extra_Classes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Security_Group_Extra_Classes")
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
