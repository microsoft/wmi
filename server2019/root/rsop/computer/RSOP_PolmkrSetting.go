// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_PolmkrSetting struct
type RSOP_PolmkrSetting struct {
	RSOP_PolicySetting

	//
	polmkrBaseCseGuid string

	//
	polmkrBaseGpeGuid string

	//
	polmkrBaseGpoDisplayName string

	//
	polmkrBaseGpoGuid string

	//
	polmkrBaseHash string

	//
	polmkrBaseInstanceXml string

	//
	polmkrBaseKeyValues []string
}

// SetpolmkrBaseCseGuid sets the value of polmkrBaseCseGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseCseGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseCseGuid", value)
}

// GetpolmkrBaseCseGuid gets the value of polmkrBaseCseGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseCseGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseCseGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseGpeGuid sets the value of polmkrBaseGpeGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpeGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpeGuid", value)
}

// GetpolmkrBaseGpeGuid gets the value of polmkrBaseGpeGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpeGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpeGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseGpoDisplayName sets the value of polmkrBaseGpoDisplayName for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpoDisplayName(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpoDisplayName", value)
}

// GetpolmkrBaseGpoDisplayName gets the value of polmkrBaseGpoDisplayName for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpoDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpoDisplayName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseGpoGuid sets the value of polmkrBaseGpoGuid for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseGpoGuid(value string) (err error) {
	return instance.SetProperty("polmkrBaseGpoGuid", value)
}

// GetpolmkrBaseGpoGuid gets the value of polmkrBaseGpoGuid for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseGpoGuid() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseGpoGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseHash sets the value of polmkrBaseHash for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseHash(value string) (err error) {
	return instance.SetProperty("polmkrBaseHash", value)
}

// GetpolmkrBaseHash gets the value of polmkrBaseHash for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseHash() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseHash")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseInstanceXml sets the value of polmkrBaseInstanceXml for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseInstanceXml(value string) (err error) {
	return instance.SetProperty("polmkrBaseInstanceXml", value)
}

// GetpolmkrBaseInstanceXml gets the value of polmkrBaseInstanceXml for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseInstanceXml() (value string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseInstanceXml")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolmkrBaseKeyValues sets the value of polmkrBaseKeyValues for the instance
func (instance *RSOP_PolmkrSetting) SetPropertypolmkrBaseKeyValues(value []string) (err error) {
	return instance.SetProperty("polmkrBaseKeyValues", value)
}

// GetpolmkrBaseKeyValues gets the value of polmkrBaseKeyValues for the instance
func (instance *RSOP_PolmkrSetting) GetPropertypolmkrBaseKeyValues() (value []string, err error) {
	retValue, err := instance.GetProperty("polmkrBaseKeyValues")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
