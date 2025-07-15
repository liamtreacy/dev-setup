## Terminal

If you are doing some piping and output, it's a good idea to add a newline at the end or the prompt might get included e.g.
`echo MY_ENCODED_STRING | base64 --decode; echo`

## TMUX

tmux ls - see active tmux sessions
tmux attach -t <name> - log in to session

Prefix + % - split current pane vertically  
Prefix + " - split pane horizontally  
Prefix + n - next window  
Prefix + p - previous window
CTRL + h,j,k,l - move between split panes
Prefix + c - new window  
Prefix + & - kill window  
Prefix + x - kill pane

### Copy mode

Prefix + [ - enter copy mode
Space - in copy mode it starts selection, enter to copy
Prefix + ] - paste from tmux buffer
Or when in copy mode you can do vim select, yank and paste

While in copy mode, naviagte using arrows.  
Space - start selection, press enter and is copied to tmux buffer.  
Prefix + ] - paste buffer

Other option is to press v, select text and then y to yank and p to paste.

## INTELLIJ

OPT+SHIFT+left/right arrow - highlight a word.  
CMD+SHIFT+left/right arrow - highlight to start/end of current line.  

## VIM

$ - go to end of line (normal mode)
dd - cut current line
G - go to end of file

### NVCHAD

Space - leader key
=======
$ - go to end of line (normal mode)  
dd - cut current line  
b - begining of word  
e - end of word  
