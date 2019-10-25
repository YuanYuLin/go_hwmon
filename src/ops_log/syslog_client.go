package ops_log

import "fmt"
//import "log/syslog"

func Info(mask uint32, format string, a ...interface{}) {
	log_str := fmt.Sprintf(format, a)
	fmt.Println(log_str)
	return
	/*
	log, _ := syslog.Dial("", "", syslog.LOG_USER|syslog.LOG_INFO, "iopc_go")
	log.Info(log_str)
	log.Close()
	*/
}

func Debug(mask uint32, format string, a ...interface{}) {
	log_str := fmt.Sprintf(format, a)
	fmt.Println(log_str)
	return
	/*
	log, _ := syslog.Dial("", "", syslog.LOG_USER|syslog.LOG_DEBUG, "iopc_go")
	log.Debug(log_str)
	log.Close()
	*/
}

func Error(mask uint32, format string, a ...interface{}) {
	log_str := fmt.Sprintf(format, a)
	fmt.Println(log_str)
	return
	/*
	log, _ := syslog.Dial("", "", syslog.LOG_USER|syslog.LOG_ERR, "iopc_go")
	log.Err(log_str)
	log.Close()
	*/
}

