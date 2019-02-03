# TorrentSize
Get total size of a folder full of .torrent

# Usage

You can find releases for various operating systems in the [releases tab](https://github.com/The-Eye-Team/TorrentSize/releases).

Download one, then make it executable:

```
chmod +x TorrentSize
```

Sample usage with a folder called `torrents` with your torrent files inside:

```
./TorrentSize -i torrents/
```

You can see the options with the `-h` flag:

```
TorrentSize [-h|--help] -i|--input "<value>" [-j|--concurrency
                   <integer>]

                   Get total size of a folder full of .torrent

Arguments:

  -h  --help         Print help information
  -i  --input        Input directory
  -j  --concurrency  Concurrency. Default: 4
  ```
 
# Build

```
git clone https://github.com/The-Eye-Team/TorrentSize.git && cd TorrentSize
```

```
go get ./...
```

```
go build .
```

[![The-Eye.eu](https://the-eye.eu/public/.css/logo3_x300.png)](https://the-eye.eu)
