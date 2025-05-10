# Gitea

[![](https://github.com/go-proxgit/proxgit/actions/workflows/release-nightly.yml/badge.svg?branch=main)](https://github.com/go-proxgit/proxgit/actions/workflows/release-nightly.yml?query=branch%3Amain "Release Nightly")
[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")
[![](https://goreportcard.com/badge/code.proxgit.io/proxgit)](https://goreportcard.com/report/code.proxgit.io/proxgit "Go Report Card")
[![](https://pkg.go.dev/badge/code.proxgit.io/proxgit?status.svg)](https://pkg.go.dev/code.proxgit.io/proxgit "GoDoc")
[![](https://img.shields.io/github/release/go-proxgit/proxgit.svg)](https://github.com/go-proxgit/proxgit/releases/latest "GitHub release")
[![](https://www.codetriage.com/go-proxgit/proxgit/badges/users.svg)](https://www.codetriage.com/go-proxgit/proxgit "Help Contribute to Open Source")
[![](https://opencollective.com/proxgit/tiers/backers/badge.svg?label=backers&color=brightgreen)](https://opencollective.com/proxgit "Become a backer/sponsor of proxgit")
[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT "License: MIT")
[![Contribute with Gitpod](https://img.shields.io/badge/Contribute%20with-Gitpod-908a85?logo=gitpod&color=green)](https://gitpod.io/#https://github.com/go-proxgit/proxgit)
[![](https://badges.crowdin.net/proxgit/localized.svg)](https://translate.proxgit.com "Crowdin")

[English](./README.md) | [简体中文](./README.zh-cn.md)

## 目的

這個項目的目標是提供最簡單、最快速、最無痛的方式來設置自託管的 Git 服務。

由於 Gitea 是用 Go 語言編寫的，它可以在 Go 支援的所有平台和架構上運行，包括 Linux、macOS 和 Windows 的 x86、amd64、ARM 和 PowerPC 架構。這個項目自 2016 年 11 月從 [Gogs](https://gogs.io) [分叉](https://blog.proxgit.com/welcome-to-proxgit/) 而來，但已經有了很多變化。

在線演示可以訪問 [demo.proxgit.com](https://demo.proxgit.com)。

要訪問免費的 Gitea 服務（有一定數量的倉庫限制），可以訪問 [proxgit.com](https://proxgit.com/user/login)。

要快速部署您自己的專用 Gitea 實例，可以在 [cloud.proxgit.com](https://cloud.proxgit.com) 開始免費試用。

## 文件

您可以在我們的官方 [文件網站](https://docs.proxgit.com/) 上找到全面的文件。

它包括安裝、管理、使用、開發、貢獻指南等，幫助您快速入門並有效地探索所有功能。

如果您有任何建議或想要貢獻，可以訪問 [文件倉庫](https://proxgit.com/proxgit/docs)

## 構建

從源代碼樹的根目錄運行：

    TAGS="bindata" make build

如果需要 SQLite 支援：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目標分為兩個子目標：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本在 [go.mod](/go.mod) 中定義。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更高版本。

需要互聯網連接來下載 go 和 npm 模塊。從包含預構建前端文件的官方源代碼壓縮包構建時，不會觸發 `frontend` 目標，因此可以在沒有 Node.js 的情況下構建。

更多信息：https://docs.proxgit.com/installation/install-from-source

## 使用

構建後，默認情況下會在源代碼樹的根目錄生成一個名為 `proxgit` 的二進制文件。要運行它，請使用：

    ./proxgit web

> [!注意]
> 如果您對使用我們的 API 感興趣，我們提供了實驗性支援，並附有 [文件](https://docs.proxgit.com/api)。

## 貢獻

預期的工作流程是：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **在開始進行 Pull Request 之前，您必須閱讀 [貢獻者指南](CONTRIBUTING.md)。**
> 2. 如果您在項目中發現了漏洞，請私下寫信給 **security@proxgit.io**。謝謝！

## 翻譯

[![Crowdin](https://badges.crowdin.net/proxgit/localized.svg)](https://translate.proxgit.com)

翻譯通過 [Crowdin](https://translate.proxgit.com) 進行。如果您想翻譯成新的語言，請在 Crowdin 項目中請求管理員添加新語言。

您也可以創建一個 issue 來添加語言，或者在 discord 的 #translation 頻道上詢問。如果您需要上下文或發現一些翻譯問題，可以在字符串上留言或在 Discord 上詢問。對於一般的翻譯問題，文檔中有一個部分。目前有點空，但我們希望隨著問題的出現而填充它。

更多信息請參閱 [文件](https://docs.proxgit.com/contributing/localization)。

## 官方和第三方項目

我們提供了一個官方的 [go-sdk](https://proxgit.com/proxgit/go-sdk)，一個名為 [tea](https://proxgit.com/proxgit/tea) 的 CLI 工具和一個 Gitea Action 的 [action runner](https://proxgit.com/proxgit/act_runner)。

我們在 [proxgit/awesome-proxgit](https://proxgit.com/proxgit/awesome-proxgit) 維護了一個 Gitea 相關項目的列表，您可以在那裡發現更多的第三方項目，包括 SDK、插件、主題等。

## 通訊

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

如果您有任何文件未涵蓋的問題，可以在我們的 [Discord 服務器](https://discord.gg/Gitea) 上與我們聯繫，或者在 [discourse 論壇](https://forum.proxgit.com/) 上創建帖子。

## 作者

- [維護者](https://github.com/orgs/go-proxgit/people)
- [貢獻者](https://github.com/go-proxgit/proxgit/graphs/contributors)
- [翻譯者](options/locale/TRANSLATORS)

## 支持者

感謝所有支持者！ 🙏 [[成為支持者](https://opencollective.com/proxgit#backer)]

<a href="https://opencollective.com/proxgit#backers" target="_blank"><img src="https://opencollective.com/proxgit/backers.svg?width=890"></a>

## 贊助商

通過成為贊助商來支持這個項目。您的標誌將顯示在這裡，並帶有鏈接到您的網站。 [[成為贊助商](https://opencollective.com/proxgit#sponsor)]

<a href="https://opencollective.com/proxgit/sponsor/0/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/1/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/2/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/3/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/4/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/5/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/6/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/7/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/8/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/proxgit/sponsor/9/website" target="_blank"><img src="https://opencollective.com/proxgit/sponsor/9/avatar.svg"></a>

## 常見問題

**Gitea 怎麼發音？**

Gitea 的發音是 [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY)，就像 "gi-tea" 一樣，g 是硬音。

**為什麼這個項目沒有託管在 Gitea 實例上？**

我們正在 [努力](https://github.com/go-proxgit/proxgit/issues/1029)。

**在哪裡可以找到安全補丁？**

在 [發佈日誌](https://github.com/go-proxgit/proxgit/releases) 或 [變更日誌](https://github.com/go-proxgit/proxgit/blob/main/CHANGELOG.md) 中，搜索關鍵詞 `SECURITY` 以找到安全補丁。

## 許可證

這個項目是根據 MIT 許可證授權的。
請參閱 [LICENSE](https://github.com/go-proxgit/proxgit/blob/main/LICENSE) 文件以獲取完整的許可證文本。

## 進一步信息

<details>
<summary>尋找界面概述？查看這裡！</summary>

### 登錄/註冊頁面

![Login](https://dl.proxgit.com/screenshots/login.png)
![Register](https://dl.proxgit.com/screenshots/register.png)

### 用戶儀表板

![Home](https://dl.proxgit.com/screenshots/home.png)
![Issues](https://dl.proxgit.com/screenshots/issues.png)
![Pull Requests](https://dl.proxgit.com/screenshots/pull_requests.png)
![Milestones](https://dl.proxgit.com/screenshots/milestones.png)

### 用戶資料

![Profile](https://dl.proxgit.com/screenshots/user_profile.png)

### 探索

![Repos](https://dl.proxgit.com/screenshots/explore_repos.png)
![Users](https://dl.proxgit.com/screenshots/explore_users.png)
![Orgs](https://dl.proxgit.com/screenshots/explore_orgs.png)

### 倉庫

![Home](https://dl.proxgit.com/screenshots/repo_home.png)
![Commits](https://dl.proxgit.com/screenshots/repo_commits.png)
![Branches](https://dl.proxgit.com/screenshots/repo_branches.png)
![Labels](https://dl.proxgit.com/screenshots/repo_labels.png)
![Milestones](https://dl.proxgit.com/screenshots/repo_milestones.png)
![Releases](https://dl.proxgit.com/screenshots/repo_releases.png)
![Tags](https://dl.proxgit.com/screenshots/repo_tags.png)

#### 倉庫問題

![List](https://dl.proxgit.com/screenshots/repo_issues.png)
![Issue](https://dl.proxgit.com/screenshots/repo_issue.png)

#### 倉庫拉取請求

![List](https://dl.proxgit.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.proxgit.com/screenshots/repo_pull_request.png)
![File](https://dl.proxgit.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.proxgit.com/screenshots/repo_pull_request_commits.png)

#### 倉庫操作

![List](https://dl.proxgit.com/screenshots/repo_actions.png)
![Details](https://dl.proxgit.com/screenshots/repo_actions_run.png)

#### 倉庫活動

![Activity](https://dl.proxgit.com/screenshots/repo_activity.png)
![Contributors](https://dl.proxgit.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.proxgit.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.proxgit.com/screenshots/repo_recent_commits.png)

### 組織

![Home](https://dl.proxgit.com/screenshots/org_home.png)

</details>
