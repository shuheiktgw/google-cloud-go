// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: google/firestore/admin/v1/field.proto

package adminpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The state of applying the TTL configuration to all documents.
type Field_TtlConfig_State int32

const (
	// The state is unspecified or unknown.
	Field_TtlConfig_STATE_UNSPECIFIED Field_TtlConfig_State = 0
	// The TTL is being applied. There is an active long-running operation to
	// track the change. Newly written documents will have TTLs applied as
	// requested. Requested TTLs on existing documents are still being
	// processed. When TTLs on all existing documents have been processed, the
	// state will move to 'ACTIVE'.
	Field_TtlConfig_CREATING Field_TtlConfig_State = 1
	// The TTL is active for all documents.
	Field_TtlConfig_ACTIVE Field_TtlConfig_State = 2
	// The TTL configuration could not be enabled for all existing documents.
	// Newly written documents will continue to have their TTL applied.
	// The LRO returned when last attempting to enable TTL for this `Field`
	// has failed, and may have more details.
	Field_TtlConfig_NEEDS_REPAIR Field_TtlConfig_State = 3
)

// Enum value maps for Field_TtlConfig_State.
var (
	Field_TtlConfig_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "NEEDS_REPAIR",
	}
	Field_TtlConfig_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"ACTIVE":            2,
		"NEEDS_REPAIR":      3,
	}
)

func (x Field_TtlConfig_State) Enum() *Field_TtlConfig_State {
	p := new(Field_TtlConfig_State)
	*p = x
	return p
}

func (x Field_TtlConfig_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Field_TtlConfig_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_firestore_admin_v1_field_proto_enumTypes[0].Descriptor()
}

func (Field_TtlConfig_State) Type() protoreflect.EnumType {
	return &file_google_firestore_admin_v1_field_proto_enumTypes[0]
}

func (x Field_TtlConfig_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Field_TtlConfig_State.Descriptor instead.
func (Field_TtlConfig_State) EnumDescriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_field_proto_rawDescGZIP(), []int{0, 1, 0}
}

// Represents a single field in the database.
//
// Fields are grouped by their "Collection Group", which represent all
// collections in the database with the same id.
type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. A field name of the form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`, which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that needs
	// to be escaped within a quoted field path is the backtick character itself,
	// escaped using a backslash. Special characters in field paths that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//
	//	additional layer of backticks to represent a code block)
	//
	// `\`address.city\“ represents a field named `address.city`, not the map key
	// `city` in the field `address`.
	// `\`*\“ represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all fields.
	// This field's resource name is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which do not
	// have their own `Field` index configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index configuration for this field. If unset, field indexing will
	// revert to the configuration defined by the `ancestor_field`. To
	// explicitly remove all indexes for this field, specify an index config
	// with an empty list of indexes.
	IndexConfig *Field_IndexConfig `protobuf:"bytes,2,opt,name=index_config,json=indexConfig,proto3" json:"index_config,omitempty"`
	// The TTL configuration for this `Field`.
	// Setting or unsetting this will enable or disable the TTL for
	// documents that have this `Field`.
	TtlConfig *Field_TtlConfig `protobuf:"bytes,3,opt,name=ttl_config,json=ttlConfig,proto3" json:"ttl_config,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_field_proto_rawDescGZIP(), []int{0}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetIndexConfig() *Field_IndexConfig {
	if x != nil {
		return x.IndexConfig
	}
	return nil
}

func (x *Field) GetTtlConfig() *Field_TtlConfig {
	if x != nil {
		return x.TtlConfig
	}
	return nil
}

// The index configuration for this field.
type Field_IndexConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The indexes supported for this field.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	// Output only. When true, the `Field`'s index configuration is set from the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `protobuf:"varint,2,opt,name=uses_ancestor_config,json=usesAncestorConfig,proto3" json:"uses_ancestor_config,omitempty"`
	// Output only. Specifies the resource name of the `Field` from which this field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `protobuf:"bytes,3,opt,name=ancestor_field,json=ancestorField,proto3" json:"ancestor_field,omitempty"`
	// Output only
	// When true, the `Field`'s index configuration is in the process of being
	// reverted. Once complete, the index config will transition to the same
	// state as the field specified by `ancestor_field`, at which point
	// `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting bool `protobuf:"varint,4,opt,name=reverting,proto3" json:"reverting,omitempty"`
}

func (x *Field_IndexConfig) Reset() {
	*x = Field_IndexConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_field_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field_IndexConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field_IndexConfig) ProtoMessage() {}

func (x *Field_IndexConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_field_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field_IndexConfig.ProtoReflect.Descriptor instead.
func (*Field_IndexConfig) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_field_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Field_IndexConfig) GetIndexes() []*Index {
	if x != nil {
		return x.Indexes
	}
	return nil
}

func (x *Field_IndexConfig) GetUsesAncestorConfig() bool {
	if x != nil {
		return x.UsesAncestorConfig
	}
	return false
}

