package assets

import (
	"github.com/richardwilkes/toolbox/xio/fs/embedded"
)

// Embedded filesystems
var (
	fs        = embedded.NewFileSystemFromEmbeddedZip("internal/assets")
	DynamicFS = embedded.NewSubFileSystem(fs, "dynamic")
	StaticFS  = embedded.NewSubFileSystem(fs, "static")
)
