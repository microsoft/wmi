// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"log"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/virtualization/core/gpu"
)

// GetPartitionableGpuCollection gets setting data for all host GPU-P partitions
func GetPartitionableGpuCollection(whost *host.WmiHost) (gpupartitionsettingcollection gpu.GpuPartitionSettingCollection, err error) {
	query := query.NewWmiQuery("Msvm_PartitionableGpu")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if err != nil {
		log.Printf("[WMI] Error getting Msvm_PartitionableGpu instances - Error details [%+v]\n", err)
		return
	}
	defer rasdcollection.Close()

	gpupartitionsettingcollection, err = gpu.NewGpuPartitionSettingCollection(rasdcollection)
	if err != nil {
		log.Printf("[WMI] Error getting new GpuPartitionSettingCollection for rasdcollection [%s] - Error details [%+v]\n", rasdcollection, err)
		return
	}

	log.Printf("[WMI] Got GPU-P partition setting collection [%s]", gpupartitionsettingcollection)
	return
}
