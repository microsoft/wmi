// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This class implements the swbemproperty class
// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemmethod

package cim

import (
	"log"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/microsoft/wmi/pkg/errors"
)

type WmiMethod struct {
	Name string
	// Reference
	session *WmiSession
	// Reference
	classInstance   *WmiInstance
	inparameterVar  *ole.VARIANT
	inparameter     *ole.IDispatch
	outparameterVar *ole.VARIANT
	outparameter    *ole.IDispatch
	method          *ole.IDispatch
	methodVar       *ole.VARIANT
}

type WmiMethodResult struct {
	ReturnValue     int32
	OutMethodParams map[string]*WmiMethodParam
}

// NewWmiMethod
func NewWmiMethod(methodName string, instance *WmiInstance) (*WmiMethod, error) {
	iDispatchInstance := instance.GetIDispatch()
	if iDispatchInstance == nil {
		return nil, errors.Wrapf(errors.InvalidInput, "InvalidInstance")
	}
	rawResult, err := iDispatchInstance.GetProperty("Methods_")
	if err != nil {
		return nil, err
	}
	defer rawResult.Clear()
	// Retrive the method
	rawMethod, err := rawResult.ToIDispatch().CallMethod("Item", methodName)
	if err != nil {
		return nil, err
	}

	inparamsRaw, err := rawMethod.ToIDispatch().GetProperty("InParameters")
	if err != nil {
		return nil, err
	}
	defer inparamsRaw.Clear()

	inparams, err := oleutil.CallMethod(inparamsRaw.ToIDispatch(), "SpawnInstance_")
	if err != nil {
		return nil, err
	}

	return &WmiMethod{
		Name:           methodName,
		classInstance:  instance,
		session:        instance.GetSession(),
		inparameter:    inparams.ToIDispatch(),
		inparameterVar: inparams,
		method:         rawMethod.ToIDispatch(),
		methodVar:      rawMethod,
	}, nil
}

func (c *WmiMethod) addInParam(paramName string, paramValue interface{}) error {
	rawProperties, err := c.inparameter.GetProperty("Properties_")
	if err != nil {
		return err
	}
	defer rawProperties.Clear()
	rawProperty, err := rawProperties.ToIDispatch().CallMethod("Item", paramName)
	if err != nil {
		return err
	}
	defer rawProperty.Clear()

	p, err := rawProperty.ToIDispatch().PutProperty("Value", paramValue)
	if err != nil {
		return err
	}
	defer p.Clear()
	return nil
}

func (c *WmiMethod) Execute(inParam, outParam WmiMethodParamCollection) (result *WmiMethodResult, err error) {
	log.Printf("[WMI] - Executing Method [%s]\n", c.Name)
	for _, inp := range inParam {
		// 	log.Printf("InParam [%s]=>[%+v]\n", inp.Name, inp.Value)
		c.addInParam(inp.Name, inp.Value)
	}

	result = &WmiMethodResult{
		OutMethodParams: map[string]*WmiMethodParam{},
	}
	outparams, err := c.classInstance.GetIDispatch().CallMethod("ExecMethod_", c.Name, c.inparameter)
	if err != nil {
		return
	}
	defer outparams.Clear()
	returnRaw, err := outparams.ToIDispatch().GetProperty("ReturnValue")
	if err != nil {
		return
	}
	defer returnRaw.Clear()
	result.ReturnValue = returnRaw.Value().(int32)
	log.Printf("[WMI] - Return [%d] ", result.ReturnValue)

	for _, outp := range outParam {
		returnRaw, err1 := outparams.ToIDispatch().GetProperty(outp.Name)
		if err1 != nil {
			err = err1
			return
		}
		defer returnRaw.Clear()

		value, err1 := GetVariantValue(returnRaw)
		if err1 != nil {
			err = err1
			return
		}
		// log.Printf("OutParam [%s]=> [%+v]\n", outp.Name, value)

		result.OutMethodParams[outp.Name] = NewWmiMethodParam(outp.Name, value)
	}
	return
}

func (c *WmiMethod) Close() error {
	if c.inparameterVar != nil {
		c.inparameterVar.Clear()
		c.inparameterVar = nil
	}
	if c.inparameter != nil {
		c.inparameter.Release()
		c.inparameter = nil
	}
	if c.outparameterVar != nil {
		c.outparameterVar.Clear()
		c.outparameterVar = nil
	}
	if c.outparameter != nil {
		c.outparameter.Release()
		c.outparameter = nil
	}

	return nil
}

type WmiMethodCollection []*WmiMethod

func (c *WmiMethodCollection) Close() error {
	var err error
	for _, p := range *c {
		err = p.Close()
	}
	return err
}
