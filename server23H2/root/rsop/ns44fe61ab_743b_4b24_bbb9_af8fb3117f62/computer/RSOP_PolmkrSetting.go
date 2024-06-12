// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS44FE61AB_743B_4B24_BBB9_AF8FB3117F62.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_PolmkrSetting struct
type RSOP_PolmkrSetting struct {
	*RSOP_PolicySetting

	//
	polmkrBaseCseGuid string

	//
	polmkrBaseGpeGuid string

	//
	polmkrBaseGpoDisplayName string

	//
	polmkrBaseGpoGuid string

	//
	polmkrBaseHash string

	//
	polmkrBaseInstanceXml string

	//
	polmkrBaseKeyValues []string
}

func NewRSOP_PolmkrSettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PolmkrSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_PolmkrSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PolmkrSetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PolmkrSetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetpolmkrBaseCseGuid sets the value of polmkrBaseCseGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseCseGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseCseGuid", (value))
}

// GetpolmkrBaseCseGuid gets the value of polmkrBaseCseGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseCseGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseCseGuid")
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

// SetpolmkrBaseGpeGuid sets the value of polmkrBaseGpeGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpeGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpeGuid", (value))
}

// GetpolmkrBaseGpeGuid gets the value of polmkrBaseGpeGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpeGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpeGuid")
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

// SetpolmkrBaseGpoDisplayName sets the value of polmkrBaseGpoDisplayName for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpoDisplayName(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpoDisplayName", (value))
}

// GetpolmkrBaseGpoDisplayName gets the value of polmkrBaseGpoDisplayName for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpoDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpoDisplayName")
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

// SetpolmkrBaseGpoGuid sets the value of polmkrBaseGpoGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpoGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpoGuid", (value))
}

// GetpolmkrBaseGpoGuid gets the value of polmkrBaseGpoGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpoGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpoGuid")
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

// SetpolmkrBaseHash sets the value of polmkrBaseHash for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseHash(value string) (err error) {
	return instance.SetProperty("polmkrBaseHash", (value))
}

// GetpolmkrBaseHash gets the value of polmkrBaseHash for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseHash() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseHash")
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

// SetpolmkrBaseInstanceXml sets the value of polmkrBaseInstanceXml for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseInstanceXml(value string) (err error) {
	return instance.SetProperty("polmkrBaseInstanceXml", (value))
}

// GetpolmkrBaseInstanceXml gets the value of polmkrBaseInstanceXml for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseInstanceXml() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseInstanceXml")
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

// SetpolmkrBaseKeyValues sets the value of polmkrBaseKeyValues for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseKeyValues(value []string) (err error) {
	return instance.SetProperty("polmkrBaseKeyValues", (value))
}

// GetpolmkrBaseKeyValues gets the value of polmkrBaseKeyValues for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseKeyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseKeyValues")
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
