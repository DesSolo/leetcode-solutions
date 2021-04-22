#!/usr/bin/bash
echo "### My leetcode solutions"
echo "#### table of contents"
echo "|leetcode problem|example solution|"
echo "|---|---|"
for file in $(find solutions/ -name *.go | sort)
do
    NAME=$(echo $file | cut -d "/" -f 2)
    URI=$(echo $NAME | tr "_" "-" | rev | cut -c 9- | rev )
    echo "|[${URI}](https://leetcode.com/problems/${URI}/)|[${NAME}]($file)|"
done
