// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// This class implements the swbemproperty class
// Documentation: https://docs.microsoft.com/en-us/windows/win32/wmisdk/swbemmethod

package cim

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/microsoft/wmi/pkg/errors"
	//"github.com/microsoft/wmi/go/wmi"
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
	MethodParams    map[string]*WmiMethodParam
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
		MethodParams:   map[string]*WmiMethodParam{},
	}, nil
}

func (c *WmiMethod) AddInParam(paramName string, paramValue interface{}) error {
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
	c.MethodParams[paramName], err = NewWmiMethodParam(paramName, paramValue, p, c.session)
	if err != nil {
		return err
	}

	return nil
}

func (c *WmiMethod) Execute() (int32, error) {
	outparams, err := c.classInstance.GetIDispatch().CallMethod("ExecMethod_", c.Name, c.inparameter)
	if err != nil {
		return 0, err
	}
	defer outparams.Clear()
	returnRaw, err := outparams.ToIDispatch().GetProperty("ReturnValue")
	if err != nil {
		return 0, err
	}
	defer returnRaw.Clear()
	outRaw, err := outparams.ToIDispatch().GetProperty("OutParameters")
	if err != nil {
		return 0, err
	}
	c.outparameterVar = outRaw
	c.outparameter = outRaw.ToIDispatch()

	return returnRaw.Value().(int32), nil
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
	// Close the MethodParams
	for _, k := range c.MethodParams {
		k.Close()
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
