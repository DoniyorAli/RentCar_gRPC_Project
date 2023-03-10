// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rentcar_protoc/brand.proto

package brand

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Manufacturer string `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	AboutBrand   string `protobuf:"bytes,4,opt,name=about_brand,json=aboutBrand,proto3" json:"about_brand,omitempty"`
}

func (x *CreateBrandRequest) Reset() {
	*x = CreateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrandRequest) ProtoMessage() {}

func (x *CreateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrandRequest.ProtoReflect.Descriptor instead.
func (*CreateBrandRequest) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBrandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBrandRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateBrandRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *CreateBrandRequest) GetAboutBrand() string {
	if x != nil {
		return x.AboutBrand
	}
	return ""
}

type Brand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId      string `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country      string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Manufacturer string `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	AboutBrand   string `protobuf:"bytes,5,opt,name=about_brand,json=aboutBrand,proto3" json:"about_brand,omitempty"`
	CreatedAt    string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Brand) Reset() {
	*x = Brand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brand) ProtoMessage() {}

func (x *Brand) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brand.ProtoReflect.Descriptor instead.
func (*Brand) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{1}
}

func (x *Brand) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *Brand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Brand) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Brand) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Brand) GetAboutBrand() string {
	if x != nil {
		return x.AboutBrand
	}
	return ""
}

func (x *Brand) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Brand) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country      string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	Manufacturer string `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	AboutBrand   string `protobuf:"bytes,5,opt,name=about_brand,json=aboutBrand,proto3" json:"about_brand,omitempty"`
}

func (x *UpdateBrandRequest) Reset() {
	*x = UpdateBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBrandRequest) ProtoMessage() {}

func (x *UpdateBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBrandRequest.ProtoReflect.Descriptor instead.
func (*UpdateBrandRequest) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateBrandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBrandRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBrandRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdateBrandRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *UpdateBrandRequest) GetAboutBrand() string {
	if x != nil {
		return x.AboutBrand
	}
	return ""
}

type DeleteBrandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId string `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *DeleteBrandRequest) Reset() {
	*x = DeleteBrandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBrandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBrandRequest) ProtoMessage() {}

func (x *DeleteBrandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBrandRequest.ProtoReflect.Descriptor instead.
func (*DeleteBrandRequest) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteBrandRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

type GetBrandByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId string `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *GetBrandByIDRequest) Reset() {
	*x = GetBrandByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandByIDRequest) ProtoMessage() {}

func (x *GetBrandByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandByIDRequest.ProtoReflect.Descriptor instead.
func (*GetBrandByIDRequest) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{4}
}

func (x *GetBrandByIDRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

type GetBrandListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetBrandListRequest) Reset() {
	*x = GetBrandListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandListRequest) ProtoMessage() {}

func (x *GetBrandListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandListRequest.ProtoReflect.Descriptor instead.
func (*GetBrandListRequest) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{5}
}

func (x *GetBrandListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetBrandListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetBrandListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetBrandListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []*Brand `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *GetBrandListResponse) Reset() {
	*x = GetBrandListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rentcar_protoc_brand_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBrandListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBrandListResponse) ProtoMessage() {}

func (x *GetBrandListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rentcar_protoc_brand_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBrandListResponse.ProtoReflect.Descriptor instead.
func (*GetBrandListResponse) Descriptor() ([]byte, []int) {
	return file_rentcar_protoc_brand_proto_rawDescGZIP(), []int{6}
}

func (x *GetBrandListResponse) GetBrands() []*Brand {
	if x != nil {
		return x.Brands
	}
	return nil
}

var File_rentcar_protoc_brand_proto protoreflect.FileDescriptor

var file_rentcar_protoc_brand_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x65, 0x6e, 0x74, 0x63, 0x61, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e,
	0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x97, 0x01, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x62, 0x6f, 0x75,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x36, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x32, 0x87,
	0x02, 0x0a, 0x0c, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x13,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x06, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x06, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x06, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rentcar_protoc_brand_proto_rawDescOnce sync.Once
	file_rentcar_protoc_brand_proto_rawDescData = file_rentcar_protoc_brand_proto_rawDesc
)

func file_rentcar_protoc_brand_proto_rawDescGZIP() []byte {
	file_rentcar_protoc_brand_proto_rawDescOnce.Do(func() {
		file_rentcar_protoc_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_rentcar_protoc_brand_proto_rawDescData)
	})
	return file_rentcar_protoc_brand_proto_rawDescData
}

var file_rentcar_protoc_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rentcar_protoc_brand_proto_goTypes = []interface{}{
	(*CreateBrandRequest)(nil),   // 0: CreateBrandRequest
	(*Brand)(nil),                // 1: Brand
	(*UpdateBrandRequest)(nil),   // 2: UpdateBrandRequest
	(*DeleteBrandRequest)(nil),   // 3: DeleteBrandRequest
	(*GetBrandByIDRequest)(nil),  // 4: GetBrandByIDRequest
	(*GetBrandListRequest)(nil),  // 5: GetBrandListRequest
	(*GetBrandListResponse)(nil), // 6: GetBrandListResponse
}
var file_rentcar_protoc_brand_proto_depIdxs = []int32{
	1, // 0: GetBrandListResponse.brands:type_name -> Brand
	0, // 1: BrandService.CreateBrand:input_type -> CreateBrandRequest
	2, // 2: BrandService.UpdateBrand:input_type -> UpdateBrandRequest
	3, // 3: BrandService.DeleteBrand:input_type -> DeleteBrandRequest
	4, // 4: BrandService.GetBrandByID:input_type -> GetBrandByIDRequest
	5, // 5: BrandService.GetBrandList:input_type -> GetBrandListRequest
	1, // 6: BrandService.CreateBrand:output_type -> Brand
	1, // 7: BrandService.UpdateBrand:output_type -> Brand
	1, // 8: BrandService.DeleteBrand:output_type -> Brand
	1, // 9: BrandService.GetBrandByID:output_type -> Brand
	6, // 10: BrandService.GetBrandList:output_type -> GetBrandListResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rentcar_protoc_brand_proto_init() }
func file_rentcar_protoc_brand_proto_init() {
	if File_rentcar_protoc_brand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rentcar_protoc_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrandRequest); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brand); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBrandRequest); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBrandRequest); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandByIDRequest); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandListRequest); i {
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
		file_rentcar_protoc_brand_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBrandListResponse); i {
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
			RawDescriptor: file_rentcar_protoc_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rentcar_protoc_brand_proto_goTypes,
		DependencyIndexes: file_rentcar_protoc_brand_proto_depIdxs,
		MessageInfos:      file_rentcar_protoc_brand_proto_msgTypes,
	}.Build()
	File_rentcar_protoc_brand_proto = out.File
	file_rentcar_protoc_brand_proto_rawDesc = nil
	file_rentcar_protoc_brand_proto_goTypes = nil
	file_rentcar_protoc_brand_proto_depIdxs = nil
}