func (x *Field_IndexConfig) GetAncestorField() string {
	if x != nil {
		return x.AncestorField
	}
	return ""
}

func (x *Field_IndexConfig) GetReverting() bool {
	if x != nil {
		return x.Reverting
	}
	return false
}

// The TTL (time-to-live) configuration for documents that have this `Field`
// set.
// Storing a timestamp value into a TTL-enabled field will be treated as
// the document's absolute expiration time. Using any other data type or
// leaving the field absent will disable the TTL for the individual document.
type Field_TtlConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The state of the TTL configuration.
	State Field_TtlConfig_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.firestore.admin.v1.Field_TtlConfig_State" json:"state,omitempty"`
}

func (x *Field_TtlConfig) Reset() {
	*x = Field_TtlConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_firestore_admin_v1_field_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field_TtlConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field_TtlConfig) ProtoMessage() {}

func (x *Field_TtlConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_firestore_admin_v1_field_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field_TtlConfig.ProtoReflect.Descriptor instead.
func (*Field_TtlConfig) Descriptor() ([]byte, []int) {
	return file_google_firestore_admin_v1_field_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Field_TtlConfig) GetState() Field_TtlConfig_State {
	if x != nil {
		return x.State
	}
	return Field_TtlConfig_STATE_UNSPECIFIED
}

var File_google_firestore_admin_v1_field_proto protoreflect.FileDescriptor

var file_google_firestore_admin_v1_field_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x05, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x74, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x54, 0x74, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x74, 0x74, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0xc0, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3a, 0x0a, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66,
	0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x07, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x75, 0x73, 0x65, 0x73, 0x41, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0xa4, 0x01, 0x0a, 0x09, 0x54, 0x74, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x54, 0x74, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x4a, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x4e, 0x45, 0x45, 0x44, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x41, 0x49, 0x52, 0x10, 0x03, 0x3a, 0x79,
	0xea, 0x41, 0x76, 0x0a, 0x1e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x54, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x2f, 0x7b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x7d, 0x42, 0xde, 0x01, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x66, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x04, 0x47, 0x43, 0x46, 0x53,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x46, 0x69, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a,
	0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_firestore_admin_v1_field_proto_rawDescOnce sync.Once
	file_google_firestore_admin_v1_field_proto_rawDescData = file_google_firestore_admin_v1_field_proto_rawDesc
)

func file_google_firestore_admin_v1_field_proto_rawDescGZIP() []byte {
	file_google_firestore_admin_v1_field_proto_rawDescOnce.Do(func() {
		file_google_firestore_admin_v1_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_firestore_admin_v1_field_proto_rawDescData)
	})
	return file_google_firestore_admin_v1_field_proto_rawDescData
}

var file_google_firestore_admin_v1_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_firestore_admin_v1_field_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_firestore_admin_v1_field_proto_goTypes = []interface{}{
	(Field_TtlConfig_State)(0), // 0: google.firestore.admin.v1.Field.TtlConfig.State
	(*Field)(nil),              // 1: google.firestore.admin.v1.Field
	(*Field_IndexConfig)(nil),  // 2: google.firestore.admin.v1.Field.IndexConfig
	(*Field_TtlConfig)(nil),    // 3: google.firestore.admin.v1.Field.TtlConfig
	(*Index)(nil),              // 4: google.firestore.admin.v1.Index
}
var file_google_firestore_admin_v1_field_proto_depIdxs = []int32{
	2, // 0: google.firestore.admin.v1.Field.index_config:type_name -> google.firestore.admin.v1.Field.IndexConfig
	3, // 1: google.firestore.admin.v1.Field.ttl_config:type_name -> google.firestore.admin.v1.Field.TtlConfig
	4, // 2: google.firestore.admin.v1.Field.IndexConfig.indexes:type_name -> google.firestore.admin.v1.Index
	0, // 3: google.firestore.admin.v1.Field.TtlConfig.state:type_name -> google.firestore.admin.v1.Field.TtlConfig.State
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_firestore_admin_v1_field_proto_init() }
func file_google_firestore_admin_v1_field_proto_init() {
	if File_google_firestore_admin_v1_field_proto != nil {
		return
	}
	file_google_firestore_admin_v1_index_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_firestore_admin_v1_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_firestore_admin_v1_field_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field_IndexConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_firestore_admin_v1_field_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field_TtlConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_firestore_admin_v1_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_firestore_admin_v1_field_proto_goTypes,
		DependencyIndexes: file_google_firestore_admin_v1_field_proto_depIdxs,
		EnumInfos:         file_google_firestore_admin_v1_field_proto_enumTypes,
		MessageInfos:      file_google_firestore_admin_v1_field_proto_msgTypes,
	}.Build()
	File_google_firestore_admin_v1_field_proto = out.File
	file_google_firestore_admin_v1_field_proto_rawDesc = nil
	file_google_firestore_admin_v1_field_proto_goTypes = nil
	file_google_firestore_admin_v1_field_proto_depIdxs = nil
}
