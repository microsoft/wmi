// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// FtpFirewallSupportSection struct
type FtpFirewallSupportSection struct {
	*ConfigurationSection

	//
	HighDataChannelPort int32

	//
	LowDataChannelPort int32
}

func NewFtpFirewallSupportSectionEx1(instance *cim.WmiInstance) (newInstance *FtpFirewallSupportSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpFirewallSupportSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewFtpFirewallSupportSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpFirewallSupportSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpFirewallSupportSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetHighDataChannelPort sets the value of HighDataChannelPort for the instance
func (instance *FtpFirewallSupportSection) SetPropertyHighDataChannelPort(value int32) (err error) {
	return instance.SetProperty("HighDataChannelPort", (value))
}

// GetHighDataChannelPort gets the value of HighDataChannelPort for the instance
func (instance *FtpFirewallSupportSection) GetPropertyHighDataChannelPort() (value int32, err error) {
	retValue, err := instance.GetProperty("HighDataChannelPort")
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

// SetLowDataChannelPort sets the value of LowDataChannelPort for the instance
func (instance *FtpFirewallSupportSection) SetPropertyLowDataChannelPort(value int32) (err error) {
	return instance.SetProperty("LowDataChannelPort", (value))
}

// GetLowDataChannelPort gets the value of LowDataChannelPort for the instance
func (instance *FtpFirewallSupportSection) GetPropertyLowDataChannelPort() (value int32, err error) {
	retValue, err := instance.GetProperty("LowDataChannelPort")
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
