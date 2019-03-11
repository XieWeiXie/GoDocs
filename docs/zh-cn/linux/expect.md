# expect 

> 自动交互，不可单独运行，需配合 shell 脚本运行 

- apt-get install expect
- spawn
- expect
- interact

1. one.sh
``` 
#!/usr/bin/expect

spawn  ssh *
expect {
  "(yes/no)" {send "yes\r"; exp_continue}
  "password:" {send "7qVvwPVZFS\r"}
}

interact

```

2. two.sh
``` 
expect ./one.sh

```