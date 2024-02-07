Tips on making your module deterministic:

- Make the output reproducible, for example for the diffusers library,
  see [here](https://huggingface.co/docs/diffusers/using-diffusers/reproducibility)
- Strip timestamps and time measurements out of the output, including to stdout/stderr
- Don't read any sources of entropy (e.g. /dev/random)
- When referencing docker images, you MUST specify their sha256 hashes, as shown in this example
