// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000 struct
type Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000 struct {
	Win32_PerfFormattedData

	//
	WorkflowsAborted uint32

	//
	WorkflowsAbortedPersec uint32

	//
	WorkflowsCompleted uint32

	//
	WorkflowsCompletedPersec uint32

	//
	WorkflowsCreated uint32

	//
	WorkflowsCreatedPersec uint32

	//
	WorkflowsExecuting uint32

	//
	WorkflowsIdlePersec uint32

	//
	WorkflowsInMemory uint32

	//
	WorkflowsLoaded uint32

	//
	WorkflowsLoadedPersec uint32

	//
	WorkflowsPending uint32

	//
	WorkflowsPersisted uint32

	//
	WorkflowsPersistedPersec uint32

	//
	WorkflowsRunnable uint32

	//
	WorkflowsSuspended uint32

	//
	WorkflowsSuspendedPersec uint32

	//
	WorkflowsTerminated uint32

	//
	WorkflowsTerminatedPersec uint32

	//
	WorkflowsUnloaded uint32

	//
	WorkflowsUnloadedPersec uint32
}

// SetWorkflowsAborted sets the value of WorkflowsAborted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsAborted(value uint32) (err error) {
	return instance.SetProperty("WorkflowsAborted", value)
}

// GetWorkflowsAborted gets the value of WorkflowsAborted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsAborted() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsAborted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsAbortedPersec sets the value of WorkflowsAbortedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsAbortedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsAbortedPersec", value)
}

// GetWorkflowsAbortedPersec gets the value of WorkflowsAbortedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsAbortedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsAbortedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsCompleted sets the value of WorkflowsCompleted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsCompleted(value uint32) (err error) {
	return instance.SetProperty("WorkflowsCompleted", value)
}

// GetWorkflowsCompleted gets the value of WorkflowsCompleted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsCompleted() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsCompleted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsCompletedPersec sets the value of WorkflowsCompletedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsCompletedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsCompletedPersec", value)
}

// GetWorkflowsCompletedPersec gets the value of WorkflowsCompletedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsCompletedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsCompletedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsCreated sets the value of WorkflowsCreated for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsCreated(value uint32) (err error) {
	return instance.SetProperty("WorkflowsCreated", value)
}

// GetWorkflowsCreated gets the value of WorkflowsCreated for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsCreated() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsCreated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsCreatedPersec sets the value of WorkflowsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsCreatedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsCreatedPersec", value)
}

// GetWorkflowsCreatedPersec gets the value of WorkflowsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsCreatedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsCreatedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsExecuting sets the value of WorkflowsExecuting for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsExecuting(value uint32) (err error) {
	return instance.SetProperty("WorkflowsExecuting", value)
}

// GetWorkflowsExecuting gets the value of WorkflowsExecuting for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsExecuting() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsExecuting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsIdlePersec sets the value of WorkflowsIdlePersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsIdlePersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsIdlePersec", value)
}

// GetWorkflowsIdlePersec gets the value of WorkflowsIdlePersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsIdlePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsIdlePersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsInMemory sets the value of WorkflowsInMemory for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsInMemory(value uint32) (err error) {
	return instance.SetProperty("WorkflowsInMemory", value)
}

// GetWorkflowsInMemory gets the value of WorkflowsInMemory for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsInMemory() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsInMemory")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsLoaded sets the value of WorkflowsLoaded for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsLoaded(value uint32) (err error) {
	return instance.SetProperty("WorkflowsLoaded", value)
}

// GetWorkflowsLoaded gets the value of WorkflowsLoaded for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsLoaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsLoadedPersec sets the value of WorkflowsLoadedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsLoadedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsLoadedPersec", value)
}

// GetWorkflowsLoadedPersec gets the value of WorkflowsLoadedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsLoadedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsLoadedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsPending sets the value of WorkflowsPending for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsPending(value uint32) (err error) {
	return instance.SetProperty("WorkflowsPending", value)
}

// GetWorkflowsPending gets the value of WorkflowsPending for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsPending")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsPersisted sets the value of WorkflowsPersisted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsPersisted(value uint32) (err error) {
	return instance.SetProperty("WorkflowsPersisted", value)
}

// GetWorkflowsPersisted gets the value of WorkflowsPersisted for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsPersisted() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsPersisted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsPersistedPersec sets the value of WorkflowsPersistedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsPersistedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsPersistedPersec", value)
}

// GetWorkflowsPersistedPersec gets the value of WorkflowsPersistedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsPersistedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsPersistedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsRunnable sets the value of WorkflowsRunnable for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsRunnable(value uint32) (err error) {
	return instance.SetProperty("WorkflowsRunnable", value)
}

// GetWorkflowsRunnable gets the value of WorkflowsRunnable for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsRunnable() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsRunnable")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsSuspended sets the value of WorkflowsSuspended for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsSuspended(value uint32) (err error) {
	return instance.SetProperty("WorkflowsSuspended", value)
}

// GetWorkflowsSuspended gets the value of WorkflowsSuspended for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsSuspended() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsSuspended")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsSuspendedPersec sets the value of WorkflowsSuspendedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsSuspendedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsSuspendedPersec", value)
}

// GetWorkflowsSuspendedPersec gets the value of WorkflowsSuspendedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsSuspendedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsSuspendedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsTerminated sets the value of WorkflowsTerminated for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsTerminated(value uint32) (err error) {
	return instance.SetProperty("WorkflowsTerminated", value)
}

// GetWorkflowsTerminated gets the value of WorkflowsTerminated for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsTerminated() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsTerminated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsTerminatedPersec sets the value of WorkflowsTerminatedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsTerminatedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsTerminatedPersec", value)
}

// GetWorkflowsTerminatedPersec gets the value of WorkflowsTerminatedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsTerminatedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsTerminatedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsUnloaded sets the value of WorkflowsUnloaded for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsUnloaded(value uint32) (err error) {
	return instance.SetProperty("WorkflowsUnloaded", value)
}

// GetWorkflowsUnloaded gets the value of WorkflowsUnloaded for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsUnloaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsUnloaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWorkflowsUnloadedPersec sets the value of WorkflowsUnloadedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) SetPropertyWorkflowsUnloadedPersec(value uint32) (err error) {
	return instance.SetProperty("WorkflowsUnloadedPersec", value)
}

// GetWorkflowsUnloadedPersec gets the value of WorkflowsUnloadedPersec for the instance
func (instance *Win32_PerfFormattedData_WindowsWorkflowFoundation4000_WFSystemWorkflow4000) GetPropertyWorkflowsUnloadedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WorkflowsUnloadedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
