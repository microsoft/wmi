// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_NETFramework_NETCLRLoading struct
type Win32_PerfFormattedData_NETFramework_NETCLRLoading struct {
	Win32_PerfFormattedData

	//
	AssemblySearchLength uint32

	//
	BytesinLoaderHeap uint32

	//
	Currentappdomains uint32

	//
	CurrentAssemblies uint32

	//
	CurrentClassesLoaded uint32

	//
	PercentTimeLoading uint64

	//
	Rateofappdomains uint32

	//
	Rateofappdomainsunloaded uint32

	//
	RateofAssemblies uint32

	//
	RateofClassesLoaded uint32

	//
	RateofLoadFailures uint32

	//
	TotalAppdomains uint32

	//
	Totalappdomainsunloaded uint32

	//
	TotalAssemblies uint32

	//
	TotalClassesLoaded uint32

	//
	TotalNumberofLoadFailures uint32
}

// SetAssemblySearchLength sets the value of AssemblySearchLength for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyAssemblySearchLength(value uint32) (err error) {
	return instance.SetProperty("AssemblySearchLength", value)
}

// GetAssemblySearchLength gets the value of AssemblySearchLength for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyAssemblySearchLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("AssemblySearchLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesinLoaderHeap sets the value of BytesinLoaderHeap for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyBytesinLoaderHeap(value uint32) (err error) {
	return instance.SetProperty("BytesinLoaderHeap", value)
}

// GetBytesinLoaderHeap gets the value of BytesinLoaderHeap for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyBytesinLoaderHeap() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesinLoaderHeap")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentappdomains sets the value of Currentappdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyCurrentappdomains(value uint32) (err error) {
	return instance.SetProperty("Currentappdomains", value)
}

// GetCurrentappdomains gets the value of Currentappdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyCurrentappdomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("Currentappdomains")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentAssemblies sets the value of CurrentAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyCurrentAssemblies(value uint32) (err error) {
	return instance.SetProperty("CurrentAssemblies", value)
}

// GetCurrentAssemblies gets the value of CurrentAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyCurrentAssemblies() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentAssemblies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentClassesLoaded sets the value of CurrentClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyCurrentClassesLoaded(value uint32) (err error) {
	return instance.SetProperty("CurrentClassesLoaded", value)
}

// GetCurrentClassesLoaded gets the value of CurrentClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyCurrentClassesLoaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentClassesLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPercentTimeLoading sets the value of PercentTimeLoading for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyPercentTimeLoading(value uint64) (err error) {
	return instance.SetProperty("PercentTimeLoading", value)
}

// GetPercentTimeLoading gets the value of PercentTimeLoading for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyPercentTimeLoading() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentTimeLoading")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRateofappdomains sets the value of Rateofappdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyRateofappdomains(value uint32) (err error) {
	return instance.SetProperty("Rateofappdomains", value)
}

// GetRateofappdomains gets the value of Rateofappdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyRateofappdomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("Rateofappdomains")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRateofappdomainsunloaded sets the value of Rateofappdomainsunloaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyRateofappdomainsunloaded(value uint32) (err error) {
	return instance.SetProperty("Rateofappdomainsunloaded", value)
}

// GetRateofappdomainsunloaded gets the value of Rateofappdomainsunloaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyRateofappdomainsunloaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("Rateofappdomainsunloaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRateofAssemblies sets the value of RateofAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyRateofAssemblies(value uint32) (err error) {
	return instance.SetProperty("RateofAssemblies", value)
}

// GetRateofAssemblies gets the value of RateofAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyRateofAssemblies() (value uint32, err error) {
	retValue, err := instance.GetProperty("RateofAssemblies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRateofClassesLoaded sets the value of RateofClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyRateofClassesLoaded(value uint32) (err error) {
	return instance.SetProperty("RateofClassesLoaded", value)
}

// GetRateofClassesLoaded gets the value of RateofClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyRateofClassesLoaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("RateofClassesLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRateofLoadFailures sets the value of RateofLoadFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyRateofLoadFailures(value uint32) (err error) {
	return instance.SetProperty("RateofLoadFailures", value)
}

// GetRateofLoadFailures gets the value of RateofLoadFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyRateofLoadFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("RateofLoadFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalAppdomains sets the value of TotalAppdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyTotalAppdomains(value uint32) (err error) {
	return instance.SetProperty("TotalAppdomains", value)
}

// GetTotalAppdomains gets the value of TotalAppdomains for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyTotalAppdomains() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalAppdomains")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalappdomainsunloaded sets the value of Totalappdomainsunloaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyTotalappdomainsunloaded(value uint32) (err error) {
	return instance.SetProperty("Totalappdomainsunloaded", value)
}

// GetTotalappdomainsunloaded gets the value of Totalappdomainsunloaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyTotalappdomainsunloaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("Totalappdomainsunloaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalAssemblies sets the value of TotalAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyTotalAssemblies(value uint32) (err error) {
	return instance.SetProperty("TotalAssemblies", value)
}

// GetTotalAssemblies gets the value of TotalAssemblies for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyTotalAssemblies() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalAssemblies")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalClassesLoaded sets the value of TotalClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyTotalClassesLoaded(value uint32) (err error) {
	return instance.SetProperty("TotalClassesLoaded", value)
}

// GetTotalClassesLoaded gets the value of TotalClassesLoaded for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyTotalClassesLoaded() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalClassesLoaded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTotalNumberofLoadFailures sets the value of TotalNumberofLoadFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) SetPropertyTotalNumberofLoadFailures(value uint32) (err error) {
	return instance.SetProperty("TotalNumberofLoadFailures", value)
}

// GetTotalNumberofLoadFailures gets the value of TotalNumberofLoadFailures for the instance
func (instance *Win32_PerfFormattedData_NETFramework_NETCLRLoading) GetPropertyTotalNumberofLoadFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalNumberofLoadFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
