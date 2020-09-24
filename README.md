# np

NP (Number Processes) is tool for run shell script in multiple processes. It's writen in pure Go.

For example you have script:
```shell script
sleep 1
sleep 2
sleep 3
echo "1"
echo "2"
echo "3"
sleep 4 && echo "4 done"
```

Run it with four processes:
```shell script
$ np -file script.sh -n 3

#3 sleep 3
#2 sleep 2
#1 sleep 1
#1 echo "1"
1
#1 echo "2"
2
#1 echo "3"
3
#1 sleep 4 && echo "4 done"
4 done
```

NP run script lines parallel with specified processes count.
Stdout and stderr of each command attaches to np main process, but you still can use redirection and pipes.
