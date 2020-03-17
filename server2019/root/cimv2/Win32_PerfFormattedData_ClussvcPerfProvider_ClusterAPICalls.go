// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls struct
type Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls struct {
	Win32_PerfFormattedData

	//
	ClusterAPICallsPersec uint64

	//
	GroupAPICallsPersec uint64

	//
	KeyAPICallsPersec uint64

	//
	NetworkAPICallsPersec uint64

	//
	NetworkInterfaceAPICallsPersec uint64

	//
	NodeAPICallsPersec uint64

	//
	NotificationAPICallsPersec uint64

	//
	NotificationBatchAPICallsPersec uint64

	//
	ResourceAPICallsPersec uint64
}

// SetClusterAPICallsPersec sets the value of ClusterAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyClusterAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("ClusterAPICallsPersec", value)
}

// GetClusterAPICallsPersec gets the value of ClusterAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyClusterAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ClusterAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupAPICallsPersec sets the value of GroupAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyGroupAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("GroupAPICallsPersec", value)
}

// GetGroupAPICallsPersec gets the value of GroupAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyGroupAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyAPICallsPersec sets the value of KeyAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyKeyAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("KeyAPICallsPersec", value)
}

// GetKeyAPICallsPersec gets the value of KeyAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyKeyAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAPICallsPersec sets the value of NetworkAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyNetworkAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("NetworkAPICallsPersec", value)
}

// GetNetworkAPICallsPersec gets the value of NetworkAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyNetworkAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetworkAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkInterfaceAPICallsPersec sets the value of NetworkInterfaceAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyNetworkInterfaceAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("NetworkInterfaceAPICallsPersec", value)
}

// GetNetworkInterfaceAPICallsPersec gets the value of NetworkInterfaceAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyNetworkInterfaceAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetworkInterfaceAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNodeAPICallsPersec sets the value of NodeAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyNodeAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("NodeAPICallsPersec", value)
}

// GetNodeAPICallsPersec gets the value of NodeAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyNodeAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NodeAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationAPICallsPersec sets the value of NotificationAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyNotificationAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("NotificationAPICallsPersec", value)
}

// GetNotificationAPICallsPersec gets the value of NotificationAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyNotificationAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NotificationAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotificationBatchAPICallsPersec sets the value of NotificationBatchAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyNotificationBatchAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("NotificationBatchAPICallsPersec", value)
}

// GetNotificationBatchAPICallsPersec gets the value of NotificationBatchAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyNotificationBatchAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("NotificationBatchAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceAPICallsPersec sets the value of ResourceAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) SetPropertyResourceAPICallsPersec(value uint64) (err error) {
	return instance.SetProperty("ResourceAPICallsPersec", value)
}

// GetResourceAPICallsPersec gets the value of ResourceAPICallsPersec for the instance
func (instance *Win32_PerfFormattedData_ClussvcPerfProvider_ClusterAPICalls) GetPropertyResourceAPICallsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceAPICallsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
