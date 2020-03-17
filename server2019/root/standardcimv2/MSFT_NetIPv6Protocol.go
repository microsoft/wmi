// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetIPv6Protocol struct
type MSFT_NetIPv6Protocol struct {
	MSFT_NetBaseIPProtocol

	//
	MaxDadAttempts uint32

	//
	MaxPreferredLifetime string

	//
	MaxRandomTime string

	//
	MaxValidLifetime string

	//
	RegenerateTime string

	//
	UseTemporaryAddresses uint32
}

// SetMaxDadAttempts sets the value of MaxDadAttempts for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxDadAttempts(value uint32) (err error) {
	return instance.SetProperty("MaxDadAttempts", value)
}

// GetMaxDadAttempts gets the value of MaxDadAttempts for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxDadAttempts() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDadAttempts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxPreferredLifetime sets the value of MaxPreferredLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxPreferredLifetime(value string) (err error) {
	return instance.SetProperty("MaxPreferredLifetime", value)
}

// GetMaxPreferredLifetime gets the value of MaxPreferredLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxPreferredLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxPreferredLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxRandomTime sets the value of MaxRandomTime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxRandomTime(value string) (err error) {
	return instance.SetProperty("MaxRandomTime", value)
}

// GetMaxRandomTime gets the value of MaxRandomTime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxRandomTime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxRandomTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxValidLifetime sets the value of MaxValidLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyMaxValidLifetime(value string) (err error) {
	return instance.SetProperty("MaxValidLifetime", value)
}

// GetMaxValidLifetime gets the value of MaxValidLifetime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyMaxValidLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("MaxValidLifetime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegenerateTime sets the value of RegenerateTime for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyRegenerateTime(value string) (err error) {
	return instance.SetProperty("RegenerateTime", value)
}

// GetRegenerateTime gets the value of RegenerateTime for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyRegenerateTime() (value string, err error) {
	retValue, err := instance.GetProperty("RegenerateTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseTemporaryAddresses sets the value of UseTemporaryAddresses for the instance
func (instance *MSFT_NetIPv6Protocol) SetPropertyUseTemporaryAddresses(value uint32) (err error) {
	return instance.SetProperty("UseTemporaryAddresses", value)
}

// GetUseTemporaryAddresses gets the value of UseTemporaryAddresses for the instance
func (instance *MSFT_NetIPv6Protocol) GetPropertyUseTemporaryAddresses() (value uint32, err error) {
	retValue, err := instance.GetProperty("UseTemporaryAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
