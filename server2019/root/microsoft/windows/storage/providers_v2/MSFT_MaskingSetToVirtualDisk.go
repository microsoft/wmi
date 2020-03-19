// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MaskingSetToVirtualDisk struct
type MSFT_MaskingSetToVirtualDisk struct {
	*cim.WmiInstance

	//
	MaskingSet MSFT_MaskingSet

	//
	VirtualDisk MSFT_VirtualDisk
}

func NewMSFT_MaskingSetToVirtualDiskEx1(instance *cim.WmiInstance) (newInstance *MSFT_MaskingSetToVirtualDisk, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSetToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_MaskingSetToVirtualDiskEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MaskingSetToVirtualDisk, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSetToVirtualDisk{
		WmiInstance: tmp,
	}
	return
}

// SetMaskingSet sets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToVirtualDisk) SetPropertyMaskingSet(value MSFT_MaskingSet) (err error) {
	return instance.SetProperty("MaskingSet", value)
}

// GetMaskingSet gets the value of MaskingSet for the instance
func (instance *MSFT_MaskingSetToVirtualDisk) GetPropertyMaskingSet() (value MSFT_MaskingSet, err error) {
	retValue, err := instance.GetProperty("MaskingSet")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_MaskingSet)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualDisk sets the value of VirtualDisk for the instance
func (instance *MSFT_MaskingSetToVirtualDisk) SetPropertyVirtualDisk(value MSFT_VirtualDisk) (err error) {
	return instance.SetProperty("VirtualDisk", value)
}

// GetVirtualDisk gets the value of VirtualDisk for the instance
func (instance *MSFT_MaskingSetToVirtualDisk) GetPropertyVirtualDisk() (value MSFT_VirtualDisk, err error) {
	retValue, err := instance.GetProperty("VirtualDisk")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_VirtualDisk)
	if !ok {
		// TODO: Set an error
	}
	return
}
