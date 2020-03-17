// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetBranchCacheDataCache struct
type MSFT_NetBranchCacheDataCache struct {
	MSFT_NetBranchCachePrimaryCache

	//
	DataCacheExtensions []MSFT_NetBranchCacheDataCacheExtension
}

// SetDataCacheExtensions sets the value of DataCacheExtensions for the instance
func (instance *MSFT_NetBranchCacheDataCache) SetPropertyDataCacheExtensions(value []MSFT_NetBranchCacheDataCacheExtension) (err error) {
	return instance.SetProperty("DataCacheExtensions", value)
}

// GetDataCacheExtensions gets the value of DataCacheExtensions for the instance
func (instance *MSFT_NetBranchCacheDataCache) GetPropertyDataCacheExtensions() (value []MSFT_NetBranchCacheDataCacheExtension, err error) {
	retValue, err := instance.GetProperty("DataCacheExtensions")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_NetBranchCacheDataCacheExtension)
	if !ok {
		// TODO: Set an error
	}
	return
}
