package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	filePathAd             = "./sensitive-word/ad.txt"
	filePathGunsExplosives = "./sensitive-word/guns_explosives.txt"
	filePathPolitic        = "./sensitive-word/politic.txt"
	filePathWebsite        = "./sensitive-word/website.txt"
)

type sensitiveWords struct {
	filePaths []string
	skipWords map[string]struct{}

	words []string
}

func newSensitiveWords(skipWords ...string) (sensitiveWords, error) {

	filePaths := []string{
		filePathAd, filePathGunsExplosives, filePathPolitic, filePathWebsite,
	}

	sw := make(map[string]struct{}, 0)
	for _, words := range skipWords {
		sw[words] = struct{}{}
	}

	s := sensitiveWords{
		filePaths: filePaths,
		skipWords: sw,
		words:     make([]string, 0),
	}

	if err := s.read(); err != nil {
		return sensitiveWords{}, err
	}

	return s, nil
}

func (s *sensitiveWords) read() error {

	for _, filePath := range s.filePaths {
		sw, err := s.readFileByLines(filePath) //#nosec G304
		if err != nil {
			return err
		}

		s.words = append(s.words, sw...)
	}

	return nil
}

func (s *sensitiveWords) readFileByLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sensitiveWords []string
	for scanner.Scan() {
		line := scanner.Text()

		if _, ok := s.skipWords[line]; ok {
			continue
		}

		sensitiveWords = append(sensitiveWords, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	return sensitiveWords, nil
}
