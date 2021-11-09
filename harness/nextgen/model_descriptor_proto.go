/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type DescriptorProto struct {
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
	Options *MessageOptions `json:"options,omitempty"`
	SerializedSize int32 `json:"serializedSize,omitempty"`
	ParserForType *ParserDescriptorProto `json:"parserForType,omitempty"`
	DefaultInstanceForType *DescriptorProto `json:"defaultInstanceForType,omitempty"`
	ExtensionRangeList []ExtensionRange `json:"extensionRangeList,omitempty"`
	ReservedRangeList []ReservedRange `json:"reservedRangeList,omitempty"`
	ReservedNameList []string `json:"reservedNameList,omitempty"`
	OneofDeclCount int32 `json:"oneofDeclCount,omitempty"`
	NestedTypeCount int32 `json:"nestedTypeCount,omitempty"`
	EnumTypeCount int32 `json:"enumTypeCount,omitempty"`
	ExtensionCount int32 `json:"extensionCount,omitempty"`
	FieldList []FieldDescriptorProto `json:"fieldList,omitempty"`
	FieldCount int32 `json:"fieldCount,omitempty"`
	EnumTypeList []EnumDescriptorProto `json:"enumTypeList,omitempty"`
	EnumTypeOrBuilderList []EnumDescriptorProtoOrBuilder `json:"enumTypeOrBuilderList,omitempty"`
	ExtensionList []FieldDescriptorProto `json:"extensionList,omitempty"`
	ExtensionOrBuilderList []FieldDescriptorProtoOrBuilder `json:"extensionOrBuilderList,omitempty"`
	FieldOrBuilderList []FieldDescriptorProtoOrBuilder `json:"fieldOrBuilderList,omitempty"`
	NestedTypeList []DescriptorProto `json:"nestedTypeList,omitempty"`
	NestedTypeOrBuilderList []DescriptorProtoOrBuilder `json:"nestedTypeOrBuilderList,omitempty"`
	ExtensionRangeOrBuilderList []ExtensionRangeOrBuilder `json:"extensionRangeOrBuilderList,omitempty"`
	OptionsOrBuilder *MessageOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	ExtensionRangeCount int32 `json:"extensionRangeCount,omitempty"`
	OneofDeclList []OneofDescriptorProto `json:"oneofDeclList,omitempty"`
	OneofDeclOrBuilderList []OneofDescriptorProtoOrBuilder `json:"oneofDeclOrBuilderList,omitempty"`
	ReservedRangeOrBuilderList []ReservedRangeOrBuilder `json:"reservedRangeOrBuilderList,omitempty"`
	ReservedRangeCount int32 `json:"reservedRangeCount,omitempty"`
	ReservedNameCount int32 `json:"reservedNameCount,omitempty"`
	NameBytes *ByteString `json:"nameBytes,omitempty"`
	Name string `json:"name,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize int32 `json:"memoizedSerializedSize,omitempty"`
}