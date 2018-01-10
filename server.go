package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/grokify/gotilla/encoding/jsonutil"
	"github.com/joho/godotenv"

	rc "github.com/grokify/go-ringcentral/client"
	rcu "github.com/grokify/go-ringcentral/clientutil"
	omr "github.com/grokify/oauth2more/ringcentral"
)

const (
	DefaultPort              = "8080"
	ValidationTokenHeader    = "Validation-Token"
	SMSEventFilter           = "/restapi/v1.0/account/~/extension/~/message-store/instant?type=SMS"
	RenewalEventFilterFormat = "/restapi/v1.0/subscription/%s?threshold=%d&interval=%d"
)

var (
	OutboundWebhookUrl           = "" // Simple inbound webhook like Zapier or Chathooks
	InboundWebhookUrl            = "" // Server URL the RingCentral API will send to
	CurrentWebhookSubscriptionId = "" // Current SubscriptionID to renew
	ExpiresIn                    = 60 * 60 * 24 * 7
	RenewalThresholdTime         = 60 * 60 * 24
	RenewalIntervalTime          = 60 * 60
)

// EventFilters determines the RingCentral events this service will subscribe to.
// Threshold is the threshold time (in seconds) remaining before subscription expiration when server should start to send renewal reminder notifications. This time is approximate. It cannot be less than the interval of reminder job execution. It also cannot be greater than a half of this subscription TTL.
// Interval is the interval (in seconds) between reminder notifications. This time is approximate. It cannot be less than the interval of reminder job execution. It also cannot be greater than a half of threshold value.
var RenewalEventFilter = GetRenewalEventFilter("~", RenewalThresholdTime, RenewalIntervalTime)
var EventFilters = []string{SMSEventFilter, RenewalEventFilter}

func GetRenewalEventFilter(subscriptionID string, threshold, interval int) string {
	return fmt.Sprintf(RenewalEventFilterFormat, subscriptionID, threshold, interval)
}

func SetEventFilters() {
	RenewalEventFilter = GetRenewalEventFilter("~", RenewalThresholdTime, RenewalIntervalTime)
	EventFilters = []string{SMSEventFilter, RenewalEventFilter}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("Handling webhook...")
	// Check to see if ValidationToken is present. If so, return
	// it immediately.
	validationToken := r.Header.Get(ValidationTokenHeader)
	if len(validationToken) > 0 {
		log.Printf("%s: %s", ValidationTokenHeader, validationToken)
		w.Header().Set(ValidationTokenHeader, validationToken)
		return
	}

	// Read the body to check if this is a renewal event
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	log.Debug(string(body))

	rcu := &rcu.Event{}
	err = json.Unmarshal(body, rcu)
	if err != nil {
		log.Warn("JSON Unmarshal Error: %s", err.Error())
		return
	}

	// If this is renewal event, renew the webhook and return.
	if rcu.Event == RenewalEventFilter {
		err := renewWebhook()
		if err != nil {
			log.Warn("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
		}
		return
	}

	// Forward the body to the Webhook URL
	resp, err := http.Post(
		OutboundWebhookUrl,
		"application/json",
		bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Downstream webhook error: %s", err.Error())
		return
	} else if resp.StatusCode >= 300 {
		log.Printf("Downstream webhook error: %v", resp.StatusCode)
		return
	}
}

func createhookHandler(w http.ResponseWriter, r *http.Request) {
	err := createWebhook()
	if err != nil {
		log.Printf(err.Error())
	}
}

func renewhookHandler(w http.ResponseWriter, r *http.Request) {
	err := renewWebhook()
	if err != nil {
		log.Printf(err.Error())
	}
}

func createWebhook() error {
	log.Info("Creating Hook...")
	apiClient, err := newRingCentralClient()
	if err != nil {
		return err
	}

	req := rc.CreateSubscriptionRequest{
		EventFilters: EventFilters,
		DeliveryMode: &rc.NotificationDeliveryModeRequest{
			TransportType: "WebHook",
			Address:       InboundWebhookUrl,
		},
		ExpiresIn: int32(ExpiresIn),
	}
	log.Info(jsonutil.MustMarshalString(req, true))

	return handleWebhookResponse(
		apiClient.PushNotificationsApi.CreateSubscription(
			context.Background(),
			req,
		),
	)
}

func renewWebhook() error {
	log.Debug("Renewing Hook Id %v ...", CurrentWebhookSubscriptionId)
	apiClient, err := newRingCentralClient()
	if err != nil {
		log.Printf("RENEW NEW RC CLIENT ERROR: %v", err.Error())
		return err
	}

	return handleWebhookResponse(
		apiClient.PushNotificationsApi.RenewSubscription(
			context.Background(),
			CurrentWebhookSubscriptionId,
		),
	)
}

func handleWebhookResponse(info rc.SubscriptionInfo, resp *http.Response, err error) error {
	if err != nil {
		return fmt.Errorf("%v: %v", "API Response Err", err.Error())
	} else if resp.StatusCode > 299 {
		return fmt.Errorf("RingCentral Subscription API request failure status code: %v", resp.StatusCode)
	}

	CurrentWebhookSubscriptionId = info.Id
	log.Info("Created/renewed Webhook with Id: %s", CurrentWebhookSubscriptionId)
	return nil
}

func newRingCentralClient() (*rc.APIClient, error) {
	return rcu.NewApiClient(
		omr.ApplicationCredentials{
			ServerURL:    os.Getenv("RINGCENTRAL_SERVER_URL"),
			ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
			ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		},
		omr.UserCredentials{
			Username:  os.Getenv("RINGCENTRAL_USERNAME"),
			Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
			Password:  os.Getenv("RINGCENTRAL_PASSWORD"),
		},
	)
}

func loadEnv() error {
	envPaths := []string{}
	if len(os.Getenv("ENV_PATH")) > 0 {
		envPaths = append(envPaths, os.Getenv("ENV_PATH"))
	}
	return godotenv.Load(envPaths...)
}

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	InboundWebhookUrl = os.Getenv("PERMAHOOKS_INBOUND_WEBHOOK_URL")
	OutboundWebhookUrl = os.Getenv("PERMAHOOKS_OUTBOUND_WEBHOOK_URL")

	shortRenewal := false // to verify if renewal is working.
	if shortRenewal {
		ExpiresIn = 180
		RenewalThresholdTime = 80
		RenewalIntervalTime = 30
	}
	SetEventFilters()

	http.Handle("/webhook", http.HandlerFunc(webhookHandler))
	http.Handle("/createhook", http.HandlerFunc(createhookHandler))
	http.Handle("/renewhook", http.HandlerFunc(renewhookHandler))

	port := os.Getenv("PORT")
	if len(strings.TrimSpace(port)) == 0 {
		port = DefaultPort
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
