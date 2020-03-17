// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000 struct
type Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000 struct {
	Win32_PerfFormattedData

	//
	Averageparticipantcommitresponsetime uint32

	//
	Averageparticipantprepareresponsetime uint32

	//
	CommitretrycountPersec uint32

	//
	FaultsreceivedcountPersec uint32

	//
	FaultssentcountPersec uint32

	//
	MessagesendfailuresPersec uint32

	//
	PreparedretrycountPersec uint32

	//
	PrepareretrycountPersec uint32

	//
	ReplayretrycountPersec uint32
}

// SetAverageparticipantcommitresponsetime sets the value of Averageparticipantcommitresponsetime for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyAverageparticipantcommitresponsetime(value uint32) (err error) {
	return instance.SetProperty("Averageparticipantcommitresponsetime", value)
}

// GetAverageparticipantcommitresponsetime gets the value of Averageparticipantcommitresponsetime for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyAverageparticipantcommitresponsetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("Averageparticipantcommitresponsetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAverageparticipantprepareresponsetime sets the value of Averageparticipantprepareresponsetime for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyAverageparticipantprepareresponsetime(value uint32) (err error) {
	return instance.SetProperty("Averageparticipantprepareresponsetime", value)
}

// GetAverageparticipantprepareresponsetime gets the value of Averageparticipantprepareresponsetime for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyAverageparticipantprepareresponsetime() (value uint32, err error) {
	retValue, err := instance.GetProperty("Averageparticipantprepareresponsetime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommitretrycountPersec sets the value of CommitretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyCommitretrycountPersec(value uint32) (err error) {
	return instance.SetProperty("CommitretrycountPersec", value)
}

// GetCommitretrycountPersec gets the value of CommitretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyCommitretrycountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("CommitretrycountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultsreceivedcountPersec sets the value of FaultsreceivedcountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyFaultsreceivedcountPersec(value uint32) (err error) {
	return instance.SetProperty("FaultsreceivedcountPersec", value)
}

// GetFaultsreceivedcountPersec gets the value of FaultsreceivedcountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyFaultsreceivedcountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultsreceivedcountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultssentcountPersec sets the value of FaultssentcountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyFaultssentcountPersec(value uint32) (err error) {
	return instance.SetProperty("FaultssentcountPersec", value)
}

// GetFaultssentcountPersec gets the value of FaultssentcountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyFaultssentcountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultssentcountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessagesendfailuresPersec sets the value of MessagesendfailuresPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyMessagesendfailuresPersec(value uint32) (err error) {
	return instance.SetProperty("MessagesendfailuresPersec", value)
}

// GetMessagesendfailuresPersec gets the value of MessagesendfailuresPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyMessagesendfailuresPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("MessagesendfailuresPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreparedretrycountPersec sets the value of PreparedretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyPreparedretrycountPersec(value uint32) (err error) {
	return instance.SetProperty("PreparedretrycountPersec", value)
}

// GetPreparedretrycountPersec gets the value of PreparedretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyPreparedretrycountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreparedretrycountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrepareretrycountPersec sets the value of PrepareretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyPrepareretrycountPersec(value uint32) (err error) {
	return instance.SetProperty("PrepareretrycountPersec", value)
}

// GetPrepareretrycountPersec gets the value of PrepareretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyPrepareretrycountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrepareretrycountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplayretrycountPersec sets the value of ReplayretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) SetPropertyReplayretrycountPersec(value uint32) (err error) {
	return instance.SetProperty("ReplayretrycountPersec", value)
}

// GetReplayretrycountPersec gets the value of ReplayretrycountPersec for the instance
func (instance *Win32_PerfFormattedData_MSDTCBridge4000_MSDTCBridge4000) GetPropertyReplayretrycountPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplayretrycountPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
