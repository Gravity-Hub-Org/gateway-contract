#!/bin/bash

docker-compose run truffle sh -c "cd /src; npm install; cd /src/ethereum; truffle deploy --network=ganache"