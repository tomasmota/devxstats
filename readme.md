# Devxstats
⚠️ Warning: Active Development ⚠️
 
Service that fetches and stores developer experience related metrics from different sources and makes them available through a rest api.

Note: These metrics should be used in combination with satisfaction, well-being, collaboration, and other measurements that are not possible to automate. See [Space Metrics](https://queue.acm.org/detail.cfm?id=3454124).

## Exposed metrics
- Pull Request review velocity (time from opened to closed)
- Time before first pull request review
- Pull Requests completed in a given period
- Pull Request size, in lines of code
- Code coverage
- Deployment frequency
- Units of work finished (e.g. Jira tickets closed)

## Supported data sources

### Git

#### Github
Configuration

| Environment Variable | Description                                                                          | Required |
| -------------------- | ------------------------------------------------------------------------------------ | -------- |
| GITHUB_URL           | URL pointing to the api (default https://api.github.com)                             | No       |
| GITHUB_TOKEN         | Token used for authentication (can be created at https://github.com/settings/tokens) | Yes      |

#### Bitbucket
Configuration

| Environment Variable | Description                                                 | Required |
| -------------------- | ----------------------------------------------------------- | -------- |
| BITBUCKET_URL        | URL pointing to the api (default https://api.bitbucket.org) | No       |
| BITBUCKET_TOKEN      | Token used for authentication                               | Yes      |

### CD

#### Octopus
Configuration

| Environment Variable | Description                   | Required |
| -------------------- | ----------------------------- | -------- |
| OCTOPUS_URL          | URL pointing to the api       | Yes      |
| OCTOPUS_TOKEN        | Token used for authentication | Yes      |

## TODO:
- document public methods
- use tomasmota/go-bitbucket client
- consider just changing all sources to implement Sync, and simply invoke a Sync method on them that does everything. This has the advantage that different sources might have different strategies on how to best sync all resources, so forcing a standard procedure can be harmful (e.g. octoups)
- different entities can be synced in different frequencies, such as repos vs commits
- abstract pagination (probably using generics)
