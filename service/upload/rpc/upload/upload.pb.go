// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: upload.proto

package upload

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

// fileMeta 文件元信息
type FileMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileHash        string `protobuf:"bytes,1,opt,name=fileHash,proto3" json:"fileHash,omitempty"`
	FileName        string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileContentType string `protobuf:"bytes,3,opt,name=fileContentType,proto3" json:"fileContentType,omitempty"`
	FileSize        int64  `protobuf:"varint,4,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	Location        string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	UploadAt        string `protobuf:"bytes,6,opt,name=uploadAt,proto3" json:"uploadAt,omitempty"`
}

func (x *FileMeta) Reset() {
	*x = FileMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMeta) ProtoMessage() {}

func (x *FileMeta) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMeta.ProtoReflect.Descriptor instead.
func (*FileMeta) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *FileMeta) GetFileHash() string {
	if x != nil {
		return x.FileHash
	}
	return ""
}

func (x *FileMeta) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileMeta) GetFileContentType() string {
	if x != nil {
		return x.FileContentType
	}
	return ""
}

func (x *FileMeta) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileMeta) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *FileMeta) GetUploadAt() string {
	if x != nil {
		return x.UploadAt
	}
	return ""
}

// fileOSSMeta 文件oss信息
type FileOSSMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketName string `protobuf:"bytes,1,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	OssPath    string `protobuf:"bytes,2,opt,name=ossPath,proto3" json:"ossPath,omitempty"`
}

func (x *FileOSSMeta) Reset() {
	*x = FileOSSMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOSSMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOSSMeta) ProtoMessage() {}

func (x *FileOSSMeta) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOSSMeta.ProtoReflect.Descriptor instead.
func (*FileOSSMeta) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *FileOSSMeta) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *FileOSSMeta) GetOssPath() string {
	if x != nil {
		return x.OssPath
	}
	return ""
}

// FileUploadRequest 文件上传
type FileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMeta    *FileMeta    `protobuf:"bytes,1,opt,name=fileMeta,proto3" json:"fileMeta,omitempty"`
	FileOSSMeta *FileOSSMeta `protobuf:"bytes,2,opt,name=fileOSSMeta,proto3" json:"fileOSSMeta,omitempty"`
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *FileUploadRequest) GetFileMeta() *FileMeta {
	if x != nil {
		return x.FileMeta
	}
	return nil
}

func (x *FileUploadRequest) GetFileOSSMeta() *FileOSSMeta {
	if x != nil {
		return x.FileOSSMeta
	}
	return nil
}

type FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket       string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key          string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	ETag         string `protobuf:"bytes,3,opt,name=eTag,proto3" json:"eTag,omitempty"`
	Size         int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	LastModified string `protobuf:"bytes,5,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	Location     string `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	VersionID    string `protobuf:"bytes,7,opt,name=versionID,proto3" json:"versionID,omitempty"`
	// Lifecycle expiry-date and ruleID associated with the expiry
	// not to be confused with `Expires` HTTP header.
	Expiration       string `protobuf:"bytes,8,opt,name=expiration,proto3" json:"expiration,omitempty"`
	ExpirationRuleID string `protobuf:"bytes,9,opt,name=expirationRuleID,proto3" json:"expirationRuleID,omitempty"`
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadResponse) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *FileUploadResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *FileUploadResponse) GetETag() string {
	if x != nil {
		return x.ETag
	}
	return ""
}

func (x *FileUploadResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileUploadResponse) GetLastModified() string {
	if x != nil {
		return x.LastModified
	}
	return ""
}

func (x *FileUploadResponse) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *FileUploadResponse) GetVersionID() string {
	if x != nil {
		return x.VersionID
	}
	return ""
}

func (x *FileUploadResponse) GetExpiration() string {
	if x != nil {
		return x.Expiration
	}
	return ""
}

func (x *FileUploadResponse) GetExpirationRuleID() string {
	if x != nil {
		return x.ExpirationRuleID
	}
	return ""
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x74, 0x22, 0x47, 0x0a, 0x0b, 0x66, 0x69, 0x6c,
	0x65, 0x4f, 0x53, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x73, 0x73, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x73, 0x73, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x78, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x53, 0x53,
	0x4d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x53, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x53, 0x53, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x90, 0x02, 0x0a,
	0x12, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x65, 0x54, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x54, 0x61,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x32,
	0x4d, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_upload_proto_goTypes = []interface{}{
	(*FileMeta)(nil),           // 0: upload.fileMeta
	(*FileOSSMeta)(nil),        // 1: upload.fileOSSMeta
	(*FileUploadRequest)(nil),  // 2: upload.FileUploadRequest
	(*FileUploadResponse)(nil), // 3: upload.FileUploadResponse
}
var file_upload_proto_depIdxs = []int32{
	0, // 0: upload.FileUploadRequest.fileMeta:type_name -> upload.fileMeta
	1, // 1: upload.FileUploadRequest.fileOSSMeta:type_name -> upload.fileOSSMeta
	2, // 2: upload.Upload.FileUpload:input_type -> upload.FileUploadRequest
	3, // 3: upload.Upload.FileUpload:output_type -> upload.FileUploadResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMeta); i {
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
		file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOSSMeta); i {
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
		file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRequest); i {
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
		file_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadResponse); i {
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
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}