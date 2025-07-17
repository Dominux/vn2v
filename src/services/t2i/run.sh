#!/bin/sh
set -e

# Downloading model files
if [ -d "/models/files/" ]; then
  echo "Model files already exist"
else
  echo "Downloading model"

  mkdir /models/files

  # ControlNet InstantID
  wget "https://huggingface.co/InstantX/InstantID/resolve/main/ip-adapter.bin?download=true" -O /models/files/ip-adapter.bin

  mkdir /models/files/ControlNetModel
  wget "https://huggingface.co/InstantX/InstantID/resolve/main/ControlNetModel/diffusion_pytorch_model.safetensors?download=true" -O /models/files/ControlNetModel/diffusion_pytorch_model.safetensors
  wget "https://huggingface.co/InstantX/InstantID/resolve/main/ControlNetModel/config.json?download=true" -O /models/files/ControlNetModel/config.json

  # insightface
  wget "https://github.com/deepinsight/insightface/releases/download/v0.7/antelopev2.zip"
  unzip ./antelopev2.zip
  rm ./antelopev2.zip
  mkdir -p /models/files/insightface/models
  mv ./antelopev2 /models/files/insightface/models/

  git lfs install

  # Installing SDXL
  git clone https://huggingface.co/stablediffusionapi/epicrealism-xl /models/files/sdxl

  echo "Downloaded model"
fi

# Running app
gunicorn main:app \
  -w "${WORKERS:-1}" \
  -k uvicorn.workers.UvicornWorker \
  -b 0.0.0.0:5000 \
  --timeout 600

# python generator.py
