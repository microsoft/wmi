// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_ClusterSharedVolume struct
type MSCluster_ClusterSharedVolume struct {
	MSCluster_LogicalElement

	//
	BackupState uint32

	//
	FaultState uint32

	//
	VolumeName string

	//
	VolumeOffset uint64
}

// SetBackupState sets the value of BackupState for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyBackupState(value uint32) (err error) {
	return instance.SetProperty("BackupState", value)
}

// GetBackupState gets the value of BackupState for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyBackupState() (value uint32, err error) {
	retValue, err := instance.GetProperty("BackupState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultState sets the value of FaultState for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyFaultState(value uint32) (err error) {
	return instance.SetProperty("FaultState", value)
}

// GetFaultState gets the value of FaultState for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyFaultState() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeName sets the value of VolumeName for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyVolumeName(value string) (err error) {
	return instance.SetProperty("VolumeName", value)
}

// GetVolumeName gets the value of VolumeName for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyVolumeName() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeOffset sets the value of VolumeOffset for the instance
func (instance *MSCluster_ClusterSharedVolume) SetPropertyVolumeOffset(value uint64) (err error) {
	return instance.SetProperty("VolumeOffset", value)
}

// GetVolumeOffset gets the value of VolumeOffset for the instance
func (instance *MSCluster_ClusterSharedVolume) GetPropertyVolumeOffset() (value uint64, err error) {
	retValue, err := instance.GetProperty("VolumeOffset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOnMaintenance() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOnMaintenance")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOffMaintenance() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOffMaintenance")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOnRedirectedAccess() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOnRedirectedAccess")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ClusterSharedVolume) TurnOffRedirectedAccess() (err error) {
	_, err = instance.InvokeMethodWithReturn("TurnOffRedirectedAccess")
	if err != nil {
		return
	}
	return

}

//

// <param name="HostName" type="string "></param>
func (instance *MSCluster_ClusterSharedVolume) MoveToNewHost( /* IN */ HostName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewHost", HostName)
	if err != nil {
		return
	}
	return

}
