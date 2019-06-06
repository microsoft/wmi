package wmi;


type WmiMethodResult struct {
	ReturnValue *WmiMethodParameter;
	OutParameters *[]WmiMethodParameters
}