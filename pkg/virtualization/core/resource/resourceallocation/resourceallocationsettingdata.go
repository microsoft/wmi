// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourceallocation

import (
	"github.com/microsoft/wmi/pkg/virtualization/core/resource"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type ResourceAllocationSettingData struct {
	*v2.Msvm_ResourceAllocationSettingData
}

// NewResourceAllocationSettingData
func NewResourceAllocationSettingData(instance *wmi.WmiInstance) (*ResourceAllocationSettingData, error) {
	wmivm, err := v2.NewMsvm_ResourceAllocationSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ResourceAllocationSettingData{wmivm}, nil
}

func (rasd *ResourceAllocationSettingData) GetResourceType() (rtype *resource.ResourceTypeValue, err error) {
	rsub, err := rasd.GetPropertyResourceSubType()
	if err != nil {
		return
	}
	rothersub, err := rasd.GetPropertyOtherResourceType()
	if err != nil {
		return
	}

	retValue, err := rasd.GetProperty("ResourceType")
	if err != nil {
		return
	}
	// v2.ResourceAllocationSettingData_ResourceType
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}

	rtype = &resource.ResourceTypeValue{
		ResourceType:      v2.ResourcePool_ResourceType(value),
		OtherResourceType: rothersub,
		ResourceSubType:   rsub,
	}
	return
}
