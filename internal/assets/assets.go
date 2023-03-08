// Copyright ©2018-2023 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package assets

import (
	"embed"
	"io/fs"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

var (
	// DynamicFS holds the dynamic templates.
	DynamicFS fs.FS
	// StaticFS holds the static files.
	StaticFS fs.FS
	//go:embed dynamic
	dynamicFS embed.FS
	//go:embed static
	staticFS embed.FS
)

func init() {
	var err error
	if DynamicFS, err = fs.Sub(dynamicFS, "dynamic"); err != nil {
		jot.Fatal(1, errs.Wrap(err))
	}
	if StaticFS, err = fs.Sub(staticFS, "static"); err != nil {
		jot.Fatal(1, errs.Wrap(err))
	}
}
