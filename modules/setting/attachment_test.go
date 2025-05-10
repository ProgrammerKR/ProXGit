// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getStorageCustomType(t *testing.T) {
	iniStr := `
[attachment]
STORAGE_TYPE = my_minio
<<<<<<< HEAD
MINIO_BUCKET = gitea-attachment
=======
MINIO_BUCKET = proxgit-attachment
>>>>>>> master

[storage.my_minio]
STORAGE_TYPE = minio
MINIO_ENDPOINT = my_minio:9000
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))

	assert.EqualValues(t, "minio", Attachment.Storage.Type)
	assert.Equal(t, "my_minio:9000", Attachment.Storage.MinioConfig.Endpoint)
<<<<<<< HEAD
	assert.Equal(t, "gitea-attachment", Attachment.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit-attachment", Attachment.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)
}

func Test_getStorageTypeSectionOverridesStorageSection(t *testing.T) {
	iniStr := `
[attachment]
STORAGE_TYPE = minio

[storage.minio]
<<<<<<< HEAD
MINIO_BUCKET = gitea-minio

[storage]
MINIO_BUCKET = gitea
=======
MINIO_BUCKET = proxgit-minio

[storage]
MINIO_BUCKET = proxgit
>>>>>>> master
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))

	assert.EqualValues(t, "minio", Attachment.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea-minio", Attachment.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit-minio", Attachment.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)
}

func Test_getStorageSpecificOverridesStorage(t *testing.T) {
	iniStr := `
[attachment]
STORAGE_TYPE = minio
<<<<<<< HEAD
MINIO_BUCKET = gitea-attachment

[storage.attachments]
MINIO_BUCKET = gitea
=======
MINIO_BUCKET = proxgit-attachment

[storage.attachments]
MINIO_BUCKET = proxgit
>>>>>>> master

[storage]
STORAGE_TYPE = local
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))

	assert.EqualValues(t, "minio", Attachment.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea-attachment", Attachment.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit-attachment", Attachment.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)
}

func Test_getStorageGetDefaults(t *testing.T) {
	cfg, err := NewConfigProviderFromData("")
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))

	// default storage is local, so bucket is empty
	assert.Empty(t, Attachment.Storage.MinioConfig.Bucket)
}

func Test_getStorageInheritNameSectionType(t *testing.T) {
	iniStr := `
[storage.attachments]
STORAGE_TYPE = minio
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))

	assert.EqualValues(t, "minio", Attachment.Storage.Type)
}

func Test_AttachmentStorage(t *testing.T) {
	iniStr := `
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
[storage]
STORAGE_TYPE            = minio
MINIO_ENDPOINT          = s3.my-domain.net
<<<<<<< HEAD
MINIO_BUCKET            = gitea
=======
MINIO_BUCKET            = proxgit
>>>>>>> master
MINIO_LOCATION          = homenet
MINIO_USE_SSL           = true
MINIO_ACCESS_KEY_ID     = correct_key
MINIO_SECRET_ACCESS_KEY = correct_key
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))
	storage := Attachment.Storage

	assert.EqualValues(t, "minio", storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", storage.MinioConfig.Bucket)
>>>>>>> master
}

func Test_AttachmentStorage1(t *testing.T) {
	iniStr := `
[storage]
STORAGE_TYPE = minio
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))
	assert.EqualValues(t, "minio", Attachment.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Attachment.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", Attachment.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)
}
