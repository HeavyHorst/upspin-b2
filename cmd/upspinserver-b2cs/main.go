// Copyright 2017 The Upspin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command upspinserver-aws is a combined DirServer and StoreServer for use on
// stand-alone machines. It provides the production implementations of the
// dir and store servers (dir/server and store/server) with support for storage
// in backblaze b2.
package main // import "backblaze.upspin.io/cmd/upspinserver-b2cs"

import (
	"upspin.io/cloud/https"
	"upspin.io/serverutil/upspinserver"

	// Storage on B2cs.
	_ "backblaze.upspin.io/cloud/storage/b2cs"
)

func main() {
	ready := upspinserver.Main()
	https.ListenAndServe(ready, https.OptionsFromFlags())
}
