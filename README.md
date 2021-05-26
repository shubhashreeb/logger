# logger

 
# Simple Logger
 
This simple Logger is a wrapper on zap SugaredLogger, provides different levels of logging, debug, info, error, fatal etc.
 
 
To Create a logger, you must provide logger config as follow
```
config := logger.Configuration{
       EnableConsole:     true,
       ConsoleLevel:      logger.InfoLevel,
       ConsoleJSONFormat: false,
       EnableFile:        true,
       FileLevel:         logger.DebugLevel,
       FileJSONFormat:    false,
       FileLocation:      "log.log",
   }
```

Instantiate logger
```
   log, err := logger.NewLogger(config)
   if err != nil {
       log.Fatalf("Could not instantiate log %s", err.Error())
   }
```
 
Using logger
```
   logger.Info("This is info")
   logger.Error("This is error")
   logger.Debug("This is debug")
   logger.Fatal("fatal Error")
```

