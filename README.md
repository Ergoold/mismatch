# mismatch

Prints all mismatched parentheses in stdin and their locations.

```
$ echo '()' >file
$ mismatch <file

```

The size of the internal read buffer (in bytes) is taken from the environment variable BUFFER_SIZE.
The default buffer size is 64 bytes.

```
$ echo '(()' >file
$ BUFFER_SIZE=1024 mismatch <file
mismatch: /dev/stdin:1:1: unclosed '('
 1 | (()
     ^ here
```

The initially-allocated stack depth for keeping mismatching parentheses is taken from the environment variable INITIAL_STACK_DEPTH.
The default stack depth is 4.

```
$ echo '())' >file
$ INITIAL_STACK_DEPTH=1 mismatch <file
mismatch: /dev/stdin:1:1: unopened ')'
1 | ())
      ^ here
```

## Gotchas

Detailed diagnostics are not printed for pipes and other non-seekable streams

```
$ echo ')(' | mismatch
mismatch: /dev/stdin:1:1: unopened ')'
mismatch: seek /dev/stdin: illegal seek

mismatch: /dev/stdin:1:2: unclosed '('
```
