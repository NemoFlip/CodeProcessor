#!/bin/bash

case "$LANG" in
  "python3")
    echo "Запуск Python-кода..."
    python3 pyCode.py
    ;;
  "clang")
    echo "Запуск C-кода..."
    clang cCode.c -o cCode && ./cCode
    ;;
  "c++")
    echo "Запуск C++-кода..."
    g++ gcc.cpp -o gcc_code && ./gcc_code
    ;;
  *)
    echo "Неизвестный язык программирования."
    exit 1
    ;;
esac