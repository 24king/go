// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import "cmd/go/internal/base"

// TODO(rsc): The "module code layout" section needs to be written.

var HelpModules = &base.Command{
	UsageLine: "modules",
	Short:     "modules, module versions, and more",
	Long: `
Modules are how Go manages dependencies.

A module is a collection of packages that are released, versioned, and
distributed together. Modules may be downloaded directly from version control
repositories or from module proxy servers.

For a series of tutorials on modules, see
https://golang.org/doc/tutorial/create-module.

For a detailed reference on modules, see https://golang.org/ref/mod.
	`,
}

var HelpGoMod = &base.Command{
	UsageLine: "go.mod",
	Short:     "the go.mod file",
	Long: `
A module version is defined by a tree of source files, with a go.mod
file in its root. When the go command is run, it looks in the current
directory and then successive parent directories to find the go.mod
marking the root of the main (current) module.

The go.mod file format is described in detail at
https://golang.org/ref/mod#go-mod-file.

To create a new go.mod file, use 'go help init'. For details see
'go help mod init' or https://golang.org/ref/mod#go-mod-init.

To add missing module requirements or remove unneeded requirements,
use 'go mod tidy'. For details, see 'go help mod tidy' or
https://golang.org/ref/mod#go-mod-tidy.

To add, upgrade, downgrade, or remove a specific module requirement, use
'go get'. For details, see 'go help module-get' or
https://golang.org/ref/mod#go-get.

To make other changes or to parse go.mod as JSON for use by other tools,
use 'go mod edit'. See 'go help mod edit' or
https://golang.org/ref/mod#go-mod-edit.
	`,
}
