#!/bin/bash

case "$LANG" in
  "python3")
    python3 pyCode.py
    ;;
  "clang")
    clang cCode.c -o cCode && ./cCode
    ;;
  "c++")
    g++ gcc.cpp -o gcc_code && ./gcc_code
    ;;
  *)
    echo "Неизвестный язык программирования."
    exit 1
    ;;
esac