from diffusers import StableDiffusionXLControlNetPipeline, ControlNetModel, AutoencoderKL
from diffusers.utils import load_image
import torch


torch_dtype = torch.float16


def generate_image(prompt):
    input_image = load_image("/data/screenshot.png")

    # initialize the models and pipeline
    controlnet_conditioning_scale = 0.5  # recommended for good generalization
    controlnet_face = ControlNetModel.from_pretrained(
        "/models/files/ip-adapter.bin", torch_dtype=torch_dtype,
    )

    controlnet_reference = ControlNetModel.from_pretrained(
        "/models/files/diffusion_pytorch_model.safetensors", torch_dtype=torch_dtype,
    )

    vae = AutoencoderKL.from_pretrained("madebyollin/sdxl-vae-fp16-fix", torch_dtype=torch_dtype)
    pipe = StableDiffusionXLControlNetPipeline.from_pretrained(
        "/models/files/sdxl.safetensors",
        controlnet=[
            controlnet_face,
            controlnet_reference
        ],
        vae=vae,
        torch_dtype=torch_dtype,
        use_safetensors=True,
        safety_checker=None,
    )
    pipe.enable_model_cpu_offload()

    output_images = pipe(
        prompt,
        controlnet_conditioning_scale=controlnet_conditioning_scale,
        image=input_image,
    ).images

    output_images[0].save("/data/output_image.png", format="PNG")
