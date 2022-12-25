#!/bin/bash

# Build script for Windows
windres -o main-res.syso res/main.rc && go build -ldflags "-H windowsgui" -o bin/$1.exe