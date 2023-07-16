#!/bin/bash -e

environment=$1
if [[ $environment == 'PROD' ]];then
    environment=$2/config.prod.yaml
fi;

ansible-playbook ${ANSIBLE_TEMPLATE_HOME}/template_push_config.yaml --extra-vars "{
  \"config_file\": \"$environment\"
}"
