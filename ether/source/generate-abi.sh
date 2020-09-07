#!/bin/bash

abigen --pkg nebula --sol Nebula/Nebula.sol > ../api/nebula/nebula.go
abigen --pkg ibport --sol IBPort/IBPort.sol > ../api/ibport/ibport.go
