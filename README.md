# mismatch

Prints the difference between the number of opening and closing parentheses in stdin.

```
$ echo '()' | mismatch
0
```

```
$ echo '(()' | mismatch
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
