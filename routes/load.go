//
// Copyright (c) 2018 Konstanin Ivanov <kostyarin.ivanov@gmail.com>.
// All rights reserved. This program is free software. It comes without
// any warranty, to the extent permitted by applicable law. You can
// redistribute it and/or modify it under the terms of the Do What
// The Fuck You Want To Public License, Version 2, as published by
// Sam Hocevar. See below for more details.
//

//
//        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//                    Version 2, December 2004
//
// Copyright (C) 2004 Sam Hocevar <sam@hocevar.net>
//
// Everyone is permitted to copy and distribute verbatim or modified
// copies of this license document, and changing it is allowed as long
// as the name is changed.
//
//            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
//   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION
//
//  0. You just DO WHAT THE FUCK YOU WANT TO.
//

// ************************************************************************** //
//                                                                            //
// This gist shows a convenient way to load and use HTML templates in Golang  //
// web-applications. The way includes:                                        //
//                                                                            //
//     - recursive loading, unlike (html/template).ParseGlob does it          //
//     - short (Rails-like) template name without shared prefix (dir)         //
//       and without file extension                                           //
//                                                                            //
// For example, there is a tree of templates                                  //
//         views/                                                             //
//           static/                                                          //
//             home.html                                                      //
//             about.html                                                     //
//             privacypolicy.html                                             //
//             help.html                                                      //
//           user/                                                            //
//             new.html                                                       //
//             edit.html                                                      //
//             show.html                                                      //
//             form.html                                                      //
//           layout/                                                          //
//             head.html                                                      //
//             foot.html                                                      //
//                                                                            //
// Thus, the home.html can include head.html and foor.html following way      //
//                                                                            //
//     {{ template "layout/head" . }}                                         //
//                                                                            //
//     <h1> Home page </h1>                                                   //
//     <!-- other content of the home.html                                    //
//                                                                            //
//     {{ template "layout/foot" . }}                                         //
//                                                                            //
// This is acceptable for user/new and user/edit which can include user/form  //
// along with the layout/head and layout/foot.                                //
//                                                                            //
// ************************************************************************** //

package routes

import (
	"fmt"
	"html/template"
	//"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// A Tmpl implements keeper, loader and reloader for HTML templates
type Tmpl struct {
	
	*template.Template // root template
}

// NewTmpl creates new Tmpl.
func NewTmpl() (tmpl *Tmpl) {
	tmpl = new(Tmpl)
	tmpl.Template = template.New("") // unnamed root template
	return
}

// SetFuncs sets template functions to underlying templates
func (t *Tmpl) SetFuncs(funcMap template.FuncMap) {
	t.Template = t.Template.Funcs(funcMap)
}

// Load templates. The dir argument is a directory to load templates from.
// The ext argument is extension of tempaltes.
func (t *Tmpl) Load(dir, ext string) (err error) {

	// get absolute path
	
	if dir, err = filepath.Abs(dir); err != nil {
		return fmt.Errorf("getting absolute path: %w", err)
	}

	var root = t.Template

	var walkFunc = func(path string, info os.FileInfo, err error) (_ error) {

		// handle walking error if any
		if err != nil {
			return err
		}

		// skip all except regular files
		// TODO (kostyarin): follow symlinks (?)
		if !info.Mode().IsRegular() {
			return
		}

		// filter by extension
		if filepath.Ext(path) != ext {
			return
		}

		// get relative path
		var rel string
		if rel, err = filepath.Rel(dir, path); err != nil {
			return err
		}

		// name of a template is its relative path
		// without extension
		rel = strings.TrimSuffix(rel, ext)
		rel = strings.Join(strings.Split(rel, string(os.PathSeparator)), "/")

		// load or reload
		var (
			nt = root.New(rel)
			b  []byte
		)

		if b, err = os.ReadFile(path); err != nil {
			return err
		}
		//if b, err = ioutil.ReadFile(path); err != nil {
		//	return err
		//}

		_, err = nt.Parse(string(b))
		return err
	}

	if err = filepath.Walk(dir, walkFunc); err != nil {
		return
	}

	t.Template = root // set or replace (does it needed?)
	return
}
