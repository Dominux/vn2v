# Downloading model files
if [ -d "/models/llama-3.2.gguf" ]; then
  echo "Model files already exist"
else
  echo "Downloading model"

  git lfs install

  git clone https://huggingface.co/bartowski/Llama-3.2-3B-Instruct-GGUF /models/llama-3.2.gguf

  echo "Downloaded model"
fi

# creating models in background
sh -c "sleep 3 &&
  ollama create "t2i" -f ./modelfiles/sdxl-prompter.Modelfile &&
  ollama create "i2v" -f ./modelfiles/wan-video-prompter.Modelfile" &

# Running Ollama
/bin/ollama serve
