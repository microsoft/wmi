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

// ads_displayspecifier struct
type ads_displayspecifier struct {
	*ds_top

	//
	DS_adminContextMenu []string

	//
	DS_adminMultiselectPropertyPages []string

	//
	DS_adminPropertyPages []string

	//
	DS_attributeDisplayNames []string

	//
	DS_classDisplayName []string

	//
	DS_contextMenu []string

	//
	DS_createDialog string

	//
	DS_createWizardExt []string

	//
	DS_creationWizard string

	//
	DS_extraColumns []string

	//
	DS_iconPath []string

	//
	DS_queryFilter string

	//
	DS_scopeFlags int32

	//
	DS_shellContextMenu []string

	//
	DS_shellPropertyPages []string

	//
	DS_treatAsLeaf bool
}

func Newads_displayspecifierEx1(instance *cim.WmiInstance) (newInstance *ads_displayspecifier, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_displayspecifier{
		ds_top: tmp,
	}
	return
}

func Newads_displayspecifierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_displayspecifier, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_displayspecifier{
		ds_top: tmp,
	}
	return
}

// SetDS_adminContextMenu sets the value of DS_adminContextMenu for the instance
func (instance *ads_displayspecifier) SetPropertyDS_adminContextMenu(value []string) (err error) {
	return instance.SetProperty("DS_adminContextMenu", (value))
}

// GetDS_adminContextMenu gets the value of DS_adminContextMenu for the instance
func (instance *ads_displayspecifier) GetPropertyDS_adminContextMenu() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_adminContextMenu")
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

// SetDS_adminMultiselectPropertyPages sets the value of DS_adminMultiselectPropertyPages for the instance
func (instance *ads_displayspecifier) SetPropertyDS_adminMultiselectPropertyPages(value []string) (err error) {
	return instance.SetProperty("DS_adminMultiselectPropertyPages", (value))
}

// GetDS_adminMultiselectPropertyPages gets the value of DS_adminMultiselectPropertyPages for the instance
func (instance *ads_displayspecifier) GetPropertyDS_adminMultiselectPropertyPages() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_adminMultiselectPropertyPages")
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

// SetDS_adminPropertyPages sets the value of DS_adminPropertyPages for the instance
func (instance *ads_displayspecifier) SetPropertyDS_adminPropertyPages(value []string) (err error) {
	return instance.SetProperty("DS_adminPropertyPages", (value))
}

// GetDS_adminPropertyPages gets the value of DS_adminPropertyPages for the instance
func (instance *ads_displayspecifier) GetPropertyDS_adminPropertyPages() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_adminPropertyPages")
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

// SetDS_attributeDisplayNames sets the value of DS_attributeDisplayNames for the instance
func (instance *ads_displayspecifier) SetPropertyDS_attributeDisplayNames(value []string) (err error) {
	return instance.SetProperty("DS_attributeDisplayNames", (value))
}

// GetDS_attributeDisplayNames gets the value of DS_attributeDisplayNames for the instance
func (instance *ads_displayspecifier) GetPropertyDS_attributeDisplayNames() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_attributeDisplayNames")
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
func (instance *ads_displayspecifier) SetPropertyDS_classDisplayName(value []string) (err error) {
	return instance.SetProperty("DS_classDisplayName", (value))
}

