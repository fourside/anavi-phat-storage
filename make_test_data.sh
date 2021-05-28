#!/bin/bash
echo -n '{ "hoge": 123 }' > ./json/hoge.`date -u "+%Y%m%d-%H%M%S"`.json
