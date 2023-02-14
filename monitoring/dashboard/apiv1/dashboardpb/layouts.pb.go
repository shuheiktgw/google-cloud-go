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
// source: google/monitoring/dashboard/v1/layouts.proto

package dashboardpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A basic layout divides the available space into vertical columns of equal
// width and arranges a list of widgets using a row-first strategy.
type GridLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of columns into which the view's width is divided. If omitted
	// or set to zero, a system default will be used while rendering.
	Columns int64 `protobuf:"varint,1,opt,name=columns,proto3" json:"columns,omitempty"`
	// The informational elements that are arranged into the columns row-first.
	Widgets []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *GridLayout) Reset() {
	*x = GridLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GridLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GridLayout) ProtoMessage() {}

func (x *GridLayout) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GridLayout.ProtoReflect.Descriptor instead.
func (*GridLayout) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{0}
}

func (x *GridLayout) GetColumns() int64 {
	if x != nil {
		return x.Columns
	}
	return 0
}

func (x *GridLayout) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

// A mosaic layout divides the available space into a grid of blocks, and
// overlays the grid with tiles. Unlike `GridLayout`, tiles may span multiple
// grid blocks and can be placed at arbitrary locations in the grid.
type MosaicLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of columns in the mosaic grid. The number of columns must be
	// between 1 and 12, inclusive.
	Columns int32 `protobuf:"varint,1,opt,name=columns,proto3" json:"columns,omitempty"`
	// The tiles to display.
	Tiles []*MosaicLayout_Tile `protobuf:"bytes,3,rep,name=tiles,proto3" json:"tiles,omitempty"`
}

func (x *MosaicLayout) Reset() {
	*x = MosaicLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MosaicLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MosaicLayout) ProtoMessage() {}

func (x *MosaicLayout) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MosaicLayout.ProtoReflect.Descriptor instead.
func (*MosaicLayout) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{1}
}

func (x *MosaicLayout) GetColumns() int32 {
	if x != nil {
		return x.Columns
	}
	return 0
}

func (x *MosaicLayout) GetTiles() []*MosaicLayout_Tile {
	if x != nil {
		return x.Tiles
	}
	return nil
}

// A simplified layout that divides the available space into rows
// and arranges a set of widgets horizontally in each row.
type RowLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The rows of content to display.
	Rows []*RowLayout_Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *RowLayout) Reset() {
	*x = RowLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowLayout) ProtoMessage() {}

func (x *RowLayout) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowLayout.ProtoReflect.Descriptor instead.
func (*RowLayout) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{2}
}

func (x *RowLayout) GetRows() []*RowLayout_Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

