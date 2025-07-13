# Downloading model files
if [ -d "/models/llama-3.2.gguf" ]; then
  echo "Model files already exist"
else
  echo "Downloading model"

  git lfs install

  git clone https://huggingface.co/bartowski/Llama-3.2-3B-Instruct-GGUF /models/llama-3.2.gguf

  echo "Downloaded model"
fi

/bin/ollama serve &
# maybe move this to the main container
sleep 5 && ollama create "t2i" -f Modelfile
wait
