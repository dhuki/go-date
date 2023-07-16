# GO-DATE

## Descriptions

This is a service for dating apps, you can find your partners within a seconds or maybe less. 

## Overview
- Using port & adapter / hexagonal architecture pattern
- Integration with consul config management
- Dependency Injection is a must for communication between object
- Redis for cache management
- Ensure quality with tests
- Rate limiter

## Getting Started

Run this service in local
1. You can use consul as a service configuration the documentation is in here [consul](https://www.consul.io/) or you can run it with local config. For local config, create a new config in local with the name `config.local.yaml` fill it with desire configuration. (notes: you can copy from `config.dev.yaml`)
2. Prepare migration script in folder `migrations` then execute the table in your local database with following order : 
a. table `user`
b. table `relation_user`
3. Run the service :
    ```sh
    make run
    ```

## Deployment
:warning: This deployment step is not running yet, but this is a explanation about the pipelines :
1. I create a new file .sh in folder scripts/ this file has functionality to trigger ansible playbook to running the template from ansible/ folder.
2. The template will running to execute some action. There are get secret key in vault, then templating the secret key to config.yaml after that the template will push config.yaml to consul.