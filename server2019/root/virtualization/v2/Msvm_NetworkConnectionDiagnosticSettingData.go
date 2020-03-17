// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_NetworkConnectionDiagnosticSettingData struct
type Msvm_NetworkConnectionDiagnosticSettingData struct {
	CIM_SettingData

	//
	IsolationId uint32

	//
	IsSender bool

	//
	PayloadSize uint32

	//
	ReceiverIP string

	//
	ReceiverMac string

	//
	SenderIP string

	//
	SequenceNumber uint32
}

// SetIsolationId sets the value of IsolationId for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyIsolationId(value uint32) (err error) {
	return instance.SetProperty("IsolationId", value)
}

// GetIsolationId gets the value of IsolationId for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyIsolationId() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsolationId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsSender sets the value of IsSender for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyIsSender(value bool) (err error) {
	return instance.SetProperty("IsSender", value)
}

// GetIsSender gets the value of IsSender for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyIsSender() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSender")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPayloadSize sets the value of PayloadSize for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyPayloadSize(value uint32) (err error) {
	return instance.SetProperty("PayloadSize", value)
}

// GetPayloadSize gets the value of PayloadSize for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PayloadSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiverIP sets the value of ReceiverIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyReceiverIP(value string) (err error) {
	return instance.SetProperty("ReceiverIP", value)
}

// GetReceiverIP gets the value of ReceiverIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyReceiverIP() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiverIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiverMac sets the value of ReceiverMac for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertyReceiverMac(value string) (err error) {
	return instance.SetProperty("ReceiverMac", value)
}

// GetReceiverMac gets the value of ReceiverMac for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertyReceiverMac() (value string, err error) {
	retValue, err := instance.GetProperty("ReceiverMac")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSenderIP sets the value of SenderIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertySenderIP(value string) (err error) {
	return instance.SetProperty("SenderIP", value)
}

// GetSenderIP gets the value of SenderIP for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertySenderIP() (value string, err error) {
	retValue, err := instance.GetProperty("SenderIP")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSequenceNumber sets the value of SequenceNumber for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) SetPropertySequenceNumber(value uint32) (err error) {
	return instance.SetProperty("SequenceNumber", value)
}

// GetSequenceNumber gets the value of SequenceNumber for the instance
func (instance *Msvm_NetworkConnectionDiagnosticSettingData) GetPropertySequenceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SequenceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
