// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getStorageMultipleName(t *testing.T) {
	iniStr := `
[lfs]
<<<<<<< HEAD
MINIO_BUCKET = gitea-lfs

[attachment]
MINIO_BUCKET = gitea-attachment

[storage]
STORAGE_TYPE = minio
MINIO_BUCKET = gitea-storage
=======
MINIO_BUCKET = proxgit-lfs

[attachment]
MINIO_BUCKET = proxgit-attachment

[storage]
STORAGE_TYPE = minio
MINIO_BUCKET = proxgit-storage
>>>>>>> master
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))
<<<<<<< HEAD
	assert.Equal(t, "gitea-attachment", Attachment.Storage.MinioConfig.Bucket)
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "gitea-lfs", LFS.Storage.MinioConfig.Bucket)
	assert.Equal(t, "lfs/", LFS.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadAvatarsFrom(cfg))
	assert.Equal(t, "gitea-storage", Avatar.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit-attachment", Attachment.Storage.MinioConfig.Bucket)
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "proxgit-lfs", LFS.Storage.MinioConfig.Bucket)
	assert.Equal(t, "lfs/", LFS.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadAvatarsFrom(cfg))
	assert.Equal(t, "proxgit-storage", Avatar.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "avatars/", Avatar.Storage.MinioConfig.BasePath)
}

func Test_getStorageUseOtherNameAsType(t *testing.T) {
	iniStr := `
[attachment]
STORAGE_TYPE = lfs

[storage.lfs]
STORAGE_TYPE = minio
<<<<<<< HEAD
MINIO_BUCKET = gitea-storage
=======
MINIO_BUCKET = proxgit-storage
>>>>>>> master
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadAttachmentFrom(cfg))
<<<<<<< HEAD
	assert.Equal(t, "gitea-storage", Attachment.Storage.MinioConfig.Bucket)
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "gitea-storage", LFS.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit-storage", Attachment.Storage.MinioConfig.Bucket)
	assert.Equal(t, "attachments/", Attachment.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "proxgit-storage", LFS.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "lfs/", LFS.Storage.MinioConfig.BasePath)
}

