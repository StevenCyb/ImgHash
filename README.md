# ImgHash

This is a simple CLI tool that can compute hashes of `jpg` and `png` as well as compare images.
Currently the following algorithms are supported:

- Average Hashing (aHash)
- Block Hashing (bHash)
- ColorMoment Hashing (cmHash)
- Difference Hashing (dHash)
- Median Hashing (mHash)
- Perceptual Hashing (pHash)
- Wavelet Hashing (wHash)

## Hash image

Hashing an image is can be done using the following command.

```bash
imghash get phash image.png
$ 00000e5e7e7f3d00%
```

With the the `format` flag, binary output format can be selected.

```bash
imghash get --f binary phash image.png
$ 0000000000000000000011100101111001111110011111110011110100000000%
```

## Compare images
Images can be compared using the following command

```bash
imghash compare phash image_a.png image_b.png
$ 0.91%
```