// A simplified layout that divides the available space into vertical columns
// and arranges a set of widgets vertically in each column.
type ColumnLayout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The columns of content to display.
	Columns []*ColumnLayout_Column `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *ColumnLayout) Reset() {
	*x = ColumnLayout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnLayout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnLayout) ProtoMessage() {}

func (x *ColumnLayout) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnLayout.ProtoReflect.Descriptor instead.
func (*ColumnLayout) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{3}
}

func (x *ColumnLayout) GetColumns() []*ColumnLayout_Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

// A single tile in the mosaic. The placement and size of the tile are
// configurable.
type MosaicLayout_Tile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The zero-indexed position of the tile in grid blocks relative to the
	// left edge of the grid. Tiles must be contained within the specified
	// number of columns. `x_pos` cannot be negative.
	XPos int32 `protobuf:"varint,1,opt,name=x_pos,json=xPos,proto3" json:"x_pos,omitempty"`
	// The zero-indexed position of the tile in grid blocks relative to the
	// top edge of the grid. `y_pos` cannot be negative.
	YPos int32 `protobuf:"varint,2,opt,name=y_pos,json=yPos,proto3" json:"y_pos,omitempty"`
	// The width of the tile, measured in grid blocks. Tiles must have a
	// minimum width of 1.
	Width int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the tile, measured in grid blocks. Tiles must have a
	// minimum height of 1.
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	// The informational widget contained in the tile. For example an `XyChart`.
	Widget *Widget `protobuf:"bytes,5,opt,name=widget,proto3" json:"widget,omitempty"`
}

func (x *MosaicLayout_Tile) Reset() {
	*x = MosaicLayout_Tile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MosaicLayout_Tile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MosaicLayout_Tile) ProtoMessage() {}

func (x *MosaicLayout_Tile) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MosaicLayout_Tile.ProtoReflect.Descriptor instead.
func (*MosaicLayout_Tile) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MosaicLayout_Tile) GetXPos() int32 {
	if x != nil {
		return x.XPos
	}
	return 0
}

func (x *MosaicLayout_Tile) GetYPos() int32 {
	if x != nil {
		return x.YPos
	}
	return 0
}

func (x *MosaicLayout_Tile) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *MosaicLayout_Tile) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *MosaicLayout_Tile) GetWidget() *Widget {
	if x != nil {
		return x.Widget
	}
	return nil
}

// Defines the layout properties and content for a row.
type RowLayout_Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative weight of this row. The row weight is used to adjust the
	// height of rows on the screen (relative to peers). Greater the weight,
	// greater the height of the row on the screen. If omitted, a value
	// of 1 is used while rendering.
	Weight int64 `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	// The display widgets arranged horizontally in this row.
	Widgets []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *RowLayout_Row) Reset() {
	*x = RowLayout_Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RowLayout_Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RowLayout_Row) ProtoMessage() {}

func (x *RowLayout_Row) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RowLayout_Row.ProtoReflect.Descriptor instead.
func (*RowLayout_Row) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{2, 0}
}

func (x *RowLayout_Row) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *RowLayout_Row) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

