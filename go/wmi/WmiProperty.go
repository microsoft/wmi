package wmi;

type WmiType string;

const (
	Int WmiType = "int";
)

type WmiPropertyFlags int;

const(
	WmiPropertyFlags None = 0,
    WmiPropertyFlags    Class = 1,
    WmiPropertyFlags    Method = 2,
    WmiPropertyFlags    Property = 4,
    WmiPropertyFlags    Parameter = 8,
	WmiPropertyFlags    Association = 16,
	WmiPropertyFlags    Indication = 32,
	WmiPropertyFlags   Reference = 64,
	WmiPropertyFlags    Any = 127,
	WmiPropertyFlags   EnableOverride = 128,
	WmiPropertyFlags   DisableOverride = 256,
	WmiPropertyFlags    Restricted = 512,
	WmiPropertyFlags    ToSubclass = 1024,
	WmiPropertyFlags    Translatable = 2048,
	WmiPropertyFlags    Key = 4096,
	WmiPropertyFlags    In = 8192,
	WmiPropertyFlags   Out = 16384,
	WmiPropertyFlags    Required = 32768,
	WmiPropertyFlags    Static = 65536,
	WmiPropertyFlags   Abstract = 131072,
	WmiPropertyFlags    Terminal = 262144,
	WmiPropertyFlags    Expensive = 524288,
	WmiPropertyFlags   Stream = 1048576,
	WmiPropertyFlags    ReadOnly = 2097152,
	WmiPropertyFlags   NotModified = 33554432,
	WmiPropertyFlags    NullValue = 536870912,
	WmiPropertyFlags    Borrow = 1073741824,
	//WmiPropertyFlags Adopt = 2147483648,
)

type WmiProperty interface {
	Name() string;
	Value() string;
	Type() WmiType;
	Flags() WmiPropertyFlags;
}