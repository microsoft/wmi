// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_PhysicalConnector struct
type CIM_PhysicalConnector struct {
	CIM_PhysicalElement

	//
	ConnectorDescription string

	//
	ConnectorElectricalCharacteristics []PhysicalConnector_ConnectorElectricalCharacteristics

	//
	ConnectorGender PhysicalConnector_ConnectorGender

	//
	ConnectorLayout PhysicalConnector_ConnectorLayout

	//
	ConnectorPinout string

	//
	ConnectorType []PhysicalConnector_ConnectorType

	//
	NumPhysicalPins uint32

	//
	OtherElectricalCharacteristics []string

	//
	OtherTypeDescription string
}

// SetConnectorDescription sets the value of ConnectorDescription for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorDescription(value string) (err error) {
	return instance.SetProperty("ConnectorDescription", value)
}

// GetConnectorDescription gets the value of ConnectorDescription for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorDescription() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectorDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorElectricalCharacteristics sets the value of ConnectorElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorElectricalCharacteristics(value []PhysicalConnector_ConnectorElectricalCharacteristics) (err error) {
	return instance.SetProperty("ConnectorElectricalCharacteristics", value)
}

// GetConnectorElectricalCharacteristics gets the value of ConnectorElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorElectricalCharacteristics() (value []PhysicalConnector_ConnectorElectricalCharacteristics, err error) {
	retValue, err := instance.GetProperty("ConnectorElectricalCharacteristics")
	if err != nil {
		return
	}
	value, ok := retValue.([]PhysicalConnector_ConnectorElectricalCharacteristics)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorGender sets the value of ConnectorGender for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorGender(value PhysicalConnector_ConnectorGender) (err error) {
	return instance.SetProperty("ConnectorGender", value)
}

// GetConnectorGender gets the value of ConnectorGender for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorGender() (value PhysicalConnector_ConnectorGender, err error) {
	retValue, err := instance.GetProperty("ConnectorGender")
	if err != nil {
		return
	}
	value, ok := retValue.(PhysicalConnector_ConnectorGender)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorLayout sets the value of ConnectorLayout for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorLayout(value PhysicalConnector_ConnectorLayout) (err error) {
	return instance.SetProperty("ConnectorLayout", value)
}

// GetConnectorLayout gets the value of ConnectorLayout for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorLayout() (value PhysicalConnector_ConnectorLayout, err error) {
	retValue, err := instance.GetProperty("ConnectorLayout")
	if err != nil {
		return
	}
	value, ok := retValue.(PhysicalConnector_ConnectorLayout)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorPinout sets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorPinout(value string) (err error) {
	return instance.SetProperty("ConnectorPinout", value)
}

// GetConnectorPinout gets the value of ConnectorPinout for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorPinout() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectorPinout")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorType sets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) SetPropertyConnectorType(value []PhysicalConnector_ConnectorType) (err error) {
	return instance.SetProperty("ConnectorType", value)
}

// GetConnectorType gets the value of ConnectorType for the instance
func (instance *CIM_PhysicalConnector) GetPropertyConnectorType() (value []PhysicalConnector_ConnectorType, err error) {
	retValue, err := instance.GetProperty("ConnectorType")
	if err != nil {
		return
	}
	value, ok := retValue.([]PhysicalConnector_ConnectorType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumPhysicalPins sets the value of NumPhysicalPins for the instance
func (instance *CIM_PhysicalConnector) SetPropertyNumPhysicalPins(value uint32) (err error) {
	return instance.SetProperty("NumPhysicalPins", value)
}

// GetNumPhysicalPins gets the value of NumPhysicalPins for the instance
func (instance *CIM_PhysicalConnector) GetPropertyNumPhysicalPins() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumPhysicalPins")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherElectricalCharacteristics sets the value of OtherElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) SetPropertyOtherElectricalCharacteristics(value []string) (err error) {
	return instance.SetProperty("OtherElectricalCharacteristics", value)
}

// GetOtherElectricalCharacteristics gets the value of OtherElectricalCharacteristics for the instance
func (instance *CIM_PhysicalConnector) GetPropertyOtherElectricalCharacteristics() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherElectricalCharacteristics")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherTypeDescription sets the value of OtherTypeDescription for the instance
func (instance *CIM_PhysicalConnector) SetPropertyOtherTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherTypeDescription", value)
}

// GetOtherTypeDescription gets the value of OtherTypeDescription for the instance
func (instance *CIM_PhysicalConnector) GetPropertyOtherTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTypeDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
