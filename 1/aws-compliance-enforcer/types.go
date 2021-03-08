package main

import "time"

//EventResponse is a response to the event
type EventResponse struct {
	Message string `json:"message"`
}

//SNSEventMessage is the message part of the SNS payload
type SNSEventMessage struct {
	S3Bucket    string   `json:"s3Bucket"`
	S3ObjectKey []string `json:"s3ObjectKey"`
}

//CloudTrailEvent is an AWS event generated by cloudtrail
type CloudTrailEvent struct {
	Records []CloudTrailEventRecord `json:"records" binding:"required"`
}

//CloudTrailEventRecord is a record of the cloud trail event
type CloudTrailEventRecord struct {
	EventVersion string `json:"eventVersion"`
	UserIdentity struct {
		Type        string `json:"type"`
		UserName    string `json:"userName"`
		PrincipalID string `json:"principalId"`
		Arn         string `json:"arn"`
		AccountID   string `json:"accountId"`
		AccessKeyID string `json:"accessKeyId"`
	} `json:"userIdentity"`
	EventTime          time.Time              `json:"eventTime"`
	EventSource        string                 `json:"eventSource"`
	EventName          string                 `json:"eventName"`
	AwsRegion          string                 `json:"awsRegion"`
	SourceIPAddress    string                 `json:"sourceIPAddress"`
	UserAgent          string                 `json:"userAgent"`
	RequestParameters  map[string]interface{} `json:"requestParameters"`
	ResponseElements   map[string]interface{} `json:"responseElements"`
	RequestID          string                 `json:"requestID"`
	EventID            string                 `json:"eventID"`
	EventType          string                 `json:"eventType"`
	RecipientAccountID string                 `json:"recipientAccountId"`
}

//CloudTrailSecurityGroupEvent are present when there is a AuthorizeSecurityGroupIngress event
type CloudTrailSecurityGroupEvent struct {
	GroupID       string `json:"groupId"`
	IPPermissions struct {
		Items []struct {
			IPProtocol string `json:"ipProtocol"`
			FromPort   int64  `json:"fromPort"`
			ToPort     int64  `json:"toPort"`
			Groups     struct {
			} `json:"groups"`
			IPRanges struct {
				Items []struct {
					CidrIP string `json:"cidrIp"`
				} `json:"items"`
			} `json:"ipRanges"`
			Ipv6Ranges struct {
			} `json:"ipv6Ranges"`
			PrefixListIds struct {
			} `json:"prefixListIds"`
		} `json:"items"`
	} `json:"ipPermissions"`
}