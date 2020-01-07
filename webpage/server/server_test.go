// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package server

import (
	"testing"
)

func TestTemplateCaching(t *testing.T) {
	s := NewTemplateServer()
	s.Cache = NewTmplCache()
	err := s.Cache.ParseAndCacheTemplate("testdata/test_template.tpl")
	if err != nil {
		t.Fail()
	}

}
