# pagespeed

pagespeed captures results regarding your webpage score as reported by [Google's PageSpeed API](https://developers.google.com/speed/docs/insights/v2/reference/pagespeedapi/runpagespeed)

## Install
```
go get -u github.com/harshavardhana/pagespeed
```

## Run

- Create `urls.txt`

List each URLs for which you need to fetch the pagespeed insights. For example:-

```txt
https://bake.minio.io:9000/
https://bake.minio.io:9000/docker.html
https://bake.minio.io:9000/features.html
https://bake.minio.io:9000/downloads.html
```

`pagespeed` will capture insights for each URL for two different strategies i.e desktop and mobile.

- Run `pagespeed`

`pagespeed` should be run under the same working directory where `urls.txt` exists.

```sh
pagespeed
--- start ---
https://bake.minio.io:9000/
https://bake.minio.io:9000/docker.html
https://bake.minio.io:9000/features.html
https://bake.minio.io:9000/downloads.html
--- end ---
```

## Verify
`pagespeed` will dump each page score along with the strategy used into `result.json`. You can
query this file using [`jq`](https://stedolan.github.io/jq/)

```
jq .score result.json
"89"
"85"
"89"
"85"
"95"
"85"
"95"
"85"
```
