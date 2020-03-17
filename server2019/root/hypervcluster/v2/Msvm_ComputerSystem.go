// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Msvm_ComputerSystem struct
type Msvm_ComputerSystem struct {
	CIM_ComputerSystem

	//
	EnhancedSessionModeState uint16

	//
	FailedOverReplicationType uint16

	//
	HwThreadsPerCoreRealized uint32

	//
	LastApplicationConsistentReplicationTime string

	//
	LastReplicationTime string

	//
	LastReplicationType uint16

	//
	LastSuccessfulBackupTime string

	//
	NumberOfNumaNodes uint16

	//
	OnTimeInMilliseconds uint64

	//
	ProcessID uint32

	//
	ReplicationHealth uint16

	//
	ReplicationMode uint16

	//
	ReplicationState uint16

	//
	TimeOfLastConfigurationChange string
}

// SetEnhancedSessionModeState sets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_ComputerSystem) SetPropertyEnhancedSessionModeState(value uint16) (err error) {
	return instance.SetProperty("EnhancedSessionModeState", value)
}

// GetEnhancedSessionModeState gets the value of EnhancedSessionModeState for the instance
func (instance *Msvm_ComputerSystem) GetPropertyEnhancedSessionModeState() (value uint16, err error) {
	retValue, err := instance.GetProperty("EnhancedSessionModeState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ComputerSystem) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", value)
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_ComputerSystem) GetPropertyFailedOverReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailedOverReplicationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHwThreadsPerCoreRealized sets the value of HwThreadsPerCoreRealized for the instance
func (instance *Msvm_ComputerSystem) SetPropertyHwThreadsPerCoreRealized(value uint32) (err error) {
	return instance.SetProperty("HwThreadsPerCoreRealized", value)
}

// GetHwThreadsPerCoreRealized gets the value of HwThreadsPerCoreRealized for the instance
func (instance *Msvm_ComputerSystem) GetPropertyHwThreadsPerCoreRealized() (value uint32, err error) {
	retValue, err := instance.GetProperty("HwThreadsPerCoreRealized")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastApplicationConsistentReplicationTime sets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastApplicationConsistentReplicationTime(value string) (err error) {
	return instance.SetProperty("LastApplicationConsistentReplicationTime", value)
}

// GetLastApplicationConsistentReplicationTime gets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastApplicationConsistentReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplicationConsistentReplicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastReplicationTime sets the value of LastReplicationTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastReplicationTime(value string) (err error) {
	return instance.SetProperty("LastReplicationTime", value)
}

// GetLastReplicationTime gets the value of LastReplicationTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastReplicationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastReplicationType sets the value of LastReplicationType for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastReplicationType(value uint16) (err error) {
	return instance.SetProperty("LastReplicationType", value)
}

// GetLastReplicationType gets the value of LastReplicationType for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastReplicationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastSuccessfulBackupTime sets the value of LastSuccessfulBackupTime for the instance
func (instance *Msvm_ComputerSystem) SetPropertyLastSuccessfulBackupTime(value string) (err error) {
	return instance.SetProperty("LastSuccessfulBackupTime", value)
}

// GetLastSuccessfulBackupTime gets the value of LastSuccessfulBackupTime for the instance
func (instance *Msvm_ComputerSystem) GetPropertyLastSuccessfulBackupTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastSuccessfulBackupTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfNumaNodes sets the value of NumberOfNumaNodes for the instance
func (instance *Msvm_ComputerSystem) SetPropertyNumberOfNumaNodes(value uint16) (err error) {
	return instance.SetProperty("NumberOfNumaNodes", value)
}

// GetNumberOfNumaNodes gets the value of NumberOfNumaNodes for the instance
func (instance *Msvm_ComputerSystem) GetPropertyNumberOfNumaNodes() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumberOfNumaNodes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOnTimeInMilliseconds sets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_ComputerSystem) SetPropertyOnTimeInMilliseconds(value uint64) (err error) {
	return instance.SetProperty("OnTimeInMilliseconds", value)
}

// GetOnTimeInMilliseconds gets the value of OnTimeInMilliseconds for the instance
func (instance *Msvm_ComputerSystem) GetPropertyOnTimeInMilliseconds() (value uint64, err error) {
	retValue, err := instance.GetProperty("OnTimeInMilliseconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProcessID sets the value of ProcessID for the instance
func (instance *Msvm_ComputerSystem) SetPropertyProcessID(value uint32) (err error) {
	return instance.SetProperty("ProcessID", value)
}

// GetProcessID gets the value of ProcessID for the instance
func (instance *Msvm_ComputerSystem) GetPropertyProcessID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", value)
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", value)
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *Msvm_ComputerSystem) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", value)
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_ComputerSystem) GetPropertyReplicationState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastConfigurationChange sets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_ComputerSystem) SetPropertyTimeOfLastConfigurationChange(value string) (err error) {
	return instance.SetProperty("TimeOfLastConfigurationChange", value)
}

// GetTimeOfLastConfigurationChange gets the value of TimeOfLastConfigurationChange for the instance
func (instance *Msvm_ComputerSystem) GetPropertyTimeOfLastConfigurationChange() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastConfigurationChange")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) RequestReplicationStateChange( /* IN */ RequestedState uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestReplicationStateChange", Action, PercentComplete, Timeout, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Job" type="CIM_ConcreteJob ">May contain a reference to the ConcreteJob created to track the status of the interrupt injection.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) InjectNonMaskableInterrupt( /* OUT */ Job CIM_ConcreteJob,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("InjectNonMaskableInterrupt", Action, PercentComplete, Timeout)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReplicationRelationship" type="string "></param>
// <param name="RequestedState" type="uint16 "></param>
// <param name="TimeoutPeriod" type="string "></param>

// <param name="Job" type="CIM_ConcreteJob "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_ComputerSystem) RequestReplicationStateChangeEx( /* IN */ ReplicationRelationship string,
	/* IN */ RequestedState uint16,
	/* OUT */ Job CIM_ConcreteJob,
	/* OPTIONAL IN */ TimeoutPeriod string,
	/*Custom IN*/ Action cim.UserAction,
	/*Custon IN*/ PercentComplete uint32,
	/*Custon IN*/ Timeout uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodAsync("RequestReplicationStateChangeEx", Action, PercentComplete, Timeout, ReplicationRelationship, RequestedState, TimeoutPeriod)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
