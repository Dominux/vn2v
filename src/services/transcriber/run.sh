#!/bin/sh
set -e

git lfs install

git clone https://huggingface.co/guillaumekln/faster-whisper-small /models

python3 main.py
