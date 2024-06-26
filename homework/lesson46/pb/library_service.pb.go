// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: library_service.proto

package pb

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId        string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	YearPublished string `protobuf:"bytes,4,opt,name=year_published,json=yearPublished,proto3" json:"year_published,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetYearPublished() string {
	if x != nil {
		return x.YearPublished
	}
	return ""
}

type GetAllBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllBooksRequest) Reset() {
	*x = GetAllBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksRequest) ProtoMessage() {}

func (x *GetAllBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBooksRequest.ProtoReflect.Descriptor instead.
func (*GetAllBooksRequest) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{1}
}

type GetAllBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetAllBooksResponse) Reset() {
	*x = GetAllBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBooksResponse) ProtoMessage() {}

func (x *GetAllBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBooksResponse.ProtoReflect.Descriptor instead.
func (*GetAllBooksResponse) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type AddBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author        string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	YearPublished string `protobuf:"bytes,4,opt,name=year_published,json=yearPublished,proto3" json:"year_published,omitempty"`
}

func (x *AddBookRequest) Reset() {
	*x = AddBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookRequest) ProtoMessage() {}

func (x *AddBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookRequest.ProtoReflect.Descriptor instead.
func (*AddBookRequest) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *AddBookRequest) GetYearPublished() string {
	if x != nil {
		return x.YearPublished
	}
	return ""
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddBookResponse) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

type SearchBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	YearPublished string `protobuf:"bytes,3,opt,name=year_published,json=yearPublished,proto3" json:"year_published,omitempty"`
}

func (x *SearchBookRequest) Reset() {
	*x = SearchBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookRequest) ProtoMessage() {}

func (x *SearchBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookRequest.ProtoReflect.Descriptor instead.
func (*SearchBookRequest) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{5}
}

func (x *SearchBookRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SearchBookRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *SearchBookRequest) GetYearPublished() string {
	if x != nil {
		return x.YearPublished
	}
	return ""
}

type SearchBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *SearchBookResponse) Reset() {
	*x = SearchBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBookResponse) ProtoMessage() {}

func (x *SearchBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBookResponse.ProtoReflect.Descriptor instead.
func (*SearchBookResponse) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{6}
}

func (x *SearchBookResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type BorrowBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId string `protobuf:"bytes,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *BorrowBookRequest) Reset() {
	*x = BorrowBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookRequest) ProtoMessage() {}

func (x *BorrowBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookRequest.ProtoReflect.Descriptor instead.
func (*BorrowBookRequest) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{7}
}

func (x *BorrowBookRequest) GetBookId() string {
	if x != nil {
		return x.BookId
	}
	return ""
}

func (x *BorrowBookRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type BorrowBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message bool `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BorrowBookResponse) Reset() {
	*x = BorrowBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_library_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowBookResponse) ProtoMessage() {}

func (x *BorrowBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_library_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowBookResponse.ProtoReflect.Descriptor instead.
func (*BorrowBookResponse) Descriptor() ([]byte, []int) {
	return file_library_service_proto_rawDescGZIP(), []int{8}
}

func (x *BorrowBookResponse) GetMessage() bool {
	if x != nil {
		return x.Message
	}
	return false
}

var File_library_service_proto protoreflect.FileDescriptor

var file_library_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x74, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x79, 0x65, 0x61, 0x72, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x79, 0x65, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x14,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x65, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x79, 0x65, 0x61, 0x72,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x79, 0x65, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22,
	0x2a, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x11, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x25,
	0x0a, 0x0e, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x79, 0x65, 0x61, 0x72, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x45, 0x0a, 0x11, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x12, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xe7, 0x02, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12, 0x4e, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1f, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f,
	0x6f, 0x6b, 0x12, 0x22, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x23, 0x2e, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x24, 0x2e, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_library_service_proto_rawDescOnce sync.Once
	file_library_service_proto_rawDescData = file_library_service_proto_rawDesc
)

func file_library_service_proto_rawDescGZIP() []byte {
	file_library_service_proto_rawDescOnce.Do(func() {
		file_library_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_library_service_proto_rawDescData)
	})
	return file_library_service_proto_rawDescData
}

var file_library_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_library_service_proto_goTypes = []any{
	(*Book)(nil),                // 0: library_service.Book
	(*GetAllBooksRequest)(nil),  // 1: library_service.GetAllBooksRequest
	(*GetAllBooksResponse)(nil), // 2: library_service.GetAllBooksResponse
	(*AddBookRequest)(nil),      // 3: library_service.AddBookRequest
	(*AddBookResponse)(nil),     // 4: library_service.AddBookResponse
	(*SearchBookRequest)(nil),   // 5: library_service.SearchBookRequest
	(*SearchBookResponse)(nil),  // 6: library_service.SearchBookResponse
	(*BorrowBookRequest)(nil),   // 7: library_service.BorrowBookRequest
	(*BorrowBookResponse)(nil),  // 8: library_service.BorrowBookResponse
}
var file_library_service_proto_depIdxs = []int32{
	0, // 0: library_service.GetAllBooksResponse.books:type_name -> library_service.Book
	0, // 1: library_service.SearchBookResponse.books:type_name -> library_service.Book
	3, // 2: library_service.Library.AddBook:input_type -> library_service.AddBookRequest
	5, // 3: library_service.Library.SearchBook:input_type -> library_service.SearchBookRequest
	7, // 4: library_service.Library.BorrowBook:input_type -> library_service.BorrowBookRequest
	1, // 5: library_service.Library.GetAllBooks:input_type -> library_service.GetAllBooksRequest
	4, // 6: library_service.Library.AddBook:output_type -> library_service.AddBookResponse
	6, // 7: library_service.Library.SearchBook:output_type -> library_service.SearchBookResponse
	8, // 8: library_service.Library.BorrowBook:output_type -> library_service.BorrowBookResponse
	2, // 9: library_service.Library.GetAllBooks:output_type -> library_service.GetAllBooksResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_library_service_proto_init() }
func file_library_service_proto_init() {
	if File_library_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_library_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Book); i {
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
		file_library_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllBooksRequest); i {
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
		file_library_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllBooksResponse); i {
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
		file_library_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookRequest); i {
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
		file_library_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddBookResponse); i {
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
		file_library_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SearchBookRequest); i {
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
		file_library_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SearchBookResponse); i {
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
		file_library_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowBookRequest); i {
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
		file_library_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowBookResponse); i {
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
			RawDescriptor: file_library_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_library_service_proto_goTypes,
		DependencyIndexes: file_library_service_proto_depIdxs,
		MessageInfos:      file_library_service_proto_msgTypes,
	}.Build()
	File_library_service_proto = out.File
	file_library_service_proto_rawDesc = nil
	file_library_service_proto_goTypes = nil
	file_library_service_proto_depIdxs = nil
}
