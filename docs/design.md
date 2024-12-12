![test](mockup.png)

# Downloading a video
This assumes that you want to use the yt-dlp defaults, essentially calling 

``yt-dlp 'url'``

- a field for a url
- the download button
- an output file location
- a loading screen

## Changing options

I would want to run a silent run, the options would load, then fill up with possible options

- quality section
    - fields for possible formats
    - and fields for conversion required formats (ffmpeg)
    - adenoidal fields for video quality separate
    - and audio quality separate
- playlist section
    - a field for start and end of a section
- parts downloader section
    - this would have an on off switch
    - then a start and end
- a custom file name builder
    - with the possible data as buttons
    - trim on off field
- thumbnail and info section
    - on off thumbnail
    - on off embedd thumbnail 
    - on off embed metadata 
- cookies section
    - cookies from browser option
- free section, to add commands

## Save profiles

a way to save and load options profiles, these should be held in localstorage, instead of a database because I don't want to run a database
