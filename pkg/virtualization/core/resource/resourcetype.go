// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resource

import (
	"fmt"

	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type ResourceTypeValue struct {
	ResourceType      v2.ResourcePool_ResourceType
	OtherResourceType string
	ResourceSubType   string
}

var ResourceSubTypeName = map[int32]string{
	0:     "Unknown",
	1:     "",
	5:     "Microsoft:Hyper-V:Emulated IDE Controller",
	6:     "Microsoft:Hyper-V:Synthetic SCSI Controller",
	10:    "Microsoft:Hyper-V:Synthetic Ethernet Port",
	17:    "Microsoft:Hyper-V:Synthetic Disk Drive",
	31:    "Microsoft:Hyper-V:Virtual Hard Disk",
	33:    "Microsoft:Hyper-V:Ethernet Connection",
	32770: "Microsoft:Hyper-V:GPU Partition",
}

var OtherResourceTypeName = map[int32]string{
	0:     "",
	1:     "Microsoft:Hyper-V:TPM",
	5:     "",
	6:     "",
	10:    "",
	17:    "",
	31:    "",
	33:    "",
	32770: "",
}

func GetResourceTypeValue(rtype v2.ResourcePool_ResourceType) *ResourceTypeValue {
	return &ResourceTypeValue{
		ResourceType:      rtype,
		OtherResourceType: OtherResourceTypeName[int32(rtype)],
		ResourceSubType:   ResourceSubTypeName[int32(rtype)],
	}
}

func (r *ResourceTypeValue) Equals(in *ResourceTypeValue) bool {
	return in.ResourceType == r.ResourceType &&
		in.ResourceSubType == r.ResourceSubType //&&
	// in.OtherResourceType == r.OtherResourceType
}
func (r *ResourceTypeValue) String() string {
	return fmt.Sprintf("ResourceType[%d] OtherResourceType[%s] ResourceSubType[%s]", r.ResourceType, r.OtherResourceType, r.ResourceSubType)
}