func Test_getStorageInheritStorageType(t *testing.T) {
	iniStr := `
[storage]
STORAGE_TYPE = minio
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadPackagesFrom(cfg))
	assert.EqualValues(t, "minio", Packages.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Packages.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", Packages.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "packages/", Packages.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.EqualValues(t, "minio", RepoArchive.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", RepoArchive.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", RepoArchive.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "repo-archive/", RepoArchive.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadActionsFrom(cfg))
	assert.EqualValues(t, "minio", Actions.LogStorage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Actions.LogStorage.MinioConfig.Bucket)
	assert.Equal(t, "actions_log/", Actions.LogStorage.MinioConfig.BasePath)

	assert.EqualValues(t, "minio", Actions.ArtifactStorage.Type)
	assert.Equal(t, "gitea", Actions.ArtifactStorage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", Actions.LogStorage.MinioConfig.Bucket)
	assert.Equal(t, "actions_log/", Actions.LogStorage.MinioConfig.BasePath)

	assert.EqualValues(t, "minio", Actions.ArtifactStorage.Type)
	assert.Equal(t, "proxgit", Actions.ArtifactStorage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "actions_artifacts/", Actions.ArtifactStorage.MinioConfig.BasePath)

	assert.NoError(t, loadAvatarsFrom(cfg))
	assert.EqualValues(t, "minio", Avatar.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Avatar.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", Avatar.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "avatars/", Avatar.Storage.MinioConfig.BasePath)

	assert.NoError(t, loadRepoAvatarFrom(cfg))
	assert.EqualValues(t, "minio", RepoAvatar.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", RepoAvatar.Storage.MinioConfig.Bucket)
=======
	assert.Equal(t, "proxgit", RepoAvatar.Storage.MinioConfig.Bucket)
>>>>>>> master
	assert.Equal(t, "repo-avatars/", RepoAvatar.Storage.MinioConfig.BasePath)
}

func Test_getStorageInheritStorageTypeAzureBlob(t *testing.T) {
	iniStr := `
[storage]
STORAGE_TYPE = azureblob
`
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)

	assert.NoError(t, loadPackagesFrom(cfg))
	assert.EqualValues(t, "azureblob", Packages.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Packages.Storage.AzureBlobConfig.Container)
=======
	assert.Equal(t, "proxgit", Packages.Storage.AzureBlobConfig.Container)
>>>>>>> master
	assert.Equal(t, "packages/", Packages.Storage.AzureBlobConfig.BasePath)

	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.EqualValues(t, "azureblob", RepoArchive.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", RepoArchive.Storage.AzureBlobConfig.Container)
=======
	assert.Equal(t, "proxgit", RepoArchive.Storage.AzureBlobConfig.Container)
>>>>>>> master
	assert.Equal(t, "repo-archive/", RepoArchive.Storage.AzureBlobConfig.BasePath)

	assert.NoError(t, loadActionsFrom(cfg))
	assert.EqualValues(t, "azureblob", Actions.LogStorage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Actions.LogStorage.AzureBlobConfig.Container)
	assert.Equal(t, "actions_log/", Actions.LogStorage.AzureBlobConfig.BasePath)

	assert.EqualValues(t, "azureblob", Actions.ArtifactStorage.Type)
	assert.Equal(t, "gitea", Actions.ArtifactStorage.AzureBlobConfig.Container)
=======
	assert.Equal(t, "proxgit", Actions.LogStorage.AzureBlobConfig.Container)
	assert.Equal(t, "actions_log/", Actions.LogStorage.AzureBlobConfig.BasePath)

	assert.EqualValues(t, "azureblob", Actions.ArtifactStorage.Type)
	assert.Equal(t, "proxgit", Actions.ArtifactStorage.AzureBlobConfig.Container)
>>>>>>> master
	assert.Equal(t, "actions_artifacts/", Actions.ArtifactStorage.AzureBlobConfig.BasePath)

	assert.NoError(t, loadAvatarsFrom(cfg))
	assert.EqualValues(t, "azureblob", Avatar.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", Avatar.Storage.AzureBlobConfig.Container)
=======
	assert.Equal(t, "proxgit", Avatar.Storage.AzureBlobConfig.Container)
>>>>>>> master
	assert.Equal(t, "avatars/", Avatar.Storage.AzureBlobConfig.BasePath)

	assert.NoError(t, loadRepoAvatarFrom(cfg))
	assert.EqualValues(t, "azureblob", RepoAvatar.Storage.Type)
<<<<<<< HEAD
	assert.Equal(t, "gitea", RepoAvatar.Storage.AzureBlobConfig.Container)
=======
	assert.Equal(t, "proxgit", RepoAvatar.Storage.AzureBlobConfig.Container)
>>>>>>> master
	assert.Equal(t, "repo-avatars/", RepoAvatar.Storage.AzureBlobConfig.BasePath)
}

type testLocalStoragePathCase struct {
	loader       func(rootCfg ConfigProvider) error
	storagePtr   **Storage
	expectedPath string
}

func testLocalStoragePath(t *testing.T, appDataPath, iniStr string, cases []testLocalStoragePathCase) {
	cfg, err := NewConfigProviderFromData(iniStr)
	assert.NoError(t, err)
	AppDataPath = appDataPath
	for _, c := range cases {
		assert.NoError(t, c.loader(cfg))
		storage := *c.storagePtr

		assert.EqualValues(t, "local", storage.Type)
		assert.True(t, filepath.IsAbs(storage.Path))
		assert.Equal(t, filepath.Clean(c.expectedPath), filepath.Clean(storage.Path))
	}
}

func Test_getStorageInheritStorageTypeLocal(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/appdata/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPath(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/gitea/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/gitea/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/gitea/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/gitea/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/data/gitea/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/gitea/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/gitea/repo-avatars"},
=======
PATH = /data/proxgit
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/proxgit/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/proxgit/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/proxgit/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/proxgit/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/data/proxgit/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/proxgit/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/proxgit/repo-avatars"},
>>>>>>> master
	})
}

func Test_getStorageInheritStorageTypeLocalRelativePath(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
PATH = storages
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/storages/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/storages/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/storages/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/storages/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/storages/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/appdata/storages/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/storages/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/storages/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea

[repo-archive]
PATH = /data/gitea/the-archives-dir
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/gitea/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/gitea/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/gitea/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/gitea/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/the-archives-dir"},
		{loadActionsFrom, &Actions.LogStorage, "/data/gitea/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/gitea/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/gitea/repo-avatars"},
=======
PATH = /data/proxgit

[repo-archive]
PATH = /data/proxgit/the-archives-dir
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/proxgit/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/proxgit/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/proxgit/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/proxgit/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/the-archives-dir"},
		{loadActionsFrom, &Actions.LogStorage, "/data/proxgit/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/proxgit/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/proxgit/repo-avatars"},
>>>>>>> master
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverrideEmpty(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea

[repo-archive]
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/gitea/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/gitea/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/gitea/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/gitea/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/data/gitea/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/gitea/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/gitea/repo-avatars"},
=======
PATH = /data/proxgit

[repo-archive]
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/data/proxgit/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/proxgit/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/proxgit/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/proxgit/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/repo-archive"},
		{loadActionsFrom, &Actions.LogStorage, "/data/proxgit/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/proxgit/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/proxgit/repo-avatars"},
>>>>>>> master
	})
}

func Test_getStorageInheritStorageTypeLocalRelativePathOverride(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea
=======
PATH = /data/proxgit
>>>>>>> master

[repo-archive]
PATH = the-archives-dir
`, []testLocalStoragePathCase{
<<<<<<< HEAD
		{loadAttachmentFrom, &Attachment.Storage, "/data/gitea/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/gitea/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/gitea/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/gitea/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/the-archives-dir"},
		{loadActionsFrom, &Actions.LogStorage, "/data/gitea/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/gitea/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/gitea/repo-avatars"},
=======
		{loadAttachmentFrom, &Attachment.Storage, "/data/proxgit/attachments"},
		{loadLFSFrom, &LFS.Storage, "/data/proxgit/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/data/proxgit/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/data/proxgit/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/the-archives-dir"},
		{loadActionsFrom, &Actions.LogStorage, "/data/proxgit/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/data/proxgit/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/data/proxgit/repo-avatars"},
>>>>>>> master
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride3(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea/archives
=======
PATH = /data/proxgit/archives
>>>>>>> master
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/packages"},
<<<<<<< HEAD
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/archives"},
=======
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/archives"},
>>>>>>> master
		{loadActionsFrom, &Actions.LogStorage, "/appdata/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride3_5(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
STORAGE_TYPE = local
PATH = a-relative-path
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/packages"},
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/a-relative-path"},
		{loadActionsFrom, &Actions.LogStorage, "/appdata/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride4(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea/archives

[repo-archive]
PATH = /tmp/gitea/archives
=======
PATH = /data/proxgit/archives

[repo-archive]
PATH = /tmp/proxgit/archives
>>>>>>> master
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/packages"},
<<<<<<< HEAD
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/tmp/gitea/archives"},
=======
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/tmp/proxgit/archives"},
>>>>>>> master
		{loadActionsFrom, &Actions.LogStorage, "/appdata/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride5(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
STORAGE_TYPE = local
<<<<<<< HEAD
PATH = /data/gitea/archives
=======
PATH = /data/proxgit/archives
>>>>>>> master

[repo-archive]
`, []testLocalStoragePathCase{
		{loadAttachmentFrom, &Attachment.Storage, "/appdata/attachments"},
		{loadLFSFrom, &LFS.Storage, "/appdata/lfs"},
		{loadActionsFrom, &Actions.ArtifactStorage, "/appdata/actions_artifacts"},
		{loadPackagesFrom, &Packages.Storage, "/appdata/packages"},
<<<<<<< HEAD
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/gitea/archives"},
=======
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/data/proxgit/archives"},
>>>>>>> master
		{loadActionsFrom, &Actions.LogStorage, "/appdata/actions_log"},
		{loadAvatarsFrom, &Avatar.Storage, "/appdata/avatars"},
		{loadRepoAvatarFrom, &RepoAvatar.Storage, "/appdata/repo-avatars"},
	})
}

func Test_getStorageInheritStorageTypeLocalPathOverride72(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[repo-archive]
STORAGE_TYPE = local
PATH = archives
`, []testLocalStoragePathCase{
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/archives"},
	})
}

func Test_getStorageConfiguration20(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = my_storage
PATH = archives
`)
	assert.NoError(t, err)

	assert.Error(t, loadRepoArchiveFrom(cfg))
}

func Test_getStorageConfiguration21(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
`, []testLocalStoragePathCase{
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/repo-archive"},
	})
}

func Test_getStorageConfiguration22(t *testing.T) {
	testLocalStoragePath(t, "/appdata", `
[storage.repo-archive]
PATH = archives
`, []testLocalStoragePathCase{
		{loadRepoArchiveFrom, &RepoArchive.Storage, "/appdata/archives"},
	})
}

func Test_getStorageConfiguration23(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
`)
	assert.NoError(t, err)

	_, err = getStorage(cfg, "", "", nil)
	assert.Error(t, err)

	assert.NoError(t, loadRepoArchiveFrom(cfg))
	cp := RepoArchive.Storage.ToShadowCopy()
	assert.Equal(t, "******", cp.MinioConfig.AccessKeyID)
	assert.Equal(t, "******", cp.MinioConfig.SecretAccessKey)
}

func Test_getStorageConfiguration24(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = my_archive

[storage.my_archive]
; unsupported, storage type should be defined explicitly
PATH = archives
`)
	assert.NoError(t, err)
	assert.Error(t, loadRepoArchiveFrom(cfg))
}

func Test_getStorageConfiguration25(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = my_archive

[storage.my_archive]
; unsupported, storage type should be known type
STORAGE_TYPE = unknown // should be local or minio
PATH = archives
`)
	assert.NoError(t, err)
	assert.Error(t, loadRepoArchiveFrom(cfg))
}

func Test_getStorageConfiguration26(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
; wrong configuration
MINIO_USE_SSL = abc
`)
	assert.NoError(t, err)
	// assert.Error(t, loadRepoArchiveFrom(cfg))
	// FIXME: this should return error but now ini package's MapTo() doesn't check type
	assert.NoError(t, loadRepoArchiveFrom(cfg))
}

func Test_getStorageConfiguration27(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[storage.repo-archive]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
MINIO_USE_SSL = true
`)
	assert.NoError(t, err)
	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.Equal(t, "my_access_key", RepoArchive.Storage.MinioConfig.AccessKeyID)
	assert.Equal(t, "my_secret_key", RepoArchive.Storage.MinioConfig.SecretAccessKey)
	assert.True(t, RepoArchive.Storage.MinioConfig.UseSSL)
	assert.Equal(t, "repo-archive/", RepoArchive.Storage.MinioConfig.BasePath)
}

func Test_getStorageConfiguration28(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
MINIO_USE_SSL = true
MINIO_BASE_PATH = /prefix
`)
	assert.NoError(t, err)
	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.Equal(t, "my_access_key", RepoArchive.Storage.MinioConfig.AccessKeyID)
	assert.Equal(t, "my_secret_key", RepoArchive.Storage.MinioConfig.SecretAccessKey)
	assert.True(t, RepoArchive.Storage.MinioConfig.UseSSL)
	assert.Equal(t, "/prefix/repo-archive/", RepoArchive.Storage.MinioConfig.BasePath)

	cfg, err = NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = minio
MINIO_IAM_ENDPOINT = 127.0.0.1
MINIO_USE_SSL = true
MINIO_BASE_PATH = /prefix
`)
	assert.NoError(t, err)
	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.Equal(t, "127.0.0.1", RepoArchive.Storage.MinioConfig.IamEndpoint)
	assert.True(t, RepoArchive.Storage.MinioConfig.UseSSL)
	assert.Equal(t, "/prefix/repo-archive/", RepoArchive.Storage.MinioConfig.BasePath)

	cfg, err = NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
MINIO_USE_SSL = true
MINIO_BASE_PATH = /prefix

[lfs]
MINIO_BASE_PATH = /lfs
`)
	assert.NoError(t, err)
	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "my_access_key", LFS.Storage.MinioConfig.AccessKeyID)
	assert.Equal(t, "my_secret_key", LFS.Storage.MinioConfig.SecretAccessKey)
	assert.True(t, LFS.Storage.MinioConfig.UseSSL)
	assert.Equal(t, "/lfs", LFS.Storage.MinioConfig.BasePath)

	cfg, err = NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = minio
