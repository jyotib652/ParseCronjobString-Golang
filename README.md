## ParseCronjobString-Golang
This is a Golang application to Parse a cronjob expression string.
Here only the standard cron format is considered with five time fields (minute, hour, day of month, month, and day of week) plus a command, and the special
time strings such as "@yearly" is not handeled. The input will be on a single line like this:
```
~$ your-program "*/15 0 1,15 * 1-5 /usr/bin/find"
```

And the output would be something like this:
```
minute          [0 15 30 45]
hour            [0]
day of month    [1 15]
month           [1 2 3 4 5 6 7 8 9 10 11 12]
day of week     [1 2 3 4 5]
command         /usr/bin/find
```

* To run the program first go to the directiory where the "main.go" is then issue this command:
  ```
  ~$ go run main.go "*/15 0 1,15 * 1-5 /usr/bin/find"
  ```

  Your cron job expression string should be in double quote to represent it in string format.
  So first the "go run main.go" command then the cron job expression in string format(within double quote "")
  
  Here, "*/15 0 1,15 * 1-5 /usr/bin/find" is the cron job expression in string format.
  

