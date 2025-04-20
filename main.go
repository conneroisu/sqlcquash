/*
Copyright © 2024 Conner Ohnesorge <conneroisu@outlook.com>
*/
package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/conneroisu/sqlcquash/cmd"
)

func main() {
	slog.SetDefault(DefaultLogger)
	cmd.Execute()
}

// DefaultLogger is a default logger.
var DefaultLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	AddSource: true,
	Level:     slog.LevelError,
	ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
		if a.Key == "time" {
			return slog.Attr{}
		}
		if a.Key == "level" {
			return slog.Attr{}
		}
		if a.Key == slog.SourceKey {
			str := a.Value.String()
			split := strings.Split(str, "/")
			if len(split) > 2 {
				a.Value = slog.StringValue(strings.Join(split[len(split)-2:], "/"))
				a.Value = slog.StringValue(strings.Replace(a.Value.String(), "}", "", -1))
			}
		}
		return a
	}}))
