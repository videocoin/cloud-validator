package service

import (
	"errors"
	"fmt"
	"image/png"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/google/uuid"
	ffprobe "github.com/vansante/go-ffprobe"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getFrames(filepath string) (int, error) {
	data, err := ffprobe.GetProbeData(filepath, 5000*time.Millisecond)
	if err != nil {
		return 0, err
	}

	stream := data.GetFirstVideoStream()
	if stream == nil {
		return 0, errors.New("no video stream")
	}

	nbFrames := 0

	frs := stream.RFrameRate
	if frs == "" {
		frs = stream.AvgFrameRate
	}

	frsParts := strings.Split(frs, "/")
	if len(frsParts) != 2 {
		return 0, errors.New("unable to calc framerate")
	}

	fr1, err := strconv.Atoi(frsParts[0])
	if err != nil {
		return 0, errors.New("unable to calc framerate")
	}

	fr2, err := strconv.Atoi(frsParts[1])
	if err != nil {
		return 0, errors.New("unable to calc framerate")
	}

	fr := fr1 / fr2

	duration := data.Format.DurationSeconds
	if duration == 0 {
		duration, err = strconv.ParseFloat(stream.Duration, 64)
		if err != nil {
			return 0, errors.New("unable to get duration")
		}
	}

	if duration == 0 {
		return 0, errors.New("unable to get duration")
	}

	nbFrames = int(duration * float64(fr))

	return nbFrames, nil
}

func extractFrame(filepath string, frame int) (string, error) {
	// http://trac.ffmpeg.org/wiki/Seeking
	// http://ffmpeg.org/ffmpeg-utils.html#time-duration-syntax
	output := fmt.Sprintf("%s.png", uuid.New().String())
	args := []string{"-i", filepath, "-vf", fmt.Sprintf(`scale=360:-2,select=eq(n\,%d)`, frame), "-frames:v", "1", output}
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
