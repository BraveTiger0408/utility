# Utilities

This program is designed to ...

Processing outline
------------------
* ...

Testing
-------
    $ go test
    $ go test -bench=.
    $ go test -cover

Compiling
---------

    $ go build rules_engine.go

(Will produce an executable `rules_engine`)

Running
-------

    $ ./rules_engine

Installing
---------

    $ go install rules_engine.go

then run `rules_engine`.  

Docker
------

All you have to do is run `./docker_build.sh` and boom!

To run the docker image locally:

    $ docker run --rm -it -p 8000:8000 dstroot/rules_engine


TODO
----
[ ] Hold rules are different between CT and PRO!  Should be easy - just build into select SQL statement
[ ] Rule to put any items > 20 days into "priority hold" repeats because it is looking at other hold records older than 20 days and they are still "seen" after the record has gone to priority hold.
[ ] Handle manual rules.  This is simple - just leave the SQL statement null.
