REST api for adding, storing and retrieving metrics around Developer Experience

# Endpoints
## /events/commit (POST)
## /events/deployment (POST)

## /events/commit?from={date}&to={date}&team={team} (GET)
- Retrieve raw events, e.g. all commits from team y over the past month

## /metrics/{metric_name}?team=y (GET)
- Retrieve an aggregated metric, e.g. average weekly commits from team y over the past year

TODO:
- read request filter from query parameters
- add deployment endpoint
- implement rest endpoints
- decide on database (nosql vs timeseries vs columnar)
