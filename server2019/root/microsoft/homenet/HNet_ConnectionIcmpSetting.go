// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// HNet_ConnectionIcmpSetting struct
type HNet_ConnectionIcmpSetting struct {
	*cim.WmiInstance

	//
	Connection HNet_Connection

	//
	IcmpSettings HNet_FwIcmpSettings
}

func NewHNet_ConnectionIcmpSettingEx1(instance *cim.WmiInstance) (newInstance *HNet_ConnectionIcmpSetting, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HNet_ConnectionIcmpSetting{
		WmiInstance: tmp,
	}
	return
}

func NewHNet_ConnectionIcmpSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HNet_ConnectionIcmpSetting, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HNet_ConnectionIcmpSetting{
		WmiInstance: tmp,
	}
	return
}

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionIcmpSetting) SetPropertyConnection(value HNet_Connection) (err error) {
	return instance.SetProperty("Connection", value)
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionIcmpSetting) GetPropertyConnection() (value HNet_Connection, err error) {
	retValue, err := instance.GetProperty("Connection")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_Connection)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIcmpSettings sets the value of IcmpSettings for the instance
func (instance *HNet_ConnectionIcmpSetting) SetPropertyIcmpSettings(value HNet_FwIcmpSettings) (err error) {
	return instance.SetProperty("IcmpSettings", value)
}

// GetIcmpSettings gets the value of IcmpSettings for the instance
func (instance *HNet_ConnectionIcmpSetting) GetPropertyIcmpSettings() (value HNet_FwIcmpSettings, err error) {
	retValue, err := instance.GetProperty("IcmpSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(HNet_FwIcmpSettings)
	if !ok {
		// TODO: Set an error
	}
	return
}
