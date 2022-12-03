package utlog

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// obj: string to be print
//
// show blue texted log
func Info(obj interface{}) {
	switch obj.(type) {
	case string:
		log.Info().Msg(obj.(string))
	case error:
		log.Info().Msg(obj.(error).Error())
	}
}

// format: string format that need to be print.
//
// show blue texted log
func Infof(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Info().Msg(msg)
}

// obj: string to be print
//
// show yellow texted log
func Warn(obj interface{}) {
	switch obj.(type) {
	case string:
		log.Warn().Msg(obj.(string))
	case error:
		log.Warn().Msg(obj.(error).Error())
	}
}

// format: string format that need to be print.
//
// show yellow texted log
func Warnf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Warn().Msg(msg)
}

// obj: string to be print
//
// show red texted log
func Error(obj interface{}) {
	switch obj.(type) {
	case string:
		log.Error().Msg(obj.(string))
	case error:
		log.Error().Msg(obj.(error).Error())
	}
}

// format: string format that need to be print.
//
// show red texted log
func Errorf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	log.Error().Msg(msg)
}
