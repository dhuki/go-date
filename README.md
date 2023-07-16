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
1. Create new config in local with name `config.local.yaml` fill it with desire configuration, you can copy from `config.dev.yaml`.
2. Prepare migration script in folder `migrations` then execute the table in your local database with following order : 
a. table `user`
b. table `relation_user`
3. Run the service :
    ```sh
    make run
    ```