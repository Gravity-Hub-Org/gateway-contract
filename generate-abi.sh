#!/bin/bash

docker-compose run abigen sh -c "cd /source/0.7; abigen --sol source/LUPort/LUPort.sol --pkg luport --solc /usr/local/bin/solc7 > api/luport/luport.go"
docker-compose run abigen sh -c "cd /source/0.7; abigen --sol source/IBPort/IBPort.sol --pkg ibport --solc /usr/local/bin/solc7 > api/ibport/ibport.go"
docker-compose run abigen sh -c "cd /source/0.7; abigen --sol source/Nebula/Nebula.sol --pkg nebula --solc /usr/local/bin/solc7 > api/nebula/nebula.go"

docker-compose run abigen sh -c "cd /source/0.5; abigen --sol source/LUPort/LUPort.sol --pkg luport --solc /usr/local/bin/solc5 > api/luport/luport.go"
docker-compose run abigen sh -c "cd /source/0.5; abigen --sol source/IBPort/IBPort.sol --pkg ibport --solc /usr/local/bin/solc5 > api/ibport/ibport.go"
docker-compose run abigen sh -c "cd /source/0.5; abigen --sol source/Nebula/Nebula.sol --pkg nebula --solc /usr/local/bin/solc5 > api/nebula/nebula.go"

