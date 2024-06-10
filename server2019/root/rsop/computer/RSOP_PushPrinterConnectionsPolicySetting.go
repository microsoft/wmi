// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_PushPrinterConnectionsPolicySetting struct
type RSOP_PushPrinterConnectionsPolicySetting struct {
	*RSOP_PolicySetting

	// Whether the the connection is applied per machine, per user1 = User, 2 = Machine.
	ConnectionType PushPrinterConnectionsPolicySetting_ConnectionType

	// Indicates whether the print connectionhas been deleted.
	deleted bool

	// Printer name,
	printerName string

	// The final result of pushed printer connection. 0 indicate success
	PushResult uint32

	// Short server name
	serverName string

	// Short server name
	uncName string
}

func NewRSOP_PushPrinterConnectionsPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_PushPrinterConnectionsPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_PushPrinterConnectionsPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_PushPrinterConnectionsPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_PushPrinterConnectionsPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_PushPrinterConnectionsPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyConnectionType(value PushPrinterConnectionsPolicySetting_ConnectionType) (err error) {
	return instance.SetProperty("ConnectionType", (value))
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyConnectionType() (value PushPrinterConnectionsPolicySetting_ConnectionType, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
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

	value = PushPrinterConnectionsPolicySetting_ConnectionType(valuetmp)

	return
}

// Setdeleted sets the value of deleted for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertydeleted(value bool) (err error) {
	return instance.SetProperty("deleted", (value))
}

// Getdeleted gets the value of deleted for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertydeleted() (value bool, err error) {
	retValue, err := instance.GetProperty("deleted")
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

// SetprinterName sets the value of printerName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyprinterName(value string) (err error) {
	return instance.SetProperty("printerName", (value))
}

// GetprinterName gets the value of printerName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyprinterName() (value string, err error) {
	retValue, err := instance.GetProperty("printerName")
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

// SetPushResult sets the value of PushResult for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyPushResult(value uint32) (err error) {
	return instance.SetProperty("PushResult", (value))
}

// GetPushResult gets the value of PushResult for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyPushResult() (value uint32, err error) {
	retValue, err := instance.GetProperty("PushResult")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetserverName sets the value of serverName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyserverName(value string) (err error) {
	return instance.SetProperty("serverName", (value))
}

// GetserverName gets the value of serverName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyserverName() (value string, err error) {
	retValue, err := instance.GetProperty("serverName")
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

// SetuncName sets the value of uncName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) SetPropertyuncName(value string) (err error) {
	return instance.SetProperty("uncName", (value))
}

// GetuncName gets the value of uncName for the instance
func (instance *RSOP_PushPrinterConnectionsPolicySetting) GetPropertyuncName() (value string, err error) {
	retValue, err := instance.GetProperty("uncName")
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
