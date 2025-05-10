#!/usr/bin/env bash
<<<<<<< HEAD
# This is an update script for gitea installed via the binary distribution
# from dl.gitea.com on linux as systemd service. It performs a backup and updates
=======
# This is an update script for proxgit installed via the binary distribution
# from dl.proxgit.com on linux as systemd service. It performs a backup and updates
>>>>>>> master
# Gitea in place.
# NOTE: This adds the GPG Signing Key of the Gitea maintainers to the keyring.
# Depends on: bash, curl, xz, sha256sum. optionally jq, gpg
#   See section below for available environment vars.
#   When no version is specified, updates to the latest release.
# Examples:
#   upgrade.sh 1.15.10
<<<<<<< HEAD
#   giteahome=/opt/gitea giteaconf=$giteahome/app.ini upgrade.sh

# Check if gitea service is running
if ! pidof gitea &> /dev/null; then
  echo "Error: gitea is not running."
  exit 1
fi

# Continue with rest of the script if gitea is running
echo "Gitea is running. Continuing with rest of script..."

# apply variables from environment
: "${giteabin:="/usr/local/bin/gitea"}"
: "${giteahome:="/var/lib/gitea"}"
: "${giteaconf:="/etc/gitea/app.ini"}"
: "${giteauser:="git"}"
: "${sudocmd:="sudo"}"
: "${arch:="linux-amd64"}"
: "${service_start:="$sudocmd systemctl start gitea"}"
: "${service_stop:="$sudocmd systemctl stop gitea"}"
: "${service_status:="$sudocmd systemctl status gitea"}"
: "${backupopts:=""}" # see `gitea dump --help` for available options

function giteacmd {
  if [[ $sudocmd = "su" ]]; then
    # `-c` only accept one string as argument.
    "$sudocmd" - "$giteauser" -c "$(printf "%q " "$giteabin" "--config" "$giteaconf" "--work-path" "$giteahome" "$@")"
  else
    "$sudocmd" --user "$giteauser" "$giteabin" --config "$giteaconf" --work-path "$giteahome" "$@"
=======
#   proxgithome=/opt/proxgit proxgitconf=$proxgithome/app.ini upgrade.sh

# Check if proxgit service is running
if ! pidof proxgit &> /dev/null; then
  echo "Error: proxgit is not running."
  exit 1
fi

# Continue with rest of the script if proxgit is running
echo "Gitea is running. Continuing with rest of script..."

# apply variables from environment
: "${proxgitbin:="/usr/local/bin/proxgit"}"
: "${proxgithome:="/var/lib/proxgit"}"
: "${proxgitconf:="/etc/proxgit/app.ini"}"
: "${proxgituser:="git"}"
: "${sudocmd:="sudo"}"
: "${arch:="linux-amd64"}"
: "${service_start:="$sudocmd systemctl start proxgit"}"
: "${service_stop:="$sudocmd systemctl stop proxgit"}"
: "${service_status:="$sudocmd systemctl status proxgit"}"
: "${backupopts:=""}" # see `proxgit dump --help` for available options

function proxgitcmd {
  if [[ $sudocmd = "su" ]]; then
    # `-c` only accept one string as argument.
    "$sudocmd" - "$proxgituser" -c "$(printf "%q " "$proxgitbin" "--config" "$proxgitconf" "--work-path" "$proxgithome" "$@")"
  else
    "$sudocmd" --user "$proxgituser" "$proxgitbin" --config "$proxgitconf" --work-path "$proxgithome" "$@"
>>>>>>> master
  fi
}

function require {
  for exe in "$@"; do
    command -v "$exe" &>/dev/null || (echo "missing dependency '$exe'"; exit 1)
  done
}

# parse command line arguments
while true; do
  case "$1" in
<<<<<<< HEAD
    -v | --version ) giteaversion="$2"; shift 2 ;;
=======
    -v | --version ) proxgitversion="$2"; shift 2 ;;
>>>>>>> master
    -y | --yes ) no_confirm="yes"; shift ;;
    --ignore-gpg) ignore_gpg="yes"; shift ;;
    "" | -- ) shift; break ;;
    * ) echo "Usage:  [<environment vars>] upgrade.sh [-v <version>] [-y] [--ignore-gpg]"; exit 1;; 
  esac
done

# exit once any command fails. this means that each step should be idempotent!
set -euo pipefail

if [[ -f /etc/os-release ]]; then
  os_release=$(cat /etc/os-release)

  if [[ "$os_release" =~ "OpenWrt" ]]; then
    sudocmd="su"
<<<<<<< HEAD
    service_start="/etc/init.d/gitea start"
    service_stop="/etc/init.d/gitea stop"
    service_status="/etc/init.d/gitea status"
