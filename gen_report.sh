#!/bin/bash

# 查找所有 Go 代码文件，排除测试文件
files=$(find . -name "*.go" -not -name "*_test.go" | sort)

# 初始化统计变量
declare -A counts=( ["E"]=0 ["M"]=0 ["H"]=0 )

# 遍历文件并统计各个难度的数量
for f in $files; do
    if [[ "$f" =~ /E[0-9.]+_ ]]; then
        ((counts["E"]++))
    elif [[ "$f" =~ /M[0-9.]+_ ]]; then
        ((counts["M"]++))
    elif [[ "$f" =~ /H[0-9.]+_ ]]; then
        ((counts["H"]++))
    fi
done

# 生成 README.md 文件
cat << EOF > README.md
#

| 难度 | 完成 |
| ---- | ----|
| 简单 | ${counts["E"]} |
| 中等 | ${counts["M"]} |
| 困难 | ${counts["H"]} |
EOF