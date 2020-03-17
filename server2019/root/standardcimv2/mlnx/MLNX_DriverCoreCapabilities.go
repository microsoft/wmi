// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_DriverCoreCapabilities struct
type MLNX_DriverCoreCapabilities struct {
	MLNX_DriverCapabilities

	//
	LogMttsPerSeg_Max uint32

	//
	LogMttsPerSeg_Min uint32

	//
	LogNumCq_Max uint32

	//
	LogNumCq_Min uint32

	//
	LogNumMac_Max uint32

	//
	LogNumMac_Min uint32

	//
	LogNumMcg_Max uint32

	//
	LogNumMcg_Min uint32

	//
	LogNumMpt_Max uint32

	//
	LogNumMpt_Min uint32

	//
	LogNumMtt_Max uint32

	//
	LogNumMtt_Min uint32

	//
	LogNumQp_Max uint32

	//
	LogNumQp_Min uint32

	//
	LogNumRdmaRc_Max uint32

	//
	LogNumRdmaRc_Min uint32

	//
	LogNumSrq_Max uint32

	//
	LogNumSrq_Min uint32

	//
	LogNumVlan_Max uint32

	//
	LogNumVlan_Min uint32

	//
	MaximumWorkingThreads_Max uint32

	//
	MaximumWorkingThreads_Min uint32

	//
	Set4kMtu_Max uint32

	//
	Set4kMtu_Min uint32
}

// SetLogMttsPerSeg_Max sets the value of LogMttsPerSeg_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogMttsPerSeg_Max(value uint32) (err error) {
	return instance.SetProperty("LogMttsPerSeg_Max", value)
}

// GetLogMttsPerSeg_Max gets the value of LogMttsPerSeg_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogMttsPerSeg_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogMttsPerSeg_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogMttsPerSeg_Min sets the value of LogMttsPerSeg_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogMttsPerSeg_Min(value uint32) (err error) {
	return instance.SetProperty("LogMttsPerSeg_Min", value)
}

// GetLogMttsPerSeg_Min gets the value of LogMttsPerSeg_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogMttsPerSeg_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogMttsPerSeg_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumCq_Max sets the value of LogNumCq_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumCq_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumCq_Max", value)
}

// GetLogNumCq_Max gets the value of LogNumCq_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumCq_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumCq_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumCq_Min sets the value of LogNumCq_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumCq_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumCq_Min", value)
}

// GetLogNumCq_Min gets the value of LogNumCq_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumCq_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumCq_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMac_Max sets the value of LogNumMac_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMac_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumMac_Max", value)
}

// GetLogNumMac_Max gets the value of LogNumMac_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMac_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMac_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMac_Min sets the value of LogNumMac_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMac_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumMac_Min", value)
}

// GetLogNumMac_Min gets the value of LogNumMac_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMac_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMac_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMcg_Max sets the value of LogNumMcg_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMcg_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumMcg_Max", value)
}

// GetLogNumMcg_Max gets the value of LogNumMcg_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMcg_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMcg_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMcg_Min sets the value of LogNumMcg_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMcg_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumMcg_Min", value)
}

// GetLogNumMcg_Min gets the value of LogNumMcg_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMcg_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMcg_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMpt_Max sets the value of LogNumMpt_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMpt_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumMpt_Max", value)
}

// GetLogNumMpt_Max gets the value of LogNumMpt_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMpt_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMpt_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMpt_Min sets the value of LogNumMpt_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMpt_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumMpt_Min", value)
}

// GetLogNumMpt_Min gets the value of LogNumMpt_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMpt_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMpt_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMtt_Max sets the value of LogNumMtt_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMtt_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumMtt_Max", value)
}

// GetLogNumMtt_Max gets the value of LogNumMtt_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMtt_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMtt_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumMtt_Min sets the value of LogNumMtt_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumMtt_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumMtt_Min", value)
}

// GetLogNumMtt_Min gets the value of LogNumMtt_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumMtt_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumMtt_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumQp_Max sets the value of LogNumQp_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumQp_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumQp_Max", value)
}

// GetLogNumQp_Max gets the value of LogNumQp_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumQp_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumQp_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumQp_Min sets the value of LogNumQp_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumQp_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumQp_Min", value)
}

// GetLogNumQp_Min gets the value of LogNumQp_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumQp_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumQp_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumRdmaRc_Max sets the value of LogNumRdmaRc_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumRdmaRc_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumRdmaRc_Max", value)
}

// GetLogNumRdmaRc_Max gets the value of LogNumRdmaRc_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumRdmaRc_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumRdmaRc_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumRdmaRc_Min sets the value of LogNumRdmaRc_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumRdmaRc_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumRdmaRc_Min", value)
}

// GetLogNumRdmaRc_Min gets the value of LogNumRdmaRc_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumRdmaRc_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumRdmaRc_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumSrq_Max sets the value of LogNumSrq_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumSrq_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumSrq_Max", value)
}

// GetLogNumSrq_Max gets the value of LogNumSrq_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumSrq_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumSrq_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumSrq_Min sets the value of LogNumSrq_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumSrq_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumSrq_Min", value)
}

// GetLogNumSrq_Min gets the value of LogNumSrq_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumSrq_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumSrq_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumVlan_Max sets the value of LogNumVlan_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumVlan_Max(value uint32) (err error) {
	return instance.SetProperty("LogNumVlan_Max", value)
}

// GetLogNumVlan_Max gets the value of LogNumVlan_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumVlan_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumVlan_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogNumVlan_Min sets the value of LogNumVlan_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyLogNumVlan_Min(value uint32) (err error) {
	return instance.SetProperty("LogNumVlan_Min", value)
}

// GetLogNumVlan_Min gets the value of LogNumVlan_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyLogNumVlan_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogNumVlan_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumWorkingThreads_Max sets the value of MaximumWorkingThreads_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyMaximumWorkingThreads_Max(value uint32) (err error) {
	return instance.SetProperty("MaximumWorkingThreads_Max", value)
}

// GetMaximumWorkingThreads_Max gets the value of MaximumWorkingThreads_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyMaximumWorkingThreads_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumWorkingThreads_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumWorkingThreads_Min sets the value of MaximumWorkingThreads_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertyMaximumWorkingThreads_Min(value uint32) (err error) {
	return instance.SetProperty("MaximumWorkingThreads_Min", value)
}

// GetMaximumWorkingThreads_Min gets the value of MaximumWorkingThreads_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertyMaximumWorkingThreads_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumWorkingThreads_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSet4kMtu_Max sets the value of Set4kMtu_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertySet4kMtu_Max(value uint32) (err error) {
	return instance.SetProperty("Set4kMtu_Max", value)
}

// GetSet4kMtu_Max gets the value of Set4kMtu_Max for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertySet4kMtu_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("Set4kMtu_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSet4kMtu_Min sets the value of Set4kMtu_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) SetPropertySet4kMtu_Min(value uint32) (err error) {
	return instance.SetProperty("Set4kMtu_Min", value)
}

// GetSet4kMtu_Min gets the value of Set4kMtu_Min for the instance
func (instance *MLNX_DriverCoreCapabilities) GetPropertySet4kMtu_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("Set4kMtu_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
