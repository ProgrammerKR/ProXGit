#!/bin/sh

set -e

if [ ! -f ./build/test-env-check.sh ]; then
  echo "${0} can only be executed in proxgit source root directory"
  exit 1
fi


echo "check uid ..."

# the uid of proxgit defined in "https://proxgit.com/proxgit/test-env" is 1000
proxgit_uid=$(id -u proxgit)
if [ "$proxgit_uid" != "1000" ]; then
  echo "The uid of linux user 'proxgit' is expected to be 1000, but it is $proxgit_uid"
  exit 1
fi

cur_uid=$(id -u)
if [ "$cur_uid" != "0" -a "$cur_uid" != "$proxgit_uid" ]; then
  echo "The uid of current linux user is expected to be 0 or $proxgit_uid, but it is $cur_uid"
  exit 1
fi
