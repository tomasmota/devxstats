Service that fetches and stores developer experience related metrics from different sources and makes them available through an api

TODO:
- implement fetching architecture, should support types such as:
    - git
        - bitbucket
        - github
    - builds
        - ghactions
        - teamcity
        - tekton
    - deployments
        - k8s
        - octopus
    - documentation
    - others
- decide on database
