package helpers

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v3.3-core/gl"
)

// taken from https://github.com/cstegel/opengl-samples-golang/blob/master/colors/gfx/texture.go
type Texture struct {
	handle  uint32
	target  uint32 // same target as gl.BindTexture(<this param>, ...)
	texUnit uint32 // Texture unit that is currently bound to ex: gl.TEXTURE0
}

type textureCreationOpts struct {
	wrapR, wrapS, minFilter, magFilter int32
	rotated                            bool
}

var errUnsupportedStride = errors.New("unsupported stride, only 32-bit colors supported")

var errTextureNotBound = errors.New("texture not bound")

type textureCreationOptsOverride func(*textureCreationOpts)

func WithWrappings(wrapR, wrapS int32) textureCreationOptsOverride {
	return func(tco *textureCreationOpts) {
		tco.wrapR = wrapR
		tco.wrapS = wrapS
	}
}

func WithFilters(min, mag int32) textureCreationOptsOverride {
	return func(tco *textureCreationOpts) {
		tco.minFilter = min
		tco.magFilter = mag
	}
}

func NewTextureForLesson(lesson int, fName string, overrides ...textureCreationOptsOverride) (*Texture, error) {
	lessonsFolder := fmt.Sprintf("lesson%d", lesson)
	file := path.Join("assets", lessonsFolder, fName)
	return NewTextureFromFile(file, overrides...)
}

func NewTextureFromFile(file string, overrides ...textureCreationOptsOverride) (*Texture, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()

	// Decode detexts the type of image as long as its image/<type> is imported
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return NewTexture(img, overrides...)
}

func NewTexture(img image.Image, overrides ...textureCreationOptsOverride) (*Texture, error) {
	opts := &textureCreationOpts{
		wrapR:     gl.REPEAT,
		wrapS:     gl.REPEAT,
		minFilter: gl.LINEAR,
		magFilter: gl.LINEAR,
		rotated:   true,
	}
	for _, override := range overrides {
		override(opts)
	}
	if opts.rotated {
		img = imaging.Rotate180(img)
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Pt(0, 0), draw.Src)
	if rgba.Stride != rgba.Rect.Size().X*4 { // TODO-cs: why?
		return nil, errUnsupportedStride
	}

	var handle uint32
	gl.GenTextures(1, &handle)

	target := uint32(gl.TEXTURE_2D)
	internalFmt := int32(gl.SRGB_ALPHA)
	format := uint32(gl.RGBA)
	width := int32(rgba.Rect.Size().X)
	height := int32(rgba.Rect.Size().Y)
	pixType := uint32(gl.UNSIGNED_BYTE)
	dataPtr := gl.Ptr(rgba.Pix)

	texture := Texture{
		handle: handle,
		target: target,
	}

	texture.Bind(gl.TEXTURE0)
	defer texture.UnBind()

	// set the texture wrapping/filtering options (applies to current bound texture obj)
	// TODO-cs
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_R, opts.wrapR)
	gl.TexParameteri(texture.target, gl.TEXTURE_WRAP_S, opts.wrapS)
	gl.TexParameteri(texture.target, gl.TEXTURE_MIN_FILTER, opts.minFilter) // minification filter
	gl.TexParameteri(texture.target, gl.TEXTURE_MAG_FILTER, opts.magFilter) // magnification filter

	gl.TexImage2D(target, 0, internalFmt, width, height, 0, format, pixType, dataPtr)

	gl.GenerateMipmap(texture.handle)

	return &texture, nil
}

func (tex *Texture) Bind(texUnit uint32) {
	gl.ActiveTexture(texUnit)
	gl.BindTexture(tex.target, tex.handle)
	tex.texUnit = texUnit
}

func (tex *Texture) UnBind() {
	tex.texUnit = 0
	gl.BindTexture(tex.target, 0)
}

func (tex *Texture) SetUniform(uniformLoc int32) error {
	if tex.texUnit == 0 {
		return errTextureNotBound
	}
	gl.Uniform1i(uniformLoc, int32(tex.texUnit-gl.TEXTURE0))
	return nil
}
