// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_DiagnosticSetting struct
type CIM_DiagnosticSetting struct {
	CIM_Setting

	//
	HaltOnError bool

	//
	Locales []string

	//
	LogOptions []DiagnosticSetting_LogOptions

	//
	LogStorage []DiagnosticSetting_LogStorage

	//
	LoopControl []DiagnosticSetting_LoopControl

	//
	LoopControlParameter []string

	//
	OtherLogOptionsDescriptions []string

	//
	OtherLogStorageDescriptions []string

	//
	OtherLoopControlDescription string

	//
	OtherLoopControlDescriptions []string

	//
	PercentOfTestCoverage uint8

	//
	QuickMode bool

	//
	ReportSoftErrors bool

	//
	ReportStatusMessages bool

	//
	ResultPersistence uint32

	//
	TestWarningLevel DiagnosticSetting_TestWarningLevel

	//
	VerbosityLevel []DiagnosticSetting_VerbosityLevel
}

// SetHaltOnError sets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyHaltOnError(value bool) (err error) {
	return instance.SetProperty("HaltOnError", value)
}

// GetHaltOnError gets the value of HaltOnError for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyHaltOnError() (value bool, err error) {
	retValue, err := instance.GetProperty("HaltOnError")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocales sets the value of Locales for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyLocales(value []string) (err error) {
	return instance.SetProperty("Locales", value)
}

// GetLocales gets the value of Locales for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyLocales() (value []string, err error) {
	retValue, err := instance.GetProperty("Locales")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogOptions sets the value of LogOptions for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyLogOptions(value []DiagnosticSetting_LogOptions) (err error) {
	return instance.SetProperty("LogOptions", value)
}

// GetLogOptions gets the value of LogOptions for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyLogOptions() (value []DiagnosticSetting_LogOptions, err error) {
	retValue, err := instance.GetProperty("LogOptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticSetting_LogOptions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogStorage sets the value of LogStorage for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyLogStorage(value []DiagnosticSetting_LogStorage) (err error) {
	return instance.SetProperty("LogStorage", value)
}

// GetLogStorage gets the value of LogStorage for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyLogStorage() (value []DiagnosticSetting_LogStorage, err error) {
	retValue, err := instance.GetProperty("LogStorage")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticSetting_LogStorage)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoopControl sets the value of LoopControl for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyLoopControl(value []DiagnosticSetting_LoopControl) (err error) {
	return instance.SetProperty("LoopControl", value)
}

// GetLoopControl gets the value of LoopControl for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyLoopControl() (value []DiagnosticSetting_LoopControl, err error) {
	retValue, err := instance.GetProperty("LoopControl")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticSetting_LoopControl)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLoopControlParameter sets the value of LoopControlParameter for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyLoopControlParameter(value []string) (err error) {
	return instance.SetProperty("LoopControlParameter", value)
}

// GetLoopControlParameter gets the value of LoopControlParameter for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyLoopControlParameter() (value []string, err error) {
	retValue, err := instance.GetProperty("LoopControlParameter")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLogOptionsDescriptions sets the value of OtherLogOptionsDescriptions for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyOtherLogOptionsDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLogOptionsDescriptions", value)
}

// GetOtherLogOptionsDescriptions gets the value of OtherLogOptionsDescriptions for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyOtherLogOptionsDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLogOptionsDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLogStorageDescriptions sets the value of OtherLogStorageDescriptions for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyOtherLogStorageDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLogStorageDescriptions", value)
}

// GetOtherLogStorageDescriptions gets the value of OtherLogStorageDescriptions for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyOtherLogStorageDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLogStorageDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLoopControlDescription sets the value of OtherLoopControlDescription for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyOtherLoopControlDescription(value string) (err error) {
	return instance.SetProperty("OtherLoopControlDescription", value)
}

// GetOtherLoopControlDescription gets the value of OtherLoopControlDescription for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyOtherLoopControlDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLoopControlDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLoopControlDescriptions sets the value of OtherLoopControlDescriptions for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyOtherLoopControlDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherLoopControlDescriptions", value)
}

// GetOtherLoopControlDescriptions gets the value of OtherLoopControlDescriptions for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyOtherLoopControlDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherLoopControlDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentOfTestCoverage sets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyPercentOfTestCoverage(value uint8) (err error) {
	return instance.SetProperty("PercentOfTestCoverage", value)
}

// GetPercentOfTestCoverage gets the value of PercentOfTestCoverage for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyPercentOfTestCoverage() (value uint8, err error) {
	retValue, err := instance.GetProperty("PercentOfTestCoverage")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQuickMode sets the value of QuickMode for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyQuickMode(value bool) (err error) {
	return instance.SetProperty("QuickMode", value)
}

// GetQuickMode gets the value of QuickMode for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyQuickMode() (value bool, err error) {
	retValue, err := instance.GetProperty("QuickMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportSoftErrors sets the value of ReportSoftErrors for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyReportSoftErrors(value bool) (err error) {
	return instance.SetProperty("ReportSoftErrors", value)
}

// GetReportSoftErrors gets the value of ReportSoftErrors for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyReportSoftErrors() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportSoftErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReportStatusMessages sets the value of ReportStatusMessages for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyReportStatusMessages(value bool) (err error) {
	return instance.SetProperty("ReportStatusMessages", value)
}

// GetReportStatusMessages gets the value of ReportStatusMessages for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyReportStatusMessages() (value bool, err error) {
	retValue, err := instance.GetProperty("ReportStatusMessages")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResultPersistence sets the value of ResultPersistence for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyResultPersistence(value uint32) (err error) {
	return instance.SetProperty("ResultPersistence", value)
}

// GetResultPersistence gets the value of ResultPersistence for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyResultPersistence() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResultPersistence")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTestWarningLevel sets the value of TestWarningLevel for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyTestWarningLevel(value DiagnosticSetting_TestWarningLevel) (err error) {
	return instance.SetProperty("TestWarningLevel", value)
}

// GetTestWarningLevel gets the value of TestWarningLevel for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyTestWarningLevel() (value DiagnosticSetting_TestWarningLevel, err error) {
	retValue, err := instance.GetProperty("TestWarningLevel")
	if err != nil {
		return
	}
	value, ok := retValue.(DiagnosticSetting_TestWarningLevel)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVerbosityLevel sets the value of VerbosityLevel for the instance
func (instance *CIM_DiagnosticSetting) SetPropertyVerbosityLevel(value []DiagnosticSetting_VerbosityLevel) (err error) {
	return instance.SetProperty("VerbosityLevel", value)
}

// GetVerbosityLevel gets the value of VerbosityLevel for the instance
func (instance *CIM_DiagnosticSetting) GetPropertyVerbosityLevel() (value []DiagnosticSetting_VerbosityLevel, err error) {
	retValue, err := instance.GetProperty("VerbosityLevel")
	if err != nil {
		return
	}
	value, ok := retValue.([]DiagnosticSetting_VerbosityLevel)
	if !ok {
		// TODO: Set an error
	}
	return
}
