// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import protostructure "github.com/mitchellh/protostructure"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Args are the common argument types that are available to many of the
// dynamic functions. The exact list of available argument types is available
// on the Go interface docs.
type Args struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}
func (*Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0}
}
func (m *Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args.Unmarshal(m, b)
}
func (m *Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args.Marshal(b, m, deterministic)
}
func (dst *Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args.Merge(dst, src)
}
func (m *Args) XXX_Size() int {
	return xxx_messageInfo_Args.Size(m)
}
func (m *Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Args proto.InternalMessageInfo

// See component.Source
type Args_Source struct {
	// app is the name of the application being deployed.
	App string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	// path is the absolute directory path to the root directory for source files.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_Source) Reset()         { *m = Args_Source{} }
func (m *Args_Source) String() string { return proto.CompactTextString(m) }
func (*Args_Source) ProtoMessage()    {}
func (*Args_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 0}
}
func (m *Args_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_Source.Unmarshal(m, b)
}
func (m *Args_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_Source.Marshal(b, m, deterministic)
}
func (dst *Args_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_Source.Merge(dst, src)
}
func (m *Args_Source) XXX_Size() int {
	return xxx_messageInfo_Args_Source.Size(m)
}
func (m *Args_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_Source.DiscardUnknown(m)
}

var xxx_messageInfo_Args_Source proto.InternalMessageInfo

func (m *Args_Source) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Args_Source) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// DataDir represents the directory where data can be stored. This is an internal
// struct and shouldn't be used directly. Use the relevant *datadir implementation
// instead.
type Args_DataDir struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_DataDir) Reset()         { *m = Args_DataDir{} }
func (m *Args_DataDir) String() string { return proto.CompactTextString(m) }
func (*Args_DataDir) ProtoMessage()    {}
func (*Args_DataDir) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 1}
}
func (m *Args_DataDir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_DataDir.Unmarshal(m, b)
}
func (m *Args_DataDir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_DataDir.Marshal(b, m, deterministic)
}
func (dst *Args_DataDir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_DataDir.Merge(dst, src)
}
func (m *Args_DataDir) XXX_Size() int {
	return xxx_messageInfo_Args_DataDir.Size(m)
}
func (m *Args_DataDir) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_DataDir.DiscardUnknown(m)
}

var xxx_messageInfo_Args_DataDir proto.InternalMessageInfo

type Args_DataDir_Project struct {
	CacheDir             string   `protobuf:"bytes,2,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
	DataDir              string   `protobuf:"bytes,3,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_DataDir_Project) Reset()         { *m = Args_DataDir_Project{} }
func (m *Args_DataDir_Project) String() string { return proto.CompactTextString(m) }
func (*Args_DataDir_Project) ProtoMessage()    {}
func (*Args_DataDir_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 1, 0}
}
func (m *Args_DataDir_Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_DataDir_Project.Unmarshal(m, b)
}
func (m *Args_DataDir_Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_DataDir_Project.Marshal(b, m, deterministic)
}
func (dst *Args_DataDir_Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_DataDir_Project.Merge(dst, src)
}
func (m *Args_DataDir_Project) XXX_Size() int {
	return xxx_messageInfo_Args_DataDir_Project.Size(m)
}
func (m *Args_DataDir_Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_DataDir_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Args_DataDir_Project proto.InternalMessageInfo

func (m *Args_DataDir_Project) GetCacheDir() string {
	if m != nil {
		return m.CacheDir
	}
	return ""
}

func (m *Args_DataDir_Project) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

type Args_DataDir_App struct {
	CacheDir             string   `protobuf:"bytes,2,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
	DataDir              string   `protobuf:"bytes,3,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_DataDir_App) Reset()         { *m = Args_DataDir_App{} }
func (m *Args_DataDir_App) String() string { return proto.CompactTextString(m) }
func (*Args_DataDir_App) ProtoMessage()    {}
func (*Args_DataDir_App) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 1, 1}
}
func (m *Args_DataDir_App) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_DataDir_App.Unmarshal(m, b)
}
func (m *Args_DataDir_App) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_DataDir_App.Marshal(b, m, deterministic)
}
func (dst *Args_DataDir_App) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_DataDir_App.Merge(dst, src)
}
func (m *Args_DataDir_App) XXX_Size() int {
	return xxx_messageInfo_Args_DataDir_App.Size(m)
}
func (m *Args_DataDir_App) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_DataDir_App.DiscardUnknown(m)
}