// GetDS_classDisplayName gets the value of DS_classDisplayName for the instance
func (instance *ads_displayspecifier) GetPropertyDS_classDisplayName() (value []string, err error) {
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

// SetDS_contextMenu sets the value of DS_contextMenu for the instance
func (instance *ads_displayspecifier) SetPropertyDS_contextMenu(value []string) (err error) {
	return instance.SetProperty("DS_contextMenu", (value))
}

// GetDS_contextMenu gets the value of DS_contextMenu for the instance
func (instance *ads_displayspecifier) GetPropertyDS_contextMenu() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_contextMenu")
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

// SetDS_createDialog sets the value of DS_createDialog for the instance
func (instance *ads_displayspecifier) SetPropertyDS_createDialog(value string) (err error) {
	return instance.SetProperty("DS_createDialog", (value))
}

// GetDS_createDialog gets the value of DS_createDialog for the instance
func (instance *ads_displayspecifier) GetPropertyDS_createDialog() (value string, err error) {
	retValue, err := instance.GetProperty("DS_createDialog")
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

// SetDS_createWizardExt sets the value of DS_createWizardExt for the instance
func (instance *ads_displayspecifier) SetPropertyDS_createWizardExt(value []string) (err error) {
	return instance.SetProperty("DS_createWizardExt", (value))
}

// GetDS_createWizardExt gets the value of DS_createWizardExt for the instance
func (instance *ads_displayspecifier) GetPropertyDS_createWizardExt() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_createWizardExt")
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

// SetDS_creationWizard sets the value of DS_creationWizard for the instance
func (instance *ads_displayspecifier) SetPropertyDS_creationWizard(value string) (err error) {
	return instance.SetProperty("DS_creationWizard", (value))
}

// GetDS_creationWizard gets the value of DS_creationWizard for the instance
func (instance *ads_displayspecifier) GetPropertyDS_creationWizard() (value string, err error) {
	retValue, err := instance.GetProperty("DS_creationWizard")
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

// SetDS_extraColumns sets the value of DS_extraColumns for the instance
func (instance *ads_displayspecifier) SetPropertyDS_extraColumns(value []string) (err error) {
	return instance.SetProperty("DS_extraColumns", (value))
}

// GetDS_extraColumns gets the value of DS_extraColumns for the instance
func (instance *ads_displayspecifier) GetPropertyDS_extraColumns() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_extraColumns")
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

// SetDS_iconPath sets the value of DS_iconPath for the instance
func (instance *ads_displayspecifier) SetPropertyDS_iconPath(value []string) (err error) {
	return instance.SetProperty("DS_iconPath", (value))
}

// GetDS_iconPath gets the value of DS_iconPath for the instance
func (instance *ads_displayspecifier) GetPropertyDS_iconPath() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_iconPath")
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

// SetDS_queryFilter sets the value of DS_queryFilter for the instance
func (instance *ads_displayspecifier) SetPropertyDS_queryFilter(value string) (err error) {
	return instance.SetProperty("DS_queryFilter", (value))
}

// GetDS_queryFilter gets the value of DS_queryFilter for the instance
func (instance *ads_displayspecifier) GetPropertyDS_queryFilter() (value string, err error) {
	retValue, err := instance.GetProperty("DS_queryFilter")
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

// SetDS_scopeFlags sets the value of DS_scopeFlags for the instance
func (instance *ads_displayspecifier) SetPropertyDS_scopeFlags(value int32) (err error) {
	return instance.SetProperty("DS_scopeFlags", (value))
}

// GetDS_scopeFlags gets the value of DS_scopeFlags for the instance
func (instance *ads_displayspecifier) GetPropertyDS_scopeFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_scopeFlags")
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

// SetDS_shellContextMenu sets the value of DS_shellContextMenu for the instance
func (instance *ads_displayspecifier) SetPropertyDS_shellContextMenu(value []string) (err error) {
	return instance.SetProperty("DS_shellContextMenu", (value))
}

// GetDS_shellContextMenu gets the value of DS_shellContextMenu for the instance
func (instance *ads_displayspecifier) GetPropertyDS_shellContextMenu() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_shellContextMenu")
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

// SetDS_shellPropertyPages sets the value of DS_shellPropertyPages for the instance
func (instance *ads_displayspecifier) SetPropertyDS_shellPropertyPages(value []string) (err error) {
	return instance.SetProperty("DS_shellPropertyPages", (value))
}

// GetDS_shellPropertyPages gets the value of DS_shellPropertyPages for the instance
func (instance *ads_displayspecifier) GetPropertyDS_shellPropertyPages() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_shellPropertyPages")
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

// SetDS_treatAsLeaf sets the value of DS_treatAsLeaf for the instance
func (instance *ads_displayspecifier) SetPropertyDS_treatAsLeaf(value bool) (err error) {
	return instance.SetProperty("DS_treatAsLeaf", (value))
}

// GetDS_treatAsLeaf gets the value of DS_treatAsLeaf for the instance
func (instance *ads_displayspecifier) GetPropertyDS_treatAsLeaf() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_treatAsLeaf")
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
