// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads struct
type Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads struct {
	Win32_PerfRawData

	//
	ContentionRatePersec uint32

	//
	CurrentQueueLength uint32

	//
	NumberofcurrentlogicalThreads uint32

	//
	NumberofcurrentphysicalThreads uint32

	//
	Numberofcurrentrecognizedthreads uint32

	//
	Numberoftotalrecognizedthreads uint32

	//
	QueueLengthPeak uint32

	//
	QueueLengthPersec uint32

	//
	rateofrecognizedthreadsPersec uint32

	//
	TotalNumberofContentions uint32
}

// SetContentionRatePersec sets the value of ContentionRatePersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyContentionRatePersec(value uint32) (err error) {
	return instance.SetProperty("ContentionRatePersec", value)
}

// GetContentionRatePersec gets the value of ContentionRatePersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyContentionRatePersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ContentionRatePersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentQueueLength sets the value of CurrentQueueLength for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyCurrentQueueLength(value uint32) (err error) {
	return instance.SetProperty("CurrentQueueLength", value)
}

// GetCurrentQueueLength gets the value of CurrentQueueLength for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyCurrentQueueLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentQueueLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofcurrentlogicalThreads sets the value of NumberofcurrentlogicalThreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyNumberofcurrentlogicalThreads(value uint32) (err error) {
	return instance.SetProperty("NumberofcurrentlogicalThreads", value)
}

// GetNumberofcurrentlogicalThreads gets the value of NumberofcurrentlogicalThreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyNumberofcurrentlogicalThreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofcurrentlogicalThreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofcurrentphysicalThreads sets the value of NumberofcurrentphysicalThreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyNumberofcurrentphysicalThreads(value uint32) (err error) {
	return instance.SetProperty("NumberofcurrentphysicalThreads", value)
}

// GetNumberofcurrentphysicalThreads gets the value of NumberofcurrentphysicalThreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyNumberofcurrentphysicalThreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofcurrentphysicalThreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberofcurrentrecognizedthreads sets the value of Numberofcurrentrecognizedthreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyNumberofcurrentrecognizedthreads(value uint32) (err error) {
	return instance.SetProperty("Numberofcurrentrecognizedthreads", value)
}

// GetNumberofcurrentrecognizedthreads gets the value of Numberofcurrentrecognizedthreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyNumberofcurrentrecognizedthreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberofcurrentrecognizedthreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberoftotalrecognizedthreads sets the value of Numberoftotalrecognizedthreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyNumberoftotalrecognizedthreads(value uint32) (err error) {
	return instance.SetProperty("Numberoftotalrecognizedthreads", value)
}

// GetNumberoftotalrecognizedthreads gets the value of Numberoftotalrecognizedthreads for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyNumberoftotalrecognizedthreads() (value uint32, err error) {
	retValue, err := instance.GetProperty("Numberoftotalrecognizedthreads")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueueLengthPeak sets the value of QueueLengthPeak for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyQueueLengthPeak(value uint32) (err error) {
	return instance.SetProperty("QueueLengthPeak", value)
}

// GetQueueLengthPeak gets the value of QueueLengthPeak for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyQueueLengthPeak() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueLengthPeak")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQueueLengthPersec sets the value of QueueLengthPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyQueueLengthPersec(value uint32) (err error) {
	return instance.SetProperty("QueueLengthPersec", value)
}

// GetQueueLengthPersec gets the value of QueueLengthPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyQueueLengthPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("QueueLengthPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrateofrecognizedthreadsPersec sets the value of rateofrecognizedthreadsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyrateofrecognizedthreadsPersec(value uint32) (err error) {
	return instance.SetProperty("rateofrecognizedthreadsPersec", value)
}

// GetrateofrecognizedthreadsPersec gets the value of rateofrecognizedthreadsPersec for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyrateofrecognizedthreadsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("rateofrecognizedthreadsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofContentions sets the value of TotalNumberofContentions for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) SetPropertyTotalNumberofContentions(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofContentions", value)
}

// GetTotalNumberofContentions gets the value of TotalNumberofContentions for the instance
func (instance *Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads) GetPropertyTotalNumberofContentions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofContentions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