MINIO_ACCESS_KEY_ID = my_access_key
MINIO_SECRET_ACCESS_KEY = my_secret_key
MINIO_USE_SSL = true
MINIO_BASE_PATH = /prefix

[storage.lfs]
MINIO_BASE_PATH = /lfs
`)
	assert.NoError(t, err)
	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "my_access_key", LFS.Storage.MinioConfig.AccessKeyID)
	assert.Equal(t, "my_secret_key", LFS.Storage.MinioConfig.SecretAccessKey)
	assert.True(t, LFS.Storage.MinioConfig.UseSSL)
	assert.Equal(t, "/lfs", LFS.Storage.MinioConfig.BasePath)
}

func Test_getStorageConfiguration29(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[repo-archive]
STORAGE_TYPE = azureblob
AZURE_BLOB_ACCOUNT_NAME = my_account_name
AZURE_BLOB_ACCOUNT_KEY = my_account_key
`)
	assert.NoError(t, err)
	// assert.Error(t, loadRepoArchiveFrom(cfg))
	// FIXME: this should return error but now ini package's MapTo() doesn't check type
	assert.NoError(t, loadRepoArchiveFrom(cfg))
}

func Test_getStorageConfiguration30(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[storage.repo-archive]
STORAGE_TYPE = azureblob
AZURE_BLOB_ACCOUNT_NAME = my_account_name
AZURE_BLOB_ACCOUNT_KEY = my_account_key
`)
	assert.NoError(t, err)
	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.Equal(t, "my_account_name", RepoArchive.Storage.AzureBlobConfig.AccountName)
	assert.Equal(t, "my_account_key", RepoArchive.Storage.AzureBlobConfig.AccountKey)
	assert.Equal(t, "repo-archive/", RepoArchive.Storage.AzureBlobConfig.BasePath)
}

func Test_getStorageConfiguration31(t *testing.T) {
	cfg, err := NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = azureblob
AZURE_BLOB_ACCOUNT_NAME = my_account_name
AZURE_BLOB_ACCOUNT_KEY = my_account_key
AZURE_BLOB_BASE_PATH = /prefix
`)
	assert.NoError(t, err)
	assert.NoError(t, loadRepoArchiveFrom(cfg))
	assert.Equal(t, "my_account_name", RepoArchive.Storage.AzureBlobConfig.AccountName)
	assert.Equal(t, "my_account_key", RepoArchive.Storage.AzureBlobConfig.AccountKey)
	assert.Equal(t, "/prefix/repo-archive/", RepoArchive.Storage.AzureBlobConfig.BasePath)

	cfg, err = NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = azureblob
