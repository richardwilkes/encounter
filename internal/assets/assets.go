package assets

// Embedded filesystems
var (
	DynamicFS = staticFS.FileSystem("internal/assets/dynamic")
	StaticFS  = staticFS.FileSystem("internal/assets/static")
)
