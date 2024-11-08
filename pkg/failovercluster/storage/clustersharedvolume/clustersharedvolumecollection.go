// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package clustersharedvolume

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type ClusterSharedVolumeCollection []*ClusterSharedVolume

func NewClusterSharedVolumeCollection(instances []*wmi.WmiInstance) (col ClusterSharedVolumeCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewClusterSharedVolume(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}

func (instances *ClusterSharedVolumeCollection) Close() (err error) {
	for _, value := range *instances {
		value.Close()
	}
	return nil
}

func (instances *ClusterSharedVolumeCollection) String() string {
	return ""
}
