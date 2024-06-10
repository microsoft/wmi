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

// ads_packageregistration struct
type ads_packageregistration struct {
	*ds_top

	//
	DS_canUpgradeScript []string

	//
	DS_categories []string

	//
	DS_cOMClassID []string

	//
	DS_cOMInterfaceID []string

	//
	DS_cOMProgID []string

	//
	DS_cOMTypelibId []string

	//
	DS_fileExtPriority []string

	//
	DS_iconPath []string

	//
	DS_installUiLevel int32

	//
	DS_lastUpdateSequence string

	//
	DS_localeID []int32

	//
	DS_machineArchitecture []int32

	//
	DS_managedBy string

	//
	DS_msiFileList []string

	//
	DS_msiScript Uint8Array

	//
	DS_msiScriptName string

	//
	DS_msiScriptPath string

	//
	DS_msiScriptSize int32

	//
	DS_packageFlags int32

	//
	DS_packageName string

	//
	DS_packageType int32

	//
	DS_productCode Uint8Array

	//
	DS_setupCommand string

	//
	DS_upgradeProductCode []Uint8Array

	//
	DS_vendor string

	//
	DS_versionNumberHi int32

	//
	DS_versionNumberLo int32
}

func Newads_packageregistrationEx1(instance *cim.WmiInstance) (newInstance *ads_packageregistration, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_packageregistration{
		ds_top: tmp,
	}
	return
}

func Newads_packageregistrationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_packageregistration, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_packageregistration{
		ds_top: tmp,
	}
	return
}

// SetDS_canUpgradeScript sets the value of DS_canUpgradeScript for the instance
func (instance *ads_packageregistration) SetPropertyDS_canUpgradeScript(value []string) (err error) {
	return instance.SetProperty("DS_canUpgradeScript", (value))
}

// GetDS_canUpgradeScript gets the value of DS_canUpgradeScript for the instance
func (instance *ads_packageregistration) GetPropertyDS_canUpgradeScript() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_canUpgradeScript")
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

// SetDS_categories sets the value of DS_categories for the instance
func (instance *ads_packageregistration) SetPropertyDS_categories(value []string) (err error) {
	return instance.SetProperty("DS_categories", (value))
}

// GetDS_categories gets the value of DS_categories for the instance
func (instance *ads_packageregistration) GetPropertyDS_categories() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_categories")
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

// SetDS_cOMClassID sets the value of DS_cOMClassID for the instance
func (instance *ads_packageregistration) SetPropertyDS_cOMClassID(value []string) (err error) {
	return instance.SetProperty("DS_cOMClassID", (value))
}

// GetDS_cOMClassID gets the value of DS_cOMClassID for the instance
func (instance *ads_packageregistration) GetPropertyDS_cOMClassID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMClassID")
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

// SetDS_cOMInterfaceID sets the value of DS_cOMInterfaceID for the instance
func (instance *ads_packageregistration) SetPropertyDS_cOMInterfaceID(value []string) (err error) {
	return instance.SetProperty("DS_cOMInterfaceID", (value))
}

// GetDS_cOMInterfaceID gets the value of DS_cOMInterfaceID for the instance
func (instance *ads_packageregistration) GetPropertyDS_cOMInterfaceID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMInterfaceID")
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

// SetDS_cOMProgID sets the value of DS_cOMProgID for the instance
func (instance *ads_packageregistration) SetPropertyDS_cOMProgID(value []string) (err error) {
	return instance.SetProperty("DS_cOMProgID", (value))
}

// GetDS_cOMProgID gets the value of DS_cOMProgID for the instance
func (instance *ads_packageregistration) GetPropertyDS_cOMProgID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMProgID")
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

// SetDS_cOMTypelibId sets the value of DS_cOMTypelibId for the instance
func (instance *ads_packageregistration) SetPropertyDS_cOMTypelibId(value []string) (err error) {
	return instance.SetProperty("DS_cOMTypelibId", (value))
}

// GetDS_cOMTypelibId gets the value of DS_cOMTypelibId for the instance
func (instance *ads_packageregistration) GetPropertyDS_cOMTypelibId() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_cOMTypelibId")
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

