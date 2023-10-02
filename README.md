# Usage
This tool can be used as a package or cli tool.
It serves as a data extractor to support ELT pipelines or any kind of process that requires a heavy data dump.

## CLI

This package allows the user to dump data into multiple json files.

### Ping database
The `ping` command does a ping in database and returns a connection check.


```bash
mongoextract ping --conn-uri "$MONGO_CONN_URI"
```

### Check if a collection exists
The `collxst` command does a ping in database and returns a connection check.


```bash
mongoextract collxst \
	--conn-uri "$MONGO_CONN_URI" \
	--db-name "$MONGO_DBNAME" \
	--collection "$MONGO_COLLECTION" \
	--app-name "$APPNAME"
```


### Extract in batches - dumping streaming (async)
The `extract-batch` command iterates over mongo cursor and dumps chunks of data into json files.


```bash
mongoextract extract-batch \
		--conn-uri "$MONGO_CONN_URI" \
		--db-name "$MONGO_DBNAME" \
		--collection "$MONGO_COLLECTION" \
		--app-name "$APPNAME" \
		--mapping $ID_NAME \
		--query '{"latitude":{"$$gte":30}}' \
		--output-path "./data" \
		--output-prefix $ID_NAME \
		--chunk-size 100 \
		--num-concurrent-files 10
```

### Extract data
The `extract` command fetches a mongo cursor and dumps the whole data into a json file.


```bash
mongoextract extract \
	--conn-uri "$MONGO_CONN_URI" \
	--db-name "$MONGO_DBNAME" \
	--collection "$MONGO_COLLECTION" \
	--app-name "$APPNAME" \
	--mapping some_mapping_name \
	--query '{"latitude":{"$$gte":30}}'
```
