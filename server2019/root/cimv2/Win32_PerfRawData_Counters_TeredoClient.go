// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Counters_TeredoClient struct
type Win32_PerfRawData_Counters_TeredoClient struct {
	Win32_PerfRawData

	//
	InTeredoBubble uint32

	//
	InTeredoData uint64

	//
	InTeredoDataKernelMode uint64

	//
	InTeredoDataUserMode uint64

	//
	InTeredoInvalid uint32

	//
	InTeredoRouterAdvertisement uint32

	//
	OutTeredoBubble uint32

	//
	OutTeredoData uint64

	//
	OutTeredoDataKernelMode uint64

	//
	OutTeredoDataUserMode uint64

	//
	OutTeredoRouterSolicitation uint32
}

// SetInTeredoBubble sets the value of InTeredoBubble for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoBubble(value uint32) (err error) {
	return instance.SetProperty("InTeredoBubble", value)
}

// GetInTeredoBubble gets the value of InTeredoBubble for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoBubble() (value uint32, err error) {
	retValue, err := instance.GetProperty("InTeredoBubble")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInTeredoData sets the value of InTeredoData for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoData(value uint64) (err error) {
	return instance.SetProperty("InTeredoData", value)
}

// GetInTeredoData gets the value of InTeredoData for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoData() (value uint64, err error) {
	retValue, err := instance.GetProperty("InTeredoData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInTeredoDataKernelMode sets the value of InTeredoDataKernelMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoDataKernelMode(value uint64) (err error) {
	return instance.SetProperty("InTeredoDataKernelMode", value)
}

// GetInTeredoDataKernelMode gets the value of InTeredoDataKernelMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoDataKernelMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("InTeredoDataKernelMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInTeredoDataUserMode sets the value of InTeredoDataUserMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoDataUserMode(value uint64) (err error) {
	return instance.SetProperty("InTeredoDataUserMode", value)
}

// GetInTeredoDataUserMode gets the value of InTeredoDataUserMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoDataUserMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("InTeredoDataUserMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInTeredoInvalid sets the value of InTeredoInvalid for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoInvalid(value uint32) (err error) {
	return instance.SetProperty("InTeredoInvalid", value)
}

// GetInTeredoInvalid gets the value of InTeredoInvalid for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoInvalid() (value uint32, err error) {
	retValue, err := instance.GetProperty("InTeredoInvalid")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInTeredoRouterAdvertisement sets the value of InTeredoRouterAdvertisement for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyInTeredoRouterAdvertisement(value uint32) (err error) {
	return instance.SetProperty("InTeredoRouterAdvertisement", value)
}

// GetInTeredoRouterAdvertisement gets the value of InTeredoRouterAdvertisement for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyInTeredoRouterAdvertisement() (value uint32, err error) {
	retValue, err := instance.GetProperty("InTeredoRouterAdvertisement")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutTeredoBubble sets the value of OutTeredoBubble for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyOutTeredoBubble(value uint32) (err error) {
	return instance.SetProperty("OutTeredoBubble", value)
}

// GetOutTeredoBubble gets the value of OutTeredoBubble for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyOutTeredoBubble() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutTeredoBubble")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutTeredoData sets the value of OutTeredoData for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyOutTeredoData(value uint64) (err error) {
	return instance.SetProperty("OutTeredoData", value)
}

// GetOutTeredoData gets the value of OutTeredoData for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyOutTeredoData() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutTeredoData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutTeredoDataKernelMode sets the value of OutTeredoDataKernelMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyOutTeredoDataKernelMode(value uint64) (err error) {
	return instance.SetProperty("OutTeredoDataKernelMode", value)
}

// GetOutTeredoDataKernelMode gets the value of OutTeredoDataKernelMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyOutTeredoDataKernelMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutTeredoDataKernelMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutTeredoDataUserMode sets the value of OutTeredoDataUserMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyOutTeredoDataUserMode(value uint64) (err error) {
	return instance.SetProperty("OutTeredoDataUserMode", value)
}

// GetOutTeredoDataUserMode gets the value of OutTeredoDataUserMode for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyOutTeredoDataUserMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutTeredoDataUserMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutTeredoRouterSolicitation sets the value of OutTeredoRouterSolicitation for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) SetPropertyOutTeredoRouterSolicitation(value uint32) (err error) {
	return instance.SetProperty("OutTeredoRouterSolicitation", value)
}

// GetOutTeredoRouterSolicitation gets the value of OutTeredoRouterSolicitation for the instance
func (instance *Win32_PerfRawData_Counters_TeredoClient) GetPropertyOutTeredoRouterSolicitation() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutTeredoRouterSolicitation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