=======
    service_start="/etc/init.d/proxgit start"
    service_stop="/etc/init.d/proxgit stop"
    service_status="/etc/init.d/proxgit status"
>>>>>>> master
  else
    require systemctl
  fi
fi

require curl xz sha256sum "$sudocmd"

# select version to install
<<<<<<< HEAD
if [[ -z "${giteaversion:-}" ]]; then
  require jq
  giteaversion=$(curl --connect-timeout 10 -sL https://dl.gitea.com/gitea/version.json | jq -r .latest.version)
  echo "Latest available version is $giteaversion"
=======
if [[ -z "${proxgitversion:-}" ]]; then
  require jq
  proxgitversion=$(curl --connect-timeout 10 -sL https://dl.proxgit.com/proxgit/version.json | jq -r .latest.version)
  echo "Latest available version is $proxgitversion"
>>>>>>> master
fi

# confirm update
echo "Checking currently installed version..."
<<<<<<< HEAD
current=$(giteacmd --version | cut -d ' ' -f 3)
[[ "$current" == "$giteaversion" ]] && echo "$current is already installed, stopping." && exit 1
if [[ -z "${no_confirm:-}"  ]]; then
  echo "Make sure to read the changelog first: https://github.com/go-gitea/gitea/blob/main/CHANGELOG.md"
  echo "Are you ready to update Gitea from ${current} to ${giteaversion}? (y/N)"
=======
current=$(proxgitcmd --version | cut -d ' ' -f 3)
[[ "$current" == "$proxgitversion" ]] && echo "$current is already installed, stopping." && exit 1
if [[ -z "${no_confirm:-}"  ]]; then
  echo "Make sure to read the changelog first: https://github.com/go-proxgit/proxgit/blob/main/CHANGELOG.md"
  echo "Are you ready to update Gitea from ${current} to ${proxgitversion}? (y/N)"
>>>>>>> master
  read -r confirm
  [[ "$confirm" == "y" ]] || [[ "$confirm" == "Y" ]] || exit 1
fi

<<<<<<< HEAD
echo "Upgrading gitea from $current to $giteaversion ..."

pushd "$(pwd)" &>/dev/null
cd "$giteahome" # needed for gitea dump later

# download new binary
binname="gitea-${giteaversion}-${arch}"
binurl="https://dl.gitea.com/gitea/${giteaversion}/${binname}.xz"
=======
echo "Upgrading proxgit from $current to $proxgitversion ..."

pushd "$(pwd)" &>/dev/null
cd "$proxgithome" # needed for proxgit dump later

# download new binary
binname="proxgit-${proxgitversion}-${arch}"
binurl="https://dl.proxgit.com/proxgit/${proxgitversion}/${binname}.xz"
>>>>>>> master
echo "Downloading $binurl..."
curl --connect-timeout 10 --silent --show-error --fail --location -O "$binurl{,.sha256,.asc}"

# validate checksum & gpg signature
sha256sum -c "${binname}.xz.sha256"
if [[ -z "${ignore_gpg:-}" ]]; then
  require gpg
  gpg --keyserver keys.openpgp.org --recv 7C9E68152594688862D62AF62D9AE806EC1592E2
  gpg --verify "${binname}.xz.asc" "${binname}.xz" || { echo 'Signature does not match'; exit 1; }
fi
rm "${binname}".xz.{sha256,asc}

# unpack binary + make executable
xz --decompress --force "${binname}.xz"
<<<<<<< HEAD
chown "$giteauser" "$binname"
chmod +x "$binname"

# stop gitea, create backup, replace binary, restart gitea
echo "Flushing gitea queues at $(date)"
giteacmd manager flush-queues
echo "Stopping gitea at $(date)"
$service_stop
echo "Creating backup in $giteahome"
giteacmd dump $backupopts
echo "Updating binary at $giteabin"
cp -f "$giteabin" "$giteabin.bak" && mv -f "$binname" "$giteabin"
$service_start
$service_status

echo "Upgrade to $giteaversion successful!"
=======
chown "$proxgituser" "$binname"
chmod +x "$binname"

# stop proxgit, create backup, replace binary, restart proxgit
echo "Flushing proxgit queues at $(date)"
proxgitcmd manager flush-queues
echo "Stopping proxgit at $(date)"
$service_stop
echo "Creating backup in $proxgithome"
proxgitcmd dump $backupopts
echo "Updating binary at $proxgitbin"
cp -f "$proxgitbin" "$proxgitbin.bak" && mv -f "$binname" "$proxgitbin"
$service_start
$service_status

echo "Upgrade to $proxgitversion successful!"
>>>>>>> master

popd
