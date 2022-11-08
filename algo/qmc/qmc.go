package qmc

//dmFyICgKCUVyclFtY0ZpbGVMZW5ndGggICAgICA9IGVycm9ycy5OZXcoImludmFsaWQgcW1jIGZpbGUgbGVuZ3RoIikKCUVyclFtY0tleURlY29kZUZhaWxlZCA9IGVycm9ycy5OZXcoImJhc2U2NCBkZWNvZGUgcW1jIGtleSBmYWlsZWQiKQoJRXJyUW1jS2V5TGVuZ3RoICAgICAgID0gZXJyb3JzLk5ldygidW5leHBlY3RlZCBkZWNvZGVkIHFtYyBrZXkgbGVuZ3RoIikKKQoKdHlwZSBEZWNvZGVyIHN0cnVjdCB7CglmaWxlICAgICAgICAgW11ieXRlCgltYXNrRGV0ZWN0b3IgZnVuYyhlbmNvZGVkRGF0YSBbXWJ5dGUpICgqS2V5MjU2TWFzaywgZXJyb3IpCgltYXNrICAgICAgICAgKktleTI1Nk1hc2sKCWF1ZGlvRXh0ICAgICBzdHJpbmcKCWtleSAgICAgICAgICBbXWJ5dGUKCWF1ZGlvICAgICAgICBbXWJ5dGUKfQoKZnVuYyBOZXdNZmxhYzI1NkRlY29kZXIoZGF0YSBbXWJ5dGUpIGNvbW1vbi5EZWNvZGVyIHsKCXJldHVybiAmRGVjb2RlcntmaWxlOiBkYXRhLCBtYXNrRGV0ZWN0b3I6IGRldGVjdE1mbGFjMjU2TWFzaywgYXVkaW9FeHQ6ICJmbGFjIn0KfQoKZnVuYyBOZXdNZ2cyNTZEZWNvZGVyKGRhdGEgW11ieXRlKSBjb21tb24uRGVjb2RlciB7CglyZXR1cm4gJkRlY29kZXJ7ZmlsZTogZGF0YSwgbWFza0RldGVjdG9yOiBkZXRlY3RNZ2cyNTZNYXNrLCBhdWRpb0V4dDogIm9nZyJ9Cn0KCmZ1bmMgKGQgKkRlY29kZXIpIFZhbGlkYXRlKCkgZXJyb3IgewoJaWYgbmlsICE9IGQubWFzayB7CgkJcmV0dXJuIG5pbAoJfQoJaWYgbmlsICE9IGQubWFza0RldGVjdG9yIHsKCQlpZiBlcnIgOj0gZC52YWxpZGF0ZUtleSgpOyBlcnIgIT0gbmlsIHsKCQkJcmV0dXJuIGVycgoJCX0KCQl2YXIgZXJyIGVycm9yCgkJZC5tYXNrLCBlcnIgPSBkLm1hc2tEZXRlY3RvcihkLmZpbGUpCgkJcmV0dXJuIGVycgoJfQoJcmV0dXJuIGVycm9ycy5OZXcoIm5vIG1hc2sgb3IgbWFzayBkZXRlY3RvciBmb3VuZCIpCn0KCmZ1bmMgKGQgKkRlY29kZXIpIHZhbGlkYXRlS2V5KCkgZXJyb3IgewoJbGVuRGF0YSA6PSBsZW4oZC5maWxlKQoJaWYgbGVuRGF0YSA8IDQgewoJCXJldHVybiBFcnJRbWNGaWxlTGVuZ3RoCgl9CgoJa2V5TGVuIDo9IGJpbmFyeS5MaXR0bGVFbmRpYW4uVWludDMyKGQuZmlsZVtsZW5EYXRhLTQ6XSkKCWlmIGxlbkRhdGEgPCBpbnQoa2V5TGVuKzQpIHsKCQlyZXR1cm4gRXJyUW1jRmlsZUxlbmd0aAoJfQoJdmFyIGVyciBlcnJvcgoJZC5rZXksIGVyciA9IGJhc2U2NC5TdGRFbmNvZGluZy5EZWNvZGVTdHJpbmcoCgkJc3RyaW5nKGQuZmlsZVtsZW5EYXRhLTQtaW50KGtleUxlbikgOiBsZW5EYXRhLTRdKSkKCWlmIGVyciAhPSBuaWwgewoJCXJldHVybiBFcnJRbWNLZXlEZWNvZGVGYWlsZWQKCX0KCglpZiBsZW4oZC5rZXkpICE9IDI3MiB7CgkJcmV0dXJuIEVyclFtY0tleUxlbmd0aAoJfQoJZC5maWxlID0gZC5maWxlWzpsZW5EYXRhLTQtaW50KGtleUxlbildCglyZXR1cm4gbmlsCgp9CgpmdW5jIChkICpEZWNvZGVyKSBEZWNvZGUoKSBlcnJvciB7CglkLmF1ZGlvID0gZC5tYXNrLkRlY3J5cHQoZC5maWxlKQoJcmV0dXJuIG5pbAp9CgpmdW5jIChkIERlY29kZXIpIEdldENvdmVySW1hZ2UoKSBbXWJ5dGUgewoJcmV0dXJuIG5pbAp9CgpmdW5jIChkIERlY29kZXIpIEdldEF1ZGlvRGF0YSgpIFtdYnl0ZSB7CglyZXR1cm4gZC5hdWRpbwp9CgpmdW5jIChkIERlY29kZXIpIEdldEF1ZGlvRXh0KCkgc3RyaW5nIHsKCWlmIGQuYXVkaW9FeHQgIT0gIiIgewoJCXJldHVybiAiLiIgKyBkLmF1ZGlvRXh0Cgl9CglyZXR1cm4gIiIKfQoKZnVuYyAoZCBEZWNvZGVyKSBHZXRNZXRhKCkgY29tbW9uLk1ldGEgewoJcmV0dXJuIG5pbAp9CgpmdW5jIERlY29kZXJGdW5jV2l0aEV4dChleHQgc3RyaW5nKSBjb21tb24uTmV3RGVjb2RlckZ1bmMgewoJcmV0dXJuIGZ1bmMoZmlsZSBbXWJ5dGUpIGNvbW1vbi5EZWNvZGVyIHsKCQlyZXR1cm4gJkRlY29kZXJ7ZmlsZTogZmlsZSwgYXVkaW9FeHQ6IGV4dCwgbWFzazogZ2V0RGVmYXVsdE1hc2soKX0KCX0KfQoKLy9nb2xhbmQ6bm9pbnNwZWN0aW9uIFNwZWxsQ2hlY2tpbmdJbnNwZWN0aW9uCmZ1bmMgaW5pdCgpIHsKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoInFtYzAiLCBEZWNvZGVyRnVuY1dpdGhFeHQoIm1wMyIpKSAvL1FRIE11c2ljIE1wMwoJY29tbW9uLlJlZ2lzdGVyRGVjb2RlcigicW1jMyIsIERlY29kZXJGdW5jV2l0aEV4dCgibXAzIikpIC8vUVEgTXVzaWMgTXAzCgoJY29tbW9uLlJlZ2lzdGVyRGVjb2RlcigicW1jMiIsIERlY29kZXJGdW5jV2l0aEV4dCgibTRhIikpIC8vUVEgTXVzaWMgTTRBCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCJxbWM0IiwgRGVjb2RlckZ1bmNXaXRoRXh0KCJtNGEiKSkgLy9RUSBNdXNpYyBNNEEKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoInFtYzYiLCBEZWNvZGVyRnVuY1dpdGhFeHQoIm00YSIpKSAvL1FRIE11c2ljIE00QQoJY29tbW9uLlJlZ2lzdGVyRGVjb2RlcigicW1jOCIsIERlY29kZXJGdW5jV2l0aEV4dCgibTRhIikpIC8vUVEgTXVzaWMgTTRBCgoJY29tbW9uLlJlZ2lzdGVyRGVjb2RlcigicW1jZmxhYyIsIERlY29kZXJGdW5jV2l0aEV4dCgiZmxhYyIpKSAvL1FRIE11c2ljIEZsYWMKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoInFtY29nZyIsIERlY29kZXJGdW5jV2l0aEV4dCgib2dnIikpICAgLy9RUSBNdXNpYyBPZ2cKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoInRrbSIsIERlY29kZXJGdW5jV2l0aEV4dCgibTRhIikpICAgICAgLy9RUSBNdXNpYyBBY2NvbXBhbmltZW50IE00YQoKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoImJrY21wMyIsIERlY29kZXJGdW5jV2l0aEV4dCgibXAzIikpICAgLy9Nb28gTXVzaWMgTXAzCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCJia2NmbGFjIiwgRGVjb2RlckZ1bmNXaXRoRXh0KCJmbGFjIikpIC8vTW9vIE11c2ljIEZsYWMKCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCI2NjZjNjE2MyIsIERlY29kZXJGdW5jV2l0aEV4dCgiZmxhYyIpKSAvL1FRIE11c2ljIFdlaXl1biBGbGFjCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCI2ZDcwMzMiLCBEZWNvZGVyRnVuY1dpdGhFeHQoIm1wMyIpKSAgICAvL1FRIE11c2ljIFdlaXl1biBNcDMKCWNvbW1vbi5SZWdpc3RlckRlY29kZXIoIjZmNjc2NyIsIERlY29kZXJGdW5jV2l0aEV4dCgib2dnIikpICAgIC8vUVEgTXVzaWMgV2VpeXVuIE9nZwoJY29tbW9uLlJlZ2lzdGVyRGVjb2RlcigiNmQzNDYxIiwgRGVjb2RlckZ1bmNXaXRoRXh0KCJtNGEiKSkgICAgLy9RUSBNdXNpYyBXZWl5dW4gTTRhCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCI3NzYxNzYiLCBEZWNvZGVyRnVuY1dpdGhFeHQoIndhdiIpKSAgICAvL1FRIE11c2ljIFdlaXl1biBXYXYKCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCJtZ2ciLCBOZXdNZ2cyNTZEZWNvZGVyKSAgICAgLy9RUSBNdXNpYyBOZXcgT2dnCgljb21tb24uUmVnaXN0ZXJEZWNvZGVyKCJtZmxhYyIsIE5ld01mbGFjMjU2RGVjb2RlcikgLy9RUSBNdXNpYyBOZXcgRmxhYwp9
import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"github.com/unlock-music/cli/algo/common"
)

