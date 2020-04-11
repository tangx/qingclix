# qingclix
青云常用操作命令行


## 使用到的库

+ `struct` 转 `url.Values`: `github.com/tangx/go-querystring/query`

```go
// The URL parameter name defaults to the struct field name but can be
// specified in the struct field's tag value.  The "url" key in the struct
// field's tag value is the key name, followed by an optional comma and
// options.  For example:
//
// 	// Field is ignored by this package.
// 	Field int `url:"-"`
//
// 	// Field appears as URL parameter "myName".
// 	Field int `url:"myName"`
//
// 	// Field appears as URL parameter "myName" and the field is omitted if
// 	// its value is empty
// 	Field int `url:"myName,omitempty"`
//
// 	// Field appears as URL parameter "Field" (the default), but the field
// 	// is skipped if empty.  Note the leading comma.
// 	Field int `url:",omitempty"`

// 通过 url 设置，可以将 驼峰 camels 转换成 snake_case
type VolumeRequest2 struct {
	Size       string          `url:"size,omitempty"`
	VolumeName string          `url:"volume_name,omitempty"`
	VolumeType string          `url:"volume_type,omitempty"`
	Zone       string          `url:"zone,omitempty"`
	Months     string          `url:"months,omitempty"`
	AutoRenew  string          `url:"auto_renew,omitempty"`
	Contract   ContractRequest `url:"contract,omitempty"`
}

```