// Defines the layout properties and content for a column.
type ColumnLayout_Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative weight of this column. The column weight is used to adjust
	// the width of columns on the screen (relative to peers).
	// Greater the weight, greater the width of the column on the screen.
	// If omitted, a value of 1 is used while rendering.
	Weight int64 `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	// The display widgets arranged vertically in this column.
	Widgets []*Widget `protobuf:"bytes,2,rep,name=widgets,proto3" json:"widgets,omitempty"`
}

func (x *ColumnLayout_Column) Reset() {
	*x = ColumnLayout_Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnLayout_Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnLayout_Column) ProtoMessage() {}

func (x *ColumnLayout_Column) ProtoReflect() protoreflect.Message {
	mi := &file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnLayout_Column.ProtoReflect.Descriptor instead.
func (*ColumnLayout_Column) Descriptor() ([]byte, []int) {
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ColumnLayout_Column) GetWeight() int64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *ColumnLayout_Column) GetWidgets() []*Widget {
	if x != nil {
		return x.Widgets
	}
	return nil
}

var File_google_monitoring_dashboard_v1_layouts_proto protoreflect.FileDescriptor

var file_google_monitoring_dashboard_v1_layouts_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x2b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0a, 0x47,
	0x72, 0x69, 0x64, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x07, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x73, 0x61, 0x69, 0x63,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x12, 0x47, 0x0a, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x73, 0x61, 0x69, 0x63, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x69,
	0x6c, 0x65, 0x52, 0x05, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x9e, 0x01, 0x0a, 0x04, 0x54, 0x69,
	0x6c, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x78, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x78, 0x50, 0x6f, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x79, 0x5f, 0x70, 0x6f, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x50, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x09, 0x52,
	0x6f, 0x77, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x1a, 0x5f, 0x0a, 0x03, 0x52,
	0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x77, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x22, 0xc1, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x4d, 0x0a,
	0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x1a, 0x62, 0x0a, 0x06,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40,
	0x0a, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x07, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x42, 0xf6, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x5c, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xea, 0x02,
	0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_monitoring_dashboard_v1_layouts_proto_rawDescOnce sync.Once
	file_google_monitoring_dashboard_v1_layouts_proto_rawDescData = file_google_monitoring_dashboard_v1_layouts_proto_rawDesc
)

func file_google_monitoring_dashboard_v1_layouts_proto_rawDescGZIP() []byte {
	file_google_monitoring_dashboard_v1_layouts_proto_rawDescOnce.Do(func() {
		file_google_monitoring_dashboard_v1_layouts_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_monitoring_dashboard_v1_layouts_proto_rawDescData)
	})
	return file_google_monitoring_dashboard_v1_layouts_proto_rawDescData
}

var file_google_monitoring_dashboard_v1_layouts_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_monitoring_dashboard_v1_layouts_proto_goTypes = []interface{}{
	(*GridLayout)(nil),          // 0: google.monitoring.dashboard.v1.GridLayout
	(*MosaicLayout)(nil),        // 1: google.monitoring.dashboard.v1.MosaicLayout
	(*RowLayout)(nil),           // 2: google.monitoring.dashboard.v1.RowLayout
	(*ColumnLayout)(nil),        // 3: google.monitoring.dashboard.v1.ColumnLayout
	(*MosaicLayout_Tile)(nil),   // 4: google.monitoring.dashboard.v1.MosaicLayout.Tile
	(*RowLayout_Row)(nil),       // 5: google.monitoring.dashboard.v1.RowLayout.Row
	(*ColumnLayout_Column)(nil), // 6: google.monitoring.dashboard.v1.ColumnLayout.Column
	(*Widget)(nil),              // 7: google.monitoring.dashboard.v1.Widget
}
var file_google_monitoring_dashboard_v1_layouts_proto_depIdxs = []int32{
	7, // 0: google.monitoring.dashboard.v1.GridLayout.widgets:type_name -> google.monitoring.dashboard.v1.Widget
	4, // 1: google.monitoring.dashboard.v1.MosaicLayout.tiles:type_name -> google.monitoring.dashboard.v1.MosaicLayout.Tile
	5, // 2: google.monitoring.dashboard.v1.RowLayout.rows:type_name -> google.monitoring.dashboard.v1.RowLayout.Row
	6, // 3: google.monitoring.dashboard.v1.ColumnLayout.columns:type_name -> google.monitoring.dashboard.v1.ColumnLayout.Column
	7, // 4: google.monitoring.dashboard.v1.MosaicLayout.Tile.widget:type_name -> google.monitoring.dashboard.v1.Widget
	7, // 5: google.monitoring.dashboard.v1.RowLayout.Row.widgets:type_name -> google.monitoring.dashboard.v1.Widget
	7, // 6: google.monitoring.dashboard.v1.ColumnLayout.Column.widgets:type_name -> google.monitoring.dashboard.v1.Widget
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_monitoring_dashboard_v1_layouts_proto_init() }
func file_google_monitoring_dashboard_v1_layouts_proto_init() {
	if File_google_monitoring_dashboard_v1_layouts_proto != nil {
		return
	}
	file_google_monitoring_dashboard_v1_widget_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GridLayout); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MosaicLayout); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowLayout); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnLayout); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MosaicLayout_Tile); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RowLayout_Row); i {
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
		file_google_monitoring_dashboard_v1_layouts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnLayout_Column); i {
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
			RawDescriptor: file_google_monitoring_dashboard_v1_layouts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_monitoring_dashboard_v1_layouts_proto_goTypes,
		DependencyIndexes: file_google_monitoring_dashboard_v1_layouts_proto_depIdxs,
		MessageInfos:      file_google_monitoring_dashboard_v1_layouts_proto_msgTypes,
	}.Build()
	File_google_monitoring_dashboard_v1_layouts_proto = out.File
	file_google_monitoring_dashboard_v1_layouts_proto_rawDesc = nil
	file_google_monitoring_dashboard_v1_layouts_proto_goTypes = nil
	file_google_monitoring_dashboard_v1_layouts_proto_depIdxs = nil
}
