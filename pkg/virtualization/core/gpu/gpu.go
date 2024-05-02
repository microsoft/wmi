// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package gpu

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
)

// GetPartitionableGpuCollection gets all host partitionable GPUs
func GetPartitionableGpuCollection(whost *host.WmiHost) (partitionablegpucollection PartitionableGpuCollection, err error) {
	query := query.NewWmiQuery("Msvm_PartitionableGpu")
	rasdcollection, err := instance.GetWmiInstancesFromHost(whost, string(constant.Virtualization), query)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get Msvm_PartitionableGpu instances")
	}
	defer func() {
		if err != nil {
			rasdcollection.Close()
		}
	}()

	partitionablegpucollection, err = NewPartitionableGpuCollection(rasdcollection)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to get new PartitionableGpuCollection for rasdcollection [%+v]", rasdcollection)
	}

	return
}