var xxx_messageInfo_Args_DataDir_App proto.InternalMessageInfo

func (m *Args_DataDir_App) GetCacheDir() string {
	if m != nil {
		return m.CacheDir
	}
	return ""
}

func (m *Args_DataDir_App) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

type Args_DataDir_Component struct {
	CacheDir             string   `protobuf:"bytes,2,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
	DataDir              string   `protobuf:"bytes,3,opt,name=data_dir,json=dataDir,proto3" json:"data_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_DataDir_Component) Reset()         { *m = Args_DataDir_Component{} }
func (m *Args_DataDir_Component) String() string { return proto.CompactTextString(m) }
func (*Args_DataDir_Component) ProtoMessage()    {}
func (*Args_DataDir_Component) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 1, 2}
}
func (m *Args_DataDir_Component) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_DataDir_Component.Unmarshal(m, b)
}
func (m *Args_DataDir_Component) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_DataDir_Component.Marshal(b, m, deterministic)
}
func (dst *Args_DataDir_Component) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_DataDir_Component.Merge(dst, src)
}
func (m *Args_DataDir_Component) XXX_Size() int {
	return xxx_messageInfo_Args_DataDir_Component.Size(m)
}
func (m *Args_DataDir_Component) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_DataDir_Component.DiscardUnknown(m)
}

var xxx_messageInfo_Args_DataDir_Component proto.InternalMessageInfo

func (m *Args_DataDir_Component) GetCacheDir() string {
	if m != nil {
		return m.CacheDir
	}
	return ""
}

func (m *Args_DataDir_Component) GetDataDir() string {
	if m != nil {
		return m.DataDir
	}
	return ""
}

// Logger is used to construct an logger for the plugin.
type Args_Logger struct {
	// name is the name of the logger
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args_Logger) Reset()         { *m = Args_Logger{} }
func (m *Args_Logger) String() string { return proto.CompactTextString(m) }
func (*Args_Logger) ProtoMessage()    {}
func (*Args_Logger) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{0, 2}
}
func (m *Args_Logger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args_Logger.Unmarshal(m, b)
}
func (m *Args_Logger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args_Logger.Marshal(b, m, deterministic)
}
func (dst *Args_Logger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args_Logger.Merge(dst, src)
}
func (m *Args_Logger) XXX_Size() int {
	return xxx_messageInfo_Args_Logger.Size(m)
}
func (m *Args_Logger) XXX_DiscardUnknown() {
	xxx_messageInfo_Args_Logger.DiscardUnknown(m)
}

var xxx_messageInfo_Args_Logger proto.InternalMessageInfo

