package assets

// Embedded filesystems
var (
	DynamicFS = dynamicFS.FileSystem("internal/assets/dynamic")
	StaticFS  = staticFS.FileSystem("internal/assets/static")
)
