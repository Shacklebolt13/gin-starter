package errs

import "github.com/rs/zerolog/log"

//nolint:ireturn
func Fatal[T any](obj T, err error) T {
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	return obj
}
