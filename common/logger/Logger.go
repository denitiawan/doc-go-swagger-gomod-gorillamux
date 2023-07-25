package logger

import (
	"github.com/rs/zerolog/log"
)

func DoLogError(class string, function string, message string, err error) {
	if err != nil {
		message := class + "/" + function + "/" + message + "/" + err.Error()
		log.Error().Msg(message)
	}
}
