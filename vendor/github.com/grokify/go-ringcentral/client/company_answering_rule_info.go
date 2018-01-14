/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type CompanyAnsweringRuleInfo struct {

	// Internal identifier of an answering rule
	Id string `json:"id,omitempty"`

	// Canonical URI of an answering rule
	Uri string `json:"uri,omitempty"`

	// Specifies if the rule is active or inactive. The default value is 'True'
	Enabled bool `json:"enabled,omitempty"`

	// Type of an answering rule, the default value is 'Custom' = ['BusinessHours', 'AfterHours', 'Custom']
	Type_ string `json:"type,omitempty"`

	// Name of an answering rule specified by user. Max number of symbols is 30. The default value is 'My Rule N' where 'N' is the first free number
	Name string `json:"name,omitempty"`

	// Answering rule will be applied when calls are received from the specified caller(s)
	Callers []CompanyAnsweringRuleCallersInfoRequest `json:"callers,omitempty"`

	// Answering rule will be applied when calling the specified number(s)
	CalledNumbers []CompanyAnsweringRuleCalledNumberInfoRequest `json:"calledNumbers,omitempty"`

	// Schedule when an answering rule should be applied ,
	Schedule *CompanyAnsweringRuleScheduleInfo `json:"schedule,omitempty"`

	// Specifies how incoming calls are forwarded. The default value is 'Operator' 'Operator' - play company greeting and forward to operator extension 'Disconnect' - play company greeting and disconnect 'Bypass' - bypass greeting to go to selected extension = ['Operator', 'Disconnect', 'Bypass']
	CallHandlingAction string `json:"callHandlingAction,omitempty"`

	// Extension to which the call is forwarded in 'Bypass' mode
	Extension *CompanyAnsweringRuleCallersInfoRequest `json:"extension,omitempty"`

	// Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List
	Greetings []GreetingInfo `json:"greetings,omitempty"`
}
