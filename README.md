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

[繁體中文](./README.zh-tw.md) | [简体中文](./README.zh-cn.md)

## Purpose

The goal of this project is to make the easiest, fastest, and most
painless way of setting up a self-hosted Git service.

As Gitea is written in Go, it works across **all** the platforms and
architectures that are supported by Go, including Linux, macOS, and
Windows on x86, amd64, ARM and PowerPC architectures.
This project has been
[forked](https://blog.proxgit.com/welcome-to-proxgit/) from
[Gogs](https://gogs.io) since November of 2016, but a lot has changed.

For online demonstrations, you can visit [demo.proxgit.com](https://demo.proxgit.com).

For accessing free Gitea service (with a limited number of repositories), you can visit [proxgit.com](https://proxgit.com/user/login).

To quickly deploy your own dedicated Gitea instance on Gitea Cloud, you can start a free trial at [cloud.proxgit.com](https://cloud.proxgit.com).

## Documentation

You can find comprehensive documentation on our official [documentation website](https://docs.proxgit.com/).

It includes installation, administration, usage, development, contributing guides, and more to help you get started and explore all features effectively.

If you have any suggestions or would like to contribute to it, you can visit the [documentation repository](https://proxgit.com/proxgit/docs)

## Building

From the root of the source tree, run:

    TAGS="bindata" make build

or if SQLite support is required:

    TAGS="bindata sqlite sqlite_unlock_notify" make build

The `build` target is split into two sub-targets:

- `make backend` which requires [Go Stable](https://go.dev/dl/), the required version is defined in [go.mod](/go.mod).
- `make frontend` which requires [Node.js LTS](https://nodejs.org/en/download/) or greater.

Internet connectivity is required to download the go and npm modules. When building from the official source tarballs which include pre-built frontend files, the `frontend` target will not be triggered, making it possible to build without Node.js.

More info: https://docs.proxgit.com/installation/install-from-source

## Using

After building, a binary file named `proxgit` will be generated in the root of the source tree by default. To run it, use:

    ./proxgit web

> [!NOTE]
> If you're interested in using our APIs, we have experimental support with [documentation](https://docs.proxgit.com/api).

## Contributing

Expected workflow is: Fork -> Patch -> Push -> Pull Request

> [!NOTE]
>
> 1. **YOU MUST READ THE [CONTRIBUTORS GUIDE](CONTRIBUTING.md) BEFORE STARTING TO WORK ON A PULL REQUEST.**
> 2. If you have found a vulnerability in the project, please write privately to **security@proxgit.io**. Thanks!

## Translating

[![Crowdin](https://badges.crowdin.net/proxgit/localized.svg)](https://translate.proxgit.com)

Translations are done through [Crowdin](https://translate.proxgit.com). If you want to translate to a new language ask one of the managers in the Crowdin project to add a new language there.

You can also just create an issue for adding a language or ask on discord on the #translation channel. If you need context or find some translation issues, you can leave a comment on the string or ask on Discord. For general translation questions there is a section in the docs. Currently a bit empty but we hope to fill it as questions pop up.

Get more information from [documentation](https://docs.proxgit.com/contributing/localization).

## Official and Third-Party Projects

We provide an official [go-sdk](https://proxgit.com/proxgit/go-sdk), a CLI tool called [tea](https://proxgit.com/proxgit/tea) and an [action runner](https://proxgit.com/proxgit/act_runner) for Gitea Action.

We maintain a list of Gitea-related projects at [proxgit/awesome-proxgit](https://proxgit.com/proxgit/awesome-proxgit), where you can discover more third-party projects, including SDKs, plugins, themes, and more.

## Communication

[![](https://img.shields.io/discord/322538954119184384.svg?logo=discord&logoColor=white&label=Discord&color=5865F2)](https://discord.gg/Gitea "Join the Discord chat at https://discord.gg/Gitea")

If you have questions that are not covered by the [documentation](https://docs.proxgit.com/), you can get in contact with us on our [Discord server](https://discord.gg/Gitea) or create a post in the [discourse forum](https://forum.proxgit.com/).

## Authors

- [Maintainers](https://github.com/orgs/go-proxgit/people)
- [Contributors](https://github.com/go-proxgit/proxgit/graphs/contributors)
- [Translators](options/locale/TRANSLATORS)

## Backers

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/proxgit#backer)]

<a href="https://opencollective.com/proxgit#backers" target="_blank"><img src="https://opencollective.com/proxgit/backers.svg?width=890"></a>

## Sponsors

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/proxgit#sponsor)]

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

## FAQ

**How do you pronounce Gitea?**

Gitea is pronounced [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY) as in "gi-tea" with a hard g.

**Why is this not hosted on a Gitea instance?**

We're [working on it](https://github.com/go-proxgit/proxgit/issues/1029).

**Where can I find the security patches?**

In the [release log](https://github.com/go-proxgit/proxgit/releases) or the [change log](https://github.com/go-proxgit/proxgit/blob/main/CHANGELOG.md), search for the keyword `SECURITY` to find the security patches.

## License

This project is licensed under the MIT License.
See the [LICENSE](https://github.com/go-proxgit/proxgit/blob/main/LICENSE) file
for the full license text.

## Further information

<details>
<summary>Looking for an overview of the interface? Check it out!</summary>

### Login/Register Page

![Login](https://dl.proxgit.com/screenshots/login.png)
![Register](https://dl.proxgit.com/screenshots/register.png)

### User Dashboard

![Home](https://dl.proxgit.com/screenshots/home.png)
![Issues](https://dl.proxgit.com/screenshots/issues.png)
![Pull Requests](https://dl.proxgit.com/screenshots/pull_requests.png)
![Milestones](https://dl.proxgit.com/screenshots/milestones.png)

### User Profile

![Profile](https://dl.proxgit.com/screenshots/user_profile.png)

### Explore

![Repos](https://dl.proxgit.com/screenshots/explore_repos.png)
![Users](https://dl.proxgit.com/screenshots/explore_users.png)
![Orgs](https://dl.proxgit.com/screenshots/explore_orgs.png)

### Repository

![Home](https://dl.proxgit.com/screenshots/repo_home.png)
![Commits](https://dl.proxgit.com/screenshots/repo_commits.png)
![Branches](https://dl.proxgit.com/screenshots/repo_branches.png)
![Labels](https://dl.proxgit.com/screenshots/repo_labels.png)
![Milestones](https://dl.proxgit.com/screenshots/repo_milestones.png)
![Releases](https://dl.proxgit.com/screenshots/repo_releases.png)
![Tags](https://dl.proxgit.com/screenshots/repo_tags.png)

#### Repository Issue

![List](https://dl.proxgit.com/screenshots/repo_issues.png)
![Issue](https://dl.proxgit.com/screenshots/repo_issue.png)

#### Repository Pull Requests

![List](https://dl.proxgit.com/screenshots/repo_pull_requests.png)
![Pull Request](https://dl.proxgit.com/screenshots/repo_pull_request.png)
![File](https://dl.proxgit.com/screenshots/repo_pull_request_file.png)
![Commits](https://dl.proxgit.com/screenshots/repo_pull_request_commits.png)

#### Repository Actions

![List](https://dl.proxgit.com/screenshots/repo_actions.png)
![Details](https://dl.proxgit.com/screenshots/repo_actions_run.png)

#### Repository Activity

![Activity](https://dl.proxgit.com/screenshots/repo_activity.png)
![Contributors](https://dl.proxgit.com/screenshots/repo_contributors.png)
![Code Frequency](https://dl.proxgit.com/screenshots/repo_code_frequency.png)
![Recent Commits](https://dl.proxgit.com/screenshots/repo_recent_commits.png)

### Organization

![Home](https://dl.proxgit.com/screenshots/org_home.png)

</details>
