// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
