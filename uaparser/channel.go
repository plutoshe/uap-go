package uaparser

import "regexp"

type Channel struct {
	IsWeibo, IsWechat, IsQQ, IsQQBrowser bool
}

type ChannelPattern struct {
	Regexp *regexp.Regexp
	Regex  string
}

func (channelPattern *ChannelPattern) Match(line string) bool {
	matches := channelPattern.Regexp.FindStringSubmatch(line)
	if len(matches) == 0 {
		return false
	}
	return true
}
