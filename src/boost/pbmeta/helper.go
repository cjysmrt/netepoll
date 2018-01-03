package pbmeta

import (
	"io/ioutil"

	pbprotos "boost/pbmeta/proto"
	pbcompiler "boost/pbmeta/proto/compiler"
	"github.com/golang/protobuf/proto"
)

// 文件格式必须是由protoc-gen-meta输出的CodeGeneratorRequest格式
func LoadFileDescriptorSet(filename string) (*pbprotos.FileDescriptorSet, error) {

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var req pbcompiler.CodeGeneratorRequest

	if err := proto.Unmarshal(content, &req); err != nil {
		return nil, err
	}

	return &pbprotos.FileDescriptorSet{
		File: req.ProtoFile,
	}, nil
}

// 根据文件描述符创建描述池
func CreatePoolByFile(filename string) (*DescriptorPool, error) {
	fds, err := LoadFileDescriptorSet(filename)
	if err != nil {
		return nil, err
	}

	return NewDescriptorPool(fds), nil
}
