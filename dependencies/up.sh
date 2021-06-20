#!/bin/bash

docker pull neo4j

docker run --detach --name=marcord-neo4j --rm --env=NEO4J_AUTH=none \
--volume=$HOME/neo4j/data:/data \
--volume=$HOME/neo4j/import:/import \
--volume=$HOME/neo4j/conf:/conf \
neo4j