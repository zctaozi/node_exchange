#!/bin/bash

project_path=$(cd `dirname $0`; pwd)

echo "directory dispatch $project_path"

function _start_build() {
      echo "start build..."

      build_command="go build -o $project_name"
      echo "$build_command"
      $build_command

      echo "packaging success"
      echo "complied directory"
}

# 开始打包
project_name="${project_path##*/}"
_start_build

exit 0