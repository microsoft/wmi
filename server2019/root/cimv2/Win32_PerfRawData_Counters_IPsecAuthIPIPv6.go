// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_IPsecAuthIPIPv6 struct
type Win32_PerfRawData_Counters_IPsecAuthIPIPv6 struct {
	Win32_PerfRawData

	//
	ActiveExtendedModeSAs uint32

	//
	ActiveMainModeSAs uint32

	//
	ActiveQuickModeSAs uint32

	//
	ExtendedModeNegotiations uint32

	//
	ExtendedModeNegotiationsPersec uint32

	//
	ExtendedModeSAsThatUsedImpersonation uint32

	//
	FailedExtendedModeNegotiations uint32

	//
	FailedExtendedModeNegotiationsPersec uint32

	//
	FailedMainModeNegotiations uint32

	//
	FailedMainModeNegotiationsPersec uint32

	//
	FailedQuickModeNegotiations uint32

	//
	FailedQuickModeNegotiationsPersec uint32

	//
	MainModeNegotiationRequestsReceived uint32

	//
	MainModeNegotiationRequestsReceivedPersec uint32

	//
	MainModeNegotiations uint32

	//
	MainModeNegotiationsPersec uint32

	//
	MainModeSAsThatUsedImpersonation uint32

	//
	MainModeSAsThatUsedImpersonationPersec uint32

	//
	PendingExtendedModeNegotiations uint32

	//
	PendingMainModeNegotiations uint32

	//
	PendingQuickModeNegotiations uint32

	//
	QuickModeNegotiations uint32

	//
	QuickModeNegotiationsPersec uint32

	//
	SuccessfulExtendedModeNegotiations uint32

	//
	SuccessfulExtendedModeNegotiationsPersec uint32

	//
	SuccessfulMainModeNegotiations uint32

	//
	SuccessfulMainModeNegotiationsPersec uint32

	//
	SuccessfulQuickModeNegotiations uint32

	//
	SuccessfulQuickModeNegotiationsPersec uint32
}

// SetActiveExtendedModeSAs sets the value of ActiveExtendedModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyActiveExtendedModeSAs(value uint32) (err error) {
	return instance.SetProperty("ActiveExtendedModeSAs", value)
}

// GetActiveExtendedModeSAs gets the value of ActiveExtendedModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyActiveExtendedModeSAs() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveExtendedModeSAs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveMainModeSAs sets the value of ActiveMainModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyActiveMainModeSAs(value uint32) (err error) {
	return instance.SetProperty("ActiveMainModeSAs", value)
}

// GetActiveMainModeSAs gets the value of ActiveMainModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyActiveMainModeSAs() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveMainModeSAs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActiveQuickModeSAs sets the value of ActiveQuickModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyActiveQuickModeSAs(value uint32) (err error) {
	return instance.SetProperty("ActiveQuickModeSAs", value)
}

// GetActiveQuickModeSAs gets the value of ActiveQuickModeSAs for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyActiveQuickModeSAs() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveQuickModeSAs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtendedModeNegotiations sets the value of ExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyExtendedModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("ExtendedModeNegotiations", value)
}

// GetExtendedModeNegotiations gets the value of ExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyExtendedModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtendedModeNegotiationsPersec sets the value of ExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyExtendedModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("ExtendedModeNegotiationsPersec", value)
}

// GetExtendedModeNegotiationsPersec gets the value of ExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyExtendedModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExtendedModeSAsThatUsedImpersonation sets the value of ExtendedModeSAsThatUsedImpersonation for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyExtendedModeSAsThatUsedImpersonation(value uint32) (err error) {
	return instance.SetProperty("ExtendedModeSAsThatUsedImpersonation", value)
}

// GetExtendedModeSAsThatUsedImpersonation gets the value of ExtendedModeSAsThatUsedImpersonation for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyExtendedModeSAsThatUsedImpersonation() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExtendedModeSAsThatUsedImpersonation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedExtendedModeNegotiations sets the value of FailedExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedExtendedModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("FailedExtendedModeNegotiations", value)
}

// GetFailedExtendedModeNegotiations gets the value of FailedExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedExtendedModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedExtendedModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedExtendedModeNegotiationsPersec sets the value of FailedExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedExtendedModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("FailedExtendedModeNegotiationsPersec", value)
}

// GetFailedExtendedModeNegotiationsPersec gets the value of FailedExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedExtendedModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedExtendedModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedMainModeNegotiations sets the value of FailedMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedMainModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("FailedMainModeNegotiations", value)
}

// GetFailedMainModeNegotiations gets the value of FailedMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedMainModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedMainModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedMainModeNegotiationsPersec sets the value of FailedMainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedMainModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("FailedMainModeNegotiationsPersec", value)
}

// GetFailedMainModeNegotiationsPersec gets the value of FailedMainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedMainModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedMainModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedQuickModeNegotiations sets the value of FailedQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedQuickModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("FailedQuickModeNegotiations", value)
}

// GetFailedQuickModeNegotiations gets the value of FailedQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedQuickModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedQuickModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedQuickModeNegotiationsPersec sets the value of FailedQuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyFailedQuickModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("FailedQuickModeNegotiationsPersec", value)
}

