// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_StorageEvent struct
type MSFT_StorageEvent struct {
	cim.WmiInstance

	// This field allows the Storage Management Provider (SMP) to provide a brief description explaining the indication.
	Description string

	// Denotes the date and time in which the event that triggered this indication occurred.
	EventTime string

	// Denotes the perceived severity of the event from the notifier's point of view.
	///0 - 'Unknown': The severity is unknown or indeterminate.
	///2 - 'Information': The event is for informative purposes.
	///3 - 'Degraded/Warning': Action may be required by the user.
	///4 - 'Minor': Action is needed, but the situation is not serious at this time.
	///5 - 'Major': Immediate action is needed.
	///6 - 'Critical': Immediate action is needed and the scope of the issue is broad.
	///7 - 'Fatal/NonRecoverable': An error has occurred, but it is too late to take remedial action.
	///
	PerceivedSeverity StorageEvent_PerceivedSeverity

	// This field indicates the class of the object which caused the indication. For example: if the object causing the indication is a storage pool, this field should be set to MSFT_StoragePool (not the vendor derived class).
	SourceClassName string

	//
	SourceInstance MSFT_StorageObject

	//
	SourceNamespace string

	// This field indicates the object that caused the indication.
	SourceObjectId string

	//
	SourceServer string

	//
	StorageSubsystemObjectId string
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_StorageEvent) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_StorageEvent) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEventTime sets the value of EventTime for the instance
func (instance *MSFT_StorageEvent) SetPropertyEventTime(value string) (err error) {
	return instance.SetProperty("EventTime", value)
}

// GetEventTime gets the value of EventTime for the instance
func (instance *MSFT_StorageEvent) GetPropertyEventTime() (value string, err error) {
	retValue, err := instance.GetProperty("EventTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPerceivedSeverity sets the value of PerceivedSeverity for the instance
func (instance *MSFT_StorageEvent) SetPropertyPerceivedSeverity(value StorageEvent_PerceivedSeverity) (err error) {
	return instance.SetProperty("PerceivedSeverity", value)
}

// GetPerceivedSeverity gets the value of PerceivedSeverity for the instance
func (instance *MSFT_StorageEvent) GetPropertyPerceivedSeverity() (value StorageEvent_PerceivedSeverity, err error) {
	retValue, err := instance.GetProperty("PerceivedSeverity")
	if err != nil {
		return
	}
	value, ok := retValue.(StorageEvent_PerceivedSeverity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceClassName sets the value of SourceClassName for the instance
func (instance *MSFT_StorageEvent) SetPropertySourceClassName(value string) (err error) {
	return instance.SetProperty("SourceClassName", value)
}

// GetSourceClassName gets the value of SourceClassName for the instance
func (instance *MSFT_StorageEvent) GetPropertySourceClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SourceClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceInstance sets the value of SourceInstance for the instance
func (instance *MSFT_StorageEvent) SetPropertySourceInstance(value MSFT_StorageObject) (err error) {
	return instance.SetProperty("SourceInstance", value)
}

// GetSourceInstance gets the value of SourceInstance for the instance
func (instance *MSFT_StorageEvent) GetPropertySourceInstance() (value MSFT_StorageObject, err error) {
	retValue, err := instance.GetProperty("SourceInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageObject)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceNamespace sets the value of SourceNamespace for the instance
func (instance *MSFT_StorageEvent) SetPropertySourceNamespace(value string) (err error) {
	return instance.SetProperty("SourceNamespace", value)
}

// GetSourceNamespace gets the value of SourceNamespace for the instance
func (instance *MSFT_StorageEvent) GetPropertySourceNamespace() (value string, err error) {
	retValue, err := instance.GetProperty("SourceNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceObjectId sets the value of SourceObjectId for the instance
func (instance *MSFT_StorageEvent) SetPropertySourceObjectId(value string) (err error) {
	return instance.SetProperty("SourceObjectId", value)
}

// GetSourceObjectId gets the value of SourceObjectId for the instance
func (instance *MSFT_StorageEvent) GetPropertySourceObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("SourceObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceServer sets the value of SourceServer for the instance
func (instance *MSFT_StorageEvent) SetPropertySourceServer(value string) (err error) {
	return instance.SetProperty("SourceServer", value)
}

// GetSourceServer gets the value of SourceServer for the instance
func (instance *MSFT_StorageEvent) GetPropertySourceServer() (value string, err error) {
	retValue, err := instance.GetProperty("SourceServer")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStorageSubsystemObjectId sets the value of StorageSubsystemObjectId for the instance
func (instance *MSFT_StorageEvent) SetPropertyStorageSubsystemObjectId(value string) (err error) {
	return instance.SetProperty("StorageSubsystemObjectId", value)
}

// GetStorageSubsystemObjectId gets the value of StorageSubsystemObjectId for the instance
func (instance *MSFT_StorageEvent) GetPropertyStorageSubsystemObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("StorageSubsystemObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
