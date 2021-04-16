// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue312.proto

package issue312

import (
	fmt "fmt"
	_ "github.com/ivansukach/protobuf/gogoproto"
	proto "github.com/ivansukach/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TaskState int32

const (
	TaskState_TASK_STAGING  TaskState = 6
	TaskState_TASK_STARTING TaskState = 0
	TaskState_TASK_RUNNING  TaskState = 1
)

var TaskState_name = map[int32]string{
	6: "TASK_STAGING",
	0: "TASK_STARTING",
	1: "TASK_RUNNING",
}

var TaskState_value = map[string]int32{
	"TASK_STAGING":  6,
	"TASK_STARTING": 0,
	"TASK_RUNNING":  1,
}

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}

func (x *TaskState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TaskState_value, data, "TaskState")
	if err != nil {
		return err
	}
	*x = TaskState(value)
	return nil
}

func (TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a64932ccacef062, []int{0}
}

func init() {
	proto.RegisterEnum("issue312.TaskState", TaskState_name, TaskState_value)
}

func init() { proto.RegisterFile("issue312.proto", fileDescriptor_8a64932ccacef062) }

var fileDescriptor_8a64932ccacef062 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x36, 0x34, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0xd4, 0x72, 0xe2, 0xe2,
	0x0c, 0x49, 0x2c, 0xce, 0x0e, 0x2e, 0x49, 0x2c, 0x49, 0x15, 0x12, 0xe0, 0xe2, 0x09, 0x71, 0x0c,
	0xf6, 0x8e, 0x0f, 0x0e, 0x71, 0x74, 0xf7, 0xf4, 0x73, 0x17, 0x60, 0x13, 0x12, 0xe4, 0xe2, 0x85,
	0x89, 0x04, 0x85, 0x80, 0x84, 0x18, 0xe0, 0x8a, 0x82, 0x42, 0xfd, 0xfc, 0x40, 0x22, 0x8c, 0x4e,
	0x52, 0x1f, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x8e, 0x47,
	0x72, 0x8c, 0x51, 0x70, 0xe7, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xdd, 0xde, 0x2a, 0xa9,
	0x00, 0x00, 0x00,
}
