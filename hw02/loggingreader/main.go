package main

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

type LoggingReader struct {
	io.Reader
	logger *slog.Logger
}

func (r LoggingReader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.logger.Info("read", slog.Int("bytes", n), slog.Any("error", err))
	return n, err
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	data := strings.NewReader("Learning is not wisdom; it is the material with which wisdom builds.")

	logReader := LoggingReader{
		Reader: data,
		logger: logger,
	}

	_, err := io.Copy(os.Stdout, logReader)

	if err != nil {
		logger.Error("copy failed", slog.Any("error", err))
	}
}