var (
	ErrQmcFileLength      = errors.New("invalid qmc file length")
	ErrQmcKeyDecodeFailed = errors.New("base64 decode qmc key failed")
	ErrQmcKeyLength       = errors.New("unexpected decoded qmc key length")
)

// type Decoder struct {
// 	file         []byte
// 	maskDetector func(encodedData []byte) (*Key256Mask, error)
// 	mask         *Key256Mask
// 	audioExt     string
// 	key          []byte
// 	audio        []byte
// }

// func NewMflac256Decoder(data []byte) common.Decoder {
// 	return &Decoder{file: data, maskDetector: detectMflac256Mask, audioExt: "flac"}
// }

// func NewMgg256Decoder(data []byte) common.Decoder {
// 	return &Decoder{file: data, maskDetector: detectMgg256Mask, audioExt: "ogg"}
// }

// func (d *Decoder) Validate() error {
// 	if nil != d.mask {
// 		return nil
// 	}
// 	if nil != d.maskDetector {
// 		if err := d.validateKey(); err != nil {
// 			return err
// 		}
// 		var err error
// 		d.mask, err = d.maskDetector(d.file)
// 		return err
// 	}
// 	return errors.New("no mask or mask detector found")
// }

// func (d *Decoder) validateKey() error {
// 	lenData := len(d.file)
// 	if lenData < 4 {
// 		return ErrQmcFileLength
// 	}

