# run the commands side by side
if [[ $TERM == *tmux* ]]; then
    tmux kill-pane -a -t $TMUX_PANE
    tmux send-keys "make runserver" Enter
    # tmux send-prefix

    tmux split-window -h -c "#{pane_current_path}"
    tmux send-keys "make watchsrc" Enter
else
    make runserver &
    make watchsrc
fi

