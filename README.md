# yj

It slurps YAML in through its standard input, and chucks out freshly ground
JSON to its standard output!

Well, that's what it does if you call it as `yj`. If you make a symlink to it
called `jy`, it will do the opposite.

It neither slices nor dices, but it's super handy for piping YAML output from
an API via [`jq`](https://stedolan.github.io/jq/):

```
$ curl https://someapi | yj | jq 'an awesome transformation' | jy
```
