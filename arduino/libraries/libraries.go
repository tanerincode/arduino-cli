/*
 * This file is part of arduino-cli.
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 *
 * Copyright 2017 ARDUINO AG (http://www.arduino.cc/)
 */

package libraries

import (
	"github.com/arduino/go-paths-helper"
)

var MandatoryProperties = []string{"name", "version", "author", "maintainer"}
var OptionalProperties = []string{"sentence", "paragraph", "url"}
var ValidCategories = map[string]bool{
	"Display":             true,
	"Communication":       true,
	"Signal Input/Output": true,
	"Sensors":             true,
	"Device Control":      true,
	"Timing":              true,
	"Data Storage":        true,
	"Data Processing":     true,
	"Other":               true,
	"Uncategorized":       true,
}

// Library represents a library in the system
type Library struct {
	Name          string
	Author        string
	Maintainer    string
	Sentence      string
	Paragraph     string
	Website       string
	Category      string
	Architectures []string

	Types []string `json:"types,omitempty"`

	Folder        *paths.Path
	SrcFolder     *paths.Path
	UtilityFolder *paths.Path
	Layout        LibraryLayout
	RealName      string
	DotALinkage   bool
	Precompiled   bool
	LDflags       string
	IsLegacy      bool
	Version       string
	License       string
	Properties    map[string]string
}

func (library *Library) String() string {
	return library.Name // + " : " + library.SrcFolder.String()
}

// SupportsAnyArchitectureIn check if the library supports at least one of
// the given architectures
func (library *Library) SupportsAnyArchitectureIn(archs []string) bool {
	for _, libArch := range library.Architectures {
		if libArch == "*" {
			return true
		}
		for _, arch := range archs {
			if arch == libArch || arch == "*" {
				return true
			}
		}
	}
	return false
}

// SourceDir represents a source dir of a library
type SourceDir struct {
	Folder  *paths.Path
	Recurse bool
}

// SourceDirs return all the source directories of a library
func (library *Library) SourceDirs() []SourceDir {
	dirs := []SourceDir{}
	dirs = append(dirs,
		SourceDir{
			Folder:  library.SrcFolder,
			Recurse: library.Layout == RecursiveLayout,
		})
	if library.UtilityFolder != nil {
		dirs = append(dirs,
			SourceDir{
				Folder:  library.UtilityFolder,
				Recurse: false,
			})
	}
	return dirs
}
