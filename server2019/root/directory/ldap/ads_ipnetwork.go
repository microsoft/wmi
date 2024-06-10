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

// ads_ipnetwork struct
type ads_ipnetwork struct {
	*ds_top

	//
	DS_ipNetmaskNumber string

	//
	DS_ipNetworkNumber string

	//
	DS_l string

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
	DS_uid []string
}

func Newads_ipnetworkEx1(instance *cim.WmiInstance) (newInstance *ads_ipnetwork, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ipnetwork{
		ds_top: tmp,
	}
	return
}

func Newads_ipnetworkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ipnetwork, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ipnetwork{
		ds_top: tmp,
	}
	return
}

// SetDS_ipNetmaskNumber sets the value of DS_ipNetmaskNumber for the instance
func (instance *ads_ipnetwork) SetPropertyDS_ipNetmaskNumber(value string) (err error) {
	return instance.SetProperty("DS_ipNetmaskNumber", (value))
}

// GetDS_ipNetmaskNumber gets the value of DS_ipNetmaskNumber for the instance
func (instance *ads_ipnetwork) GetPropertyDS_ipNetmaskNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipNetmaskNumber")
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

// SetDS_ipNetworkNumber sets the value of DS_ipNetworkNumber for the instance
func (instance *ads_ipnetwork) SetPropertyDS_ipNetworkNumber(value string) (err error) {
	return instance.SetProperty("DS_ipNetworkNumber", (value))
}

// GetDS_ipNetworkNumber gets the value of DS_ipNetworkNumber for the instance
func (instance *ads_ipnetwork) GetPropertyDS_ipNetworkNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipNetworkNumber")
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

// SetDS_l sets the value of DS_l for the instance
func (instance *ads_ipnetwork) SetPropertyDS_l(value string) (err error) {
	return instance.SetProperty("DS_l", (value))
}

// GetDS_l gets the value of DS_l for the instance
func (instance *ads_ipnetwork) GetPropertyDS_l() (value string, err error) {
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

// SetDS_manager sets the value of DS_manager for the instance
func (instance *ads_ipnetwork) SetPropertyDS_manager(value string) (err error) {
	return instance.SetProperty("DS_manager", (value))
}

// GetDS_manager gets the value of DS_manager for the instance
func (instance *ads_ipnetwork) GetPropertyDS_manager() (value string, err error) {
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
func (instance *ads_ipnetwork) SetPropertyDS_msSFU30Aliases(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30Aliases", (value))
}

// GetDS_msSFU30Aliases gets the value of DS_msSFU30Aliases for the instance
func (instance *ads_ipnetwork) GetPropertyDS_msSFU30Aliases() (value []string, err error) {
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
func (instance *ads_ipnetwork) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_ipnetwork) GetPropertyDS_msSFU30Name() (value string, err error) {
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
func (instance *ads_ipnetwork) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_ipnetwork) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
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
func (instance *ads_ipnetwork) SetPropertyDS_nisMapName(value string) (err error) {
	return instance.SetProperty("DS_nisMapName", (value))
}

// GetDS_nisMapName gets the value of DS_nisMapName for the instance
func (instance *ads_ipnetwork) GetPropertyDS_nisMapName() (value string, err error) {
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ads_ipnetwork) SetPropertyDS_uid(value []string) (err error) {
	return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ads_ipnetwork) GetPropertyDS_uid() (value []string, err error) {
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
