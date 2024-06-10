// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetBranchCacheDataCache struct
type MSFT_NetBranchCacheDataCache struct {
	*MSFT_NetBranchCachePrimaryCache

	//
	DataCacheExtensions []MSFT_NetBranchCacheDataCacheExtension
}

func NewMSFT_NetBranchCacheDataCacheEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheDataCache, err error) {
	tmp, err := NewMSFT_NetBranchCachePrimaryCacheEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheDataCache{
		MSFT_NetBranchCachePrimaryCache: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheDataCacheEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheDataCache, err error) {
	tmp, err := NewMSFT_NetBranchCachePrimaryCacheEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheDataCache{
		MSFT_NetBranchCachePrimaryCache: tmp,
	}
	return
}

// SetDataCacheExtensions sets the value of DataCacheExtensions for the instance
func (instance *MSFT_NetBranchCacheDataCache) SetPropertyDataCacheExtensions(value []MSFT_NetBranchCacheDataCacheExtension) (err error) {
	return instance.SetProperty("DataCacheExtensions", (value))
}

// GetDataCacheExtensions gets the value of DataCacheExtensions for the instance
func (instance *MSFT_NetBranchCacheDataCache) GetPropertyDataCacheExtensions() (value []MSFT_NetBranchCacheDataCacheExtension, err error) {
	retValue, err := instance.GetProperty("DataCacheExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetBranchCacheDataCacheExtension)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetBranchCacheDataCacheExtension is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetBranchCacheDataCacheExtension(valuetmp))
	}

	return
}
