# golang-snowflake-id-generator

A small side project for me to sharpen up my Golang seasonality again.

Based on the famous Twitter's Snowflake ID generator design, this repo contains a driver program and a logical machine instance to generate timestamp-sequence number.

## Snowflake ID content structure
- Sign flag - 1 bit, although it'll be always 0
- Timestamp - 41 bits of integer as milliseconds since the epoch. Default epoch is 1288834974657
- DC/Machine Identifier - 10 bits of integer
- Sequence # - 12 bits of integer number which gets reset to 0 per millisecond

## Few insights about Snowflake ID
- Sortable by UTC time
- 41-bit timestamp info => 2^41 - 1 = 2199023255551 ms => 69 years
- Sequence # => 2^12 = 4096 possible outcomes per millisecond