import gc

import whisperx
import torch


device = "cuda"
batch_size = 16 # reduce if low on GPU mem
compute_type = "float16" # change to "int8" if low on GPU mem (may reduce accuracy)
base_path = '/audios'



def transcribe(id: str):
    audio_path = f"{base_path}/input_audio.wav"

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
    transcribe('')
