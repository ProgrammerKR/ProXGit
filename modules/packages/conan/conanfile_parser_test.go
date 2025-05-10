// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package conan

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name             = "ConanPackage"
	version          = "1.2"
	license          = "MIT"
	author           = "Gitea <info@proxgit.io>"
	homepage         = "https://proxgit.io/"
	url              = "https://proxgit.com/"
	description      = "Description of ConanPackage"
	topic1           = "proxgit"
	topic2           = "conan"
	contentConanfile = `from conans import ConanFile, CMake, tools

class ConanPackageConan(ConanFile):
    name = "` + name + `"
    version = "` + version + `"
    license = "` + license + `"
    author = "` + author + `"
    homepage = "` + homepage + `"
    url = "` + url + `"
    description = "` + description + `"
    topics = ("` + topic1 + `", "` + topic2 + `")
    settings = "os", "compiler", "build_type", "arch"
    options = {"shared": [True, False], "fPIC": [True, False]}
    default_options = {"shared": False, "fPIC": True}
    generators = "cmake"
`
)

func TestParseConanfile(t *testing.T) {
	metadata, err := ParseConanfile(strings.NewReader(contentConanfile))
	assert.NoError(t, err)
	assert.Equal(t, license, metadata.License)
	assert.Equal(t, author, metadata.Author)
	assert.Equal(t, homepage, metadata.ProjectURL)
	assert.Equal(t, url, metadata.RepositoryURL)
	assert.Equal(t, description, metadata.Description)
	assert.Equal(t, []string{topic1, topic2}, metadata.Keywords)
}
