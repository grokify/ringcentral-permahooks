/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

type MessageStoreCallerInfoResponse struct {

	// Extension short number (usually 3 or 4 digits). This property is filled when parties communicate by means of short internal numbers, for example when calling to other extension or sending/receiving Company Pager message
	ExtensionNumber string `json:"extensionNumber"`

	// Contains party location (city, state) if one can be determined from phoneNumber. This property is filled only when phoneNumber is not empty and server can calculate location information from it (for example, this information is unavailable for US toll-free numbers)
	Location string `json:"location,omitempty"`

	// Status of a message. Returned for outbound fax messages only
	MessageStatus string `json:"messageStatus,omitempty"`

	// Fax only. Error code returned in case of fax sending failure. Returned if messageStatus value is 'SendingFailed'
	FaxErrorCode string `json:"faxErrorCode,omitempty"`

	// Symbolic name associated with a party. If the phone does not belong to the known extension, only the location is returned, the name is not determined then
	Name string `json:"name,omitempty"`

	// Phone number of a party. Usually it is a plain number including country and area code like 18661234567. But sometimes it could be returned from database with some formatting applied, for example (866)123-4567. This property is filled in all cases where parties communicate by means of global phone numbers, for example when calling to direct numbers or sending/receiving SMS
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
