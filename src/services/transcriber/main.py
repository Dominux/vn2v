import gc
import os
import time

import whisperx
import torch


device = "cuda"
batch_size = 16 # reduce if low on GPU mem
compute_type = "float16" # change to "int8" if low on GPU mem (may reduce accuracy)
audio_path = '/audios/input_audio.wav'


def main():
    while True:
        if not os.path.exists(audio_path):
            print('waiting for input file')
            time.sleep(1)
            continue

        transcribe()

        # removing file
        os.remove(audio_path)


def transcribe():
    model = whisperx.load_model(
        "/models",
        device,
        # language="ru",
        compute_type=compute_type,
        # download_root="/models/",
    )

    audio = whisperx.load_audio(audio_path)
    result = model.transcribe(audio, batch_size=batch_size)

    entire_transcribe = " ".join(s.get('text') for s in result['segments'])

    print(f"\n[TEXT]\n{entire_transcribe}")

    # delete model if low on GPU resources
    gc.collect()
    torch.cuda.empty_cache()
    del model


if __name__ == '__main__':
    main()
