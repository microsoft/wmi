// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_CacheMemory struct
type CIM_CacheMemory struct {
	CIM_Memory

	//
	Associativity uint16

	//
	CacheType uint16

	//
	FlushTimer uint32

	//
	Level uint16

	//
	LineSize uint32

	//
	ReadPolicy uint16

	//
	ReplacementPolicy uint16

	//
	WritePolicy uint16
}

// SetAssociativity sets the value of Associativity for the instance
func (instance *CIM_CacheMemory) SetPropertyAssociativity(value uint16) (err error) {
	return instance.SetProperty("Associativity", value)
}

// GetAssociativity gets the value of Associativity for the instance
func (instance *CIM_CacheMemory) GetPropertyAssociativity() (value uint16, err error) {
	retValue, err := instance.GetProperty("Associativity")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCacheType sets the value of CacheType for the instance
func (instance *CIM_CacheMemory) SetPropertyCacheType(value uint16) (err error) {
	return instance.SetProperty("CacheType", value)
}

// GetCacheType gets the value of CacheType for the instance
func (instance *CIM_CacheMemory) GetPropertyCacheType() (value uint16, err error) {
	retValue, err := instance.GetProperty("CacheType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFlushTimer sets the value of FlushTimer for the instance
func (instance *CIM_CacheMemory) SetPropertyFlushTimer(value uint32) (err error) {
	return instance.SetProperty("FlushTimer", value)
}

// GetFlushTimer gets the value of FlushTimer for the instance
func (instance *CIM_CacheMemory) GetPropertyFlushTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("FlushTimer")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLevel sets the value of Level for the instance
func (instance *CIM_CacheMemory) SetPropertyLevel(value uint16) (err error) {
	return instance.SetProperty("Level", value)
}

// GetLevel gets the value of Level for the instance
func (instance *CIM_CacheMemory) GetPropertyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("Level")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLineSize sets the value of LineSize for the instance
func (instance *CIM_CacheMemory) SetPropertyLineSize(value uint32) (err error) {
	return instance.SetProperty("LineSize", value)
}

// GetLineSize gets the value of LineSize for the instance
func (instance *CIM_CacheMemory) GetPropertyLineSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("LineSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadPolicy sets the value of ReadPolicy for the instance
func (instance *CIM_CacheMemory) SetPropertyReadPolicy(value uint16) (err error) {
	return instance.SetProperty("ReadPolicy", value)
}

// GetReadPolicy gets the value of ReadPolicy for the instance
func (instance *CIM_CacheMemory) GetPropertyReadPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReadPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplacementPolicy sets the value of ReplacementPolicy for the instance
func (instance *CIM_CacheMemory) SetPropertyReplacementPolicy(value uint16) (err error) {
	return instance.SetProperty("ReplacementPolicy", value)
}

// GetReplacementPolicy gets the value of ReplacementPolicy for the instance
func (instance *CIM_CacheMemory) GetPropertyReplacementPolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplacementPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWritePolicy sets the value of WritePolicy for the instance
func (instance *CIM_CacheMemory) SetPropertyWritePolicy(value uint16) (err error) {
	return instance.SetProperty("WritePolicy", value)
}

// GetWritePolicy gets the value of WritePolicy for the instance
func (instance *CIM_CacheMemory) GetPropertyWritePolicy() (value uint16, err error) {
	retValue, err := instance.GetProperty("WritePolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
