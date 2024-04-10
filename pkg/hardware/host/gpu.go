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

// GetPartitionableGpuCollection gets all host partitionable GPUs
func GetPartitionableGpuCollection(whost *host.WmiHost) (partitionablegpucollection gpu.PartitionableGpuCollection, err error) {
	query := query.NewWmiQuery("Msvm_PartitionableGpu")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if err != nil {
		log.Printf("[WMI] Error getting Msvm_PartitionableGpu instances - Error details [%+v]\n", err)
		return
	}
	defer rasdcollection.Close()

	partitionablegpucollection, err = gpu.NewPartitionableGpuCollection(rasdcollection)
	if err != nil {
		log.Printf("[WMI] Error getting new PartitionableGpuCollection for rasdcollection [%s] - Error details [%+v]\n", rasdcollection, err)
		return
	}

	log.Printf("[WMI] Got partitionable GPU collection [%s]", partitionablegpucollection)
	return
}
