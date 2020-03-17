// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_HealthAction struct
type MSFT_HealthAction struct {
	MSFT_StorageObject

	// The full description of the action.
	Description string

	// This is the error code if the action failed. This should be set to 0 if the action was successful. This should be NULL until the action has completed.
	ErrorCode uint32

	// A free-form string that can contain a vendor-specific error description if the action failed.
	ErrorDescription string

	// The message parameters that are plugged into the default message template for the Health Action type. These can be plugged into a custom message.
	MessageParameters []string

	// The percentage of the action that has completed at the time that this value is requested.
	PercentComplete uint16

	// The short summary description of the action.
	Reason string

	// A identifier of the object that triggered the action.
	ReportingObjectId string

	// A string that uniquely identifies the type of the object that triggered the action.
	ReportingObjectType string

	// A unique identifier of the object that triggered the action.
	ReportingObjectUniqueId string

	// The time that the action was started.
	StartTime string

	// Current high level state of the action.
	State uint16

	// A free-form string that provides implementation-specific status of the action.
	Status string

	// The type of Health Action.
	Type string
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_HealthAction) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_HealthAction) GetPropertyDescription() (value string, err error) {
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

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *MSFT_HealthAction) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *MSFT_HealthAction) GetPropertyErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorDescription sets the value of ErrorDescription for the instance
func (instance *MSFT_HealthAction) SetPropertyErrorDescription(value string) (err error) {
	return instance.SetProperty("ErrorDescription", value)
}

// GetErrorDescription gets the value of ErrorDescription for the instance
func (instance *MSFT_HealthAction) GetPropertyErrorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ErrorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessageParameters sets the value of MessageParameters for the instance
func (instance *MSFT_HealthAction) SetPropertyMessageParameters(value []string) (err error) {
	return instance.SetProperty("MessageParameters", value)
}

// GetMessageParameters gets the value of MessageParameters for the instance
func (instance *MSFT_HealthAction) GetPropertyMessageParameters() (value []string, err error) {
	retValue, err := instance.GetProperty("MessageParameters")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentComplete sets the value of PercentComplete for the instance
func (instance *MSFT_HealthAction) SetPropertyPercentComplete(value uint16) (err error) {
	return instance.SetProperty("PercentComplete", value)
}

// GetPercentComplete gets the value of PercentComplete for the instance
func (instance *MSFT_HealthAction) GetPropertyPercentComplete() (value uint16, err error) {
	retValue, err := instance.GetProperty("PercentComplete")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReason sets the value of Reason for the instance
func (instance *MSFT_HealthAction) SetPropertyReason(value string) (err error) {
	return instance.SetProperty("Reason", value)
}

// GetReason gets the value of Reason for the instance
func (instance *MSFT_HealthAction) GetPropertyReason() (value string, err error) {
	retValue, err := instance.GetProperty("Reason")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportingObjectId sets the value of ReportingObjectId for the instance
func (instance *MSFT_HealthAction) SetPropertyReportingObjectId(value string) (err error) {
	return instance.SetProperty("ReportingObjectId", value)
}

// GetReportingObjectId gets the value of ReportingObjectId for the instance
func (instance *MSFT_HealthAction) GetPropertyReportingObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("ReportingObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportingObjectType sets the value of ReportingObjectType for the instance
func (instance *MSFT_HealthAction) SetPropertyReportingObjectType(value string) (err error) {
	return instance.SetProperty("ReportingObjectType", value)
}

// GetReportingObjectType gets the value of ReportingObjectType for the instance
func (instance *MSFT_HealthAction) GetPropertyReportingObjectType() (value string, err error) {
	retValue, err := instance.GetProperty("ReportingObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportingObjectUniqueId sets the value of ReportingObjectUniqueId for the instance
func (instance *MSFT_HealthAction) SetPropertyReportingObjectUniqueId(value string) (err error) {
	return instance.SetProperty("ReportingObjectUniqueId", value)
}

// GetReportingObjectUniqueId gets the value of ReportingObjectUniqueId for the instance
func (instance *MSFT_HealthAction) GetPropertyReportingObjectUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("ReportingObjectUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartTime sets the value of StartTime for the instance
func (instance *MSFT_HealthAction) SetPropertyStartTime(value string) (err error) {
	return instance.SetProperty("StartTime", value)
}

// GetStartTime gets the value of StartTime for the instance
func (instance *MSFT_HealthAction) GetPropertyStartTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_HealthAction) SetPropertyState(value uint16) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_HealthAction) GetPropertyState() (value uint16, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_HealthAction) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_HealthAction) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_HealthAction) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_HealthAction) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
