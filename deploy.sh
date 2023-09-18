#!/usr/bin/env sh

# 确保脚本抛出遇到的错误
set -e

git init
git config user.name "zagss"
git config user.email ""
git add -A
git commit -m 'deploy'

git push -f https://github.com/zagss/valet_driv.git main:main