// SetDS_fileExtPriority sets the value of DS_fileExtPriority for the instance
func (instance *ads_packageregistration) SetPropertyDS_fileExtPriority(value []string) (err error) {
	return instance.SetProperty("DS_fileExtPriority", (value))
}

// GetDS_fileExtPriority gets the value of DS_fileExtPriority for the instance
func (instance *ads_packageregistration) GetPropertyDS_fileExtPriority() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_fileExtPriority")
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
func (instance *ads_packageregistration) SetPropertyDS_iconPath(value []string) (err error) {
	return instance.SetProperty("DS_iconPath", (value))
}

// GetDS_iconPath gets the value of DS_iconPath for the instance
func (instance *ads_packageregistration) GetPropertyDS_iconPath() (value []string, err error) {
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

// SetDS_installUiLevel sets the value of DS_installUiLevel for the instance
func (instance *ads_packageregistration) SetPropertyDS_installUiLevel(value int32) (err error) {
	return instance.SetProperty("DS_installUiLevel", (value))
}

// GetDS_installUiLevel gets the value of DS_installUiLevel for the instance
func (instance *ads_packageregistration) GetPropertyDS_installUiLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_installUiLevel")
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

// SetDS_lastUpdateSequence sets the value of DS_lastUpdateSequence for the instance
func (instance *ads_packageregistration) SetPropertyDS_lastUpdateSequence(value string) (err error) {
	return instance.SetProperty("DS_lastUpdateSequence", (value))
}

// GetDS_lastUpdateSequence gets the value of DS_lastUpdateSequence for the instance
func (instance *ads_packageregistration) GetPropertyDS_lastUpdateSequence() (value string, err error) {
	retValue, err := instance.GetProperty("DS_lastUpdateSequence")
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

// SetDS_localeID sets the value of DS_localeID for the instance
func (instance *ads_packageregistration) SetPropertyDS_localeID(value []int32) (err error) {
	return instance.SetProperty("DS_localeID", (value))
}

// GetDS_localeID gets the value of DS_localeID for the instance
func (instance *ads_packageregistration) GetPropertyDS_localeID() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_localeID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_machineArchitecture sets the value of DS_machineArchitecture for the instance
func (instance *ads_packageregistration) SetPropertyDS_machineArchitecture(value []int32) (err error) {
	return instance.SetProperty("DS_machineArchitecture", (value))
}

// GetDS_machineArchitecture gets the value of DS_machineArchitecture for the instance
func (instance *ads_packageregistration) GetPropertyDS_machineArchitecture() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_machineArchitecture")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_packageregistration) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_packageregistration) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_msiFileList sets the value of DS_msiFileList for the instance
func (instance *ads_packageregistration) SetPropertyDS_msiFileList(value []string) (err error) {
	return instance.SetProperty("DS_msiFileList", (value))
}

// GetDS_msiFileList gets the value of DS_msiFileList for the instance
func (instance *ads_packageregistration) GetPropertyDS_msiFileList() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msiFileList")
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

// SetDS_msiScript sets the value of DS_msiScript for the instance
func (instance *ads_packageregistration) SetPropertyDS_msiScript(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msiScript", (value))
}

// GetDS_msiScript gets the value of DS_msiScript for the instance
func (instance *ads_packageregistration) GetPropertyDS_msiScript() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msiScript")
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

// SetDS_msiScriptName sets the value of DS_msiScriptName for the instance
func (instance *ads_packageregistration) SetPropertyDS_msiScriptName(value string) (err error) {
	return instance.SetProperty("DS_msiScriptName", (value))
}

// GetDS_msiScriptName gets the value of DS_msiScriptName for the instance
func (instance *ads_packageregistration) GetPropertyDS_msiScriptName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msiScriptName")
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

// SetDS_msiScriptPath sets the value of DS_msiScriptPath for the instance
func (instance *ads_packageregistration) SetPropertyDS_msiScriptPath(value string) (err error) {
	return instance.SetProperty("DS_msiScriptPath", (value))
}

// GetDS_msiScriptPath gets the value of DS_msiScriptPath for the instance
func (instance *ads_packageregistration) GetPropertyDS_msiScriptPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msiScriptPath")
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

// SetDS_msiScriptSize sets the value of DS_msiScriptSize for the instance
func (instance *ads_packageregistration) SetPropertyDS_msiScriptSize(value int32) (err error) {
	return instance.SetProperty("DS_msiScriptSize", (value))
}

// GetDS_msiScriptSize gets the value of DS_msiScriptSize for the instance
func (instance *ads_packageregistration) GetPropertyDS_msiScriptSize() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msiScriptSize")
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

// SetDS_packageFlags sets the value of DS_packageFlags for the instance
func (instance *ads_packageregistration) SetPropertyDS_packageFlags(value int32) (err error) {
	return instance.SetProperty("DS_packageFlags", (value))
}

// GetDS_packageFlags gets the value of DS_packageFlags for the instance
func (instance *ads_packageregistration) GetPropertyDS_packageFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_packageFlags")
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

// SetDS_packageName sets the value of DS_packageName for the instance
func (instance *ads_packageregistration) SetPropertyDS_packageName(value string) (err error) {
	return instance.SetProperty("DS_packageName", (value))
}

// GetDS_packageName gets the value of DS_packageName for the instance
func (instance *ads_packageregistration) GetPropertyDS_packageName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_packageName")
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

// SetDS_packageType sets the value of DS_packageType for the instance
func (instance *ads_packageregistration) SetPropertyDS_packageType(value int32) (err error) {
	return instance.SetProperty("DS_packageType", (value))
}

// GetDS_packageType gets the value of DS_packageType for the instance
func (instance *ads_packageregistration) GetPropertyDS_packageType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_packageType")
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

// SetDS_productCode sets the value of DS_productCode for the instance
func (instance *ads_packageregistration) SetPropertyDS_productCode(value Uint8Array) (err error) {
	return instance.SetProperty("DS_productCode", (value))
}

// GetDS_productCode gets the value of DS_productCode for the instance
func (instance *ads_packageregistration) GetPropertyDS_productCode() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_productCode")
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

// SetDS_setupCommand sets the value of DS_setupCommand for the instance
func (instance *ads_packageregistration) SetPropertyDS_setupCommand(value string) (err error) {
	return instance.SetProperty("DS_setupCommand", (value))
}

// GetDS_setupCommand gets the value of DS_setupCommand for the instance
func (instance *ads_packageregistration) GetPropertyDS_setupCommand() (value string, err error) {
	retValue, err := instance.GetProperty("DS_setupCommand")
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

// SetDS_upgradeProductCode sets the value of DS_upgradeProductCode for the instance
func (instance *ads_packageregistration) SetPropertyDS_upgradeProductCode(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_upgradeProductCode", (value))
}

// GetDS_upgradeProductCode gets the value of DS_upgradeProductCode for the instance
func (instance *ads_packageregistration) GetPropertyDS_upgradeProductCode() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_upgradeProductCode")
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

// SetDS_vendor sets the value of DS_vendor for the instance
func (instance *ads_packageregistration) SetPropertyDS_vendor(value string) (err error) {
	return instance.SetProperty("DS_vendor", (value))
}

// GetDS_vendor gets the value of DS_vendor for the instance
func (instance *ads_packageregistration) GetPropertyDS_vendor() (value string, err error) {
	retValue, err := instance.GetProperty("DS_vendor")
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

// SetDS_versionNumberHi sets the value of DS_versionNumberHi for the instance
func (instance *ads_packageregistration) SetPropertyDS_versionNumberHi(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberHi", (value))
}

// GetDS_versionNumberHi gets the value of DS_versionNumberHi for the instance
func (instance *ads_packageregistration) GetPropertyDS_versionNumberHi() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumberHi")
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

// SetDS_versionNumberLo sets the value of DS_versionNumberLo for the instance
func (instance *ads_packageregistration) SetPropertyDS_versionNumberLo(value int32) (err error) {
	return instance.SetProperty("DS_versionNumberLo", (value))
}

// GetDS_versionNumberLo gets the value of DS_versionNumberLo for the instance
func (instance *ads_packageregistration) GetPropertyDS_versionNumberLo() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_versionNumberLo")
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
