// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_BITS02 struct
type MDM_Policy_Result01_BITS02 struct {
	cim.WmiInstance

	//
	BandwidthThrottlingEndTime int32

	//
	BandwidthThrottlingStartTime int32

	//
	BandwidthThrottlingTransferRate int32

	//
	CostedNetworkBehaviorBackgroundPriority int32

	//
	CostedNetworkBehaviorForegroundPriority int32

	//
	InstanceID string

	//
	JobInactivityTimeout int32

	//
	ParentID string
}

// SetBandwidthThrottlingEndTime sets the value of BandwidthThrottlingEndTime for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyBandwidthThrottlingEndTime(value int32) (err error) {
	return instance.SetProperty("BandwidthThrottlingEndTime", value)
}

// GetBandwidthThrottlingEndTime gets the value of BandwidthThrottlingEndTime for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyBandwidthThrottlingEndTime() (value int32, err error) {
	retValue, err := instance.GetProperty("BandwidthThrottlingEndTime")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBandwidthThrottlingStartTime sets the value of BandwidthThrottlingStartTime for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyBandwidthThrottlingStartTime(value int32) (err error) {
	return instance.SetProperty("BandwidthThrottlingStartTime", value)
}

// GetBandwidthThrottlingStartTime gets the value of BandwidthThrottlingStartTime for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyBandwidthThrottlingStartTime() (value int32, err error) {
	retValue, err := instance.GetProperty("BandwidthThrottlingStartTime")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBandwidthThrottlingTransferRate sets the value of BandwidthThrottlingTransferRate for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyBandwidthThrottlingTransferRate(value int32) (err error) {
	return instance.SetProperty("BandwidthThrottlingTransferRate", value)
}

// GetBandwidthThrottlingTransferRate gets the value of BandwidthThrottlingTransferRate for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyBandwidthThrottlingTransferRate() (value int32, err error) {
	retValue, err := instance.GetProperty("BandwidthThrottlingTransferRate")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCostedNetworkBehaviorBackgroundPriority sets the value of CostedNetworkBehaviorBackgroundPriority for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyCostedNetworkBehaviorBackgroundPriority(value int32) (err error) {
	return instance.SetProperty("CostedNetworkBehaviorBackgroundPriority", value)
}

// GetCostedNetworkBehaviorBackgroundPriority gets the value of CostedNetworkBehaviorBackgroundPriority for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyCostedNetworkBehaviorBackgroundPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("CostedNetworkBehaviorBackgroundPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCostedNetworkBehaviorForegroundPriority sets the value of CostedNetworkBehaviorForegroundPriority for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyCostedNetworkBehaviorForegroundPriority(value int32) (err error) {
	return instance.SetProperty("CostedNetworkBehaviorForegroundPriority", value)
}

// GetCostedNetworkBehaviorForegroundPriority gets the value of CostedNetworkBehaviorForegroundPriority for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyCostedNetworkBehaviorForegroundPriority() (value int32, err error) {
	retValue, err := instance.GetProperty("CostedNetworkBehaviorForegroundPriority")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobInactivityTimeout sets the value of JobInactivityTimeout for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyJobInactivityTimeout(value int32) (err error) {
	return instance.SetProperty("JobInactivityTimeout", value)
}

// GetJobInactivityTimeout gets the value of JobInactivityTimeout for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyJobInactivityTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("JobInactivityTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_BITS02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_BITS02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
