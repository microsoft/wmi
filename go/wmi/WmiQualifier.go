package wmi;

type WmiQualifier interface {
	Name() string;
	Value() string;
}