// 	keyLen := binary.LittleEndian.Uint32(d.file[lenData-4:])
// 	if lenData < int(keyLen+4) {
// 		return ErrQmcFileLength
// 	}
// 	var err error
// 	d.key, err = base64.StdEncoding.DecodeString(
// 		string(d.file[lenData-4-int(keyLen) : lenData-4]))
// 	if err != nil {
// 		return ErrQmcKeyDecodeFailed
// 	}

// 	if len(d.key) != 272 {
// 		return ErrQmcKeyLength
// 	}
// 	d.file = d.file[:lenData-4-int(keyLen)]
// 	return nil

// }

// func (d *Decoder) Decode() error {
// 	d.audio = d.mask.Decrypt(d.file)
// 	return nil
// }

// func (d Decoder) GetCoverImage() []byte {
// 	return nil
// }

// func (d Decoder) GetAudioData() []byte {
// 	return d.audio
// }

// func (d Decoder) GetAudioExt() string {
// 	if d.audioExt != "" {
// 		return "." + d.audioExt
// 	}
// 	return ""
// }

// func (d Decoder) GetMeta() common.Meta {
// 	return nil
// }

// func DecoderFuncWithExt(ext string) common.NewDecoderFunc {
// 	return func(file []byte) common.Decoder {
// 		return &Decoder{file: file, audioExt: ext, mask: getDefaultMask()}
// 	}
// }

