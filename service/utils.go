package service

import (
	"errors"
	"fmt"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/google/uuid"
	ffprobe "github.com/vansante/go-ffprobe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getDuration(filepath string) (float64, error) {
	data, err := ffprobe.GetProbeData(filepath, 5000*time.Millisecond)
	if err != nil {
		return 0, err
	}

	return data.Format.DurationSeconds, nil
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

func checkSource(url string) error {
	if strings.HasPrefix(url, "file://") || strings.HasPrefix(url, "/") {
		fp := strings.TrimPrefix(url, "file://")
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			return err
		}
	} else if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		hc := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := hc.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp != nil && resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get %s, return status %s", url, resp.Status)
		}
	} else {
		return errors.New("unknown source type")
	}

	return nil
}
