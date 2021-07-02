// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/v1alpha1/extauth.proto

package v1alpha1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *ExtAuthConfigSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAuthConfigRefName(), target.GetAuthConfigRefName()) != 0 {
		return false
	}

	if len(m.GetConfigs()) != len(target.GetConfigs()) {
		return false
	}
	for idx, v := range m.GetConfigs() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetConfigs()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetConfigs()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetBooleanExpr()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBooleanExpr()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBooleanExpr(), target.GetBooleanExpr()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *AuthPlugin) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AuthPlugin)
	if !ok {
		that2, ok := that.(AuthPlugin)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetPluginFileName(), target.GetPluginFileName()) != 0 {
		return false
	}

	if strings.Compare(m.GetExportedSymbolName(), target.GetExportedSymbolName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *BasicAuth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BasicAuth)
	if !ok {
		that2, ok := that.(BasicAuth)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetRealm(), target.GetRealm()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetApr()).(equality.Equalizer); ok {
		if !h.Equal(target.GetApr()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetApr(), target.GetApr()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RedisOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RedisOptions)
	if !ok {
		that2, ok := that.(RedisOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if m.GetDb() != target.GetDb() {
		return false
	}

	if m.GetPoolSize() != target.GetPoolSize() {
		return false
	}

	return true
}

// Equal function
func (m *UserSession) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UserSession)
	if !ok {
		that2, ok := that.(UserSession)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetFailOnFetchFailure() != target.GetFailOnFetchFailure() {
		return false
	}

	if h, ok := interface{}(m.GetCookieOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCookieOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCookieOptions(), target.GetCookieOptions()) {
			return false
		}
	}

	switch m.Session.(type) {

	case *UserSession_Cookie:
		if _, ok := target.Session.(*UserSession_Cookie); !ok {
			return false
		}

		if h, ok := interface{}(m.GetCookie()).(equality.Equalizer); ok {
			if !h.Equal(target.GetCookie()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetCookie(), target.GetCookie()) {
				return false
			}
		}

	case *UserSession_Redis:
		if _, ok := target.Session.(*UserSession_Redis); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedis()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedis()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedis(), target.GetRedis()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Session != target.Session {
			return false
		}
	}

	return true
}

