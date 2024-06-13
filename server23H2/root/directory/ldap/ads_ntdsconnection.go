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

// ads_ntdsconnection struct
type ads_ntdsconnection struct {
	*ds_leaf

	//
	DS_enabledConnection bool

	//
	DS_fromServer string

	//
	DS_generatedConnection bool

	//
	DS_mS_DS_ReplicatesNCReason []DN_With_Binary

	//
	DS_options int32

	//
	DS_schedule Uint8Array

	//
	DS_transportType string
}

func Newads_ntdsconnectionEx1(instance *cim.WmiInstance) (newInstance *ads_ntdsconnection, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsconnection{
		ds_leaf: tmp,
	}
	return
}

func Newads_ntdsconnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntdsconnection, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsconnection{
		ds_leaf: tmp,
	}
	return
}

// SetDS_enabledConnection sets the value of DS_enabledConnection for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_enabledConnection(value bool) (err error) {
	return instance.SetProperty("DS_enabledConnection", (value))
}

// GetDS_enabledConnection gets the value of DS_enabledConnection for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_enabledConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_enabledConnection")
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

// SetDS_fromServer sets the value of DS_fromServer for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_fromServer(value string) (err error) {
	return instance.SetProperty("DS_fromServer", (value))
}

// GetDS_fromServer gets the value of DS_fromServer for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_fromServer() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fromServer")
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

// SetDS_generatedConnection sets the value of DS_generatedConnection for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_generatedConnection(value bool) (err error) {
	return instance.SetProperty("DS_generatedConnection", (value))
}

// GetDS_generatedConnection gets the value of DS_generatedConnection for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_generatedConnection() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_generatedConnection")
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

// SetDS_mS_DS_ReplicatesNCReason sets the value of DS_mS_DS_ReplicatesNCReason for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_mS_DS_ReplicatesNCReason(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_mS_DS_ReplicatesNCReason", (value))
}

// GetDS_mS_DS_ReplicatesNCReason gets the value of DS_mS_DS_ReplicatesNCReason for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_mS_DS_ReplicatesNCReason() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_mS_DS_ReplicatesNCReason")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_options sets the value of DS_options for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_options(value int32) (err error) {
	return instance.SetProperty("DS_options", (value))
}

// GetDS_options gets the value of DS_options for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_options")
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

// SetDS_schedule sets the value of DS_schedule for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_schedule(value Uint8Array) (err error) {
	return instance.SetProperty("DS_schedule", (value))
}

// GetDS_schedule gets the value of DS_schedule for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_schedule() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_schedule")
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

// SetDS_transportType sets the value of DS_transportType for the instance
func (instance *ads_ntdsconnection) SetPropertyDS_transportType(value string) (err error) {
	return instance.SetProperty("DS_transportType", (value))
}

// GetDS_transportType gets the value of DS_transportType for the instance
func (instance *ads_ntdsconnection) GetPropertyDS_transportType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_transportType")
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
