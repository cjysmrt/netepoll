package main

import "boost/protoplus/model"

type Context struct {
	*model.DescriptorSet
	OutputFileName string
	PackageName    string
}
