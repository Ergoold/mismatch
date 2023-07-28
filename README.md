# mismatch

Prints the difference between the number of opening and closing parentheses in stdin.

```
$ echo '()' | mismatch
0
```

The size of the internal read buffer (in bytes) is taken from the environment variable BUFFER_SIZE.
The default buffer size is 64 bytes.

```
$ BUFFER_SIZE=1024 echo '(()' | mismatch
1
```

```
$ echo '())' | mismatch
-1
```

```
$ echo ')(' | mismatch
0
```
