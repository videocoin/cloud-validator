package service

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/dwbuiten/go-mediainfo/mediainfo"
	"github.com/google/uuid"
)

func init() {
	mediainfo.Init()
	rand.Seed(time.Now().UnixNano())
}

func getDuration(filepath string) (float64, error) {
	info, err := mediainfo.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer info.Close()

	field, err := info.Get("Duration", 0, mediainfo.General)
	if err != nil {
		return 0, err
	}

	duration, err := strconv.ParseFloat(field, 64)

	return duration / 1000, err
}

func extractFrame(filepath string, atTime float64) (string, error) {
	// http://trac.ffmpeg.org/wiki/Seeking
	// http://ffmpeg.org/ffmpeg-utils.html#time-duration-syntax
	output := fmt.Sprintf("%s.png", uuid.New().String())
	args := []string{"-i", filepath, "-ss", fmt.Sprintf("%f", atTime), "-vf", "scale=360:-2", "-frames:v", "1", output}
	cmd := exec.Command("ffmpeg", args...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return output, nil
}

func getHash(framepath string) (*goimagehash.ImageHash, error) {
	img, err := os.Open(framepath)
	if err != nil {
		return nil, err
	}

	data, err := png.Decode(img)
	if err != nil {
		return nil, err
	}

	return goimagehash.PerceptionHash(data)
}
