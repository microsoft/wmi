// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_NetAdapterGenStatistics struct
type MLNX_NetAdapterGenStatistics struct {
	MLNX_NetAdapterStatistics

	//
	ifHCInBroadcastOctets uint64

	//
	ifHCInBroadcastPkts uint64

	//
	ifHCInMulticastOctets uint64

	//
	ifHCInMulticastPkts uint64

	//
	ifHCInOctets uint64

	//
	ifHCInUcastOctets uint64

	//
	ifHCInUcastPkts uint64

	//
	ifHCOutBroadcastOctets uint64

	//
	ifHCOutBroadcastPkts uint64

	//
	ifHCOutMulticastOctets uint64

	//
	ifHCOutMulticastPkts uint64

	//
	ifHCOutOctets uint64

	//
	ifHCOutUcastOctets uint64

	//
	ifHCOutUcastPkts uint64

	//
	ifHInCrcError uint64

	//
	ifHInCrcOverRun uint64

	//
	ifInDiscards uint64

	//
	ifInErrors uint64

	//
	ifOutDiscards uint64

	//
	ifOutErrors uint64
}

// SetifHCInBroadcastOctets sets the value of ifHCInBroadcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInBroadcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInBroadcastOctets", value)
}

// GetifHCInBroadcastOctets gets the value of ifHCInBroadcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInBroadcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInBroadcastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInBroadcastPkts sets the value of ifHCInBroadcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInBroadcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInBroadcastPkts", value)
}

// GetifHCInBroadcastPkts gets the value of ifHCInBroadcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInBroadcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInBroadcastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInMulticastOctets sets the value of ifHCInMulticastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInMulticastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInMulticastOctets", value)
}

// GetifHCInMulticastOctets gets the value of ifHCInMulticastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInMulticastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInMulticastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInMulticastPkts sets the value of ifHCInMulticastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInMulticastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInMulticastPkts", value)
}

// GetifHCInMulticastPkts gets the value of ifHCInMulticastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInMulticastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInMulticastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInOctets sets the value of ifHCInOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInOctets", value)
}

// GetifHCInOctets gets the value of ifHCInOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInUcastOctets sets the value of ifHCInUcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInUcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCInUcastOctets", value)
}

// GetifHCInUcastOctets gets the value of ifHCInUcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInUcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInUcastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCInUcastPkts sets the value of ifHCInUcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCInUcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCInUcastPkts", value)
}

// GetifHCInUcastPkts gets the value of ifHCInUcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCInUcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCInUcastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutBroadcastOctets sets the value of ifHCOutBroadcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutBroadcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutBroadcastOctets", value)
}

// GetifHCOutBroadcastOctets gets the value of ifHCOutBroadcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutBroadcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutBroadcastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutBroadcastPkts sets the value of ifHCOutBroadcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutBroadcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutBroadcastPkts", value)
}

// GetifHCOutBroadcastPkts gets the value of ifHCOutBroadcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutBroadcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutBroadcastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutMulticastOctets sets the value of ifHCOutMulticastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutMulticastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutMulticastOctets", value)
}

// GetifHCOutMulticastOctets gets the value of ifHCOutMulticastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutMulticastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutMulticastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutMulticastPkts sets the value of ifHCOutMulticastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutMulticastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutMulticastPkts", value)
}

// GetifHCOutMulticastPkts gets the value of ifHCOutMulticastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutMulticastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutMulticastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutOctets sets the value of ifHCOutOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutOctets", value)
}

// GetifHCOutOctets gets the value of ifHCOutOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutUcastOctets sets the value of ifHCOutUcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutUcastOctets(value uint64) (err error) {
	return instance.SetProperty("ifHCOutUcastOctets", value)
}

// GetifHCOutUcastOctets gets the value of ifHCOutUcastOctets for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutUcastOctets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutUcastOctets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHCOutUcastPkts sets the value of ifHCOutUcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHCOutUcastPkts(value uint64) (err error) {
	return instance.SetProperty("ifHCOutUcastPkts", value)
}

// GetifHCOutUcastPkts gets the value of ifHCOutUcastPkts for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHCOutUcastPkts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHCOutUcastPkts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHInCrcError sets the value of ifHInCrcError for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHInCrcError(value uint64) (err error) {
	return instance.SetProperty("ifHInCrcError", value)
}

// GetifHInCrcError gets the value of ifHInCrcError for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHInCrcError() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHInCrcError")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifHInCrcOverRun sets the value of ifHInCrcOverRun for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifHInCrcOverRun(value uint64) (err error) {
	return instance.SetProperty("ifHInCrcOverRun", value)
}

// GetifHInCrcOverRun gets the value of ifHInCrcOverRun for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifHInCrcOverRun() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifHInCrcOverRun")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifInDiscards sets the value of ifInDiscards for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifInDiscards(value uint64) (err error) {
	return instance.SetProperty("ifInDiscards", value)
}

// GetifInDiscards gets the value of ifInDiscards for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifInDiscards() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifInDiscards")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifInErrors sets the value of ifInErrors for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifInErrors(value uint64) (err error) {
	return instance.SetProperty("ifInErrors", value)
}

// GetifInErrors gets the value of ifInErrors for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifInErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifInErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifOutDiscards sets the value of ifOutDiscards for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifOutDiscards(value uint64) (err error) {
	return instance.SetProperty("ifOutDiscards", value)
}

// GetifOutDiscards gets the value of ifOutDiscards for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifOutDiscards() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifOutDiscards")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetifOutErrors sets the value of ifOutErrors for the instance
func (instance *MLNX_NetAdapterGenStatistics) SetPropertyifOutErrors(value uint64) (err error) {
	return instance.SetProperty("ifOutErrors", value)
}

// GetifOutErrors gets the value of ifOutErrors for the instance
func (instance *MLNX_NetAdapterGenStatistics) GetPropertyifOutErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ifOutErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
