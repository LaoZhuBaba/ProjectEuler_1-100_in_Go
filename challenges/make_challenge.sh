#!/bin/bash

# Do no convert tabs below to spaces!

if [ "$#" != "1" ]
then
	echo "missing required parameter"
	exit 255
fi 

CHALLENGE=challenge${1}

cat <<-EOF > ${CHALLENGE}.go
	//go:build c${1}
	
	package challenges

	import "euler/${CHALLENGE}"
	
	func Challenge() {
	    ${CHALLENGE}.Challenge${1}()
	}
	EOF

mkdir ../${CHALLENGE}
cat <<-EOF > ../${CHALLENGE}/${CHALLENGE}.go
	package $CHALLENGE
	
	func Challenge${1}() {
	}
	EOF
