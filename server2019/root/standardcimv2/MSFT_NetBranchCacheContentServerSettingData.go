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

// MSFT_NetBranchCacheContentServerSettingData struct
type MSFT_NetBranchCacheContentServerSettingData struct {
	*MSFT_NetBranchCacheSettingData

	//
	ContentServerIsEnabled bool
}

func NewMSFT_NetBranchCacheContentServerSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheContentServerSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheContentServerSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheContentServerSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheContentServerSettingData, err error) {
	tmp, err := NewMSFT_NetBranchCacheSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheContentServerSettingData{
		MSFT_NetBranchCacheSettingData: tmp,
	}
	return
}

// SetContentServerIsEnabled sets the value of ContentServerIsEnabled for the instance
func (instance *MSFT_NetBranchCacheContentServerSettingData) SetPropertyContentServerIsEnabled(value bool) (err error) {
	return instance.SetProperty("ContentServerIsEnabled", value)
}

// GetContentServerIsEnabled gets the value of ContentServerIsEnabled for the instance
func (instance *MSFT_NetBranchCacheContentServerSettingData) GetPropertyContentServerIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ContentServerIsEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
