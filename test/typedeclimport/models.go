package typedeclimport

import subpkg "github.com/ivansukach/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
