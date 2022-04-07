package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type formatter struct {
	serviceName string
	logrus.TextFormatter
}

func (f *formatter) getColors(entry *logrus.Entry) (int, int) {
	var (
		blueColor  = 34
		levelColor int
	)

	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 32 // green
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // cyan
	}

	return blueColor, levelColor
}

// Format 로그 format화
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	blueColor, levelColor := f.getColors(entry)

	switch {
	case f.serviceName == "":
		if len(entry.Data) > 0 {
			return []byte(
				fmt.Sprintf(
					"\x1b[%dm[%s]\x1b[0m \x1b[%dm%s\x1b[0m - %s (%s) \n",
					blueColor,
					entry.Time.Format(f.TimestampFormat),
					levelColor,
					strings.ToUpper(entry.Level.String()),
					entry.Message,
					entry.Data,
				),
			), nil
		}
		return []byte(
			fmt.Sprintf(
				"\x1b[%dm[%s]\x1b[0m \x1b[%dm%s\x1b[0m - %s\n",
				blueColor,
				entry.Time.Format(f.TimestampFormat),
				levelColor,
				strings.ToUpper(entry.Level.String()),
				entry.Message,
			),
		), nil
	case len(entry.Data) > 0:
		return []byte(
			fmt.Sprintf(
				"\x1b[%dm[%s] %s\x1b[0m \x1b[%dm%s\x1b[0m - %s (%s) \n",
				blueColor,
				entry.Time.Format(f.TimestampFormat),
				strings.ToUpper(f.serviceName),
				levelColor,
				strings.ToUpper(entry.Level.String()),
				entry.Message,
				entry.Data,
			),
		), nil
	default:
		return []byte(
			fmt.Sprintf(
				"\x1b[%dm[%s] %s\x1b[0m \x1b[%dm%s\x1b[0m - %s\n",
				blueColor,
				entry.Time.Format(f.TimestampFormat),
				strings.ToUpper(f.serviceName),
				levelColor,
				strings.ToUpper(entry.Level.String()),
				entry.Message,
			),
		), nil
	}
}