AZURE_BLOB_ACCOUNT_NAME = my_account_name
AZURE_BLOB_ACCOUNT_KEY = my_account_key
AZURE_BLOB_BASE_PATH = /prefix

[lfs]
AZURE_BLOB_BASE_PATH = /lfs
`)
	assert.NoError(t, err)
	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "my_account_name", LFS.Storage.AzureBlobConfig.AccountName)
	assert.Equal(t, "my_account_key", LFS.Storage.AzureBlobConfig.AccountKey)
	assert.Equal(t, "/lfs", LFS.Storage.AzureBlobConfig.BasePath)

	cfg, err = NewConfigProviderFromData(`
[storage]
STORAGE_TYPE = azureblob
AZURE_BLOB_ACCOUNT_NAME = my_account_name
AZURE_BLOB_ACCOUNT_KEY = my_account_key
AZURE_BLOB_BASE_PATH = /prefix

[storage.lfs]
AZURE_BLOB_BASE_PATH = /lfs
`)
	assert.NoError(t, err)
	assert.NoError(t, loadLFSFrom(cfg))
	assert.Equal(t, "my_account_name", LFS.Storage.AzureBlobConfig.AccountName)
	assert.Equal(t, "my_account_key", LFS.Storage.AzureBlobConfig.AccountKey)
	assert.Equal(t, "/lfs", LFS.Storage.AzureBlobConfig.BasePath)
}
