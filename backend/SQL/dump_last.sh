cd dump
if test -f "./last.sql"; then
    echo "rewrite last dump?(y/n)"
    read rewrite_flag
    if [ $rewrite_flag == "y" ]; then
        echo "removing last.sql"
        rm -f "./last.sql"
        echo "making new dump"
        docker exec -t postgres pg_dumpall -c -U postgres > last.sql
        echo "dump complete"
    else
        echo "exit the program"
        exit 0
    fi
else
    echo "making new dump"
    docker exec -t postgres pg_dumpall -c -U postgres > last.sql
    echo "dump complete"
fi
