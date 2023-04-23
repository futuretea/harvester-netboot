package log

import "github.com/rs/zerolog"

func Setup() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}
