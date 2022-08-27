REST api for adding, storing and retrieving metrics around Developer Experience

# Endpoints
## /events (POST)
- Report an event, e.g. commit by user x, from team y, at 3:15pm

## /events?from={date}&to={date}&team={team} (GET)
- Retrieve raw events, e.g. all commits from team y over the past month

## /metrics/{metric_name}?team=y (GET)
- Retrieve an aggregated metric, e.g. average weekly commits from team y over the past year

TODO:
- implement rest endpoints
- decide on database (nosql vs timeseries vs columnar)
