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

RUN apt update && apt install -y build-essential unzip wget

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
COPY --from=builder /usr/bin/wget /usr/bin/wget
RUN apt update && apt install -y wget

WORKDIR /app
COPY ./main.py ./main.py
COPY ./schemas.py ./schemas.py
COPY ./generator.py ./generator.py
COPY ./run.sh ./
RUN chmod +x run.sh

# CMD [ "sh" ]
CMD ["sh", "./run.sh"]
