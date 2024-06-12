// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS4C15617D_7EC1_44F0_9878_0B4CCDB5F690.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionSettings struct
type RSOP_IEConnectionSettings struct {
	*RSOP_IEProxySettings

	//
	autoConfigEnable bool

	//
	autoConfigTime int32

	//
	autoConfigURL string

	//
	autoConfigUseLocal bool

	//
	autoDetectConfigSettings bool

	//
	autoProxyURL string

	//
	defaultDialUpConnection string

	//
	deleteExistingConnSettings bool

	//
	dialUpConnections []string

	//
	dialUpState uint8

	//
	importCurrentConnSettings bool
}

func NewRSOP_IEConnectionSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionSettings, err error) {
	tmp, err := NewRSOP_IEProxySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionSettings{
		RSOP_IEProxySettings: tmp,
	}
	return
}

func NewRSOP_IEConnectionSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionSettings, err error) {
	tmp, err := NewRSOP_IEProxySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionSettings{
		RSOP_IEProxySettings: tmp,
	}
	return
}

// SetautoConfigEnable sets the value of autoConfigEnable for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoConfigEnable(value bool) (err error) {
	return instance.SetProperty("autoConfigEnable", (value))
}

// GetautoConfigEnable gets the value of autoConfigEnable for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoConfigEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("autoConfigEnable")
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

// SetautoConfigTime sets the value of autoConfigTime for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoConfigTime(value int32) (err error) {
	return instance.SetProperty("autoConfigTime", (value))
}

// GetautoConfigTime gets the value of autoConfigTime for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoConfigTime() (value int32, err error) {
	retValue, err := instance.GetProperty("autoConfigTime")
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

// SetautoConfigURL sets the value of autoConfigURL for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoConfigURL(value string) (err error) {
	return instance.SetProperty("autoConfigURL", (value))
}

// GetautoConfigURL gets the value of autoConfigURL for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoConfigURL() (value string, err error) {
	retValue, err := instance.GetProperty("autoConfigURL")
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

// SetautoConfigUseLocal sets the value of autoConfigUseLocal for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoConfigUseLocal(value bool) (err error) {
	return instance.SetProperty("autoConfigUseLocal", (value))
}

// GetautoConfigUseLocal gets the value of autoConfigUseLocal for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoConfigUseLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("autoConfigUseLocal")
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

// SetautoDetectConfigSettings sets the value of autoDetectConfigSettings for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoDetectConfigSettings(value bool) (err error) {
	return instance.SetProperty("autoDetectConfigSettings", (value))
}

// GetautoDetectConfigSettings gets the value of autoDetectConfigSettings for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoDetectConfigSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("autoDetectConfigSettings")
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

// SetautoProxyURL sets the value of autoProxyURL for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyautoProxyURL(value string) (err error) {
	return instance.SetProperty("autoProxyURL", (value))
}

// GetautoProxyURL gets the value of autoProxyURL for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyautoProxyURL() (value string, err error) {
	retValue, err := instance.GetProperty("autoProxyURL")
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

// SetdefaultDialUpConnection sets the value of defaultDialUpConnection for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertydefaultDialUpConnection(value string) (err error) {
	return instance.SetProperty("defaultDialUpConnection", (value))
}

// GetdefaultDialUpConnection gets the value of defaultDialUpConnection for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertydefaultDialUpConnection() (value string, err error) {
	retValue, err := instance.GetProperty("defaultDialUpConnection")
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

// SetdeleteExistingConnSettings sets the value of deleteExistingConnSettings for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertydeleteExistingConnSettings(value bool) (err error) {
	return instance.SetProperty("deleteExistingConnSettings", (value))
}

// GetdeleteExistingConnSettings gets the value of deleteExistingConnSettings for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertydeleteExistingConnSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("deleteExistingConnSettings")
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

// SetdialUpConnections sets the value of dialUpConnections for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertydialUpConnections(value []string) (err error) {
	return instance.SetProperty("dialUpConnections", (value))
}

// GetdialUpConnections gets the value of dialUpConnections for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertydialUpConnections() (value []string, err error) {
	retValue, err := instance.GetProperty("dialUpConnections")
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

// SetdialUpState sets the value of dialUpState for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertydialUpState(value uint8) (err error) {
	return instance.SetProperty("dialUpState", (value))
}

// GetdialUpState gets the value of dialUpState for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertydialUpState() (value uint8, err error) {
	retValue, err := instance.GetProperty("dialUpState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetimportCurrentConnSettings sets the value of importCurrentConnSettings for the instance
func (instance *RSOP_IEConnectionSettings) SetPropertyimportCurrentConnSettings(value bool) (err error) {
	return instance.SetProperty("importCurrentConnSettings", (value))
}

// GetimportCurrentConnSettings gets the value of importCurrentConnSettings for the instance
func (instance *RSOP_IEConnectionSettings) GetPropertyimportCurrentConnSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("importCurrentConnSettings")
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
