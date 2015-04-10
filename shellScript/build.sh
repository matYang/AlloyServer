#!/bin/bash
echo "building golang server"

HOME="/home/ubuntu"
SOURCEFOLDERPATH=$GOPATH"/src/github.com/matYang/AlloyServer"
INSTALLPATH=$SOURCEFOLDERPATH"/alloyServer"
PYTHONSOURCEPATH=$SOURCEFOLDERPATH"/pythonScript"
ALLOYSOURCEPATH=$SOURCEFOLDERPATH"/alloyScript"

SERVERFOLDERPATH=$HOME"/server"
ALSPATHBASE=$SERVERFOLDERPATH"/als"

WORKLOAD="3"
if [ -z "$1" ]
then
    WORKLOAD = "3"
else
    WORKLOAD = $1
fi

#build latest golang source
cd $SOURCEFOLDERPATH
git pull origin master
cd $INSTALLPATH
go install

#check for install result, if failed to install do to proceed
if [[ $? -ne 0 ]] ; then
    exit 1
fi

#kill the old server process
ps aux | grep -ie alloyServer | awk '{print $2}' | xargs kill -9

#prepare file structure
for (( i=0; i<${WORKLOAD}; i++ ));
do
       rm -r $ALSPATHBASE$i
       mkdir $ALSPATHBASE$i
       cp $PYTHONSOURCEPATH/* $ALSPATHBASE$i
       cp $ALLOYSOURCEPATH/*  $ALSPATHBASE$i
done

#restart the server process
cd $SERVERFOLDERPATH
nohup alloyServer & > log
