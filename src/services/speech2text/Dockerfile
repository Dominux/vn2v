FROM python:3.11-slim AS base

# Building stage
FROM base AS builder

ENV PYTHONFAULTHANDLER=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONHASHSEED=random \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PATH="$PATH:/runtime/bin" \
  PYTHONPATH="$PYTHONPATH:/runtime/lib/python3.11/site-packages"

RUN apt update && apt install -y build-essential unzip wget

LABEL com.nvidia.volumes.needed="nvidia_driver"
ENV CUDA_LAUNCH_BLOCKING=1
ENV NVIDIA_VISIBLE_DEVICES=all
ENV NVIDIA_DRIVER_CAPABILITIES=compute,utility

WORKDIR /src
COPY requirements.txt /src/
RUN pip install --prefix=/runtime --force-reinstall -r requirements.txt
# ЕЬАЛ ВРОТ ПИДОРОВ, ЧТО ПРИВЯЗАЛИ СТАРУЮ ВЕРСИЮ CUDNN К БЕСПОЛЕЗНОЙ ПАРАШЕ, НАЗЫВАЕМОЙ ПАКЕТОM, ИЗ-ЗА ЧЕГО Я ПОТРАТИЛ МИЛЛИАРД ЧАСОВ ДЛЯ ВЫЯСНЕНИЯ СУКА
RUN pip install --prefix=/runtime --force-reinstall ctranslate2==4.5.0

# Runtime stage
FROM base AS runtime
COPY --from=builder /runtime /usr/local
RUN apt update && apt install -y curl git-lfs ffmpeg

ENV LD_LIBRARY_PATH=/usr/local/lib/python3.11/site-packages/nvidia/cudnn/lib

WORKDIR /

CMD ["sh"]
