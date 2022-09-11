Service that fetches and stores developer experience related metrics from different sources and makes them available through an api

# TODO:
- define a separate model for each client communicating to a source
- decide on database
- add application configuration file
- define the code for a specific source
- restructure storage package to mirror sources structure
- use drone/go-scm to interact with git sources
- make envs required
- document public methods
- handle pagination

## Bitbucket todos
- go-scm does not support listing projects, add that capability

# Supported data sources

## Git

### Github
Configuration

| Environment Variable | Description                                                                          | Required |
| -------------------- | ------------------------------------------------------------------------------------ | -------- |
| GITHUB_URL           | URL pointing to the api (default https://api.github.com)                             | No       |
| GITHUB_TOKEN         | Token used for authentication (can be created at https://github.com/settings/tokens) | Yes      |

### Bitbucket
Configuration

| Environment Variable | Description                                                 | Required |
| -------------------- | ----------------------------------------------------------- | -------- |
| BITBUCKET_URL        | URL pointing to the api (default https://api.bitbucket.org) | No       |
| BITBUCKET_TOKEN      | Token used for authentication                               | Yes      |

## CD

### Octopus
Configuration

| Environment Variable | Description                   | Required |
| -------------------- | ----------------------------- | -------- |
| OCTOPUS_URL          | URL pointing to the api       | Yes      |
| OCTOPUS_TOKEN        | Token used for authentication | Yes      |