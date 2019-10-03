#!/bin/sh
$rval = shuf -i 6001-9999 -n 1
sh /code/main/main $rval 10.0.0.2:6000
