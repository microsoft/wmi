// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_EventClusterCallback struct
type MSCluster_EventClusterCallback struct {
	MSCluster_Event

	//
	ObjectName string

	//
	PercentComplete int32

	//
	PhaseSeverity int32

	//
	PhaseType int32

	//
	SetupPhase int32

	//
	Status int32
}

// SetObjectName sets the value of ObjectName for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyObjectName(value string) (err error) {
	return instance.SetProperty("ObjectName", value)
}

// GetObjectName gets the value of ObjectName for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("ObjectName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPercentComplete(value int32) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPercentComplete() (value int32, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhaseSeverity sets the value of PhaseSeverity for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPhaseSeverity(value int32) (err error) {
	return instance.SetProperty("PhaseSeverity", value)
}

// GetPhaseSeverity gets the value of PhaseSeverity for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPhaseSeverity() (value int32, err error) {
	retValue, err := instance.GetProperty("PhaseSeverity")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPhaseType sets the value of PhaseType for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyPhaseType(value int32) (err error) {
	return instance.SetProperty("PhaseType", value)
}

// GetPhaseType gets the value of PhaseType for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyPhaseType() (value int32, err error) {
	retValue, err := instance.GetProperty("PhaseType")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetupPhase sets the value of SetupPhase for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertySetupPhase(value int32) (err error) {
	return instance.SetProperty("SetupPhase", value)
}

// GetSetupPhase gets the value of SetupPhase for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertySetupPhase() (value int32, err error) {
	retValue, err := instance.GetProperty("SetupPhase")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSCluster_EventClusterCallback) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSCluster_EventClusterCallback) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