func (m *Args_Logger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Empty is just an empty message useful with some RPC endpoints.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// FuncSpec describes a function and is used by the dependency-injection
// framework to provide the function with the proper values.
//
// Value types are specified using strings. Built-in framework types
// are defined via constants in the `sdk` package. For custom types, you
// can use whatever string as long as it is unique. We recommend using a
// unique prefix plus the Go type name.
type FuncSpec struct {
	// name of the function. This is used for improved logging.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// args is the list of arguments by protobuf Any types.
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// result is the resulting type. This is only critically important to be
	// set for functions that may chain to other functions. It can be set to
	// blank in which case it will not be used.
	Result               string   `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FuncSpec) Reset()         { *m = FuncSpec{} }
func (m *FuncSpec) String() string { return proto.CompactTextString(m) }
func (*FuncSpec) ProtoMessage()    {}
func (*FuncSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{2}
}
func (m *FuncSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FuncSpec.Unmarshal(m, b)
}
func (m *FuncSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FuncSpec.Marshal(b, m, deterministic)
}
func (dst *FuncSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FuncSpec.Merge(dst, src)
}
func (m *FuncSpec) XXX_Size() int {
	return xxx_messageInfo_FuncSpec.Size(m)
}
func (m *FuncSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FuncSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FuncSpec proto.InternalMessageInfo

func (m *FuncSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FuncSpec) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *FuncSpec) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

// Args is the standard argument type for an RPC that is calling a FuncSpec.
type FuncSpec_Args struct {
	// args is the list of argument types. This will include some of the
	// standard types in this file (in the Args message namespace) as well
	// as custom types declared by the FuncSpec that the plugin is expected
	// to understand how to decode.
	Args                 []*any.Any `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FuncSpec_Args) Reset()         { *m = FuncSpec_Args{} }
func (m *FuncSpec_Args) String() string { return proto.CompactTextString(m) }
func (*FuncSpec_Args) ProtoMessage()    {}
func (*FuncSpec_Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{2, 0}
}
func (m *FuncSpec_Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FuncSpec_Args.Unmarshal(m, b)
}
func (m *FuncSpec_Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FuncSpec_Args.Marshal(b, m, deterministic)
}
func (dst *FuncSpec_Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FuncSpec_Args.Merge(dst, src)
}
func (m *FuncSpec_Args) XXX_Size() int {
	return xxx_messageInfo_FuncSpec_Args.Size(m)
}
func (m *FuncSpec_Args) XXX_DiscardUnknown() {
	xxx_messageInfo_FuncSpec_Args.DiscardUnknown(m)
}

var xxx_messageInfo_FuncSpec_Args proto.InternalMessageInfo

func (m *FuncSpec_Args) GetArgs() []*any.Any {
	if m != nil {
		return m.Args
	}
	return nil
}

// Config is the namespace of messages related to configuration.
//
// All components that take configuration are expected to have two RPC calls:
//
//   * ConfigStruct - Returns the configuration structure.
//   * Configure - Sends the configuration data back to the plugin and the
//       plugin is also expected to perform any validation at this stage.
//
type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{3}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

// ConfigureRequest is the request sent once the configuration decoding
// is complete to actually assign the values to the structure.
type Config_ConfigureRequest struct {
	// json is the json data for the structure returned in the StructResp.
	// It is guaranteed to decode cleanly into the target structure.
	Json                 []byte   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config_ConfigureRequest) Reset()         { *m = Config_ConfigureRequest{} }
func (m *Config_ConfigureRequest) String() string { return proto.CompactTextString(m) }
func (*Config_ConfigureRequest) ProtoMessage()    {}
func (*Config_ConfigureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{3, 0}
}
func (m *Config_ConfigureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_ConfigureRequest.Unmarshal(m, b)
}
func (m *Config_ConfigureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_ConfigureRequest.Marshal(b, m, deterministic)
}
func (dst *Config_ConfigureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_ConfigureRequest.Merge(dst, src)
}
func (m *Config_ConfigureRequest) XXX_Size() int {
	return xxx_messageInfo_Config_ConfigureRequest.Size(m)
}
func (m *Config_ConfigureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_ConfigureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Config_ConfigureRequest proto.InternalMessageInfo

func (m *Config_ConfigureRequest) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

// StructResp returns the struct for configuration.
type Config_StructResp struct {
	// struct is the configuration structure (or empty/nil if one doesn't exist).
	// This struct should have all the proper struct tags for HCL decoding
	// You should do validation on the Configure call.
	Struct               *protostructure.Struct `protobuf:"bytes,1,opt,name=struct,proto3" json:"struct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Config_StructResp) Reset()         { *m = Config_StructResp{} }
func (m *Config_StructResp) String() string { return proto.CompactTextString(m) }
func (*Config_StructResp) ProtoMessage()    {}
func (*Config_StructResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{3, 1}
}
func (m *Config_StructResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_StructResp.Unmarshal(m, b)
}
func (m *Config_StructResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_StructResp.Marshal(b, m, deterministic)
}
func (dst *Config_StructResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_StructResp.Merge(dst, src)
}
func (m *Config_StructResp) XXX_Size() int {
	return xxx_messageInfo_Config_StructResp.Size(m)
}
func (m *Config_StructResp) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_StructResp.DiscardUnknown(m)
}

var xxx_messageInfo_Config_StructResp proto.InternalMessageInfo

func (m *Config_StructResp) GetStruct() *protostructure.Struct {
	if m != nil {
		return m.Struct
	}
	return nil
}

// ImplementsResp returns true if the component implements an additional interface.
type ImplementsResp struct {
	Implements           bool     `protobuf:"varint,1,opt,name=implements,proto3" json:"implements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImplementsResp) Reset()         { *m = ImplementsResp{} }
func (m *ImplementsResp) String() string { return proto.CompactTextString(m) }
func (*ImplementsResp) ProtoMessage()    {}
func (*ImplementsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_plugin_58543b634987d0a9, []int{4}
}
func (m *ImplementsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImplementsResp.Unmarshal(m, b)
}
func (m *ImplementsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImplementsResp.Marshal(b, m, deterministic)
}
func (dst *ImplementsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplementsResp.Merge(dst, src)
}
func (m *ImplementsResp) XXX_Size() int {
	return xxx_messageInfo_ImplementsResp.Size(m)
}
func (m *ImplementsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplementsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ImplementsResp proto.InternalMessageInfo

func (m *ImplementsResp) GetImplements() bool {
	if m != nil {
		return m.Implements
	}
	return false
}

func init() {
	proto.RegisterType((*Args)(nil), "proto.Args")
	proto.RegisterType((*Args_Source)(nil), "proto.Args.Source")
	proto.RegisterType((*Args_DataDir)(nil), "proto.Args.DataDir")
	proto.RegisterType((*Args_DataDir_Project)(nil), "proto.Args.DataDir.Project")
	proto.RegisterType((*Args_DataDir_App)(nil), "proto.Args.DataDir.App")
	proto.RegisterType((*Args_DataDir_Component)(nil), "proto.Args.DataDir.Component")
	proto.RegisterType((*Args_Logger)(nil), "proto.Args.Logger")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*FuncSpec)(nil), "proto.FuncSpec")
	proto.RegisterType((*FuncSpec_Args)(nil), "proto.FuncSpec.Args")
	proto.RegisterType((*Config)(nil), "proto.Config")
	proto.RegisterType((*Config_ConfigureRequest)(nil), "proto.Config.ConfigureRequest")
	proto.RegisterType((*Config_StructResp)(nil), "proto.Config.StructResp")
	proto.RegisterType((*ImplementsResp)(nil), "proto.ImplementsResp")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_plugin_58543b634987d0a9) }

var fileDescriptor_plugin_58543b634987d0a9 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0xf1, 0xae, 0xed, 0xcc, 0x2e, 0x65, 0x11, 0xcb, 0x92, 0x55, 0x4b, 0x09, 0x3e, 0x94,
	0x9c, 0xbc, 0xcb, 0xf6, 0xda, 0x1e, 0x4c, 0xd2, 0x42, 0xa1, 0x87, 0xe2, 0x3c, 0x40, 0xd1, 0x3a,
	0x13, 0xc5, 0xc1, 0x96, 0x54, 0xfd, 0x1c, 0x02, 0x7d, 0x9f, 0xbe, 0x4e, 0x1f, 0xa9, 0x78, 0xac,
	0x96, 0x52, 0x7a, 0xca, 0x49, 0x9f, 0x3e, 0x7d, 0xf3, 0xcd, 0x68, 0x66, 0xe0, 0xda, 0xf4, 0x41,
	0x76, 0xaa, 0x32, 0x56, 0x7b, 0xcd, 0x2e, 0xe9, 0xe0, 0xf7, 0x52, 0x6b, 0xd9, 0xe3, 0x03, 0xdd,
	0x9e, 0xc3, 0xfe, 0x41, 0xa8, 0xd3, 0xa4, 0xe0, 0xb7, 0x74, 0x38, 0x6f, 0x43, 0xeb, 0x83, 0xc5,
	0x89, 0x2d, 0x7f, 0xcc, 0xe0, 0xa2, 0xb6, 0xd2, 0xf1, 0x0a, 0xb2, 0xad, 0x0e, 0xb6, 0x45, 0x76,
	0x03, 0xa9, 0x30, 0x66, 0x91, 0x2c, 0x93, 0xd5, 0xbc, 0x19, 0x21, 0x63, 0x70, 0x61, 0x84, 0x3f,
	0x2c, 0x66, 0x44, 0x11, 0xe6, 0x3f, 0x13, 0xc8, 0x37, 0xc2, 0x8b, 0x4d, 0x67, 0x79, 0x0d, 0xf9,
	0x17, 0xab, 0x8f, 0xd8, 0x7a, 0xf6, 0x12, 0xe6, 0xad, 0x68, 0x0f, 0xf8, 0x75, 0xd7, 0xd9, 0xa8,
	0x2f, 0x88, 0xd8, 0x74, 0x96, 0xdd, 0x43, 0xb1, 0x13, 0x5e, 0xd0, 0x5b, 0x4a, 0x6f, 0xf9, 0x2e,
	0x5a, 0xbc, 0x87, 0xb4, 0x36, 0xe6, 0xec, 0xf0, 0x35, 0xcc, 0xd7, 0x7a, 0x30, 0x5a, 0xa1, 0x3a,
	0xbf, 0x86, 0x57, 0x90, 0x7d, 0xd6, 0x52, 0xa2, 0x1d, 0x3f, 0xac, 0xc4, 0x80, 0xb1, 0x07, 0x84,
	0xcb, 0x1c, 0x2e, 0x3f, 0x0c, 0xc6, 0x9f, 0xca, 0xef, 0x50, 0x7c, 0x0c, 0xaa, 0xdd, 0x1a, 0x6c,
	0xff, 0x27, 0x1c, 0x39, 0x61, 0xa5, 0x5b, 0xcc, 0x96, 0xe9, 0xc8, 0x8d, 0x98, 0xdd, 0x41, 0x66,
	0xd1, 0x85, 0xde, 0xc7, 0x9c, 0xf1, 0xc6, 0x1f, 0xa7, 0xee, 0xb3, 0x55, 0x8c, 0x49, 0x96, 0xe9,
	0xea, 0xea, 0xe9, 0xb6, 0x9a, 0xc6, 0x58, 0xfd, 0x1e, 0x63, 0x55, 0xab, 0xd3, 0xe4, 0x54, 0x2a,
	0xc8, 0xd6, 0x5a, 0xed, 0x3b, 0xc9, 0xdf, 0xc0, 0xcd, 0x84, 0x82, 0xc5, 0x06, 0xbf, 0x05, 0x74,
	0x7e, 0xcc, 0x7d, 0x74, 0x5a, 0x51, 0x3d, 0xd7, 0x0d, 0x61, 0xfe, 0x0e, 0x60, 0x4b, 0x53, 0x6f,
	0xd0, 0x19, 0x56, 0x41, 0x36, 0xed, 0x00, 0x69, 0xae, 0x9e, 0xee, 0xaa, 0x7f, 0xf6, 0x22, 0x6a,
	0xa3, 0xaa, 0x7c, 0x84, 0x17, 0x9f, 0x06, 0xd3, 0xe3, 0x80, 0xca, 0x3b, 0x72, 0x78, 0x0d, 0xd0,
	0xfd, 0x61, 0xc8, 0xa5, 0x68, 0xfe, 0x62, 0x9e, 0x33, 0x32, 0x7c, 0xfb, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xa8, 0x34, 0x02, 0xa1, 0x02, 0x00, 0x00,
}