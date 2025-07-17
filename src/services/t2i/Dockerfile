FROM python:3.11-slim AS base

# Building stage
FROM base AS builder

ENV PYTHONFAULTHANDLER=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONHASHSEED=random \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PIP_DEFAULT_TIMEOUT=100 \
  PATH="$PATH:/runtime/bin" \
  PYTHONPATH="$PYTHONPATH:/runtime/lib/python3.11/site-packages"

RUN apt update && apt install -y build-essential unzip

LABEL com.nvidia.volumes.needed="nvidia_driver"
ENV CUDA_LAUNCH_BLOCKING 1
ENV NVIDIA_VISIBLE_DEVICES all
ENV NVIDIA_DRIVER_CAPABILITIES compute,utility

WORKDIR /src
COPY requirements.txt /src/
RUN pip install --prefix=/runtime --force-reinstall -r requirements.txt

# Runtime stage
FROM base AS runtime
COPY --from=builder /runtime /usr/local
RUN apt update && apt install -y git-lfs wget libgl1 libglib2.0-0 unzip

WORKDIR /app

RUN mkdir ./ip_adapter && \
  wget https://raw.githubusercontent.com/instantX-research/InstantID/e6b1f21b6edbcc26e6ffa151aba603982c15c92c/pipeline_stable_diffusion_xl_instantid.py -O ./pipeline_stable_diffusion_xl_instantid.py && \
  wget https://raw.githubusercontent.com/instantX-research/InstantID/e6b1f21b6edbcc26e6ffa151aba603982c15c92c/ip_adapter/attention_processor.py -O ./ip_adapter/attention_processor.py && \
  wget https://raw.githubusercontent.com/instantX-research/InstantID/e6b1f21b6edbcc26e6ffa151aba603982c15c92c/ip_adapter/resampler.py -O ./ip_adapter/resampler.py && \
  wget https://raw.githubusercontent.com/instantX-research/InstantID/e6b1f21b6edbcc26e6ffa151aba603982c15c92c/ip_adapter/utils.py -O ./ip_adapter/utils.py

COPY ./main.py ./main.py
COPY ./schemas.py ./schemas.py
COPY ./generator.py ./generator.py
COPY ./run.sh ./
RUN chmod +x run.sh

# CMD [ "sh" ]
CMD ["sh", "./run.sh"]