// Equal function
func (m *HeaderConfiguration) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HeaderConfiguration)
	if !ok {
		that2, ok := that.(HeaderConfiguration)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetIdTokenHeader(), target.GetIdTokenHeader()) != 0 {
		return false
	}

	if strings.Compare(m.GetAccessTokenHeader(), target.GetAccessTokenHeader()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *DiscoveryOverride) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DiscoveryOverride)
	if !ok {
		that2, ok := that.(DiscoveryOverride)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAuthEndpoint(), target.GetAuthEndpoint()) != 0 {
		return false
	}

	if strings.Compare(m.GetTokenEndpoint(), target.GetTokenEndpoint()) != 0 {
		return false
	}

	if strings.Compare(m.GetJwksUri(), target.GetJwksUri()) != 0 {
		return false
	}

	if len(m.GetScopes()) != len(target.GetScopes()) {
		return false
	}
	for idx, v := range m.GetScopes() {

		if strings.Compare(v, target.GetScopes()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetResponseTypes()) != len(target.GetResponseTypes()) {
		return false
	}
	for idx, v := range m.GetResponseTypes() {

		if strings.Compare(v, target.GetResponseTypes()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetSubjects()) != len(target.GetSubjects()) {
		return false
	}
	for idx, v := range m.GetSubjects() {

		if strings.Compare(v, target.GetSubjects()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetIdTokenAlgs()) != len(target.GetIdTokenAlgs()) {
		return false
	}
	for idx, v := range m.GetIdTokenAlgs() {

		if strings.Compare(v, target.GetIdTokenAlgs()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAuthMethods()) != len(target.GetAuthMethods()) {
		return false
	}
	for idx, v := range m.GetAuthMethods() {

		if strings.Compare(v, target.GetAuthMethods()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetClaims()) != len(target.GetClaims()) {
		return false
	}
	for idx, v := range m.GetClaims() {

		if strings.Compare(v, target.GetClaims()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *JwksOnDemandCacheRefreshPolicy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JwksOnDemandCacheRefreshPolicy)
	if !ok {
		that2, ok := that.(JwksOnDemandCacheRefreshPolicy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Policy.(type) {

	case *JwksOnDemandCacheRefreshPolicy_Never:
		if _, ok := target.Policy.(*JwksOnDemandCacheRefreshPolicy_Never); !ok {
			return false
		}

		if h, ok := interface{}(m.GetNever()).(equality.Equalizer); ok {
			if !h.Equal(target.GetNever()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetNever(), target.GetNever()) {
				return false
			}
		}

	case *JwksOnDemandCacheRefreshPolicy_Always:
		if _, ok := target.Policy.(*JwksOnDemandCacheRefreshPolicy_Always); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAlways()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAlways()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAlways(), target.GetAlways()) {
				return false
			}
		}

	case *JwksOnDemandCacheRefreshPolicy_MaxIdpReqPerPollingInterval:
		if _, ok := target.Policy.(*JwksOnDemandCacheRefreshPolicy_MaxIdpReqPerPollingInterval); !ok {
			return false
		}

		if m.GetMaxIdpReqPerPollingInterval() != target.GetMaxIdpReqPerPollingInterval() {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Policy != target.Policy {
			return false
		}
	}

	return true
}

// Equal function
func (m *Ldap) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Ldap)
	if !ok {
		that2, ok := that.(Ldap)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if strings.Compare(m.GetUserDnTemplate(), target.GetUserDnTemplate()) != 0 {
		return false
	}

	if strings.Compare(m.GetMembershipAttributeName(), target.GetMembershipAttributeName()) != 0 {
		return false
	}

	if len(m.GetAllowedGroups()) != len(target.GetAllowedGroups()) {
		return false
	}
	for idx, v := range m.GetAllowedGroups() {

		if strings.Compare(v, target.GetAllowedGroups()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetPool()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPool()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPool(), target.GetPool()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *PassThroughAuth) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PassThroughAuth)
	if !ok {
		that2, ok := that.(PassThroughAuth)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConfig(), target.GetConfig()) {
			return false
		}
	}

	switch m.Protocol.(type) {

	case *PassThroughAuth_Grpc:
		if _, ok := target.Protocol.(*PassThroughAuth_Grpc); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGrpc()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpc()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpc(), target.GetGrpc()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Protocol != target.Protocol {
			return false
		}
	}

	return true
}

// Equal function
func (m *PassThroughGrpc) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PassThroughGrpc)
	if !ok {
		that2, ok := that.(PassThroughGrpc)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetAddress(), target.GetAddress()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetConnectionTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnectionTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnectionTimeout(), target.GetConnectionTimeout()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigStatus)
	if !ok {
		that2, ok := that.(ExtAuthConfigStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_OidcAuthorizationCodeConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_OidcAuthorizationCodeConfig)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_OidcAuthorizationCodeConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetClientId(), target.GetClientId()) != 0 {
		return false
	}

	if strings.Compare(m.GetClientSecret(), target.GetClientSecret()) != 0 {
		return false
	}

	if strings.Compare(m.GetIssuerUrl(), target.GetIssuerUrl()) != 0 {
		return false
	}

	if len(m.GetAuthEndpointQueryParams()) != len(target.GetAuthEndpointQueryParams()) {
		return false
	}
	for k, v := range m.GetAuthEndpointQueryParams() {

		if strings.Compare(v, target.GetAuthEndpointQueryParams()[k]) != 0 {
			return false
		}

	}

	if len(m.GetTokenEndpointQueryParams()) != len(target.GetTokenEndpointQueryParams()) {
		return false
	}
	for k, v := range m.GetTokenEndpointQueryParams() {

		if strings.Compare(v, target.GetTokenEndpointQueryParams()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetAppUrl(), target.GetAppUrl()) != 0 {
		return false
	}

	if strings.Compare(m.GetCallbackPath(), target.GetCallbackPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetLogoutPath(), target.GetLogoutPath()) != 0 {
		return false
	}

	if strings.Compare(m.GetAfterLogoutUrl(), target.GetAfterLogoutUrl()) != 0 {
		return false
	}

	if len(m.GetScopes()) != len(target.GetScopes()) {
		return false
	}
	for idx, v := range m.GetScopes() {

		if strings.Compare(v, target.GetScopes()[idx]) != 0 {
			return false
		}

	}

	if h, ok := interface{}(m.GetSession()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSession()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSession(), target.GetSession()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaders()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaders()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaders(), target.GetHeaders()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDiscoveryOverride()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDiscoveryOverride()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDiscoveryOverride(), target.GetDiscoveryOverride()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDiscoveryPollInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDiscoveryPollInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDiscoveryPollInterval(), target.GetDiscoveryPollInterval()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetJwksCacheRefreshPolicy()).(equality.Equalizer); ok {
		if !h.Equal(target.GetJwksCacheRefreshPolicy()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetJwksCacheRefreshPolicy(), target.GetJwksCacheRefreshPolicy()) {
			return false
		}
	}

	if strings.Compare(m.GetSessionIdHeaderName(), target.GetSessionIdHeaderName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetUserinfoUrl(), target.GetUserinfoUrl()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetCacheTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCacheTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCacheTimeout(), target.GetCacheTimeout()) {
			return false
		}
	}

	switch m.ValidationType.(type) {

	case *ExtAuthConfigSpec_AccessTokenValidationConfig_Jwt:
		if _, ok := target.ValidationType.(*ExtAuthConfigSpec_AccessTokenValidationConfig_Jwt); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwt()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwt()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwt(), target.GetJwt()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_AccessTokenValidationConfig_Introspection:
		if _, ok := target.ValidationType.(*ExtAuthConfigSpec_AccessTokenValidationConfig_Introspection); !ok {
			return false
		}

		if h, ok := interface{}(m.GetIntrospection()).(equality.Equalizer); ok {
			if !h.Equal(target.GetIntrospection()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetIntrospection(), target.GetIntrospection()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ValidationType != target.ValidationType {
			return false
		}
	}

	switch m.ScopeValidation.(type) {

	case *ExtAuthConfigSpec_AccessTokenValidationConfig_RequiredScopes:
		if _, ok := target.ScopeValidation.(*ExtAuthConfigSpec_AccessTokenValidationConfig_RequiredScopes); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRequiredScopes()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRequiredScopes()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRequiredScopes(), target.GetRequiredScopes()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ScopeValidation != target.ScopeValidation {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_OAuth2Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_OAuth2Config)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_OAuth2Config)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.OauthType.(type) {

	case *ExtAuthConfigSpec_OAuth2Config_OidcAuthorizationCode:
		if _, ok := target.OauthType.(*ExtAuthConfigSpec_OAuth2Config_OidcAuthorizationCode); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOidcAuthorizationCode()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOidcAuthorizationCode()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOidcAuthorizationCode(), target.GetOidcAuthorizationCode()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_OAuth2Config_AccessTokenValidationConfig:
		if _, ok := target.OauthType.(*ExtAuthConfigSpec_OAuth2Config_AccessTokenValidationConfig); !ok {
			return false
		}

		if h, ok := interface{}(m.GetAccessTokenValidationConfig()).(equality.Equalizer); ok {
			if !h.Equal(target.GetAccessTokenValidationConfig()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetAccessTokenValidationConfig(), target.GetAccessTokenValidationConfig()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.OauthType != target.OauthType {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_ApiKeyAuthConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_ApiKeyAuthConfig)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_ApiKeyAuthConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetValidApiKeys()) != len(target.GetValidApiKeys()) {
		return false
	}
	for k, v := range m.GetValidApiKeys() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetValidApiKeys()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetValidApiKeys()[k]) {
				return false
			}
		}

	}

	if strings.Compare(m.GetHeaderName(), target.GetHeaderName()) != 0 {
		return false
	}

	if len(m.GetHeadersFromKeyMetadata()) != len(target.GetHeadersFromKeyMetadata()) {
		return false
	}
	for k, v := range m.GetHeadersFromKeyMetadata() {

		if strings.Compare(v, target.GetHeadersFromKeyMetadata()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_OpaAuthConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_OpaAuthConfig)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_OpaAuthConfig)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetModules()) != len(target.GetModules()) {
		return false
	}
	for k, v := range m.GetModules() {

		if strings.Compare(v, target.GetModules()[k]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetQuery(), target.GetQuery()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_Config) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_Config)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_Config)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetName()).(equality.Equalizer); ok {
		if !h.Equal(target.GetName()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetName(), target.GetName()) {
			return false
		}
	}

	switch m.AuthConfig.(type) {

	case *ExtAuthConfigSpec_Config_Oauth2:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_Oauth2); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOauth2()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOauth2()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOauth2(), target.GetOauth2()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_BasicAuth:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_BasicAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetBasicAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetBasicAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetBasicAuth(), target.GetBasicAuth()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_ApiKeyAuth:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_ApiKeyAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetApiKeyAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetApiKeyAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetApiKeyAuth(), target.GetApiKeyAuth()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_PluginAuth:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_PluginAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPluginAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPluginAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPluginAuth(), target.GetPluginAuth()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_OpaAuth:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_OpaAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetOpaAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetOpaAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetOpaAuth(), target.GetOpaAuth()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_Ldap:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_Ldap); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLdap()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLdap()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLdap(), target.GetLdap()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_Jwt:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_Jwt); !ok {
			return false
		}

		if h, ok := interface{}(m.GetJwt()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJwt()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJwt(), target.GetJwt()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_Config_PassThroughAuth:
		if _, ok := target.AuthConfig.(*ExtAuthConfigSpec_Config_PassThroughAuth); !ok {
			return false
		}

		if h, ok := interface{}(m.GetPassThroughAuth()).(equality.Equalizer); ok {
			if !h.Equal(target.GetPassThroughAuth()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetPassThroughAuth(), target.GetPassThroughAuth()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.AuthConfig != target.AuthConfig {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetIssuer(), target.GetIssuer()) != 0 {
		return false
	}

	switch m.JwksSourceSpecifier.(type) {

	case *ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_RemoteJwks_:
		if _, ok := target.JwksSourceSpecifier.(*ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_RemoteJwks_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRemoteJwks()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRemoteJwks()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRemoteJwks(), target.GetRemoteJwks()) {
				return false
			}
		}

	case *ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_LocalJwks_:
		if _, ok := target.JwksSourceSpecifier.(*ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_LocalJwks_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocalJwks()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocalJwks()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocalJwks(), target.GetLocalJwks()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.JwksSourceSpecifier != target.JwksSourceSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig_IntrospectionValidation) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig_IntrospectionValidation)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig_IntrospectionValidation)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetIntrospectionUrl(), target.GetIntrospectionUrl()) != 0 {
		return false
	}

	if strings.Compare(m.GetClientId(), target.GetClientId()) != 0 {
		return false
	}

	if strings.Compare(m.GetClientSecret(), target.GetClientSecret()) != 0 {
		return false
	}

	if strings.Compare(m.GetUserIdAttributeName(), target.GetUserIdAttributeName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig_ScopeList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig_ScopeList)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig_ScopeList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetScope()) != len(target.GetScope()) {
		return false
	}
	for idx, v := range m.GetScope() {

		if strings.Compare(v, target.GetScope()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_RemoteJwks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_RemoteJwks)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_RemoteJwks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetUrl(), target.GetUrl()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetRefreshInterval()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRefreshInterval()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRefreshInterval(), target.GetRefreshInterval()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_LocalJwks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_LocalJwks)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_AccessTokenValidationConfig_JwtValidation_LocalJwks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetInlineString(), target.GetInlineString()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ExtAuthConfigSpec_ApiKeyAuthConfig_KeyMetadata) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ExtAuthConfigSpec_ApiKeyAuthConfig_KeyMetadata)
	if !ok {
		that2, ok := that.(ExtAuthConfigSpec_ApiKeyAuthConfig_KeyMetadata)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetUsername(), target.GetUsername()) != 0 {
		return false
	}

	if len(m.GetMetadata()) != len(target.GetMetadata()) {
		return false
	}
	for k, v := range m.GetMetadata() {

		if strings.Compare(v, target.GetMetadata()[k]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *BasicAuth_Apr) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BasicAuth_Apr)
	if !ok {
		that2, ok := that.(BasicAuth_Apr)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetUsers()) != len(target.GetUsers()) {
		return false
	}
	for k, v := range m.GetUsers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetUsers()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetUsers()[k]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *BasicAuth_Apr_SaltedHashedPassword) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*BasicAuth_Apr_SaltedHashedPassword)
	if !ok {
		that2, ok := that.(BasicAuth_Apr_SaltedHashedPassword)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetSalt(), target.GetSalt()) != 0 {
		return false
	}

	if strings.Compare(m.GetHashedPassword(), target.GetHashedPassword()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *UserSession_InternalSession) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UserSession_InternalSession)
	if !ok {
		that2, ok := that.(UserSession_InternalSession)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	return true
}

