#!/bin/sh

set -e

if [ ! -f ./build/test-env-check.sh ]; then
<<<<<<< HEAD
  echo "${0} can only be executed in gitea source root directory"
=======
  echo "${0} can only be executed in proxgit source root directory"
>>>>>>> master
  exit 1
fi


echo "check uid ..."

<<<<<<< HEAD
# the uid of gitea defined in "https://gitea.com/gitea/test-env" is 1000
gitea_uid=$(id -u gitea)
if [ "$gitea_uid" != "1000" ]; then
  echo "The uid of linux user 'gitea' is expected to be 1000, but it is $gitea_uid"
=======
# the uid of proxgit defined in "https://proxgit.com/proxgit/test-env" is 1000
proxgit_uid=$(id -u proxgit)
if [ "$proxgit_uid" != "1000" ]; then
  echo "The uid of linux user 'proxgit' is expected to be 1000, but it is $proxgit_uid"
>>>>>>> master
  exit 1
fi

cur_uid=$(id -u)
<<<<<<< HEAD
if [ "$cur_uid" != "0" -a "$cur_uid" != "$gitea_uid" ]; then
  echo "The uid of current linux user is expected to be 0 or $gitea_uid, but it is $cur_uid"
=======
if [ "$cur_uid" != "0" -a "$cur_uid" != "$proxgit_uid" ]; then
  echo "The uid of current linux user is expected to be 0 or $proxgit_uid, but it is $cur_uid"
>>>>>>> master
  exit 1
fi
