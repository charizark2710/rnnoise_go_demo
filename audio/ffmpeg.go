package audio

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ConvertToRaw48kHz(inputName string) (string, error) {
	if (!strings.HasSuffix(inputName, ".wav") && !strings.HasSuffix(inputName, ".mp3")) || inputName == "" {
		return "", fmt.Errorf("invalid input file")
	} else if strings.HasSuffix(inputName, ".raw") {
		return "", nil
	}
	outputName := strings.Split(inputName, ".")[0] + "_48kHz.raw"
	cmd := exec.Command("ffmpeg", "-i", inputName, "-ar", "48000", "-ac", "1", "-f", "s16le", outputName)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error converting audio file: " + err.Error())
		return "", err
	}
	return outputName, nil
}

func ConvertToWav(inputName string) (string, error) {
	if !strings.HasSuffix(inputName, ".raw") || inputName == "" {
		return "", fmt.Errorf("invalid input file")
	}
	outputName := strings.Split(inputName, ".")[0] + ".wav"
	cmd := exec.Command("ffmpeg", "-f", "s16le", "-y", "-ac", "1", "-i", inputName, "-ar", "8000", "-ac", "1", outputName)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error converting audio file to WAV: " + err.Error())
		return "", err
	}
	return outputName, nil
}
