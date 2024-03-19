# Get_Lrc
show current lrc of YesPlayMusic on status bar for Ubuntu gnome.

> Only work for YesPlayMusic.

This project is written in Go.

## Steps for use
1. Download release file get_lrc-xxx-linux-x64.tar.xz and extract it.

2. Grant it executable permission and move it to `/usr/local/bin` dir.

3. Install gnome extension `executor@raujonas.github.io` and config it add lrc output cmd: `echo "   🎤 $(get_lrc yesplay)    "`, like this:
![截图 2024-03-19 14-30-59](https://github.com/MarsSwimmer/get_lrc/assets/146618222/6a0ce857-9951-4806-95ce-b72c4af59f1d)

4. Open YesPlayMusic and play a song with LRC, you will see current lrc show on topbar, like this:
![preview](https://github.com/MarsSwimmer/get_lrc/assets/146618222/b6e43f94-e9b3-41cb-9e08-06621d936c56)


## Command Usage And Flags
```
get_lrc [command] [flags]
```

## Available Commands
```
cmd         flags                                           

yesplay                                    get YesPlayMusic's current lrc.
            --currentTrackUr               optional, specific the currentTrackUrl of yesplaymusic (default "http://127.0.0.1:10754/lyric?id=").
            --playerUrl string             optional, specific the playerUrl of yesplaymusic (default "http://127.0.0.1:27232/player").
            -d, --defaultValue             optional, specific the default value when internal error and not show error message.
            --delay                        optional, specific the progress delay time for current lrc's sentence show, ms unit (default 1000).
            --prefix                       optional, specific the prefix of lrc's sentence (default "歌词：").
            --sentenceCount                optional, specific show current sentence count (default 1).
            --showAll                      optional, show all sentence of lrc.
            --showErrMsg                   optional, show error message when internal error.
            --split                        optional, specific the split to merge multi lrc's sentence (default "|").

help                                       Help about any command.
```
