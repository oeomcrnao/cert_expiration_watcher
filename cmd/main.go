package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/oeomcrnao/domain_expiration_watcher/pkg/alerts/telegram"
	"github.com/oeomcrnao/domain_expiration_watcher/pkg/tools/environment"
)

func main() {

	var targetDomain, _ = environment.GetStringFromEnv("TARGET_DOMAIN", "")
	if len(targetDomain) == 0 {
		panic("You must specify the TARGET_DOMAIN variable to verify!")
	}

	var targetPort, _ = environment.GetStringFromEnv("TARGET_PORT", "443")
	if len(targetPort) == 0 {
		panic("You must specify the TARGET_PORT variable to verify!")
	}

	var target = fmt.Sprintf("%s:%s", targetDomain, targetPort)

	conn, err := tls.Dial("tcp", target, nil)
	if err != nil {
		log.Fatal("Server doesn't support SSL certificate err: " + err.Error())
	}

	err = conn.VerifyHostname(targetDomain)
	if err != nil {
		log.Fatal("Hostname doesn't match with certificate: " + err.Error())
	}

	var tresholdInDays, _ = environment.GetIntFromEnv("TRESHOLD_IN_DAYS")

	today := time.Now()
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	difference := int(expiry.Sub(today).Hours() / 24)

	if difference > tresholdInDays {
		return
	}

	var message = fmt.Sprintf("Domain: %s\nExpiry: %v\n", targetDomain, expiry.Format(time.RFC850))
	fmt.Println(message)

	var telegramEnabled, _ = environment.GetBoolFromEnv("TELEGRAM_ENABLED")
	if telegramEnabled {
		telegram.SendMessage(message)
	}
}
