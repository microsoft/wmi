// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_PotsModem struct
type CIM_PotsModem struct {
	CIM_LogicalDevice

	//
	AnswerMode uint16

	//
	CompressionInfo uint16

	//
	CountriesSupported []string

	//
	CountrySelected string

	//
	CurrentPasswords []string

	//
	DialType uint16

	//
	ErrorControlInfo uint16

	//
	InactivityTimeout uint32

	//
	MaxBaudRateToPhone uint32

	//
	MaxBaudRateToSerialPort uint32

	//
	MaxNumberOfPasswords uint16

	//
	ModulationScheme uint16

	//
	RingsBeforeAnswer uint8

	//
	SpeakerVolumeInfo uint16

	//
	SupportsCallback bool

	//
	SupportsSynchronousConnect bool

	//
	TimeOfLastReset string
}

// SetAnswerMode sets the value of AnswerMode for the instance
func (instance *CIM_PotsModem) SetPropertyAnswerMode(value uint16) (err error) {
	return instance.SetProperty("AnswerMode", value)
}

// GetAnswerMode gets the value of AnswerMode for the instance
func (instance *CIM_PotsModem) GetPropertyAnswerMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("AnswerMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCompressionInfo sets the value of CompressionInfo for the instance
func (instance *CIM_PotsModem) SetPropertyCompressionInfo(value uint16) (err error) {
	return instance.SetProperty("CompressionInfo", value)
}

// GetCompressionInfo gets the value of CompressionInfo for the instance
func (instance *CIM_PotsModem) GetPropertyCompressionInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("CompressionInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCountriesSupported sets the value of CountriesSupported for the instance
func (instance *CIM_PotsModem) SetPropertyCountriesSupported(value []string) (err error) {
	return instance.SetProperty("CountriesSupported", value)
}

// GetCountriesSupported gets the value of CountriesSupported for the instance
func (instance *CIM_PotsModem) GetPropertyCountriesSupported() (value []string, err error) {
	retValue, err := instance.GetProperty("CountriesSupported")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCountrySelected sets the value of CountrySelected for the instance
func (instance *CIM_PotsModem) SetPropertyCountrySelected(value string) (err error) {
	return instance.SetProperty("CountrySelected", value)
}

// GetCountrySelected gets the value of CountrySelected for the instance
func (instance *CIM_PotsModem) GetPropertyCountrySelected() (value string, err error) {
	retValue, err := instance.GetProperty("CountrySelected")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentPasswords sets the value of CurrentPasswords for the instance
func (instance *CIM_PotsModem) SetPropertyCurrentPasswords(value []string) (err error) {
	return instance.SetProperty("CurrentPasswords", value)
}

// GetCurrentPasswords gets the value of CurrentPasswords for the instance
func (instance *CIM_PotsModem) GetPropertyCurrentPasswords() (value []string, err error) {
	retValue, err := instance.GetProperty("CurrentPasswords")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDialType sets the value of DialType for the instance
func (instance *CIM_PotsModem) SetPropertyDialType(value uint16) (err error) {
	return instance.SetProperty("DialType", value)
}

// GetDialType gets the value of DialType for the instance
func (instance *CIM_PotsModem) GetPropertyDialType() (value uint16, err error) {
	retValue, err := instance.GetProperty("DialType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetErrorControlInfo sets the value of ErrorControlInfo for the instance
func (instance *CIM_PotsModem) SetPropertyErrorControlInfo(value uint16) (err error) {
	return instance.SetProperty("ErrorControlInfo", value)
}

// GetErrorControlInfo gets the value of ErrorControlInfo for the instance
func (instance *CIM_PotsModem) GetPropertyErrorControlInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("ErrorControlInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInactivityTimeout sets the value of InactivityTimeout for the instance
func (instance *CIM_PotsModem) SetPropertyInactivityTimeout(value uint32) (err error) {
	return instance.SetProperty("InactivityTimeout", value)
}

// GetInactivityTimeout gets the value of InactivityTimeout for the instance
func (instance *CIM_PotsModem) GetPropertyInactivityTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("InactivityTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxBaudRateToPhone sets the value of MaxBaudRateToPhone for the instance
func (instance *CIM_PotsModem) SetPropertyMaxBaudRateToPhone(value uint32) (err error) {
	return instance.SetProperty("MaxBaudRateToPhone", value)
}

// GetMaxBaudRateToPhone gets the value of MaxBaudRateToPhone for the instance
func (instance *CIM_PotsModem) GetPropertyMaxBaudRateToPhone() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxBaudRateToPhone")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxBaudRateToSerialPort sets the value of MaxBaudRateToSerialPort for the instance
func (instance *CIM_PotsModem) SetPropertyMaxBaudRateToSerialPort(value uint32) (err error) {
	return instance.SetProperty("MaxBaudRateToSerialPort", value)
}

// GetMaxBaudRateToSerialPort gets the value of MaxBaudRateToSerialPort for the instance
func (instance *CIM_PotsModem) GetPropertyMaxBaudRateToSerialPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxBaudRateToSerialPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxNumberOfPasswords sets the value of MaxNumberOfPasswords for the instance
func (instance *CIM_PotsModem) SetPropertyMaxNumberOfPasswords(value uint16) (err error) {
	return instance.SetProperty("MaxNumberOfPasswords", value)
}

// GetMaxNumberOfPasswords gets the value of MaxNumberOfPasswords for the instance
func (instance *CIM_PotsModem) GetPropertyMaxNumberOfPasswords() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxNumberOfPasswords")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModulationScheme sets the value of ModulationScheme for the instance
func (instance *CIM_PotsModem) SetPropertyModulationScheme(value uint16) (err error) {
	return instance.SetProperty("ModulationScheme", value)
}

// GetModulationScheme gets the value of ModulationScheme for the instance
func (instance *CIM_PotsModem) GetPropertyModulationScheme() (value uint16, err error) {
	retValue, err := instance.GetProperty("ModulationScheme")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRingsBeforeAnswer sets the value of RingsBeforeAnswer for the instance
func (instance *CIM_PotsModem) SetPropertyRingsBeforeAnswer(value uint8) (err error) {
	return instance.SetProperty("RingsBeforeAnswer", value)
}

// GetRingsBeforeAnswer gets the value of RingsBeforeAnswer for the instance
func (instance *CIM_PotsModem) GetPropertyRingsBeforeAnswer() (value uint8, err error) {
	retValue, err := instance.GetProperty("RingsBeforeAnswer")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpeakerVolumeInfo sets the value of SpeakerVolumeInfo for the instance
func (instance *CIM_PotsModem) SetPropertySpeakerVolumeInfo(value uint16) (err error) {
	return instance.SetProperty("SpeakerVolumeInfo", value)
}

// GetSpeakerVolumeInfo gets the value of SpeakerVolumeInfo for the instance
func (instance *CIM_PotsModem) GetPropertySpeakerVolumeInfo() (value uint16, err error) {
	retValue, err := instance.GetProperty("SpeakerVolumeInfo")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsCallback sets the value of SupportsCallback for the instance
func (instance *CIM_PotsModem) SetPropertySupportsCallback(value bool) (err error) {
	return instance.SetProperty("SupportsCallback", value)
}

// GetSupportsCallback gets the value of SupportsCallback for the instance
func (instance *CIM_PotsModem) GetPropertySupportsCallback() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCallback")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsSynchronousConnect sets the value of SupportsSynchronousConnect for the instance
func (instance *CIM_PotsModem) SetPropertySupportsSynchronousConnect(value bool) (err error) {
	return instance.SetProperty("SupportsSynchronousConnect", value)
}

// GetSupportsSynchronousConnect gets the value of SupportsSynchronousConnect for the instance
func (instance *CIM_PotsModem) GetPropertySupportsSynchronousConnect() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsSynchronousConnect")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeOfLastReset sets the value of TimeOfLastReset for the instance
func (instance *CIM_PotsModem) SetPropertyTimeOfLastReset(value string) (err error) {
	return instance.SetProperty("TimeOfLastReset", value)
}

// GetTimeOfLastReset gets the value of TimeOfLastReset for the instance
func (instance *CIM_PotsModem) GetPropertyTimeOfLastReset() (value string, err error) {
	retValue, err := instance.GetProperty("TimeOfLastReset")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
