package main

import "github.com/cihub/seelog"

func main() {
	defer seelog.Flush()
	//seelog.Info("Hello from Seelog!")
	// [date] [level] message
	// /var/log

	config := `
    <seelog type="sync">
        <outputs formatid="application">
			<console />
            <rollingfile type="date" filename="/tmp/seelog.log" datepattern="20060102" maxrolls="7" />
        </outputs>
        <formats>
            <format id="application" format="%UTCDate %UTCTime [%LEVEL] %Msg%n" />
        </formats>
    </seelog>
    `

	logger, err := seelog.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		panic(err)
	}
	seelog.ReplaceLogger(logger)

	seelog.Trace("Test message!")
	seelog.Tracef("Tracef: %s", "hello")
	//TRACE 2014/12/15 21:04:35 controller.go:375: Registered controller: UserSite
	// logs.Tracef("hoge")
}
