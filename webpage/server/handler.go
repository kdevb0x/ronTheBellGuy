// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var idx = template.New("index")

// TemplateServer listens on Addr for incomming web connections and server a
// posibly already parsed and executed, then cached static file. It uses a mux
// .Router to route and multiplex html traffic.
type TemplateServer struct {
	Addr string // listen address and port
	Mux  *mux.Router
	mux  *http.ServeMux // maybe remove this or Muxer

	// cache of parsed templates
	Cache *TmplCache
}

func NewTemplateServer() *TemplateServer {
	return &TemplateServer{Mux: mux.NewRouter()}
}

func (s *TemplateServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	s.Mux.Handle("/", s.Mux) // ?
	s.Mux.ServeHTTP(w, r)

}

func (s *TemplateServer) AddRoute(route string, handler http.HandlerFunc) error {
	s.Mux.HandleFunc(route, handler)
	return nil
}

// TmplCache is an in memory cache of parsed templates and their rendered static file
// once executed.
type TmplCache struct {
	// cache size in bytes
	Count int

	// maps filename to the bytes of the parsed file
	Tmpl map[string]*template.Template

	// maps name to the static file bytes after template exec.
	Static map[string][]byte
}

func NewTmplCache() *TmplCache {
	return &TmplCache{Tmpl: make(map[string]*template.Template), Static: make(map[string][]byte)}
}

func (c *TmplCache) ParseAndCacheTemplate(filepath string) error {
	t, err := template.ParseFiles(filepath)
	if err != nil {
		return err
	}
	c.Tmpl[filepath] = t
	c.Count++
	return nil
}

func (c *TmplCache) ExecuteCachedTemplate(name string, data interface{}) error {
	t, ok := c.Tmpl[name]
	if !ok {
		return fmt.Errorf("error template %s not found in cache\n", name)
	}
	var buf bytes.Buffer

	err := t.Execute(buf, data)
	if err != nil {
		return err
	}
	c.Static[t.Name()] = buf.Bytes()
	return nil
}
