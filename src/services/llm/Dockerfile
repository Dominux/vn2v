FROM ollama/ollama:0.9.7-rc0-rocm

RUN apt update && apt install -y curl git-lfs

COPY ./run.sh /
COPY ./modelfiles /modelfiles

ENTRYPOINT [ "sh" ]
CMD [ "/run.sh" ]
