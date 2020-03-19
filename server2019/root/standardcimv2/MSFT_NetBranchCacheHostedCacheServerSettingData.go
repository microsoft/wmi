// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheHostedCacheServerSettingData struct
type MSFT_NetBranchCacheHostedCacheServerSettingData struct {
	*MSFT_NetBranchCacheSettingData

	//
	ClientAuthenticationMode uint32

	//
	HostedCacheScpRegistrationEnabled bool

	//
	HostedCacheServerIsEnabled bool
}

func NewMSFT_NetBranchCacheHostedCacheServerSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheHostedCacheServerSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheHostedCacheServerSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheHostedCacheServerSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheHostedCacheServerSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheHostedCacheServerSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

// SetClientAuthenticationMode sets the value of ClientAuthenticationMode for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) SetPropertyClientAuthenticationMode(value uint32) (err error) {
	return instance.SetProperty("ClientAuthenticationMode", value)
}

// GetClientAuthenticationMode gets the value of ClientAuthenticationMode for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) GetPropertyClientAuthenticationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClientAuthenticationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheScpRegistrationEnabled sets the value of HostedCacheScpRegistrationEnabled for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) SetPropertyHostedCacheScpRegistrationEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheScpRegistrationEnabled", value)
}

// GetHostedCacheScpRegistrationEnabled gets the value of HostedCacheScpRegistrationEnabled for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) GetPropertyHostedCacheScpRegistrationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheScpRegistrationEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheServerIsEnabled sets the value of HostedCacheServerIsEnabled for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) SetPropertyHostedCacheServerIsEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheServerIsEnabled", value)
}

// GetHostedCacheServerIsEnabled gets the value of HostedCacheServerIsEnabled for the instance
func (instance *MSFT_NetBranchCacheHostedCacheServerSettingData) GetPropertyHostedCacheServerIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheServerIsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
