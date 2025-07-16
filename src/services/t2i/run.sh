#!/bin/sh
set -e

# Downloading model files
if [ -d "/models/files/" ]; then
  echo "Model files already exist"
else
  echo "Downloading model"

  # SDXL
  wget "https://civitai.com/api/download/models/1920523?type=Model&format=SafeTensor&size=pruned&fp=fp16" -O /models/files/sdxl.safetensors

  # ControlNet InstantID
  wget "https://huggingface.co/InstantX/InstantID/resolve/main/ip-adapter.bin?download=true" -O /models/files/ip-adapter.bin
  wget "https://huggingface.co/InstantX/InstantID/resolve/main/ControlNetModel/diffusion_pytorch_model.safetensors?download=true" -O /models/files/diffusion_pytorch_model.safetensors

  echo "Downloaded model"
fi

# Running app
gunicorn main:app \
  -w "${WORKERS:-1}" \
  -k uvicorn.workers.UvicornWorker \
  -b 0.0.0.0:5000 \
  --timeout 600
