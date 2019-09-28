#!/bin/bash

cmd="./a.out"
cmd="python3 ./lcs.py"

function test()
{
    echo -e "$(tput setaf 3)testcase $1 $2$(tput sgr0)"
    set -x
    lcs=$($cmd $1 $2)
    set +x
    if [ $(echo "$lcs" | tail -n 1) == "$3" ]; then
        echo -e "$(tput setaf 2)OK$(tput sgr0)\n"
    else
        echo -e "$(tput setaf 1)NG$(tput sgr0)\n"
    fi
}

test 1234 1234 1234
test 1234 1254 124
test XMJYAU MZJAWXU MJAU