//goland:noinspection SpellCheckingInspection
func init() {
	common.RegisterDecoder("qmc0", DecoderFuncWithExt("mp3")) //QQ Music Mp3
	common.RegisterDecoder("qmc3", DecoderFuncWithExt("mp3")) //QQ Music Mp3

	common.RegisterDecoder("qmc2", DecoderFuncWithExt("m4a")) //QQ Music M4A
	common.RegisterDecoder("qmc4", DecoderFuncWithExt("m4a")) //QQ Music M4A
	common.RegisterDecoder("qmc6", DecoderFuncWithExt("m4a")) //QQ Music M4A
	common.RegisterDecoder("qmc8", DecoderFuncWithExt("m4a")) //QQ Music M4A

	common.RegisterDecoder("qmcflac", DecoderFuncWithExt("flac")) //QQ Music Flac
	common.RegisterDecoder("qmcogg", DecoderFuncWithExt("ogg"))   //QQ Music Ogg
	common.RegisterDecoder("tkm", DecoderFuncWithExt("m4a"))      //QQ Music Accompaniment M4a

	common.RegisterDecoder("bkcmp3", DecoderFuncWithExt("mp3"))   //Moo Music Mp3
	common.RegisterDecoder("bkcflac", DecoderFuncWithExt("flac")) //Moo Music Flac

	common.RegisterDecoder("666c6163", DecoderFuncWithExt("flac")) //QQ Music Weiyun Flac
	common.RegisterDecoder("6d7033", DecoderFuncWithExt("mp3"))    //QQ Music Weiyun Mp3
	common.RegisterDecoder("6f6767", DecoderFuncWithExt("ogg"))    //QQ Music Weiyun Ogg
	common.RegisterDecoder("6d3461", DecoderFuncWithExt("m4a"))    //QQ Music Weiyun M4a
	common.RegisterDecoder("776176", DecoderFuncWithExt("wav"))    //QQ Music Weiyun Wav

	common.RegisterDecoder("mgg", NewMgg256Decoder)     //QQ Music New Ogg
	common.RegisterDecoder("mflac", NewMflac256Decoder) //QQ Music New Flac
}
