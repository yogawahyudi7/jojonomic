package config

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerText() logger.Config {

	t := time.Now()
	formatDate := t.Format("20060102")
	logJoin := []string{".", "/", "log", "/", formatDate, ".log"}
	logFile := strings.Join(logJoin, "")
	_, err := os.Stat(logFile)

	//check exist file log
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_APPEND, 0755)
	if os.IsNotExist(err) {
		file, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	}

	logger := (logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ROUTE:[${path}]\nparam : [${queryParams}]\nrequest : [${body}]\nresponse : [${resBody}]\n\n",
		Output: file,
	})

	return logger
}

func LoggerTermial() logger.Config {

	logger := (logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ROUTE:[${path}]\nparam : [${queryParams}]\nrequest : [${body}]\nresponse : [${resBody}]\n\n",
	})

	return logger
}
