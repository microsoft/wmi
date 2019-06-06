package wmi;

type WmiMethodParameter struct {
	Name() string;
	Value() string;
	Type() WmiType;
}