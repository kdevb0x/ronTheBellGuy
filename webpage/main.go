// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package main

import (
	"log"
	"time"

	"github.com/kdevb0x/ronTheBellGuy/webpage/server"
)

const static = "assets/"

func main() {
	s := server.NewTemplateServer()
	log.Printf("started %s\n", time.Now())

}