// Equal function
func (m *UserSession_RedisSession) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UserSession_RedisSession)
	if !ok {
		that2, ok := that.(UserSession_RedisSession)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if strings.Compare(m.GetKeyPrefix(), target.GetKeyPrefix()) != 0 {
		return false
	}

	if strings.Compare(m.GetCookieName(), target.GetCookieName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetAllowRefreshing()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAllowRefreshing()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAllowRefreshing(), target.GetAllowRefreshing()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *UserSession_CookieOptions) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UserSession_CookieOptions)
	if !ok {
		that2, ok := that.(UserSession_CookieOptions)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMaxAge()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxAge()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxAge(), target.GetMaxAge()) {
			return false
		}
	}

	if m.GetNotSecure() != target.GetNotSecure() {
		return false
	}

	if h, ok := interface{}(m.GetPath()).(equality.Equalizer); ok {
		if !h.Equal(target.GetPath()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetPath(), target.GetPath()) {
			return false
		}
	}

	if strings.Compare(m.GetDomain(), target.GetDomain()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Ldap_ConnectionPool) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Ldap_ConnectionPool)
	if !ok {
		that2, ok := that.(Ldap_ConnectionPool)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMaxSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMaxSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMaxSize(), target.GetMaxSize()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetInitialSize()).(equality.Equalizer); ok {
		if !h.Equal(target.GetInitialSize()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetInitialSize(), target.GetInitialSize()) {
			return false
		}
	}

	return true
}