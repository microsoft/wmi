// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetBranchCacheContentServerSettingData struct
type MSFT_NetBranchCacheContentServerSettingData struct {
	MSFT_NetBranchCacheSettingData

	//
	ContentServerIsEnabled bool
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
