// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_Net6to4Configuration struct
type MSFT_Net6to4Configuration struct {
	MSFT_NetSettingData

	//
	AutoSharing uint32

	//
	PolicyStore string

	//
	RelayName string

	//
	RelayState uint32

	//
	ResolutionInterval uint32

	//
	State uint32
}

// SetAutoSharing sets the value of AutoSharing for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyAutoSharing(value uint32) (err error) {
	return instance.SetProperty("AutoSharing", value)
}

// GetAutoSharing gets the value of AutoSharing for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyAutoSharing() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoSharing")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", value)
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRelayName sets the value of RelayName for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyRelayName(value string) (err error) {
	return instance.SetProperty("RelayName", value)
}

// GetRelayName gets the value of RelayName for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyRelayName() (value string, err error) {
	retValue, err := instance.GetProperty("RelayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRelayState sets the value of RelayState for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyRelayState(value uint32) (err error) {
	return instance.SetProperty("RelayState", value)
}

// GetRelayState gets the value of RelayState for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyRelayState() (value uint32, err error) {
	retValue, err := instance.GetProperty("RelayState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolutionInterval sets the value of ResolutionInterval for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyResolutionInterval(value uint32) (err error) {
	return instance.SetProperty("ResolutionInterval", value)
}

// GetResolutionInterval gets the value of ResolutionInterval for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyResolutionInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResolutionInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_Net6to4Configuration) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_Net6to4Configuration) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="AutoSharing" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="RelayName" type="bool "></param>
// <param name="RelayState" type="bool "></param>
// <param name="ResolutionInterval" type="bool "></param>
// <param name="State" type="bool "></param>

// <param name="OutputObject" type="MSFT_Net6to4Configuration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Net6to4Configuration) Reset( /* IN */ State bool,
	/* IN */ AutoSharing bool,
	/* IN */ RelayName bool,
	/* IN */ RelayState bool,
	/* IN */ ResolutionInterval bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_Net6to4Configuration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", State, AutoSharing, RelayName, RelayState, ResolutionInterval, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
