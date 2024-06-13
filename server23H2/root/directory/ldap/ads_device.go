// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_device struct
type ads_device struct {
	*ds_top

	//
	DS_bootFile []string

	//
	DS_bootParameter []string

	//
	DS_ipHostNumber []string

	//
	DS_l string

	//
	DS_macAddress []string

	//
	DS_manager string

	//
	DS_msSFU30Aliases []string

	//
	DS_msSFU30Name string

	//
	DS_msSFU30NisDomain string

	//
	DS_nisMapName string

	//
	DS_o []string

	//
	DS_ou []string

	//
	DS_owner string

	//
	DS_seeAlso []string

	//
	DS_serialNumber []string

	//
	DS_uid []string
}

func Newads_deviceEx1(instance *cim.WmiInstance) (newInstance *ads_device, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_device{
		ds_top: tmp,
	}
	return
}

func Newads_deviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_device, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_device{
		ds_top: tmp,
	}
	return
}

// SetDS_bootFile sets the value of DS_bootFile for the instance
func (instance *ads_device) SetPropertyDS_bootFile(value []string) (err error) {
	return instance.SetProperty("DS_bootFile", (value))
}

// GetDS_bootFile gets the value of DS_bootFile for the instance
func (instance *ads_device) GetPropertyDS_bootFile() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_bootFile")
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

// SetDS_bootParameter sets the value of DS_bootParameter for the instance
func (instance *ads_device) SetPropertyDS_bootParameter(value []string) (err error) {
	return instance.SetProperty("DS_bootParameter", (value))
}

// GetDS_bootParameter gets the value of DS_bootParameter for the instance
func (instance *ads_device) GetPropertyDS_bootParameter() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_bootParameter")
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

// SetDS_ipHostNumber sets the value of DS_ipHostNumber for the instance
func (instance *ads_device) SetPropertyDS_ipHostNumber(value []string) (err error) {
	return instance.SetProperty("DS_ipHostNumber", (value))
}

// GetDS_ipHostNumber gets the value of DS_ipHostNumber for the instance
func (instance *ads_device) GetPropertyDS_ipHostNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ipHostNumber")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_device) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_device) GetPropertyDS_l() (value string, err error) {
	retValue, err := instance.GetProperty("DS_l")
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

// SetDS_macAddress sets the value of DS_macAddress for the instance
func (instance *ads_device) SetPropertyDS_macAddress(value []string) (err error) {
	return instance.SetProperty("DS_macAddress", (value))
}

// GetDS_macAddress gets the value of DS_macAddress for the instance
func (instance *ads_device) GetPropertyDS_macAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_macAddress")
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

// SetDS_manager sets the value of DS_manager for the instance
func (instance *ads_device) SetPropertyDS_manager(value string) (err error) {
	return instance.SetProperty("DS_manager", (value))
}

// GetDS_manager gets the value of DS_manager for the instance
func (instance *ads_device) GetPropertyDS_manager() (value string, err error) {
	retValue, err := instance.GetProperty("DS_manager")
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

// SetDS_msSFU30Aliases sets the value of DS_msSFU30Aliases for the instance
func (instance *ads_device) SetPropertyDS_msSFU30Aliases(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30Aliases", (value))
}

// GetDS_msSFU30Aliases gets the value of DS_msSFU30Aliases for the instance
func (instance *ads_device) GetPropertyDS_msSFU30Aliases() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Aliases")
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

// SetDS_msSFU30Name sets the value of DS_msSFU30Name for the instance
func (instance *ads_device) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_device) GetPropertyDS_msSFU30Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Name")
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

// SetDS_msSFU30NisDomain sets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_device) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_device) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NisDomain")
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

// SetDS_nisMapName sets the value of DS_nisMapName for the instance
func (instance *ads_device) SetPropertyDS_nisMapName(value string) (err error) {
	return instance.SetProperty("DS_nisMapName", (value))
}

// GetDS_nisMapName gets the value of DS_nisMapName for the instance
func (instance *ads_device) GetPropertyDS_nisMapName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nisMapName")
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

// SetDS_o sets the value of DS_o for the instance
func (instance *ads_device) SetPropertyDS_o(value []string) (err error) {
	return instance.SetProperty("DS_o", (value))
}

// GetDS_o gets the value of DS_o for the instance
func (instance *ads_device) GetPropertyDS_o() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_o")
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

// SetDS_ou sets the value of DS_ou for the instance
func (instance *ads_device) SetPropertyDS_ou(value []string) (err error) {
	return instance.SetProperty("DS_ou", (value))
}

// GetDS_ou gets the value of DS_ou for the instance
func (instance *ads_device) GetPropertyDS_ou() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ou")
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

// SetDS_owner sets the value of DS_owner for the instance
func (instance *ads_device) SetPropertyDS_owner(value string) (err error) {
	return instance.SetProperty("DS_owner", (value))
}

// GetDS_owner gets the value of DS_owner for the instance
func (instance *ads_device) GetPropertyDS_owner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_owner")
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

// SetDS_seeAlso sets the value of DS_seeAlso for the instance
func (instance *ads_device) SetPropertyDS_seeAlso(value []string) (err error) {
	return instance.SetProperty("DS_seeAlso", (value))
}

// GetDS_seeAlso gets the value of DS_seeAlso for the instance
func (instance *ads_device) GetPropertyDS_seeAlso() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_seeAlso")
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

// SetDS_serialNumber sets the value of DS_serialNumber for the instance
func (instance *ads_device) SetPropertyDS_serialNumber(value []string) (err error) {
	return instance.SetProperty("DS_serialNumber", (value))
}

// GetDS_serialNumber gets the value of DS_serialNumber for the instance
func (instance *ads_device) GetPropertyDS_serialNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_serialNumber")
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ads_device) SetPropertyDS_uid(value []string) (err error) {
	return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ads_device) GetPropertyDS_uid() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_uid")
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
