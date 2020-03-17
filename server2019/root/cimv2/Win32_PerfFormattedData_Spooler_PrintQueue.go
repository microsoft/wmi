// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Spooler_PrintQueue struct
type Win32_PerfFormattedData_Spooler_PrintQueue struct {
	Win32_PerfFormattedData

	//
	AddNetworkPrinterCalls uint32

	//
	BytesPrintedPersec uint64

	//
	EnumerateNetworkPrinterCalls uint32

	//
	JobErrors uint32

	//
	Jobs uint32

	//
	JobsSpooling uint32

	//
	MaxJobsSpooling uint32

	//
	MaxReferences uint32

	//
	NotReadyErrors uint32

	//
	OutofPaperErrors uint32

	//
	References uint32

	//
	TotalJobsPrinted uint32

	//
	TotalPagesPrinted uint32
}

// SetAddNetworkPrinterCalls sets the value of AddNetworkPrinterCalls for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyAddNetworkPrinterCalls(value uint32) (err error) {
	return instance.SetProperty("AddNetworkPrinterCalls", value)
}

// GetAddNetworkPrinterCalls gets the value of AddNetworkPrinterCalls for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyAddNetworkPrinterCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("AddNetworkPrinterCalls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesPrintedPersec sets the value of BytesPrintedPersec for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyBytesPrintedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesPrintedPersec", value)
}

// GetBytesPrintedPersec gets the value of BytesPrintedPersec for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyBytesPrintedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesPrintedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnumerateNetworkPrinterCalls sets the value of EnumerateNetworkPrinterCalls for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyEnumerateNetworkPrinterCalls(value uint32) (err error) {
	return instance.SetProperty("EnumerateNetworkPrinterCalls", value)
}

// GetEnumerateNetworkPrinterCalls gets the value of EnumerateNetworkPrinterCalls for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyEnumerateNetworkPrinterCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnumerateNetworkPrinterCalls")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobErrors sets the value of JobErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyJobErrors(value uint32) (err error) {
	return instance.SetProperty("JobErrors", value)
}

// GetJobErrors gets the value of JobErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyJobErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobs sets the value of Jobs for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyJobs(value uint32) (err error) {
	return instance.SetProperty("Jobs", value)
}

// GetJobs gets the value of Jobs for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyJobs() (value uint32, err error) {
	retValue, err := instance.GetProperty("Jobs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobsSpooling sets the value of JobsSpooling for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyJobsSpooling(value uint32) (err error) {
	return instance.SetProperty("JobsSpooling", value)
}

// GetJobsSpooling gets the value of JobsSpooling for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyJobsSpooling() (value uint32, err error) {
	retValue, err := instance.GetProperty("JobsSpooling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxJobsSpooling sets the value of MaxJobsSpooling for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyMaxJobsSpooling(value uint32) (err error) {
	return instance.SetProperty("MaxJobsSpooling", value)
}

// GetMaxJobsSpooling gets the value of MaxJobsSpooling for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyMaxJobsSpooling() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxJobsSpooling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReferences sets the value of MaxReferences for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyMaxReferences(value uint32) (err error) {
	return instance.SetProperty("MaxReferences", value)
}

// GetMaxReferences gets the value of MaxReferences for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyMaxReferences() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReferences")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotReadyErrors sets the value of NotReadyErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyNotReadyErrors(value uint32) (err error) {
	return instance.SetProperty("NotReadyErrors", value)
}

// GetNotReadyErrors gets the value of NotReadyErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyNotReadyErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("NotReadyErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutofPaperErrors sets the value of OutofPaperErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyOutofPaperErrors(value uint32) (err error) {
	return instance.SetProperty("OutofPaperErrors", value)
}

// GetOutofPaperErrors gets the value of OutofPaperErrors for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyOutofPaperErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutofPaperErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReferences sets the value of References for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyReferences(value uint32) (err error) {
	return instance.SetProperty("References", value)
}

// GetReferences gets the value of References for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyReferences() (value uint32, err error) {
	retValue, err := instance.GetProperty("References")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalJobsPrinted sets the value of TotalJobsPrinted for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyTotalJobsPrinted(value uint32) (err error) {
	return instance.SetProperty("TotalJobsPrinted", value)
}

// GetTotalJobsPrinted gets the value of TotalJobsPrinted for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyTotalJobsPrinted() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalJobsPrinted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalPagesPrinted sets the value of TotalPagesPrinted for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) SetPropertyTotalPagesPrinted(value uint32) (err error) {
	return instance.SetProperty("TotalPagesPrinted", value)
}

// GetTotalPagesPrinted gets the value of TotalPagesPrinted for the instance
func (instance *Win32_PerfFormattedData_Spooler_PrintQueue) GetPropertyTotalPagesPrinted() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalPagesPrinted")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
