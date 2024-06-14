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

// ads_nisnetgroup struct
type ads_nisnetgroup struct {
	*ds_top

	//
	DS_memberNisNetgroup []string

	//
	DS_msSFU30Name string

	//
	DS_msSFU30NetgroupHostAtDomain []string

	//
	DS_msSFU30NetgroupUserAtDomain []string

	//
	DS_msSFU30NisDomain string

	//
	DS_nisMapName string

	//
	DS_nisNetgroupTriple []string
}

func Newads_nisnetgroupEx1(instance *cim.WmiInstance) (newInstance *ads_nisnetgroup, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_nisnetgroup{
		ds_top: tmp,
	}
	return
}

func Newads_nisnetgroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_nisnetgroup, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_nisnetgroup{
		ds_top: tmp,
	}
	return
}

// SetDS_memberNisNetgroup sets the value of DS_memberNisNetgroup for the instance
func (instance *ads_nisnetgroup) SetPropertyDS_memberNisNetgroup(value []string) (err error) {
	return instance.SetProperty("DS_memberNisNetgroup", (value))
}

// GetDS_memberNisNetgroup gets the value of DS_memberNisNetgroup for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_memberNisNetgroup() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_memberNisNetgroup")
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
func (instance *ads_nisnetgroup) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_msSFU30Name() (value string, err error) {
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

// SetDS_msSFU30NetgroupHostAtDomain sets the value of DS_msSFU30NetgroupHostAtDomain for the instance
func (instance *ads_nisnetgroup) SetPropertyDS_msSFU30NetgroupHostAtDomain(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30NetgroupHostAtDomain", (value))
}

// GetDS_msSFU30NetgroupHostAtDomain gets the value of DS_msSFU30NetgroupHostAtDomain for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_msSFU30NetgroupHostAtDomain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NetgroupHostAtDomain")
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

// SetDS_msSFU30NetgroupUserAtDomain sets the value of DS_msSFU30NetgroupUserAtDomain for the instance
func (instance *ads_nisnetgroup) SetPropertyDS_msSFU30NetgroupUserAtDomain(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30NetgroupUserAtDomain", (value))
}

// GetDS_msSFU30NetgroupUserAtDomain gets the value of DS_msSFU30NetgroupUserAtDomain for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_msSFU30NetgroupUserAtDomain() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NetgroupUserAtDomain")
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

// SetDS_msSFU30NisDomain sets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_nisnetgroup) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
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
func (instance *ads_nisnetgroup) SetPropertyDS_nisMapName(value string) (err error) {
	return instance.SetProperty("DS_nisMapName", (value))
}

// GetDS_nisMapName gets the value of DS_nisMapName for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_nisMapName() (value string, err error) {
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

// SetDS_nisNetgroupTriple sets the value of DS_nisNetgroupTriple for the instance
func (instance *ads_nisnetgroup) SetPropertyDS_nisNetgroupTriple(value []string) (err error) {
	return instance.SetProperty("DS_nisNetgroupTriple", (value))
}

// GetDS_nisNetgroupTriple gets the value of DS_nisNetgroupTriple for the instance
func (instance *ads_nisnetgroup) GetPropertyDS_nisNetgroupTriple() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_nisNetgroupTriple")
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
