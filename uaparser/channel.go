package uaparser

import (
	"regexp"
	"strings"
)

type Channel struct {
	Params map[string]bool
}

type ChannelPattern struct {
	Model           string
	Regexp          *regexp.Regexp
	Regex           string
	ExactMatch      string
	UserAgentFamily string
	MajorVersion    string
}

func (channelPattern *ChannelPattern) Match(cli *Client, line string) bool {
	if channelPattern.ExactMatch != "" && strings.Contains(line, channelPattern.ExactMatch) {
		return true
	}
	if channelPattern.Regex != "" {
		matches := channelPattern.Regexp.FindStringSubmatch(line)
		if len(matches) != 0 {
			return true
		}
	}
	if channelPattern.UserAgentFamily != "" && strings.Contains(cli.UserAgent.Family, channelPattern.UserAgentFamily) {
		return true
	}
	return false
}
