#!/bin/bash

files=$(find . -name *.go| grep -v _test  | cut -d'/' -f2-| sort)

declare -i e=0
declare -i m=0
declare -i h=0

for f in $files; do
    if [[ "$f" =~ ^.*/E[0-9.]+_.*$ ]]; then
        e+=1
    fi
    
    if [[ "$f" =~ ^.*/M[0-9.]+_.*$ ]]; then
        m+=1
    fi

    if [[ "$f" =~ ^.*/H[0-9.]+_.*$ ]]; then
        h+=1
    fi
done

cat << EOF > README.md
#

| 难度 | 完成 |
| ---- | ----|
| 简单 | $e |
| 中等 | $m |
| 困难 | $h |
EOF