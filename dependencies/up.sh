#!/bin/bash

docker pull neo4j

docker run --detach --name=marcord-neo4j --rm --env=NEO4J_AUTH=none \
--publish=7474:7474 --publish=7473:7473 --publish=7687:7687 \
--volume=$HOME/neo4j/data:/data \
--volume=$HOME/neo4j/import:/import \
--volume=$HOME/neo4j/conf:/conf \
neo4j