/**
 * ���{���̐����Ńt���N�^���������_�����O����ɂ͍����Z�p���x�����߂��܂��B
 * complex64, complex128, big.Float, big.Rat �̎l�̈قȂ鐔�l�̕\�����g���ē����t���N�^�����������Ȃ����B
 * (�Ō�̓�̌^�� math/big �p�b�P�[�W�ɂ���܂��BFloat �͔C�Ӑ��x�ł����L�E���x�̕��������_���g���Ă��܂��B
 *  Rat �͔�L�E���x�̗L�������g���Ă��܂��B) ���\�ƃ������g�p�ʂɊւ��Ăǂ̂悤�Ȕ�r���ʂɂȂ�܂����B
 *  �ǂ̔{���̐����ɂȂ�ƃ����_�����O�̌��ʂ����o�I�ɕ�����悤�ɂȂ�܂����B
 */
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			px := ((real(z) - xmin) * width / (xmax - xmin))
			r := uint8(px / 4)
			g := uint8(math.MaxUint8 * cmplx.Abs(v) / math.Hypot(xmax, ymax))
			b := uint8(math.MaxUint8 * cmplx.Abs(v) / math.Hypot(xmax, ymax))
			return color.RGBA{r, g, b, 128}
		}
	}
	return color.Black
}
