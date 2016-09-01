#!/bin/bash
init(){
	echo "$title"
    mkdir $title
	echo "package leetcode " >"$title"/"$title".go
	echo "package leetcode 
	
import (
	\"testing\"
)

func Test(t *testing.T) {
	
}
" >"$title"/"$title"_test.go
}

remove(){
	rm -rf $title
}

case $1 in
	"init")
	title=$2
	init
	;;
	"rm")
	title=$2
	remove
	;;
	*)
	echo "add a cmd to run "
	;;
esac

