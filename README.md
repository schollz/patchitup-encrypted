<p align="center">
<img
    src="https://raw.githubusercontent.com/schollz/patchitup-encrypted/master/.github/logo.png"
    width="260px" border="0" alt="patchitup-encrypted">
<br>
<a href="https://github.com/schollz/patchitup-encrypted/releases/latest"><img src="https://img.shields.io/badge/version-0.1.0-brightgreen.svg?style=flat-square" alt="Version"></a>
<img src="https://img.shields.io/badge/coverage-75%25-green.svg?style=flat-square" alt="Code Coverage">
<a href="https://godoc.org/github.com/schollz/patchitup-encrypted/patchitup-encrypted"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square" alt="Code Coverage"></a>
<a href="https://www.paypal.me/ZackScholl/5.00"><img src="https://img.shields.io/badge/donate-$5-brown.svg?style=flat-square" alt="Donate"></a>
</p>

<p align="center">Backup your file to a cloud server using minimum bandwidth.</p>

*patchitup-encrypted* is a way to keep the cloud up-to-date through incremental patches. In a nutshell, this is a pure-Golang library and a CLI tool for creating a client+server that exchange incremental gzipped patches to overwrite a remote copy to keep it up-to-date with the client's local file. The files on the server stay encrypted so that only the client ever has access to them.

<em><strong>Why?</strong></em> I wrote this program to reduce the bandwidth usage when backing up SQLite databases to a remote server from Raspberry Pis. I have deployed some software on Raspberry Pis that periodically [dumps the database to SQL text](http://www.sqlitetutorial.net/sqlite-dump/). Since Raspberry Pi's can die sometimes, I want to keep their data stored remotely. As the databases can get fairly large, a patch from SQL text will only ever be the changed/new records. *patchitup-encrypted* allows the client to just send to the cloud only the changed/new records and still maintain the exact copy on the cloud. This can massively reduce bandwidth between the client and the cloud. 

<em><strong>Why not git?</strong></em> While *git* basically  does this already, its not terribly easy to setup a *git* server to support multiple users (though [gitolite](https://github.com/sitaramc/gitolite) does a great job of simplifying the process).  Also, most of the features of *git* are not necessary for my use-case.

# Quickstart

In addition to being a Golang library, the *patchitup-encrypted* is a server+client. To try it, first install *patchitup-encrypted* with Go:

```
$ go install -u -v github.com/schollz/patchitup-encrypted/...
```

Then start a *patchitup-encrypted* server:

```
$ patchitup-encrypted -host
Running at http://0.0.0.0:8002
```

Then you can patch a file:

```
$ patchitup-encrypted -u me -s http://localhost:8002 -f SOMEFILE
2018-02-23 08:56:44 [INFO] patched 2.4 kB (62.8%) to remote 'SOMEFILE' for 'me'

$ vim SOMEFILE # make some edits

$ patchitup-encrypted -u me -s http://localhost:8002 -f SOMEFILE
2018-02-23 08:57:40 [INFO] patched 408 B (9.9%) to remote 'SOMEFILE' for 'me'
```

The first time you patch will basically just send up the gzipped file. Subsequent edits will just send up the patches. The percentage (e.g. `9.9%`) specifies the percentage of the entire file size that is being sent (to get an idea of bandwidth savings). The server also will log bandwidth usage.


# Roadmap

I would love PRs.

Some ideas I'd like to add:

- [x] Built-in security (authentication tokens?)
- [x] Encryption option (to keep data on server private)

# License

MIT

# Thanks

Logo designed by <a rel="nofollow" target="_blank" href="https://www.vecteezy.com">www.Vecteezy.com</a>
