#!/bin/bash

i="hello"
trap "echo $i" SIGINT
sleep 3600
