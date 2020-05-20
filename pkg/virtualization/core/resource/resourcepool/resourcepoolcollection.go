// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package resourcepool

type ResourcePoolCollection []*ResourcePool

func (vms *ResourcePoolCollection) Close() (err error) {
	for _, value := range *vms {
		value.Close()
	}
	return nil
}

func (vms *ResourcePoolCollection) String() string {
	return ""
}
