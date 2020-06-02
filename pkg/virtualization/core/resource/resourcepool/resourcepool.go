// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcepool

import (
	"fmt"
	//"log"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource"
	"github.com/microsoft/wmi/pkg/virtualization/core/resource/resourceallocation"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type ResourcePool struct {
	*v2.Msvm_ResourcePool
}

// NewResourcePool
func NewResourcePool(instance *wmi.WmiInstance) (*ResourcePool, error) {
	wmivm, err := v2.NewMsvm_ResourcePoolEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ResourcePool{wmivm}, nil
}

// GetResourceAllocationSettingData
func (rp *ResourcePool) GetResourceAllocationSettingData(crole v2.SettingsDefineCapabilities_ValueRole, crange v2.SettingsDefineCapabilities_ValueRange) (setting *resourceallocation.ResourceAllocationSettingData, err error) {
	col, err := rp.GetAllRelated("Cim_AllocationCapabilities")
	if err != nil {
		return
	}
	defer col.Close()
	for _, inst := range col {
		col2, err1 := inst.GetReferences("Cim_SettingsDefineCapabilities")
		if err1 != nil {
			err = err1
			return
		}
		defer col2.Close()
		for _, dc := range col2 {
			tmpInstance, err2 := v2.NewCIM_SettingsDefineCapabilitiesEx1(dc)
			if err2 != nil {
				err = err2
				return
			}

			tmp, err2 := tmpInstance.GetProperty("ValueRange")
			if err2 != nil {
				err = err2
				return
			}
			valueRange := tmp.(int32)
			tmp, err2 = tmpInstance.GetProperty("ValueRole")
			if err != nil {
				err = err2
				return
			}
			valueRole := tmp.(int32)
			if valueRange != int32(crange) || valueRole != int32(crole) {
				continue
			}
			// Found the match
			tmp, err1 := tmpInstance.GetProperty("PartComponent")
			if err1 != nil {
				err = err1
				return
			}
			tmpInstancePath := tmp.(string)
			instnew, err1 := instance.GetWmiInstanceFromPath(rp.GetWmiHost(), string(constant.Virtualization), tmpInstancePath)
			if err1 != nil {
				err = err1
				return
			}
			return resourceallocation.NewResourceAllocationSettingData(instnew)
		}
	}
	return nil, errors.Wrapf(errors.NotFound, "GetResourceAllocationSettingData [%d] [%d]", crole, crange)
}

// GetDefaultResourceAllocationSettingData
func (rp *ResourcePool) GetDefaultResourceAllocationSettingData() (setting *resourceallocation.ResourceAllocationSettingData, err error) {
	return rp.GetResourceAllocationSettingData(v2.SettingsDefineCapabilities_ValueRole_Default, v2.SettingsDefineCapabilities_ValueRange_Point)
}

func GetPrimordialResourcePool(whost *host.WmiHost, rtype v2.ResourcePool_ResourceType) (rp *ResourcePool, err error) {
	rptype := resource.GetResourceTypeValue(rtype)
	col, err := GetResourcePools(whost, true, rptype)
	if err != nil {
		return
	}
	defer col.Close()
	if len(col) == 0 {
		err = errors.Wrapf(errors.NotFound, "ResourcePool [%s]", rptype.String())
		return
	}
	inst, err := col[0].Clone()
	if err != nil {
		return
	}
	return NewResourcePool(inst)
}
func GetResourcePools(whost *host.WmiHost, isPrimordial bool, resType *resource.ResourceTypeValue) (col wmi.WmiInstanceCollection, err error) {
	query := query.NewWmiQuery("Cim_ResourcePool")
	if isPrimordial {
		query.AddFilter("Primordial", "True")
	}
	if resType != nil {
		if int32(resType.ResourceType) > 0 {
			query.AddFilter("ResourceType", fmt.Sprintf("%d", int32(resType.ResourceType)))
		}
		if len(resType.ResourceSubType) > 0 {
			query.AddFilter("ResourceSubType", resType.ResourceSubType)
		}
		if len(resType.OtherResourceType) > 0 {
			query.AddFilter("OtherResourceType", resType.OtherResourceType)
		}
	}
	col, err = instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if len(col) == 0 {
		err = errors.Wrapf(errors.NotFound, "Cim_ResourcePool Primordial[%s] Type[%s]", isPrimordial, resType.String())
	}
	return
}
