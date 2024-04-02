package cmd

import (
	log "github.com/sirupsen/logrus"
)

func writeLog(sc int, target, method string, errorMsg string) {
	if sc != 200 {
		if sc == 401 || sc == 403 {
			logUnauth(sc, target, method, errorMsg)
		} else if sc == 301 || sc == 302 {
			logRedirect(sc, target, method)
		} else if sc == 0 {
			logSkipped(sc, target, method)
		} else if sc == 404 {
			logNotFound(sc, target, method, errorMsg)
		} else {
			logManual(sc, target, method, errorMsg)
		}
	} else {
		logAccessible(sc, target, method)
	}
}

func logAccessible(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Print("Endpoint accessible!")
	message := "Endpoint accessible!"
	writeToExcel(status, target, method, message)
}

func logManual(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Manual testing may be required."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Warn(errorMsg)
	message := "Manual testing may be required."
	writeToExcel(status, target, method, message)
}

func logNotFound(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Endpoint not found."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error(errorMsg)
	message := "Endpoint not found."
	writeToExcel(status, target, method, message)
}

func logRedirect(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error("Redirect detected. This likely requires authentication.")
	message := "Redirect detected. This likely requires authentication."
	writeToExcel(status, target, method, message)
}

func logSkipped(status int, target, method string) {
	log.WithFields(log.Fields{
		"Status": "N/A",
		"Target": target,
		"Method": method,
	}).Warn("Request skipped (dangerous keyword found).")
	message := "Request skipped (dangerous keyword found)."
	writeToExcel(status, target, method, message)
}

func logUnauth(status int, target, method, errorMsg string) {
	if errorMsg == "" {
		errorMsg = "Unauthorized."
	}
	log.WithFields(log.Fields{
		"Status": status,
		"Target": target,
		"Method": method,
	}).Error(errorMsg)
	message := "Unauthorized."
	writeToExcel(status, target, method, message)
}
