# Get_Lrc
show current lrc of Music Player such as Yesplaymusic and LXMusic on status bar for Ubuntu gnome.

> Work for YesPlayMusic and LXMusic.

This project is written in Go.

## Steps for use
1. Download release file get_lrc-xxx-linux-x64.tar.xz and extract it.

2. Grant it executable permission and move it to `/usr/local/bin` dir.

3. Install gnome extension `executor@raujonas.github.io` and config it add lrc output cmd: `echo "   🎤 $(get_lrc)    "`, like this:
![截图 2024-03-19 14-30-59](https://github.com/MarsSwimmer/get_lrc/assets/146618222/6a0ce857-9951-4806-95ce-b72c4af59f1d)

4. Open YesPlayMusic and play a song with LRC, you will see current lrc show on topbar, like this:
![preview](https://github.com/MarsSwimmer/get_lrc/assets/146618222/b6e43f94-e9b3-41cb-9e08-06621d936c56)

5. Lx Music need enable open api from settings.
![截图 2024-03-30 01-57-23](https://github.com/MarsSwimmer/get_lrc/assets/146618222/802df09c-f1fc-49f3-8f11-558fce85dc43)


## Command Usage And Flags
- Get current lrc sentence, auto choose source from all music player.
```
get_lrc [flags]
```

- Specific music player
```
get_lrc [command] [flags] 
```

## Available Commands
```
cmd         flags  
--                                         get current lrc sentence, it will auto choose source from all music player.
            --yesplayPlayerUrl string      optional, specific the playerUrl of yesplaymusic (default "http://127.0.0.1:27232/player")
            --lxPlayerUrl string           optional, specific the playerUrl of lx music (default "http://localhost:23330/status")                            
            --showErrMsg                   optional, show error message when internal error.
            -d, --defaultValue string      optional, specific the default value when internal error and not show error message

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

lx                                         get lx's current lrc.
            --playerUrl string             optional, specific the playerUrl of yesplaymusic (default "http://127.0.0.1:23330/status").
            --showErrMsg                   optional, show error message when internal error.
            -d, --defaultValue string      optional, specific the default value when internal error and not show error message


help                                       Help about any command.
```