// GetFailedQuickModeNegotiationsPersec gets the value of FailedQuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyFailedQuickModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailedQuickModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeNegotiationRequestsReceived sets the value of MainModeNegotiationRequestsReceived for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeNegotiationRequestsReceived(value uint32) (err error) {
	return instance.SetProperty("MainModeNegotiationRequestsReceived", value)
}

// GetMainModeNegotiationRequestsReceived gets the value of MainModeNegotiationRequestsReceived for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeNegotiationRequestsReceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeNegotiationRequestsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeNegotiationRequestsReceivedPersec sets the value of MainModeNegotiationRequestsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeNegotiationRequestsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("MainModeNegotiationRequestsReceivedPersec", value)
}

// GetMainModeNegotiationRequestsReceivedPersec gets the value of MainModeNegotiationRequestsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeNegotiationRequestsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeNegotiationRequestsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeNegotiations sets the value of MainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("MainModeNegotiations", value)
}

// GetMainModeNegotiations gets the value of MainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeNegotiationsPersec sets the value of MainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("MainModeNegotiationsPersec", value)
}

// GetMainModeNegotiationsPersec gets the value of MainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeSAsThatUsedImpersonation sets the value of MainModeSAsThatUsedImpersonation for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeSAsThatUsedImpersonation(value uint32) (err error) {
	return instance.SetProperty("MainModeSAsThatUsedImpersonation", value)
}

// GetMainModeSAsThatUsedImpersonation gets the value of MainModeSAsThatUsedImpersonation for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeSAsThatUsedImpersonation() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeSAsThatUsedImpersonation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMainModeSAsThatUsedImpersonationPersec sets the value of MainModeSAsThatUsedImpersonationPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyMainModeSAsThatUsedImpersonationPersec(value uint32) (err error) {
	return instance.SetProperty("MainModeSAsThatUsedImpersonationPersec", value)
}

// GetMainModeSAsThatUsedImpersonationPersec gets the value of MainModeSAsThatUsedImpersonationPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyMainModeSAsThatUsedImpersonationPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MainModeSAsThatUsedImpersonationPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPendingExtendedModeNegotiations sets the value of PendingExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyPendingExtendedModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("PendingExtendedModeNegotiations", value)
}

// GetPendingExtendedModeNegotiations gets the value of PendingExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyPendingExtendedModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingExtendedModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPendingMainModeNegotiations sets the value of PendingMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyPendingMainModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("PendingMainModeNegotiations", value)
}

// GetPendingMainModeNegotiations gets the value of PendingMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyPendingMainModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingMainModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPendingQuickModeNegotiations sets the value of PendingQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyPendingQuickModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("PendingQuickModeNegotiations", value)
}

// GetPendingQuickModeNegotiations gets the value of PendingQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyPendingQuickModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingQuickModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickModeNegotiations sets the value of QuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyQuickModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("QuickModeNegotiations", value)
}

// GetQuickModeNegotiations gets the value of QuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyQuickModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuickModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickModeNegotiationsPersec sets the value of QuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertyQuickModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("QuickModeNegotiationsPersec", value)
}

// GetQuickModeNegotiationsPersec gets the value of QuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertyQuickModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuickModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulExtendedModeNegotiations sets the value of SuccessfulExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulExtendedModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("SuccessfulExtendedModeNegotiations", value)
}

// GetSuccessfulExtendedModeNegotiations gets the value of SuccessfulExtendedModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulExtendedModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulExtendedModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulExtendedModeNegotiationsPersec sets the value of SuccessfulExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulExtendedModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("SuccessfulExtendedModeNegotiationsPersec", value)
}

// GetSuccessfulExtendedModeNegotiationsPersec gets the value of SuccessfulExtendedModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulExtendedModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulExtendedModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulMainModeNegotiations sets the value of SuccessfulMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulMainModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("SuccessfulMainModeNegotiations", value)
}

// GetSuccessfulMainModeNegotiations gets the value of SuccessfulMainModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulMainModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulMainModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulMainModeNegotiationsPersec sets the value of SuccessfulMainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulMainModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("SuccessfulMainModeNegotiationsPersec", value)
}

// GetSuccessfulMainModeNegotiationsPersec gets the value of SuccessfulMainModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulMainModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulMainModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulQuickModeNegotiations sets the value of SuccessfulQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulQuickModeNegotiations(value uint32) (err error) {
	return instance.SetProperty("SuccessfulQuickModeNegotiations", value)
}

// GetSuccessfulQuickModeNegotiations gets the value of SuccessfulQuickModeNegotiations for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulQuickModeNegotiations() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulQuickModeNegotiations")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuccessfulQuickModeNegotiationsPersec sets the value of SuccessfulQuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) SetPropertySuccessfulQuickModeNegotiationsPersec(value uint32) (err error) {
	return instance.SetProperty("SuccessfulQuickModeNegotiationsPersec", value)
}

// GetSuccessfulQuickModeNegotiationsPersec gets the value of SuccessfulQuickModeNegotiationsPersec for the instance
func (instance *Win32_PerfRawData_Counters_IPsecAuthIPIPv6) GetPropertySuccessfulQuickModeNegotiationsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SuccessfulQuickModeNegotiationsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
