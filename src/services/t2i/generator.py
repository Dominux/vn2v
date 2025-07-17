from diffusers.utils import load_image
from diffusers.models import ControlNetModel
from diffusers.schedulers import DPMSolverMultistepScheduler

import cv2
import torch
import numpy as np
from insightface.app import FaceAnalysis

from pipeline_stable_diffusion_xl_instantid import StableDiffusionXLInstantIDPipeline, draw_kps


def generate_image(prompt):
    app = FaceAnalysis(
        name='antelopev2',
        root='/models/files/insightface/',
        providers=['CUDAExecutionProvider', 'CPUExecutionProvider'],
    )
    app.prepare(ctx_id=0, det_size=(512, 512))

    # load an image
    face_image = load_image("/data/screenshot.png")

    # prepare face emb
    face_info = app.get(cv2.cvtColor(np.array(face_image), cv2.COLOR_RGB2BGR))
    face_info = sorted(face_info, key=lambda x:(x['bbox'][2]-x['bbox'][0])*(x['bbox'][3]-x['bbox'][1]))[-1]  # only use the maximum face
    face_emb = face_info['embedding']
    face_kps = draw_kps(face_image, face_info['kps'])

    # prepare models under ./checkpoints
    face_adapter = '/models/files/ip-adapter.bin'
    controlnet_path = '/models/files/ControlNetModel'

    # load IdentityNet
    controlnet = ControlNetModel.from_pretrained(controlnet_path, torch_dtype=torch.float16)

    base_model = '/models/files/sdxl'
    pipe = StableDiffusionXLInstantIDPipeline.from_pretrained(
        base_model,
        controlnet=controlnet,
        torch_dtype=torch.float16
    )

    # scheduler
    scheduler = DPMSolverMultistepScheduler.from_config(pipe.scheduler.config)
    scheduler.use_karras_sigmas = True
    scheduler.algorithm_type = "sde-dpmsolver++"
    pipe.scheduler = scheduler

    pipe.cuda()

    # load adapter
    pipe.load_ip_adapter_instantid(face_adapter)

    # generate image
    image = pipe(
        prompt,
        image_embeds=face_emb,
        image=face_kps,
        controlnet_conditioning_scale=0.9,
        ip_adapter_scale=0.9,
        guidance_scale=3,
        num_inference_steps=60,
        width=1016,
        height=1016,
    ).images[0]

    image.save("/data/output_image.png")


if __name__ == '__main__':
    generate_image("a king")
